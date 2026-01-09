package routes

import "github.com/gofiber/fiber/v2"

// =================== Dummy User model untuk Swagger ===================

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Provider string `json:"provider"`
}
type RefreshToken struct {
	RefreshToken string `json:"refresh_token"`
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
// @Router /v1/auth/generate-token [post]
func AuthHandler(c *fiber.Ctx) error { return nil }

// @Summary Refresh Tokens
// @Description Refresh token
// @Tags Authentication
// @Security X-Signature
// @Security X-Timestamp
// @Security X-DeviceId
// @Accept json
// @Produce json
// @Param auth body routes.RefreshToken true "RefreshToken info"
// @Router /v1/auth/refresh-token [post]
func RefreshTokenHandler(c *fiber.Ctx) error { return nil }

// @Summary Logout
// @Description Logout APP
// @Tags Authentication
// @Security X-Signature
// @Security X-Timestamp
// @Security X-DeviceId
// @Accept json
// @Produce json
// @Param auth body routes.RefreshToken true "RefreshToken info"
// @Success 200 {object} map[string]bool
// @Router /v1/auth/logout [post]
func LogoutHandler(c *fiber.Ctx) error { return nil }

// @Summary Logout All Device
// @Description Logout All Device
// @Tags Authentication
// @Security X-Signature
// @Security X-Timestamp
// @Security X-DeviceId
// @Success 200 {object} map[string]bool
// @Accept json
// @Produce json
// @Param auth body routes.RefreshToken true "RefreshToken info"
// @Router /v1/auth/logout-all [post]
func LogoutAllHandler(c *fiber.Ctx) error { return nil }
