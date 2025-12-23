package routes

import "github.com/gofiber/fiber/v2"

// =================== Dummy AutoCaption model untuk Swagger ===================

type AutoCaption struct {
	Url string `json:"url"`
}

// =================== Dummy Handlers untuk Swagger ===================

// @Summary Get all autocaption
// @Description Get list of autocaption
// @Tags AutoCaption
// @Security BearerAuth
// @Security X-Signature
// @Security X-Timestamp
// @Produce json
// @Success 200 {array} routes.AutoCaption
// @Router /v1/astro-zenith/auto-caption [get]
func GetAutoCaptionsHandler(c *fiber.Ctx) error { return nil }

// @Summary Get autocaption by ID
// @Description Get single autocaption
// @Tags AutoCaption
// @Security BearerAuth
// @Security X-Signature
// @Security X-Timestamp
// @Produce json
// @Param id path string true "AutoCaption ID"
// @Success 200 {object} routes.AutoCaption
// @Router /v1/astro-zenith/auto-caption/{id} [get]
func GetAutoCaptionHandler(c *fiber.Ctx) error { return nil }

// @Summary Create autocaption
// @Description Create new autocaption
// @Tags AutoCaption
// @Security X-Signature
// @Security X-Timestamp
// @Accept json
// @Produce json
// @Param autocaption body routes.AutoCaption true "AutoCaption info"
// @Success 201 {object} routes.AutoCaption
// @Router /v1/astro-zenith/auto-caption [post]
func CreateAutoCaptionHandler(c *fiber.Ctx) error { return nil }

// @Summary Update autocaption
// @Description Update autocaption by ID
// @Tags AutoCaption
// @Security BearerAuth
// @Security X-Signature
// @Security X-Timestamp
// @Accept json
// @Produce json
// @Param id path string true "AutoCaption ID"
// @Param autocaption body routes.AutoCaption true "AutoCaption info"
// @Success 200 {object} routes.AutoCaption
// @Router /v1/astro-zenith/auto-caption/{id} [put]
func UpdateAutoCaptionHandler(c *fiber.Ctx) error { return nil }

// @Summary Delete autocaption
// @Description Delete autocaption by ID
// @Tags AutoCaption
// @Security BearerAuth
// @Security X-Signature
// @Security X-Timestamp
// @Param id path string true "AutoCaption ID"
// @Success 200 {object} map[string]bool
// @Router /v1/astro-zenith/auto-caption/{id} [delete]
func DeleteAutoCaptionHandler(c *fiber.Ctx) error { return nil }
