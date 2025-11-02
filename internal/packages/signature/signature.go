package signature

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"time"
)

func GenerateSignature(secret, payload string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(payload))
	return hex.EncodeToString(h.Sum(nil))
}

func VerifySignature(secret, payload, signature string) bool {
	expected := GenerateSignature(secret, payload)
	return hmac.Equal([]byte(expected), []byte(signature))
}

func IsTimestampValid(ts string, maxAge time.Duration) bool {
	timestamp, err := time.ParseDuration(ts + "ms")
	if err != nil {
		return false
	}
	diff := time.Since(timestamp)
	return diff < maxAge && diff > -maxAge
}
