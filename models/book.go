package models

import (
	"time"
)

type Book struct {
	ID            string    `json:"id"`
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	Price         float64   `json:"price"`
	Stock         int       `json:"stock"`
	ISBN          string    `json:"isbn"`
	Category      string    `json:"category"`
	PublishedDate *time.Time `json:"published_date"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
} 