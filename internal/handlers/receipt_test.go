package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joshguarino/receipt-processor-challenge/internal/models"
	"github.com/stretchr/testify/assert"
)

func setup() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/receipts/process", ProcessReceipt)
	router.GET("/receipts/:id/points", GetReceiptPoints)
	return router
}

func request(router *gin.Engine, method, url string, body []byte) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	return rec
}

func TestProcessReceipt(t *testing.T) {
	router := setup()
	payload := `{
		"retailer": "Test",
		"purchaseDate": "2025-01-22",
		"purchaseTime": "22:00",
		"items": [{"shortDescription": "Test", "price": "5.00"}],
		"total": "7.00"
	}`
	rec := request(router, http.MethodPost, "/receipts/process", []byte(payload))
	var response map[string]string
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NoError(t, err)
	assert.Contains(t, response, "id")
	assert.NotEmpty(t, response["id"])
}

func TestProcessReceiptBadRequest(t *testing.T) {
	router := setup()
	payload := `{
		"retailer": "Test",
		"total": "test" // Invalid total value
	}`
	rec := request(router, http.MethodPost, "/receipts/process", []byte(payload))
	var response map[string]string
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.NoError(t, err)
	assert.Contains(t, response, "error")
	assert.NotEmpty(t, response["error"])
}

func TestGetReceiptPoints(t *testing.T) {
	router := setup()
	receipts = make(map[string]*models.Receipt)
	receiptID := "test-id"
	receipts[receiptID] = &models.Receipt{
		ID:     receiptID,
		Points: 22,
	}
	rec := request(router, http.MethodGet, "/receipts/test-id/points", nil)
	var response map[string]int
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NoError(t, err)
	assert.Contains(t, response, "points")
	assert.Equal(t, 22, response["points"])
}

func TestGetReceiptPoints_NotFound(t *testing.T) {
	router := setup()
	receipts = make(map[string]*models.Receipt)
	rec := request(router, http.MethodGet, "/receipts/test-id/points", nil)
	var response map[string]string
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(t, http.StatusNotFound, rec.Code)
	assert.NoError(t, err)
	assert.Contains(t, response, "error")
	assert.Equal(t, "Receipt not found", response["error"])
}
