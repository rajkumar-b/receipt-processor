package main

import (
	"fmt"
	"net/http"

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

	// Start server and bind to all network interfaces
	port := 8080
	fmt.Printf("\nReceipt Processor Server is running on port %d...\n", port)
	fmt.Printf("Access the API via localhost: http://localhost:%d/<endpoint>\n\n", port)
	err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), router)
	if err != nil {
		fmt.Println("Failed to start the server:", err)
	}
}