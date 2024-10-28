	package test

	import (
		"bytes"
		"encoding/json"
		"net/http"
		"net/http/httptest"
		"testing"

		"github.com/gin-gonic/gin"
		"golang.org/x/crypto/bcrypt"
		"main.go/handlers"
		"main.go/initializer"
		"main.go/model"
	)

	func TestLogin(t *testing.T) {
		gin.SetMode(gin.TestMode) // Set gin to test mode
		router := gin.Default()

		// Test handler function
		router.POST("/login", handlers.Login)

		t.Run("Successful Login", func(t *testing.T) {
			// Mock a user and hash password
			password := "validpassword"
			hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

			// Add mock user to the database (mock or in-memory)
			initializer.DB.Create(&model.UserModel{
				Email:    "test@example.com",
				Password: string(hashedPassword),
			})

			// Prepare the request
			reqBody, _ := json.Marshal(gin.H{"email": "test@example.com", "password": password})
			req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(reqBody))
			req.Header.Set("Content-Type", "application/json")

			// Record the response
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			// Validate response
			if w.Code != http.StatusOK {
				t.Errorf("Expected status %v; got %v", http.StatusOK, w.Code)
			}

			var response map[string]string
			json.Unmarshal(w.Body.Bytes(), &response)
			if message, ok := response["message"]; !ok || message != "Login successful" {
				t.Errorf("Expected message 'Login successful'; got %v", response)
			}
		})

		t.Run("Invalid Email", func(t *testing.T) {
			reqBody, _ := json.Marshal(gin.H{"email": "invalid@example.com", "password": "any"})
			req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(reqBody))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != http.StatusUnauthorized {
				t.Errorf("Expected status %v; got %v", http.StatusUnauthorized, w.Code)
			}

			var response map[string]string
			json.Unmarshal(w.Body.Bytes(), &response)
			if errorMsg, ok := response["error"]; !ok || errorMsg != "Invalid email or password" {
				t.Errorf("Expected error 'Invalid email or password'; got %v", response)
			}
		})

		t.Run("Invalid Password", func(t *testing.T) {
			// Prepare the request with correct email but incorrect password
			reqBody, _ := json.Marshal(gin.H{"email": "test@example.com", "password": "wrongpassword"})
			req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(reqBody))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != http.StatusUnauthorized {
				t.Errorf("Expected status %v; got %v", http.StatusUnauthorized, w.Code)
			}

			var response map[string]string
			json.Unmarshal(w.Body.Bytes(), &response)
			if errorMsg, ok := response["error"]; !ok || errorMsg != "Invalid email or password" {
				t.Errorf("Expected error 'Invalid email or password'; got %v", response)
			}
		})
	}
