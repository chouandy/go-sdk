package middleware

import (
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// NewCORSMiddleware cors middleware
func NewCORSMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: strings.Split(os.Getenv("ALLOW_ORIGINS"), ","),
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"PATCH",
			"DELETE",
			"HEAD",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Length",
			"Content-Type",
			"Authorization",
			"X-CSRF-TOKEN",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
