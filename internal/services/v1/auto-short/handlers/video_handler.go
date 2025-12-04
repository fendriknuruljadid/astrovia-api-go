package handlers

import (
	"app/internal/packages/errors"
	"app/internal/packages/response"
	"app/internal/packages/utils"
	"app/internal/services/v1/auto-short/dto"
	"app/internal/services/v1/auto-short/models"
	"app/internal/services/v1/auto-short/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateVideo(c *gin.Context) {
	var req dto.CreateDTO

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
	video := models.Video{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := repository.CreateVideo(&video); err != nil {
		c.Error(errors.NewInternal(utils.ParseDBError(err)))
		return
	}

	res := dto.ToResponseDTO(&video)

	c.JSON(201, response.Success(201, "success", res))
}

func GetVideo(c *gin.Context) {
	video, err := repository.GetVideo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(500, "Server internal error", err.Error()))
		return
	}

	res := dto.ToResponseDTOs(video)
	c.JSON(200, response.Success(200, "success", res))
}

func GetVideoByID(c *gin.Context) {
	id := c.Param("id")
	video, err := repository.GetVideoByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Video not found"})
		return
	}
	res := dto.ToResponseDTO(video)
	c.JSON(200, response.Success(200, "success", res))
}

func UpdateVideo(c *gin.Context) {
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

	video, err := repository.GetVideoByID(id)
	if err != nil {
		c.Error(errors.NewNotFound("video not found"))
		return
	}

	// Apply updates safely
	if req.Name != nil {
		video.Name = *req.Name
	}
	if req.Email != nil {
		video.Email = *req.Email
	}
	if req.Password != nil {
		video.Password = *req.Password
	}

	if err := repository.UpdateVideo(video); err != nil {
		c.Error(errors.NewInternal(utils.ParseDBError(err)))
		return
	}

	c.JSON(200, response.Success(200, "updated", dto.ToResponseDTO(video)))
}

func DeleteVideo(c *gin.Context) {
	id := c.Param("id")

	// Cek dulu apakah video ada
	video, err := repository.GetVideoByID(id)
	if err != nil || video == nil {
		c.Error(errors.NewNotFound("video not found"))
		return
	}

	// Jika ada, lanjut hapus
	if err := repository.DeleteVideo(id); err != nil {
		c.Error(errors.NewInternal(utils.ParseDBError(err)))
		return
	}

	c.JSON(200, response.Success(200, "deleted successfully", gin.H{
		"deleted": true,
		"id":      id,
	}))
}
