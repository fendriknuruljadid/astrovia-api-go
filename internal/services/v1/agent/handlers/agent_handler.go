package handlers

import (
	"app/internal/packages/errors"
	"app/internal/packages/response"
	"app/internal/packages/utils"
	"app/internal/services/v1/agent/dto"
	"app/internal/services/v1/agent/models"
	"app/internal/services/v1/agent/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func CreateAgent(c *gin.Context) {
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
	agent := models.Agent{
		Name: req.Name,
	}

	if err := repository.CreateAgent(&agent); err != nil {
		c.Error(errors.NewInternal(utils.ParseDBError(err)))
		return
	}

	res := dto.ToResponseDTO(&agent)

	c.JSON(201, response.Success(201, "success", res))
}

// func GetAgents(c *gin.Context) {
// 	agents, err := repository.GetAgents()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, response.Error(500, "Server internal error", err.Error()))
// 		return
// 	}

// 	res := dto.ToResponseDTOs(agents)
// 	c.JSON(200, response.Success(200, "success", res))
// }

func GetAgents(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, response.Error(
			401, "Unauthorized", nil,
		))
		return
	}

	agents, err := repository.GetAgentsWithUserAccess(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(
			500, "Server internal error", err.Error(),
		))
		return
	}

	res := dto.ToAgentResponseDTOs(agents)
	c.JSON(http.StatusOK, response.Success(
		200, "success", res,
	))
}

func GetAgentsPublic(c *gin.Context) {
	agents, err := repository.GetAgents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(500, "Server internal error", err.Error()))
		return
	}

	res := dto.ToResponseDTOs(agents)
	c.JSON(200, response.Success(200, "success", res))
}

func GetAgentByID(c *gin.Context) {
	id := c.Param("id")
	agent, err := repository.GetAgentByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Agent not found"})
		return
	}
	res := dto.ToResponseDTO(agent)
	c.JSON(200, response.Success(200, "success", res))
}

func UpdateAgent(c *gin.Context) {
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

	agent, err := repository.GetAgentByID(id)
	if err != nil {
		c.Error(errors.NewNotFound("agent not found"))
		return
	}

	// Apply updates safely
	if req.Name != nil {
		agent.Name = *req.Name
	}
	if err := repository.UpdateAgent(agent); err != nil {
		c.Error(errors.NewInternal(utils.ParseDBError(err)))
		return
	}

	c.JSON(200, response.Success(200, "updated", dto.ToResponseDTO(agent)))
}

func DeleteAgent(c *gin.Context) {
	id := c.Param("id")

	// Cek dulu apakah agent ada
	agent, err := repository.GetAgentByID(id)
	if err != nil || agent == nil {
		c.Error(errors.NewNotFound("agent not found"))
		return
	}

	// Jika ada, lanjut hapus
	if err := repository.DeleteAgent(id); err != nil {
		c.Error(errors.NewInternal(utils.ParseDBError(err)))
		return
	}

	c.JSON(200, response.Success(200, "deleted successfully", gin.H{
		"deleted": true,
		"id":      id,
	}))
}
