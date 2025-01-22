package services

import (
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/google/uuid"
	"github.com/joshguarino/receipt-processor-challenge/internal/models"
)

func CalculatePoints(receipt *models.Receipt) int {
	points := 0
	points += calcRetailerNamePoints(receipt.Retailer)
	points += calcNumItemsPoints(receipt.Items)
	points += calcRoundTotalPoints(receipt.Total)
	points += calcTotalMultPoint25(receipt.Total)
	for _, item := range receipt.Items {
		points += calcItemDescriptionPoints(item)
	}
	points += calcDateOddPoints(receipt.PurchaseDate)
	points += calcTime2To4Points(receipt.PurchaseTime)

	return points
}

func GenerateReceiptID() string {
	return uuid.New().String()
}

func calcRetailerNamePoints(str string) int {
	count := 0
	for _, char := range str {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			count++
		}
	}
	return count
}

func calcNumItemsPoints(items []models.Item) int {
	floatNum := float64(len(items) / 2)
	return int(math.Floor(floatNum)) * 5
}

func calcRoundTotalPoints(total string) int {
	floatTotal, _ := strconv.ParseFloat(total, 64)
	if math.Floor(floatTotal) == floatTotal {
		return 50
	}
	return 0
}

func calcTotalMultPoint25(total string) int {
	floatTotal, _ := strconv.ParseFloat(total, 64)
	divTotal := floatTotal / .25
	if math.Floor(divTotal) == divTotal {
		return 25
	}
	return 0
}

func calcItemDescriptionPoints(item models.Item) int {
	lenDescription := len(strings.TrimSpace(item.ShortDescription))
	if lenDescription%3 == 0 {
		price, _ := strconv.ParseFloat(item.Price, 64)
		return int(math.Ceil(price * .2))
	}
	return 0
}

func calcDateOddPoints(purchaseDate string) int {
	date, _ := time.Parse("2006-01-2", purchaseDate)
	if date.Day()%2 == 1 {
		return 6
	}
	return 0
}

func calcTime2To4Points(purchaseTime string) int {
	time, _ := time.Parse("15:04", purchaseTime)
	if time.Hour() >= 14 && time.Hour() <= 16 {
		return 10
	}
	return 0
}
