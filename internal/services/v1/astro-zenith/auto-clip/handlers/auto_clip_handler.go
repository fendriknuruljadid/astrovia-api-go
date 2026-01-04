package handlers

import (
	"app/internal/packages/errors"
	"app/internal/packages/redis"
	"app/internal/packages/response"
	"app/internal/packages/utils"
	"app/internal/services/v1/astro-zenith/auto-clip/dto"
	"app/internal/services/v1/astro-zenith/auto-clip/models"
	"app/internal/services/v1/astro-zenith/auto-clip/repository"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateVideos(c *gin.Context) {
	var req dto.CreateDTO
	user, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, response.Error(401, "user ID not found", nil))
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		// Deteksi apakah ini validation error dari validator bawaan Gin
		if ve, ok := err.(validator.ValidationErrors); ok {
			c.JSON(402, response.Error(402, "validation failed", utils.ValidationErrors(ve)))
			return
		}

		// Kalau bukan validasi, baru fallback ke invalid JSON
		c.JSON(400, response.Error(400, "invalid request payload", err.Error()))
		return
	}
	uid, ok := user.(string)
	if !ok {
		c.JSON(500, response.Error(500, "invalid user id type", nil))
		return
	}

	videos := models.Videos{
		VideoUrl:     req.VideoUrl,
		Thumbnail:    req.Thumbnail,
		VideoTitle:   req.VideoTitle,
		UsersID:      uid,
		UserAgentsID: "usr-agn-01kd52hz2wpw5vsm306m2dj674",
	}

	if err := repository.CreateVideos(&videos); err != nil {
		c.Error(errors.NewInternal(utils.ParseDBError(err)))
		return
	}
	payload, _ := json.Marshal(gin.H{
		"id":        videos.ID,
		"video_url": videos.VideoUrl,
	})
	redis.Rdb.RPush(
		redis.Ctx,
		"queue:auto_clips",
		payload,
	)
	res := dto.ToResponseDTO(&videos)
	c.JSON(201, response.Success(201, "success", res))
}

func GetVideos(c *gin.Context) {
	videos, err := repository.GetVideos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(500, "Server internal error", err.Error()))
		return
	}

	res := dto.ToResponseDTOs(videos)
	c.JSON(200, response.Success(200, "success", res))
}

func GetVideosByID(c *gin.Context) {
	id := c.Param("id")
	videos, err := repository.GetVideosByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Videos not found"})
		return
	}
	res := dto.ToResponseDTO(videos)
	c.JSON(200, response.Success(200, "success", res))
}

func UpdateVideos(c *gin.Context) {
	id := c.Param("id")

	var req dto.UpdateDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		// Deteksi apakah ini validation error dari validator bawaan Gin
		if ve, ok := err.(validator.ValidationErrors); ok {
			c.JSON(402, response.Error(402, "validation failed", utils.ValidationErrors(ve)))
			return
		}

		// Kalau bukan validasi, baru fallback ke invalid JSON
		c.JSON(400, response.Error(400, "invalid request payload", err.Error()))
		return
	}

	videos, err := repository.GetVideosByID(id)
	if err != nil {
		c.Error(errors.NewNotFound("videos not found"))
		return
	}

	if err := repository.UpdateVideos(videos); err != nil {
		c.Error(errors.NewInternal(utils.ParseDBError(err)))
		return
	}

	c.JSON(200, response.Success(200, "updated", dto.ToResponseDTO(videos)))
}

func DeleteVideos(c *gin.Context) {
	id := c.Param("id")

	// Cek dulu apakah videos ada
	videos, err := repository.GetVideosByID(id)
	if err != nil || videos == nil {
		c.Error(errors.NewNotFound("videos not found"))
		return
	}

	// Jika ada, lanjut hapus
	if err := repository.DeleteVideos(id); err != nil {
		c.Error(errors.NewInternal(utils.ParseDBError(err)))
		return
	}

	c.JSON(200, response.Success(200, "deleted successfully", gin.H{
		"deleted": true,
		"id":      id,
	}))
}
