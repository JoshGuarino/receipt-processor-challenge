package services

import (
	"github.com/google/uuid"
	"github.com/joshguarino/receipt-processor-challenge/internal/models"
)

func CalculatePoints(receipt *models.Receipt) {
	// implment logic
}

func GenerateReceiptID() string {
	return uuid.New().String()
}
