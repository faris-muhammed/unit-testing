package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"main.go/handlers"
)

func TestLogout(t *testing.T) {
	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Create a new Gin router
	router := gin.Default()

	// Register the Logout handler
	router.GET("/logout", handlers.Logout)

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodGet, "/logout", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	// Create a ResponseRecorder to record the response
	w := httptest.NewRecorder()

	// Serve the HTTP request
	router.ServeHTTP(w, req)

	// Assert the status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Assert the response body
	expectedResponse := gin.H{"message": "Logout successful"}
	var responseBody map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &responseBody)
	if err != nil {
		t.Fatalf("Could not parse response body: %v", err)
	}
	assert.Equal(t, expectedResponse["message"], responseBody["message"])
}
