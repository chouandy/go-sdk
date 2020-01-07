package auth

import "github.com/gin-gonic/gin"

// GetAccessToken get access token
func GetAccessToken(c *gin.Context) string {
	// Get from header
	accessToken := c.GetHeader("Authorization")

	// Get from query string
	if len(accessToken) == 0 {
		accessToken = c.Query("access_token")
	}

	return accessToken
}
