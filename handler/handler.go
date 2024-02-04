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

// Calculate Points for given receipt and store it if valid receipt with id
func calcPoints(receipt model.Receipt) {
	point := 0
	receipt.SetPoints(point)
}

// GetPointsForReceipt responds with the points of a receipt by its ID.
func GetPointsForReceipt(c *gin.Context) {
	receiptID := c.Param("id")

	receipt, err := model.GetReceiptByID(receiptID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "No receipt found for that id"})
        return
    }

    c.IndentedJSON(http.StatusOK, gin.H{"points": receipt.Points})
}

// AddNewReceipt responds with the ID of the given receipt, if created.
func AddNewReceipt(c *gin.Context) {
	var receipt model.Receipt

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&receipt); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "The receipt is invalid"})
        return
    }

	if valerr := receipt.Validate(); valerr != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "The receipt is invalid"})
        return
    }

	calcPoints(receipt)

    // Add the new receipt to the slice.
    model.Receipts = append(model.Receipts, receipt)

    c.IndentedJSON(http.StatusOK, gin.H{"id":receipt.ID})
}
