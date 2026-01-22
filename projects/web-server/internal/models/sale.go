package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Sale struct {
	ID            int             `json:"sale_id"`
	CustomerID    *int            `json:"customer_id"`
	TotalAmount   decimal.Decimal `json:"total_amount"`
	PaymentMethod string          `json:"payment_method"`
	SaleDate      time.Time       `json:"sale_date"`
	Items         []SaleItem      `json:"items,omitempty"`
}

type SaleItem struct {
	ID        int             `json:"sale_item_id"`
	SaleID    int             `json:"sale_id"`
	ProductID int             `json:"product_id"`
	Quantity  int             `json:"quantity"`
	UnitPrice decimal.Decimal `json:"unit_price"`
	Subtotal  decimal.Decimal `json:"subtotal"`
}

type CreateSaleRequest struct {
	CustomerID    *int                    `json:"customer_id"`
	PaymentMethod string                  `json:"payment_method"`
	Items         []CreateSaleItemRequest `json:"items"`
}

type CreateSaleItemRequest struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}
