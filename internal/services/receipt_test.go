package services

import (
	"testing"

	"github.com/joshguarino/receipt-processor-challenge/internal/models"
	"github.com/stretchr/testify/assert"
)

var testReceipt = models.Receipt{
	Retailer:     "Test Retailer",
	PurchaseDate: "2025-1-22",
	PurchaseTime: "15:30",
	Items: []models.Item{
		{ShortDescription: "Item 1", Price: "10.00"},
		{ShortDescription: "Item 22", Price: "3.00"},
		{ShortDescription: "Item 5", Price: "2.22"},
	},
	Total: "15.00",
}

func TestCalculatePoints(t *testing.T) {
	points := CalculatePoints(&testReceipt)
	assert.Equal(t, 111, points, "Points are being calculated incorrectly")
}

func TestGenerateReceiptID(t *testing.T) {
	id := GenerateReceiptID()
	assert.IsType(t, "", id, "Should be returning string")
}

func TestCalcRetailerNamePoints(t *testing.T) {
	points := calcRetailerNamePoints(*&testReceipt.Retailer)
	assert.Equal(t, 12, points, "Points are being calculated incorrectly on calcRetailerNamePoints")
}

func TestCalcNumItemsPoints(t *testing.T) {
	points := calcNumItemsPoints(*&testReceipt.Items)
	assert.Equal(t, 5, points, "Points are being calculated incorrectly on calcNumItemsPoints")
}

func TestCalcRoundTotalPoints(t *testing.T) {
	roundPoints := calcRoundTotalPoints(*&testReceipt.Total)
	notRoundPoints := calcRoundTotalPoints("5.22")
	assert.Equal(t, 50, roundPoints, "Points are being calculated incorrectly on calcRoundTotalPoints")
	assert.Equal(t, 0, notRoundPoints, "Points are being calculated incorrectly on calcRoundTotalPoints")
}

func TestCalcTotalMultPoint25(t *testing.T) {
	multPoints := calcTotalMultPoint25(*&testReceipt.Total)
	notMultPoints := calcTotalMultPoint25("5.22")
	assert.Equal(t, 25, multPoints, "Points are being calculated incorrectly on calcTotalMultPoint25")
	assert.Equal(t, 0, notMultPoints, "Points are being calculated incorrectly on calcTotalMultPoint25")
}

func TestCalcItemDescriptionPoints(t *testing.T) {
	divPoints := calcItemDescriptionPoints(*&testReceipt.Items[0])
	notDivPoints := calcItemDescriptionPoints(*&testReceipt.Items[1])
	assert.Equal(t, 2, divPoints, "Points are being calculated incorrectly on calcItemDescriptionPoints")
	assert.Equal(t, 0, notDivPoints, "Points are being calculated incorrectly on calcItemDescriptionPoints")
}

func TestCalcDateOddPoints(t *testing.T) {
	oddPoints := calcDateOddPoints("2025-01-5")
	notOddPoints := calcDateOddPoints("2025-01-22")
	assert.Equal(t, 6, oddPoints, "Points are being calculated incorrectly on calcDateOddPoints")
	assert.Equal(t, 0, notOddPoints, "Points are being calculated incorrectly on calcDateOddPoints")
}

func TestCalcTime2To4Points(t *testing.T) {
	inPoints := calcTime2To4Points("15:00")
	outPoints := calcTime2To4Points("22:00")
	assert.Equal(t, 10, inPoints, "Points are being calculated incorrectly on calcTime2To4Points")
	assert.Equal(t, 0, outPoints, "Points are being calculated incorrectly on calcTime2To4Points")
}
