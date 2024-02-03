package server

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStartServer(t *testing.T) {
	port := 8080

	// Start the server in a subroutine
	go func() {
		err := StartServer(port)
		assert.NoError(t, err, "StartServer returned an error")
	}()

	// Allow some time for the server to start
	time.Sleep(1 * time.Second)

	// Make a request to the server
	resp, err := http.Get(fmt.Sprintf("http://localhost:%d/ping", port))
	assert.NoError(t, err, "Failed to make a request to the server")
	defer resp.Body.Close()

	// Check the response status code
	assert.Equal(t, resp.StatusCode, http.StatusOK)
}
