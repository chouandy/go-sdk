package rand

import (
	"encoding/hex"
	"math/rand"
	"time"
)

// Hex generate random hex
func Hex(size int) []byte {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, size)
	r.Read(b)

	return b
}

// HexString generate random hex string
func HexString(size int) string {
	h := Hex(size)
	enc := make([]byte, len(h)*2+2)
	copy(enc, "0x")
	hex.Encode(enc[2:], h)

	return string(enc)
}
