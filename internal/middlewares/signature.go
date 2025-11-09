package middlewares

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"time"

	"astrovia-api-go/internal/packages/response"
	"github.com/gofiber/fiber/v2"

)

const (
	CLIENT_SECRET = "4sTrovia53cretProd"
	MAX_TIME_DIFF = 3000 // detik
)

func generateSignature(secret, timestamp string) string {
	msg := timestamp
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(msg))
	return hex.EncodeToString(mac.Sum(nil))
}

func SignatureClientMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		signature := c.Get("X-Signature")
		timestamp := c.Get("X-Timestamp")

		if signature == "" || timestamp == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(response.Error(401, "missing signature or timestamp", nil))
		}

		// validasi timestamp
		ts, err := time.Parse(time.RFC3339, timestamp)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				response.Error(401, "invalid client timestamp format (use RFC3339)", nil))
		}

		if time.Since(ts).Seconds() > MAX_TIME_DIFF {
			return c.Status(fiber.StatusUnauthorized).JSON(
				response.Error(401, "client signature expired", nil))
		}

		expected := generateSignature(
			CLIENT_SECRET,
			timestamp,
		)
		

		if signature != expected {
			return c.Status(fiber.StatusUnauthorized).JSON(
				
				response.Error(401, "invalid client signature", nil))
		}

		// kalau valid, lanjutkan
		return c.Next()
	}
}
