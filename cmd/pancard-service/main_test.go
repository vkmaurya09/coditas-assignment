package main

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"pan-service/internal/handlers"
	"pan-service/internal/middleware"
	"pan-service/internal/models"
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
	"gotest.tools/v3/assert"
)

func TestSubmitHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Create log files for testing
	latencyFile, err := os.OpenFile("test_latency.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		t.Fatalf("Failed to open latency log file: %v", err)
	}
	defer latencyFile.Close()

	errorFile, err := os.OpenFile("test_error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		t.Fatalf("Failed to open error log file: %v", err)
	}
	defer errorFile.Close()

	// Initialize loggers for testing
	latencyLog := log.New(latencyFile, "", 0)
	errorLog := log.New(errorFile, "", 0)

	router.Use(middleware.LatencyLogger(latencyLog))

	var inMemoryData []models.Payload
	var mu sync.Mutex
	router.POST("/submit", handlers.SubmitHandler(&inMemoryData, &mu, errorLog))

	tests := []struct {
		name         string
		payload      string
		expectedCode int
	}{
		{
			name:         "Valid payload",
			payload:      `{"name":"John Doe","pan":"ABCDE1234F","mobile":"1234567890","email":"john@example.com"}`,
			expectedCode: http.StatusOK,
		},
		{
			name:         "Invalid PAN",
			payload:      `{"name":"John Doe","pan":"ABCDE12345","mobile":"1234567890","email":"john@example.com"}`,
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "Invalid mobile",
			payload:      `{"name":"John Doe","pan":"ABCDE1234F","mobile":"12345","email":"john@example.com"}`,
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "Invalid email",
			payload:      `{"name":"John Doe","pan":"ABCDE1234F","mobile":"1234567890","email":"john@com"}`,
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("POST", "/submit", bytes.NewBufferString(tt.payload))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)
		})
	}
}
