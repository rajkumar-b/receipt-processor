package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"rajkumar.app/receipt-processor/handler"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.ForwardedByClientIP = true

	router.GET("/ping", handler.SendPing)
	router.POST("/receipts/process", handler.AddNewReceipt)
	router.GET("/receipts/:id/points", handler.GetPointsForReceipt)

	return router
}

func StartServer(port int) error {
	router := setupRouter()

	// Check if the environment variable is set for binding address
	bindAddress := os.Getenv("BIND_ADDRESS")
	if bindAddress == "" {
		bindAddress = "127.0.0.1"
	}
	addr := fmt.Sprintf("%s:%d", bindAddress, port)

	fmt.Printf("\nReceipt Processor Server is running on port %d...\n", port)
	fmt.Printf("Access the API via localhost: http://localhost:%d/<endpoint>\n\n", port)
	fmt.Printf("To test a simple ping, use: http://localhost:%d/ping\n\n", port)

	return http.ListenAndServe(addr, router)
}
