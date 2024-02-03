package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"rajkumar.app/receipt-processor/handler"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1"})

	router.GET("/ping", handler.SendPing)
	router.GET("/items", handler.GetItems)

	return router
}

func StartServer(port int) error {
	router := setupRouter()

	fmt.Printf("\nReceipt Processor Server is running on port %d...\n", port)
	fmt.Printf("Access the API via localhost: http://localhost:%d/<endpoint>\n\n", port)
	fmt.Printf("To test a simple ping, use: http://localhost:%d/ping\n\n", port)

	return http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), router)
}
