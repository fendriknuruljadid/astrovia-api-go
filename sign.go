package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

func main() {
	secret := "4sTrovia53cretProd"
	timestamp := time.Now().UTC().Format(time.RFC3339)

	msg := timestamp
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(msg))
	signature := hex.EncodeToString(mac.Sum(nil))

	fmt.Println("X-Timestamp:", timestamp)
	fmt.Println("X-Signature:", signature)
}
