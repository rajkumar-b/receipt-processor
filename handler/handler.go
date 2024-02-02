package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func SendPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
