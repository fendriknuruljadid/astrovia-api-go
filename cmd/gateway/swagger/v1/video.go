package routes

import "github.com/gofiber/fiber/v2"

// =================== Dummy Video model untuk Swagger ===================

type Video struct {
	Url string `json:"url"`
}

// =================== Dummy Handlers untuk Swagger ===================

// @Summary Get all video
// @Description Get list of video
// @Tags Videos
// @Security BearerAuth
// @Security X-Signature
// @Security X-Timestamp
// @Produce json
// @Success 200 {array} routes.Video
// @Router /v1/video [get]
func GetVideosHandler(c *fiber.Ctx) error { return nil }

// @Summary Get video by ID
// @Description Get single video
// @Tags Videos
// @Security BearerAuth
// @Security X-Signature
// @Security X-Timestamp
// @Produce json
// @Param id path string true "Video ID"
// @Success 200 {object} routes.Video
// @Router /v1/video/{id} [get]
func GetVideoHandler(c *fiber.Ctx) error { return nil }

// @Summary Create video
// @Description Create new video
// @Tags Videos
// @Security X-Signature
// @Security X-Timestamp
// @Accept json
// @Produce json
// @Param video body routes.Video true "Video info"
// @Success 201 {object} routes.Video
// @Router /v1/video [post]
func CreateVideoHandler(c *fiber.Ctx) error { return nil }

// @Summary Update video
// @Description Update video by ID
// @Tags Videos
// @Security BearerAuth
// @Security X-Signature
// @Security X-Timestamp
// @Accept json
// @Produce json
// @Param id path string true "Video ID"
// @Param video body routes.Video true "Video info"
// @Success 200 {object} routes.Video
// @Router /v1/video/{id} [put]
func UpdateVideoHandler(c *fiber.Ctx) error { return nil }

// @Summary Delete video
// @Description Delete video by ID
// @Tags Videos
// @Security BearerAuth
// @Security X-Signature
// @Security X-Timestamp
// @Param id path string true "Video ID"
// @Success 200 {object} map[string]bool
// @Router /v1/video/{id} [delete]
func DeleteVideoHandler(c *fiber.Ctx) error { return nil }
