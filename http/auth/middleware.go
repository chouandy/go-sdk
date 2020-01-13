package auth

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	logex "github.com/chouandy/go-sdk/log"
)

// FindUserFunc find user function
type FindUserFunc func(userID uint64) (interface{}, error)

// NewMiddleware new middleware
func NewMiddleware(findUserFunc FindUserFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		// New task
		task := &AuthenticateUserTask{c: c, FindUserFunc: findUserFunc}

		// Get and validate access token
		if err := task.GetAndValidateAccessToken(); err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Get user id by access token
		if err := task.GetUserIDByAccessToken(); err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Find user
		if err := task.FindUser(); err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Reset access token ttl
		if err := task.ResetAccessTokenTTL(); err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}

// --------------------------
// -          Task          -
// --------------------------

// AuthenticateUserTask authenticate user task
type AuthenticateUserTask struct {
	c *gin.Context

	AccessToken  string
	UserID       uint64
	FindUserFunc FindUserFunc
}

// GetAndValidateAccessToken get and validate access token
func (t *AuthenticateUserTask) GetAndValidateAccessToken() error {
	// Get from header
	t.AccessToken = GetAccessToken(t.c)

	if len(t.AccessToken) == 0 {
		return errors.New("access_token not found")
	}

	return nil
}

// GetUserIDByAccessToken get user id by access token
func (t *AuthenticateUserTask) GetUserIDByAccessToken() error {
	// Get user id by access token
	userID, err := GetUserIDByAccessToken(t.AccessToken)
	if err != nil {
		logex.Log.Error(err)
		return err
	}

	// Set user id
	t.UserID = userID
	t.c.Set("user_id", t.UserID)

	return nil
}

// FindUser find user
func (t *AuthenticateUserTask) FindUser() error {
	// Execute find user func
	user, err := t.FindUserFunc(t.UserID)
	if err != nil {
		logex.Log.Error(err)
		return err
	}

	// Set user
	t.c.Set("user", user)

	return nil
}

// ResetAccessTokenTTL reset access token ttl
func (t *AuthenticateUserTask) ResetAccessTokenTTL() error {
	if err := ResetAccessTokenTTL(t.AccessToken, t.UserID); err != nil {
		logex.Log.Error(err)
		return err
	}

	return nil
}
