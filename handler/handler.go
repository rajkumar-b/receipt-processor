package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rajkumar.app/receipt-processor/model"
)

// sendPing responds with a simple JSON reply.
func SendPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// getItems responds with the list of all items as JSON.
func GetItems(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, model.Items)
}
