package dto

import "app/internal/services/v1/pricing/models"

type CreateDTO struct {
	Duration     int    `json:"duration" binding:"required,min=1"`
	AgentsID     string `json:"agents_id" binding:"required"`
	MonthlyPrice int64  `json:"monthly_price" binding:"required,min=0"`
	YearlyPrice  int64  `json:"yearly_price" binding:"required,min=0"`
	TokenMonthly int64  `json:"token_monthly" binding:"required,min=0"`
}

type UpdateDTO struct {
	Duration     *int    `json:"duration,omitempty" binding:"omitempty,min=1"`
	AgentsID     *string `json:"agents_id,omitempty"`
	MonthlyPrice *int64  `json:"monthly_price,omitempty" binding:"omitempty,min=0"`
	YearlyPrice  *int64  `json:"yearly_price,omitempty" binding:"omitempty,min=0"`
	TokenMonthly *int64  `json:"token_monthly,omitempty" binding:"omitempty,min=0"`
}

type ResponseDTO struct {
	ID           string `json:"id"`
	Duration     int    `json:"duration"`
	AgentsID     string `json:"agents_id"`
	MonthlyPrice int64  `json:"monthly_price"`
	YearlyPrice  int64  `json:"yearly_price"`
	TokenMonthly int64  `json:"token_monthly"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

func ToResponseDTO(p *models.Pricing) ResponseDTO {
	return ResponseDTO{
		ID:           p.ID,
		Duration:     p.Duration,
		AgentsID:     p.AgentsID,
		MonthlyPrice: p.MonthlyPrice,
		YearlyPrice:  p.YearlyPrice,
		TokenMonthly: p.TokenMonthly,
		CreatedAt:    p.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    p.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func ToResponseDTOs(pricing []models.Pricing) []ResponseDTO {
	res := make([]ResponseDTO, len(pricing))
	for i, p := range pricing {
		res[i] = ToResponseDTO(&p)
	}
	return res
}
