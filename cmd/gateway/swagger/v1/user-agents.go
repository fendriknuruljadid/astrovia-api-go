package routes

import "github.com/gofiber/fiber/v2"

// =================== Dummy UserAgent model untuk Swagger ===================

type UserAgent struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// =================== Dummy Handlers untuk Swagger ===================

// @Summary Get all user agent
// @Description Get list of user agent
// @Tags UserAgents
// @Security BearerAuth
// @Security X-Signature
// @Security X-Timestamp
// @Produce json
// @Success 200 {array} routes.UserAgent
// @Router /v1/user-agents [get]
func GetUserAgentsHandler(c *fiber.Ctx) error { return nil }

// @Summary Get user agent by ID
// @Description Get single user agent
// @Tags UserAgents
// @Security BearerAuth
// @Security X-Signature
// @Security X-Timestamp
// @Produce json
// @Param id path string true "UserAgent ID"
// @Success 200 {object} routes.UserAgent
// @Router /v1/user-agents/{id} [get]
func GetUserAgentHandler(c *fiber.Ctx) error { return nil }

// @Summary Create user agent
// @Description Create new user agent
// @Tags UserAgents
// @Security X-Signature
// @Security X-Timestamp
// @Accept json
// @Produce json
// @Param useragent body routes.UserAgent true "UserAgent info"
// @Success 201 {object} routes.UserAgent
// @Router /v1/user-agents [post]
func CreateUserAgentHandler(c *fiber.Ctx) error { return nil }

// @Summary Update user agent
// @Description Update user agent by ID
// @Tags UserAgents
// @Security BearerAuth
// @Security X-Signature
// @Security X-Timestamp
// @Accept json
// @Produce json
// @Param id path string true "UserAgent ID"
// @Param useragent body routes.UserAgent true "UserAgent info"
// @Success 200 {object} routes.UserAgent
// @Router /v1/user-agents/{id} [put]
func UpdateUserAgentHandler(c *fiber.Ctx) error { return nil }

// @Summary Delete user agent
// @Description Delete user agent by ID
// @Tags UserAgents
// @Security BearerAuth
// @Security X-Signature
// @Security X-Timestamp
// @Param id path string true "UserAgent ID"
// @Success 200 {object} map[string]bool
// @Router /v1/user-agents/{id} [delete]
func DeleteUserAgentHandler(c *fiber.Ctx) error { return nil }
