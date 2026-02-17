package handlers

import (
	"net/http"

	"app/internal/packages/errors"
	"app/internal/packages/response"
	"app/internal/packages/utils"
	"app/internal/services/v1/user-agent/dto"
	"app/internal/services/v1/user-agent/models"
	"app/internal/services/v1/user-agent/repository"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func CreateUserAgent(c *gin.Context) {
	var req dto.CreateDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			c.JSON(402, response.Error(402, "validation failed", utils.ValidationErrors(ve)))
			return
		}

		c.JSON(400, response.Error(400, "invalid request payload", err.Error()))
		return
	}

	userAgent := models.UserAgent{
		UsersID:  req.UsersID,
		AgentsID: req.AgentsID,
		Active:   req.Active,
		// Expired:  req.Expired,
	}

	if err := repository.CreateUserAgent(&userAgent); err != nil {
		c.Error(errors.NewInternal(utils.ParseDBError(err)))
		return
	}

	c.JSON(201, response.Success(201, "success", dto.ToResponseDTO(&userAgent)))
}

func GetUserAgents(c *gin.Context) {
	items, err := repository.GetUserAgents()
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			response.Error(500, "server internal error", err.Error()))
		return
	}

	c.JSON(200, response.Success(200, "success", dto.ToResponseDTOs(items)))
}

func GetUserAgentByID(c *gin.Context) {
	id := c.Param("id")

	item, err := repository.GetUserAgentByID(id)
	if err != nil || item == nil {
		c.Error(errors.NewNotFound("user agent not found"))
		return
	}

	c.JSON(200, response.Success(200, "success", dto.ToResponseDTO(item)))
}

func UpdateUserAgent(c *gin.Context) {
	id := c.Param("id")

	var req dto.UpdateDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			c.JSON(402, response.Error(402, "validation failed", utils.ValidationErrors(ve)))
			return
		}

		c.JSON(400, response.Error(400, "invalid request payload", err.Error()))
		return
	}

	userAgent, err := repository.GetUserAgentByID(id)
	if err != nil || userAgent == nil {
		c.Error(errors.NewNotFound("user agent not found"))
		return
	}

	// Apply updates
	if req.UsersID != nil {
		userAgent.UsersID = *req.UsersID
	}
	if req.AgentsID != nil {
		userAgent.AgentsID = *req.AgentsID
	}

	if req.Active != nil {
		userAgent.Active = *req.Active
	}
	// if req.Expired != nil {
	// 	userAgent.Expired = *req.Expired
	// }

	if err := repository.UpdateUserAgent(userAgent); err != nil {
		c.Error(errors.NewInternal(utils.ParseDBError(err)))
		return
	}

	c.JSON(200, response.Success(200, "updated", dto.ToResponseDTO(userAgent)))
}

func DeleteUserAgent(c *gin.Context) {
	id := c.Param("id")

	userAgent, err := repository.GetUserAgentByID(id)
	if err != nil || userAgent == nil {
		c.Error(errors.NewNotFound("user agent not found"))
		return
	}

	if err := repository.DeleteUserAgent(id); err != nil {
		c.Error(errors.NewInternal(utils.ParseDBError(err)))
		return
	}

	c.JSON(200, response.Success(200, "deleted successfully", gin.H{
		"deleted": true,
		"id":      id,
	}))
}
