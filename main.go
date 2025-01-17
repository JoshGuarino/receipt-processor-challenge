package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joshguarino/receipt-processor-challenge/internal/handlers"
)

func main() {
	r := gin.Default()

	r.POST("/receipts/process", handlers.ProcessReceipt)
	r.GET("/receipts/:id/points", handlers.GetReceiptPoints)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
