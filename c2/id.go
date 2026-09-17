package c2

import (
	"crypto/rand"
	"encoding/hex"
)

//nolint:unused
func generateChannelID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}