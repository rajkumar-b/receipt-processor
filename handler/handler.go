package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rajkumar.app/receipt-processor/model"
)

// SendPing responds with a simple JSON reply.
func SendPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// GetItems responds with the list of all items as JSON.
func GetItems(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, model.Items)
}

// GetPointsForReceipt responds with the points of a receipt by its ID.
func GetPointsForReceipt(c *gin.Context) {
	receiptID := c.Param("id")

	receipt, err := model.GetReceiptByID(receiptID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "No receipt found for that id"})
        return
    }

    c.IndentedJSON(http.StatusOK, receipt.Points)
}
