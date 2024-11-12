package main

import (
	"log"
	"os"
	"pan-service/internal/handlers"
	"pan-service/internal/middleware"
	"pan-service/internal/models"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	inMemoryData []models.Payload
	mu           sync.Mutex
	latencyLog   *log.Logger
	errorLog     *log.Logger
)

func init() {
	// Create log files
	latencyFile, err := os.OpenFile("latency.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open latency log file: %v", err)
	}

	errorFile, err := os.OpenFile("error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open error log file: %v", err)
	}

	// Initialize loggers
	latencyLog = log.New(latencyFile, "LATENCY: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLog = log.New(errorFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
func main() {
	r := gin.Default()
	r.Use(middleware.LatencyLogger(latencyLog))

	r.POST("/submit", handlers.SubmitHandler(&inMemoryData, &mu, errorLog))

	r.Run(":8080")
}
