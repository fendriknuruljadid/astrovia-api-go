package dto

import (
	"app/internal/services/v1/pricing/dto"
	"app/internal/services/v1/user-agent/models"
)

type CreateDTO struct {
	UsersID  string `json:"users_id" binding:"required"`
	AgentsID string `json:"agents_id" binding:"required"`
	Active   bool   `json:"active"`
	Expired  bool   `json:"expired"`
}

type UpdateDTO struct {
	UsersID  *string `json:"users_id,omitempty"`
	AgentsID *string `json:"agents_id,omitempty"`
	Active   *bool   `json:"active,omitempty"`
	Expired  *bool   `json:"expired,omitempty"`
}

type ResponseDTO struct {
	ID        string `json:"id"`
	UsersID   string `json:"users_id"`
	AgentsID  string `json:"agents_id"`
	Active    bool   `json:"active"`
	Expired   string `json:"expired"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Pricing   dto.ResponseDTO
}

func ToResponseDTO(ua *models.UserAgent) ResponseDTO {
	return ResponseDTO{
		ID:        ua.ID,
		UsersID:   ua.UsersID,
		AgentsID:  ua.AgentsID,
		Active:    ua.Active,
		Expired:   ua.Expired.Format("2006-01-02"),
		CreatedAt: ua.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: ua.UpdatedAt.Format("2006-01-02 15:04:05"),
		Pricing:   dto.ToResponseDTO(ua.Pricing),
	}
}

func ToResponseDTOs(items []models.UserAgent) []ResponseDTO {
	res := make([]ResponseDTO, len(items))
	for i, ua := range items {
		res[i] = ToResponseDTO(&ua)
	}
	return res
}
