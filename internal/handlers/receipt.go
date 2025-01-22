package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshguarino/receipt-processor-challenge/internal/models"
	"github.com/joshguarino/receipt-processor-challenge/internal/services"
)

var receipts = make(map[string]*models.Receipt)

func ProcessReceipt(c *gin.Context) {
	var receipt models.Receipt
	if err := c.BindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	receipt.ID = services.GenerateReceiptID()
	receipt.Points = services.CalculatePoints(&receipt)
	receipts[receipt.ID] = &receipt
	c.JSON(http.StatusOK, gin.H{"id": receipt.ID})
}

func GetReceiptPoints(c *gin.Context) {
	receiptID := c.Param("id")
	receipt, exists := receipts[receiptID]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Receipt not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"points": receipt.Points})
}
