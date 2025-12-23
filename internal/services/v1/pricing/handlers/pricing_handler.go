package handlers

import (
	"app/internal/packages/errors"
	"app/internal/packages/response"
	"app/internal/packages/utils"
	"app/internal/services/v1/pricing/dto"
	"app/internal/services/v1/pricing/models"
	"app/internal/services/v1/pricing/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func CreatePricing(c *gin.Context) {
	var req dto.CreateDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			c.JSON(402, response.Error(402, "validation failed", utils.ValidationErrors(ve)))
			return
		}

		c.JSON(400, response.Error(400, "invalid request payload", err.Error()))
		return
	}

	pricing := models.Pricing{
		Duration:     req.Duration,
		AgentsID:     req.AgentsID,
		MonthlyPrice: req.MonthlyPrice,
		YearlyPrice:  req.YearlyPrice,
		TokenMonthly: req.TokenMonthly,
	}

	if err := repository.CreatePricing(&pricing); err != nil {
		c.Error(errors.NewInternal(utils.ParseDBError(err)))
		return
	}

	c.JSON(201, response.Success(201, "success", dto.ToResponseDTO(&pricing)))
}

func GetPricings(c *gin.Context) {
	pricings, err := repository.GetPricings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(500, "server internal error", err.Error()))
		return
	}

	c.JSON(200, response.Success(200, "success", dto.ToResponseDTOs(pricings)))
}

func GetPricingByID(c *gin.Context) {
	id := c.Param("id")

	pricing, err := repository.GetPricingByID(id)
	if err != nil || pricing == nil {
		c.Error(errors.NewNotFound("pricing not found"))
		return
	}

	c.JSON(200, response.Success(200, "success", dto.ToResponseDTO(pricing)))
}

func UpdatePricing(c *gin.Context) {
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

	pricing, err := repository.GetPricingByID(id)
	if err != nil || pricing == nil {
		c.Error(errors.NewNotFound("pricing not found"))
		return
	}

	// Apply updates
	if req.Duration != nil {
		pricing.Duration = *req.Duration
	}
	if req.AgentsID != nil {
		pricing.AgentsID = *req.AgentsID
	}
	if req.MonthlyPrice != nil {
		pricing.MonthlyPrice = *req.MonthlyPrice
	}
	if req.YearlyPrice != nil {
		pricing.YearlyPrice = *req.YearlyPrice
	}
	if req.TokenMonthly != nil {
		pricing.TokenMonthly = *req.TokenMonthly
	}

	if err := repository.UpdatePricing(pricing); err != nil {
		c.Error(errors.NewInternal(utils.ParseDBError(err)))
		return
	}

	c.JSON(200, response.Success(200, "updated", dto.ToResponseDTO(pricing)))
}

func DeletePricing(c *gin.Context) {
	id := c.Param("id")

	pricing, err := repository.GetPricingByID(id)
	if err != nil || pricing == nil {
		c.Error(errors.NewNotFound("pricing not found"))
		return
	}

	if err := repository.DeletePricing(id); err != nil {
		c.Error(errors.NewInternal(utils.ParseDBError(err)))
		return
	}

	c.JSON(200, response.Success(200, "deleted successfully", gin.H{
		"deleted": true,
		"id":      id,
	}))
}
