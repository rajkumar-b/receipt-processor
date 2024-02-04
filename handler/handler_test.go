package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"rajkumar.app/receipt-processor/model"
)

func TestSendPing(t *testing.T) {
	router := gin.New()
	router.GET("/ping", SendPing)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.Contains(t, w.Body.String(),  "pong")
}

func TestGetPointsForReceipt(t *testing.T) {
	router := gin.New()
	router.GET("/receipts/:id/points", GetPointsForReceipt)
	validReceipts := []model.Receipt{
		{
			ID:				"abc-def-ghij",
			Retailer:		"Valid Retailer 1",
			PurchaseDate:	"2022-09-20",
			PurchaseTime:	"13:01",
			PurchasedItems: []model.Item{
				{
					Description: "Valid Item 1",
					Price:       "2.99",
				},
				{
					Description: "Valid Item 2",
					Price:       "1.99",
				},
			},
			Total: "4.98",
		},
		{
			ID:				"xyz-hijk-lmnop",
			Retailer:		"Valid Retailer 2",
			PurchaseDate:	"2022-09-12",
			PurchaseTime:	"03:01",
			PurchasedItems: []model.Item{
				{
					Description: "Valid Item 3",
					Price:       "5.99",
				},
			},
			Total: "5.99",
		},
		{
			Retailer:		"Valid Retailer 1",
			PurchaseDate:	"2022-02-20",
			PurchaseTime:	"23:01",
			PurchasedItems: []model.Item{
				{
					Description: "Valid Item 4",
					Price:       "1.99",
				},
				{
					Description: "Valid Item 5",
					Price:       "5.99",
				},
			},
			Total: "7.98",
		},
	}

	model.Receipts = append(model.Receipts, validReceipts...)
	err := model.ValidateReceipts()
	assert.NoError(t, err, "ValidateReceipts should pass with valid receipts")

	t.Run("Receipt Not Found", func(t *testing.T) {

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/receipts/abc-def/points", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, w.Code, http.StatusNotFound)
	})

	
	t.Run("Receipt Found", func(t *testing.T) {

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/receipts/abc-def-ghij/points", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, w.Code, http.StatusOK)
	})

	t.Run("Receipt Found - Zero Point Check", func(t *testing.T) {

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/receipts/abc-def-ghij/points", nil)

		router.ServeHTTP(w, req)

		assert.Equal(t, w.Code, http.StatusOK)
		
		// Parse the response body into a map
		var responseBody map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &responseBody)
		assert.NoError(t, err)

		points, ok := responseBody["points"]
		assert.True(t, ok, "No 'points' field found in the response")

		// Perform the assertion for points
		assert.Equal(t, 0, int(points.(float64)), "Points should be 0")
	})

	t.Run("Receipt Found - Actual Point Check - single receipt", func(t *testing.T) {

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/receipts/xyz-hijk-lmnop/points", nil)
		calcPoints(validReceipts[1])
		router.ServeHTTP(w, req)

		assert.Equal(t, w.Code, http.StatusOK)
		
		// Parse the response body into a map
		var responseBody map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &responseBody)
		assert.NoError(t, err)

		points, ok := responseBody["points"]
		assert.True(t, ok, "No 'points' field found in the response")

		// Perform the assertion for points
		assert.Equal(t, 0, int(points.(float64)), "Points should be 0")
	})

	t.Run("Receipt Found - Actual Point Check - multi receipt", func(t *testing.T) {

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/receipts/abc-def-ghij/points", nil)
		calcPoints(validReceipts[0])
		router.ServeHTTP(w, req)

		assert.Equal(t, w.Code, http.StatusOK)
		
		// Parse the response body into a map
		var responseBody map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &responseBody)
		assert.NoError(t, err)

		points, ok := responseBody["points"]
		assert.True(t, ok, "No 'points' field found in the response")

		// Perform the assertion for points
		assert.Equal(t, 0, int(points.(float64)), "Points should be 0")
	})
}


func TestAddNewReceipt(t *testing.T) {
	router := gin.New()
	router.POST("/receipts/process", AddNewReceipt)

	t.Run("Receipt created successfully", func(t *testing.T) {
		rawJSON := `{
			"retailer": "Target",
			"purchaseDate": "2022-01-02",
			"purchaseTime": "13:13",
			"total": "1.25",
			"items": [
				{"shortDescription": "Pepsi - 12-oz", "price": "1.25"}
			]
		}`

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/receipts/process", bytes.NewBufferString(rawJSON))
		router.ServeHTTP(w, req)

		assert.Equal(t, w.Code, http.StatusOK)
	})

	t.Run("Receipt created successfully - multiple items", func(t *testing.T) {
		rawJSON := `{
			"retailer": "Walgreens",
			"purchaseDate": "2022-01-02",
			"purchaseTime": "08:13",
			"total": "2.65",
			"items": [
				{"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
				{"shortDescription": "Dasani", "price": "1.40"}
			]
		}}`

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/receipts/process", bytes.NewBufferString(rawJSON))
		router.ServeHTTP(w, req)

		assert.Equal(t, w.Code, http.StatusOK)
	})
	
	t.Run("Receipt creation failed - no payload", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/receipts/process", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, w.Code, http.StatusBadRequest)
	})

	
	t.Run("Receipt creation failed - invalid values", func(t *testing.T) {
		rawJSON := `{
			"retailer": "Target",
			"purchaseDate": "2022-01-02",
			"purchaseTime": "3:3",
			"total": "-1.25",
			"items": [
				{"shortDescription": "Pepsi - 12-oz", "price": "1.25"}
			]
		}`

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/receipts/process", bytes.NewBufferString(rawJSON))
		router.ServeHTTP(w, req)

		assert.Equal(t, w.Code, http.StatusBadRequest)
	})

	
	t.Run("Receipt creation failed - missing columns", func(t *testing.T) {
		rawJSON := `{
			"retailer": "Target",
			"purchaseDate": "2022-01-02",
			"purchaseTime": "3:33",
			"items": [
				{"shortDescription": "Pepsi - 12-oz", "price": "1.25"}
			]
		}`

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/receipts/process", bytes.NewBufferString(rawJSON))
		router.ServeHTTP(w, req)

		assert.Equal(t, w.Code, http.StatusBadRequest)
	})

}
