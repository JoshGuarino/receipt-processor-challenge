package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joshguarino/receipt-processor-challenge/internal/handlers"
)

func main() {
	router := gin.Default()

	router.POST("/receipts/process", handlers.ProcessReceipt)
	router.GET("/receipts/:id/points", handlers.GetReceiptPoints)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
