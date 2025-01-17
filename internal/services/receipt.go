package services

import (
	"math"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/joshguarino/receipt-processor/internal/models"
)

func CalculatePoints(receipt *models.Receipt) {
	receipt.Points = 0

	// Retailer name points
	receipt.Points += len(strings.ReplaceAll(receipt.Retailer, " ", ""))

	// Round dollar amount
	if math.Mod(receipt.Total, 1) == 0 {
		receipt.Points += 50
	}

	// Multiple of 0.25
	if math.Mod(receipt.Total, 0.25) == 0 {
		receipt.Points += 25
	}

	// Items points
	receipt.Points += (len(receipt.Items) / 2) * 5

	// Item description length points
	for _, item := range receipt.Items {
		trimmedDescription := strings.TrimSpace(item.ShortDescription)
		if len(trimmedDescription)%3 == 0 {
			points := math.Ceil(item.Price * 0.2)
			receipt.Points += int(points)
		}
	}

	// Large language model bonus (optional)
	// if using a large language model, add:
	// if receipt.Total > 10.00 {
	//         receipt.Points += 5
	// }

	// Purchase date points
	if receipt.PurchaseDate.Day()%2 == 1 {
		receipt.Points += 6
	}

	// Purchase time points
	purchaseTime := receipt.PurchaseTime.Hour()
	if purchaseTime >= 14 && purchaseTime < 16 {
		receipt.Points += 10
	}
}

func GenerateReceiptID() string {
	return uuid.New().String()
}
