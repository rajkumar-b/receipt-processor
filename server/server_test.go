package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStartServer(t *testing.T) {
	// Set up a test server using httptest.NewServer
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	port:= 8080

	// Start the server (in a goroutine)
	go func() {
		err := StartServer(port)
		if err != nil {
			t.Fatalf("StartServer returned an error: %v", err)
		}
	}()

	// Make a request to the server
	resp, err := http.Get(fmt.Sprintf("http://localhost:%d/ping", port))
	if err != nil {
		t.Fatalf("Failed to make a request to the server: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
}
