package database

import (
	"context"
	"fmt"
	"time"
	"web-server/internal/models"

	"github.com/shopspring/decimal"
)

func (s *service) CreateSale(ctx context.Context, req *models.CreateSaleRequest) (*models.Sale, error) {
	// Start transaction
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	// 1. Create Sale Header
	var saleID int
	var saleDate time.Time

	// Initial total amount is 0, will be updated after inserting items
	headerQuery := `
		INSERT INTO sales (customer_id, payment_method, total_amount)
		VALUES ($1, $2, 0)
		RETURNING sale_id, sale_date
	`
	err = tx.QueryRow(ctx, headerQuery, req.CustomerID, req.PaymentMethod).Scan(&saleID, &saleDate)
	if err != nil {
		return nil, fmt.Errorf("failed to create sale header: %w", err)
	}

	totalAmount := decimal.NewFromInt(0)
	var saleItems []models.SaleItem

	// 2. Process Items
	for _, itemReq := range req.Items {
		// Get product current price and check stock
		// Use FOR UPDATE to lock the row and prevent race conditions
		var currentPrice decimal.Decimal
		var currentStock int

		prodQuery := `SELECT price, stock FROM products WHERE product_id = $1 FOR UPDATE`
		err = tx.QueryRow(ctx, prodQuery, itemReq.ProductID).Scan(&currentPrice, &currentStock)
		if err != nil {
			return nil, fmt.Errorf("product %d not found or error fetching: %w", itemReq.ProductID, err)
		}

		if currentStock < itemReq.Quantity {
			return nil, fmt.Errorf("insufficient stock for product %d", itemReq.ProductID)
		}

		// Insert Sale Item
		// Note: subtotal is generated always, so we don't insert it, but we can read it back if needed.
		// For simplicity in this response, we calculate it here for the totalAmount logic.
		quantityDecimal := decimal.NewFromInt(int64(itemReq.Quantity))
		itemSubtotal := currentPrice.Mul(quantityDecimal)
		totalAmount = totalAmount.Add(itemSubtotal)

		itemQuery := `
			INSERT INTO sale_items (sale_id, product_id, quantity, unit_price)
			VALUES ($1, $2, $3, $4)
			RETURNING sale_item_id, subtotal
		`

		var saleItemID int
		var dbSubtotal decimal.Decimal
		err = tx.QueryRow(ctx, itemQuery, saleID, itemReq.ProductID, itemReq.Quantity, currentPrice).Scan(&saleItemID, &dbSubtotal)
		if err != nil {
			return nil, fmt.Errorf("failed to insert sale item: %w", err)
		}

		// Update Stock
		updateStockQuery := `UPDATE products SET stock = stock - $1 WHERE product_id = $2`
		_, err = tx.Exec(ctx, updateStockQuery, itemReq.Quantity, itemReq.ProductID)
		if err != nil {
			return nil, fmt.Errorf("failed to update stock for product %d: %w", itemReq.ProductID, err)
		}

		saleItems = append(saleItems, models.SaleItem{
			ID:        saleItemID,
			SaleID:    saleID,
			ProductID: itemReq.ProductID,
			Quantity:  itemReq.Quantity,
			UnitPrice: currentPrice,
			Subtotal:  dbSubtotal,
		})
	}

	// 3. Update Sale Total Amount
	updateTotalQuery := `UPDATE sales SET total_amount = $1 WHERE sale_id = $2`
	_, err = tx.Exec(ctx, updateTotalQuery, totalAmount, saleID)
	if err != nil {
		return nil, fmt.Errorf("failed to update sale total: %w", err)
	}

	// Commit transaction
	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return &models.Sale{
		ID:            saleID,
		CustomerID:    req.CustomerID,
		TotalAmount:   totalAmount,
		PaymentMethod: req.PaymentMethod,
		SaleDate:      saleDate,
		Items:         saleItems,
	}, nil
}

func (s *service) GetSale(ctx context.Context, id int) (*models.Sale, error) {
	// Get Header
	saleQuery := `
		SELECT sale_id, customer_id, total_amount, payment_method, sale_date
		FROM sales WHERE sale_id = $1
	`
	var sale models.Sale
	err := s.db.QueryRow(ctx, saleQuery, id).Scan(
		&sale.ID,
		&sale.CustomerID,
		&sale.TotalAmount,
		&sale.PaymentMethod,
		&sale.SaleDate,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get sale header: %w", err)
	}

	// Get Items
	itemsQuery := `
		SELECT sale_item_id, sale_id, product_id, quantity, unit_price, subtotal
		FROM sale_items WHERE sale_id = $1
	`
	rows, err := s.db.Query(ctx, itemsQuery, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get sale items: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var item models.SaleItem
		if err := rows.Scan(
			&item.ID,
			&item.SaleID,
			&item.ProductID,
			&item.Quantity,
			&item.UnitPrice,
			&item.Subtotal,
		); err != nil {
			return nil, err
		}
		sale.Items = append(sale.Items, item)
	}

	return &sale, nil
}
