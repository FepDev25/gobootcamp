package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Product struct {
	ID         int             `json:"product_id"`
	CategoryID *int            `json:"category_id"`
	Barcode    *string         `json:"barcode"`
	Name       string          `json:"name"`
	Price      decimal.Decimal `json:"price"`
	Stock      int             `json:"stock"`
	IsActive   bool            `json:"is_active"`
	CreatedAt  time.Time       `json:"created_at"`
}
