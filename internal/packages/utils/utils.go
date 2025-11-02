package utils

import (
    "crypto/rand"
    "github.com/oklog/ulid/v2"
    "time"
)

func NewULID() string {
    t := time.Now().UTC() // disarankan UTC
    entropy := ulid.Monotonic(rand.Reader, 0)
    return ulid.MustNew(ulid.Timestamp(t), entropy).String()
}
