package handler

import (
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
		var point int
		err := json.Unmarshal(w.Body.Bytes(), &point)
		assert.NoError(t, err)

		// Perform the assertion for points
		assert.Equal(t, 0, point, "Points should be 0")
	})

	t.Run("Receipt Found - Actual Point Check - single receipt", func(t *testing.T) {

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/receipts/xyz-hijk-lmnop/points", nil)
		calcPoints(validReceipts[1])
		router.ServeHTTP(w, req)

		assert.Equal(t, w.Code, http.StatusOK)
		
		// Parse the response body into a map
		var point int
		err := json.Unmarshal(w.Body.Bytes(), &point)
		assert.NoError(t, err)

		// Perform the assertion for points
		assert.Equal(t, 0, point, "Points should be 0")
	})

	t.Run("Receipt Found - Actual Point Check - multi receipt", func(t *testing.T) {

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/receipts/abc-def-ghij/points", nil)
		calcPoints(validReceipts[0])
		router.ServeHTTP(w, req)

		assert.Equal(t, w.Code, http.StatusOK)
		
		// Parse the response body into a map
		var point int
		err := json.Unmarshal(w.Body.Bytes(), &point)
		assert.NoError(t, err)

		// Perform the assertion for points
		assert.Equal(t, 0, point, "Points should be 0")
	})
}
