package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joshguarino/receipt-processor/internal/models"
	"github.com/joshguarino/receipt-processor/internal/services"
)

func ProcessReceipt(c *gin.Context) {
	var receipt models.Receipt
	if err := c.BindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	receipt.PurchaseDate, _ = time.Parse("2006-01-02", receipt.PurchaseDate)
	receipt.PurchaseTime, _ = time.Parse("15:04", receipt.PurchaseTime)
	receipt.ID = services.GenerateReceiptID()

	services.CalculatePoints(&receipt)

	c.JSON(http.StatusOK, gin.H{"id": receipt.ID})
}

func GetReceiptPoints(c *gin.Context) {
	// Implement logic to retrieve receipt by ID and return points
}
