package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LatencyLogger(logger *log.Logger) gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        c.Next()
        latency := time.Since(start)
        logger.Printf("Path: %s, Latency: %v", c.Request.URL.Path, latency)
    }
}