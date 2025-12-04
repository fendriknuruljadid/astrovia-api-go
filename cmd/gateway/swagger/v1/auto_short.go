package routes

import "github.com/gofiber/fiber/v2"

// =================== Dummy AutoShort model untuk Swagger ===================

type AutoShort struct {
	Url string `json:"url"`
}

// =================== Dummy Handlers untuk Swagger ===================

// @Summary Get all video
// @Description Get list of video
// @Tags AutoShort
// @Security BearerAuth
// @Security X-Signature
// @Security X-Timestamp
// @Produce json
// @Success 200 {array} routes.AutoShort
// @Router /v1/video [get]
func GetAutoShortsHandler(c *fiber.Ctx) error { return nil }

// @Summary Get video by ID
// @Description Get single video
// @Tags AutoShort
// @Security BearerAuth
// @Security X-Signature
// @Security X-Timestamp
// @Produce json
// @Param id path string true "AutoShort ID"
// @Success 200 {object} routes.AutoShort
// @Router /v1/video/{id} [get]
func GetAutoShortHandler(c *fiber.Ctx) error { return nil }

// @Summary Create video
// @Description Create new video
// @Tags AutoShort
// @Security X-Signature
// @Security X-Timestamp
// @Accept json
// @Produce json
// @Param video body routes.AutoShort true "AutoShort info"
// @Success 201 {object} routes.AutoShort
// @Router /v1/video [post]
func CreateAutoShortHandler(c *fiber.Ctx) error { return nil }

// @Summary Update video
// @Description Update video by ID
// @Tags AutoShort
// @Security BearerAuth
// @Security X-Signature
// @Security X-Timestamp
// @Accept json
// @Produce json
// @Param id path string true "AutoShort ID"
// @Param video body routes.AutoShort true "AutoShort info"
// @Success 200 {object} routes.AutoShort
// @Router /v1/video/{id} [put]
func UpdateAutoShortHandler(c *fiber.Ctx) error { return nil }

// @Summary Delete video
// @Description Delete video by ID
// @Tags AutoShort
// @Security BearerAuth
// @Security X-Signature
// @Security X-Timestamp
// @Param id path string true "AutoShort ID"
// @Success 200 {object} map[string]bool
// @Router /v1/video/{id} [delete]
func DeleteAutoShortHandler(c *fiber.Ctx) error { return nil }
