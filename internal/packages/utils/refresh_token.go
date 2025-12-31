package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateRefreshToken() string {
	b := make([]byte, 64)
	rand.Read(b)
	return hex.EncodeToString(b)
}
