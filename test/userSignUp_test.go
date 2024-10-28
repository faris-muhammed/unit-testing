package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"main.go/handlers"
	"main.go/initializer"
	"main.go/model"
)

func SetupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the in-memory database")
	}
	db.AutoMigrate(&model.UserModel{})
	initializer.DB = db
	return db
}

func TestMain(m *testing.M) {
	SetupTestDB()
	m.Run()
}

func TestSignUp(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.POST("/signup", handlers.SignUp)

	tests := []struct {
		name           string
		input          model.UserModel
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "Successful Signup",
			input: model.UserModel{
				Email:    "testuser",
				Password: "password123",
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"message":"User created successfully"}`,
		},
		{
			name:           "Bad Request - Invalid JSON",
			input:          model.UserModel{},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"Key: 'UserModel.Email' Error:Field validation for 'Email' failed on the 'required' tag"}`, // adjust based on actual validation error
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.input)
			req, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()
			r.ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)
			assert.JSONEq(t, tt.expectedBody, rr.Body.String())
		})
	}
}
