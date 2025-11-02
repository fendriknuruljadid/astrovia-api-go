package middleware

import (
	"astrovia-api-go/internal/packages/auth"
	"astrovia-api-go/internal/packages/response"
	"github.com/gofiber/fiber/v2"
	"os"
	"time"
)

var clientSecret = os.Getenv("CLIENT_SECRET")

// Verifikasi signature request dari client
func VerifyClientSignature(c *fiber.Ctx) error {
	signature := c.Get("X-Signature")
	timestamp := c.Get("X-Timestamp")
	body := string(c.Body())

	if signature == "" || timestamp == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(response.Error(401, "missing signature or timestamp", nil))
	}

	// Periksa timestamp (maks 5 menit)
	if !auth.IsTimestampValid(timestamp, 5*time.Minute) {
		return c.Status(fiber.StatusUnauthorized).JSON(response.Error(401, "expired request", nil))
	}

	// Verifikasi signature
	if !auth.VerifySignature(clientSecret, body, signature) {
		return c.Status(fiber.StatusUnauthorized).JSON(response.Error(401, "invalid signature", nil))
	}

	return c.Next()
}
