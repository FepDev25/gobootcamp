package database

import (
	"context"
	"fmt"
	"web-server/internal/models"
)

func (s *service) CreateProduct(ctx context.Context, product *models.Product) error {
	query := `
		INSERT INTO products (category_id, barcode, name, price, stock, is_active)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING product_id, created_at
	`
	err := s.db.QueryRow(ctx, query,
		product.CategoryID,
		product.Barcode,
		product.Name,
		product.Price,
		product.Stock,
		product.IsActive,
	).Scan(&product.ID, &product.CreatedAt)

	if err != nil {
		return fmt.Errorf("failed to create product: %w", err)
	}
	return nil
}

func (s *service) GetProduct(ctx context.Context, id int) (*models.Product, error) {
	query := `
		SELECT product_id, category_id, barcode, name, price, stock, is_active, created_at
		FROM products
		WHERE product_id = $1
	`
	var p models.Product
	err := s.db.QueryRow(ctx, query, id).Scan(
		&p.ID,
		&p.CategoryID,
		&p.Barcode,
		&p.Name,
		&p.Price,
		&p.Stock,
		&p.IsActive,
		&p.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get product: %w", err)
	}
	return &p, nil
}

func (s *service) GetAllProducts(ctx context.Context) ([]models.Product, error) {
	query := `
		SELECT product_id, category_id, barcode, name, price, stock, is_active, created_at
		FROM products
	`
	rows, err := s.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get products: %w", err)
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(
			&p.ID,
			&p.CategoryID,
			&p.Barcode,
			&p.Name,
			&p.Price,
			&p.Stock,
			&p.IsActive,
			&p.CreatedAt,
		); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (s *service) UpdateProduct(ctx context.Context, product *models.Product) error {
	query := `
		UPDATE products
		SET category_id = $1, barcode = $2, name = $3, price = $4, stock = $5, is_active = $6
		WHERE product_id = $7
	`
	_, err := s.db.Exec(ctx, query,
		product.CategoryID,
		product.Barcode,
		product.Name,
		product.Price,
		product.Stock,
		product.IsActive,
		product.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to update product: %w", err)
	}
	return nil
}

func (s *service) DeleteProduct(ctx context.Context, id int) error {
	query := `DELETE FROM products WHERE product_id = $1`
	_, err := s.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}
	return nil
}
