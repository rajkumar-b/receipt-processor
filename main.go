package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"rajkumar.app/receipt-processor/handler"
)

func main() {
	// Configure router
	router := gin.Default()
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1"})

	// Add routes
	router.GET("/ping", handler.SendPing)
	router.GET("/items", handler.GetItems)

	// Start server
	port := 8080
	fmt.Printf("\nReceipt Processor Server is running on port %d...\n", port)
	fmt.Printf("Access the API via localhost: http://localhost:%d/<endpoint>\n\n", port)
	router.Run(fmt.Sprintf("localhost:%d", port))
}