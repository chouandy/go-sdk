package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RequestIDMiddleware request id middleware
func RequestIDMiddleware(c *gin.Context) {
	// Generate request id
	requestID := uuid.New().String()
	// Set request id
	c.Set("request_id", requestID)
	// Set request id to response
	c.Writer.Header().Set("X-Request-Id", requestID)

	c.Next()
}
