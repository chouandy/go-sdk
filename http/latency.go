package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// LatencyRoutePath latency route path
const LatencyRoutePath = "/health/latency"

// LatencyHandler latency handler
func LatencyHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
