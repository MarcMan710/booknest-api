package models

import (
	"time"
)

type OrderItem struct {
	ID          string    `json:"id"`
	OrderID     string    `json:"order_id"`
	BookID      string    `json:"book_id"`
	Quantity    int       `json:"quantity"`
	PriceAtOrder float64   `json:"price_at_order"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
} 