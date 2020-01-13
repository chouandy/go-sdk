package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// NewSessionMiddleware new session middleware
func NewSessionMiddleware(name, secret string) gin.HandlerFunc {
	return sessions.Sessions(name, cookie.NewStore([]byte(secret)))
}
