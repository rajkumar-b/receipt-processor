package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	swaggerfiles "github.com/swaggo/files"
	docs "rajkumar.app/receipt-processor/docs"
	"rajkumar.app/receipt-processor/handler"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.ForwardedByClientIP = true

	// Define Routes
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

	// Add Swagger UI
	docs.SwaggerInfo.BasePath = "/"
	url := ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", port))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, url))


	fmt.Printf("\nReceipt Processor Server is running on port %d...\n", port)
	fmt.Printf("Access the REST API via Swagger UI: http://localhost:%d/swagger/index.html\n\n", port)
	fmt.Printf("To test a simple ping, use: http://localhost:%d/ping\n\n", port)

	return http.ListenAndServe(addr, router)
}
