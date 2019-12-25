package middleware

import (
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// SessionMiddleware session middleware
func SessionMiddleware(name string) gin.HandlerFunc {
	return sessions.Sessions(
		name,
		cookie.NewStore([]byte(os.Getenv("SECRET_KEY_BASE"))),
	)
}
