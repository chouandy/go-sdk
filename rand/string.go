package rand

import (
	"math/rand"
	"time"
)

// Alphanumeric alphanumeric characters
var Alphanumeric = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// String random generate string
func String(characters string, n int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, n)
	for i := range b {
		b[i] = characters[r.Intn(len(characters))]
	}

	return string(b)
}
