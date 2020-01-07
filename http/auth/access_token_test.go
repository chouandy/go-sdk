package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAccessToken(t *testing.T) {
	accessToken := NewAccessToken()
	assert.Equal(t, 64, len(accessToken))
}

func TestNewAccessTokenUniq(t *testing.T) {
	accessTokens := make([]string, 0)

	for i := 0; i < 10000; i++ {
		accessToken := NewAccessToken()
		assert.NotContains(t, accessTokens, accessToken)
		accessTokens = append(accessTokens, accessToken)
	}
}
