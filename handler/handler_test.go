package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSendPing(t *testing.T) {
	router := gin.New()
	router.GET("/ping", SendPing)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.Contains(t, w.Body.String(),  "pong")
}

func TestGetItems(t *testing.T) {
	router := gin.New()
	router.GET("/items", GetItems)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/items", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusOK)
}
