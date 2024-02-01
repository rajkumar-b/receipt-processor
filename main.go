package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Define routes
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Start server
	port := 8080
	fmt.Printf("Receipt Processor Server is running on port %d...\n", port)
	fmt.Printf("Access the API via localhost: http://localhost:%d/ping\n", port)
	router.Run(fmt.Sprintf(":%d", port))
}