package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshguarino/receipt-processor-challenge/internal/models"
	"github.com/joshguarino/receipt-processor-challenge/internal/services"
)

func ProcessReceipt(c *gin.Context) {
	var receipt models.Receipt
	if err := c.BindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	receipt.ID = services.GenerateReceiptID()
	receipt.Points = services.CalculatePoints(&receipt)

	c.JSON(http.StatusOK, gin.H{"id": receipt.ID})
}

func GetReceiptPoints(c *gin.Context) {
	// Implement logic to retrieve receipt by ID and return points
}
