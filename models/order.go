package models

import (
	"time"
)

type Order struct {
	ID         string    `json:"id"`
	UserID     string    `json:"user_id"`
	OrderDate  time.Time `json:"order_date"`
	TotalAmount float64   `json:"total_amount"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
} 