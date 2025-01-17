package models

import (
	"time"
)

type Receipt struct {
	ID           string    `json:"id"`
	Retailer     string    `json:"retailer"`
	PurchaseDate time.Time `json:"purchaseDate"`
	PurchaseTime time.Time `json:"purchaseTime"`
	Items        []Item    `json:"items"`
	Total        float64   `json:"total"`
	Points       int       `json:"points"`
}

type Item struct {
	ShortDescription string  `json:"shortDescription"`
	Price            float64 `json:"price"`
}
