package main

import (
	"fmt"

	"rajkumar.app/receipt-processor/server"
)

func main() {
	port := 8080
	err := server.StartServer(port)
	if err != nil {
		fmt.Println("Failed to start the server:", err)
	}
}
