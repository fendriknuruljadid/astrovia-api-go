package routes

import "github.com/gofiber/fiber/v2"

// =================== Dummy Pricing model untuk Swagger ===================

type Pricing struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// =================== Dummy Handlers untuk Swagger ===================

// @Summary Get all pricing
// @Description Get list of pricing
// @Tags Pricing
// @Security BearerAuth
// @Security X-Signature
// @Security X-Timestamp
// @Produce json
// @Success 200 {array} routes.Pricing
// @Router /v1/pricing [get]
func GetPricingsHandler(c *fiber.Ctx) error { return nil }

// @Summary Get user by ID
// @Description Get single user
// @Tags Pricing
// @Security BearerAuth
// @Security X-Signature
// @Security X-Timestamp
// @Produce json
// @Param id path string true "Pricing ID"
// @Success 200 {object} routes.Pricing
// @Router /v1/pricing/{id} [get]
func GetPricingHandler(c *fiber.Ctx) error { return nil }

// @Summary Create user
// @Description Create new user
// @Tags Pricing
// @Security X-Signature
// @Security X-Timestamp
// @Accept json
// @Produce json
// @Param user body routes.Pricing true "Pricing info"
// @Success 201 {object} routes.Pricing
// @Router /v1/pricing [post]
func CreatePricingHandler(c *fiber.Ctx) error { return nil }

// @Summary Update user
// @Description Update user by ID
// @Tags Pricing
// @Security BearerAuth
// @Security X-Signature
// @Security X-Timestamp
// @Accept json
// @Produce json
// @Param id path string true "Pricing ID"
// @Param user body routes.Pricing true "Pricing info"
// @Success 200 {object} routes.Pricing
// @Router /v1/pricing/{id} [put]
func UpdatePricingHandler(c *fiber.Ctx) error { return nil }

// @Summary Delete user
// @Description Delete user by ID
// @Tags Pricing
// @Security BearerAuth
// @Security X-Signature
// @Security X-Timestamp
// @Param id path string true "Pricing ID"
// @Success 200 {object} map[string]bool
// @Router /v1/pricing/{id} [delete]
func DeletePricingHandler(c *fiber.Ctx) error { return nil }
