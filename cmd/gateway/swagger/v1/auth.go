
package routes

import "github.com/gofiber/fiber/v2"

// =================== Dummy User model untuk Swagger ===================

type Auth struct {
	Email string `json:"email"`
	Password string `json:"password"`
	Provider string `json:"provider"`
}
// =================== Dummy Handlers untuk Swagger ===================

// @Summary Authentication
// @Description Generate token
// @Tags Authentication
// @Security X-Signature
// @Security X-Timestamp
// @Accept json
// @Produce json
// @Param auth body routes.Auth true "Auth info"
// @Router /v1/generate-token [post]
func AuthHandler(c *fiber.Ctx) error { return nil }
