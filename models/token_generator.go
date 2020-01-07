package models

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"math"
	"os"

	cryptoex "github.com/chouandy/go-sdk/crypto"
	dbex "github.com/chouandy/go-sdk/db"
	randex "github.com/chouandy/go-sdk/rand"
)

// TokenGenerator token generator
var TokenGenerator = &tokenGenerator{
	KeyGenerator: &cryptoex.KeyGenerator{
		SecretKey:  []byte(os.Getenv("SECRET_KEY_BASE")),
		Iterations: int(math.Pow(2, 16)),
	},
}

// tokenGenerator token generator struct
type tokenGenerator struct {
	KeyGenerator *cryptoex.KeyGenerator
}

// Digest digest
func (g *tokenGenerator) Digest(column, raw string) string {
	hash := hmac.New(sha256.New, g.keyFor(column))
	hash.Write([]byte(raw))

	return hex.EncodeToString(hash.Sum(nil))
}

// Generate generate
func (g *tokenGenerator) Generate(model interface{}, column string) (string, string) {
	// New key by column pkcs5 pbkdf2 hmac sha1
	key := g.keyFor(column)
	// New raw, enc and break without taking
	var raw, enc string
	for {
		// Rand raw
		raw = randex.GenerateFriendlyToken(20)
		// hmac sha256 raw with key
		hash := hmac.New(sha256.New, key)
		hash.Write([]byte(raw))
		// Encode to hex string
		enc = hex.EncodeToString(hash.Sum(nil))

		// Check enc is token or not
		var count int
		dbex.GORM().Model(model).Where(column+" = ?", enc).Count(&count)
		if count == 0 {
			break
		}
	}

	return raw, enc
}

func (g *tokenGenerator) keyFor(column string) []byte {
	return g.KeyGenerator.GenerateKey([]byte("Devise "+column), 64)
}
