package routes

import "github.com/gofiber/fiber/v2"

// =================== Dummy AutoClip model untuk Swagger ===================

type AutoClip struct {
	VideoUrl   string `json:"video_url"`
	VideoTitle string `json:"video_title"`
	Thumbnail  string `json:"thumbnail"`
}

// =================== Dummy Handlers untuk Swagger ===================

// @Summary Get all autoclip
// @Description Get list of autoclip
// @Tags AutoClips
// @Security BearerAuth
// @Security X-Signature
// @Security X-Timestamp
// @Produce json
// @Success 200 {array} routes.AutoClip
// @Router /v1/astro-zenith/auto-clip [get]
func GetAutoClipsHandler(c *fiber.Ctx) error { return nil }

// @Summary Get autoclip by ID
// @Description Get single autoclip
// @Tags AutoClips
// @Security BearerAuth
// @Security X-Signature
// @Security X-Timestamp
// @Produce json
// @Param id path string true "AutoClip ID"
// @Success 200 {object} routes.AutoClip
// @Router /v1/astro-zenith/auto-clip/{id} [get]
func GetAutoClipHandler(c *fiber.Ctx) error { return nil }

// @Summary Create autoclip
// @Description Create new autoclip
// @Tags AutoClips
// @Security BearerAuth
// @Security X-Signature
// @Security X-Timestamp
// @Accept json
// @Produce json
// @Param autoclip body routes.AutoClip true "AutoClip info"
// @Success 201 {object} routes.AutoClip
// @Router /v1/astro-zenith/auto-clip [post]
func CreateAutoClipHandler(c *fiber.Ctx) error { return nil }

// @Summary Update autoclip
// @Description Update autoclip by ID
// @Tags AutoClips
// @Security BearerAuth
// @Security X-Signature
// @Security X-Timestamp
// @Accept json
// @Produce json
// @Param id path string true "AutoClip ID"
// @Param autoclip body routes.AutoClip true "AutoClip info"
// @Success 200 {object} routes.AutoClip
// @Router /v1/astro-zenith/auto-clip/{id} [put]
func UpdateAutoClipHandler(c *fiber.Ctx) error { return nil }

// @Summary Delete autoclip
// @Description Delete autoclip by ID
// @Tags AutoClips
// @Security BearerAuth
// @Security X-Signature
// @Security X-Timestamp
// @Param id path string true "AutoClip ID"
// @Success 200 {object} map[string]bool
// @Router /v1/astro-zenith/auto-clip/{id} [delete]
func DeleteAutoClipHandler(c *fiber.Ctx) error { return nil }
