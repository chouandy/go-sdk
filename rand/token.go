package rand

import (
	"crypto/rand"
	"encoding/base64"
	"strings"
)

// GenerateFriendlyToken Generate a friendly string randomly to be used as token.
func GenerateFriendlyToken(length int) string {
	// Calculate length for base64 url safe encode
	rlength := (length * 3) / 4
	// New token
	token := make([]byte, rlength)
	// Rand token
	rand.Read(token)
	// Base64 url safe encode token
	encodedToken := base64.RawURLEncoding.EncodeToString(token)
	// Replace 'lIO0' => 'sxyz'
	encodedToken = strings.Replace(encodedToken, "lIO0", "sxyz", -1)

	return encodedToken
}
