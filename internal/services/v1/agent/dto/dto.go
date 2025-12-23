package dto

import (
	"app/internal/services/v1/agent/models"
	"time"
)

type CreateDTO struct {
	Name        string `json:"name" binding:"required,min=3"`
	Description string `json:"description" binding:"omitempty"`
	Logo        string `json:"logo" binding:"omitempty"`
	URL         string `json:"url" binding:"omitempty,url"`
}

type UpdateDTO struct {
	Name        *string `json:"name,omitempty" binding:"omitempty,min=3"`
	Description *string `json:"description,omitempty"`
	Logo        *string `json:"logo,omitempty"`
	URL         *string `json:"url,omitempty" binding:"omitempty,url"`
}

type ResponseDTO struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Logo        string `json:"logo"`
	URL         string `json:"url"`
	HasAccess   bool   `json:"has_access"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func ToResponseDTO(a *models.Agent) ResponseDTO {
	return ResponseDTO{
		ID:          a.ID,
		Name:        a.Name,
		Description: a.Description,
		Logo:        a.Logo,
		URL:         a.URL,
		CreatedAt:   a.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   a.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func ToResponseDTOs(agents []models.Agent) []ResponseDTO {
	res := make([]ResponseDTO, len(agents))
	for i, a := range agents {
		res[i] = ToResponseDTO(&a)
	}
	return res
}

type AgentResponseDTO struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Logo        string     `json:"logo"`
	URL         string     `json:"url"`
	HasAccess   bool       `json:"has_access"`
	Expired     bool       `json:"expired"`
	ExpiredAt   *time.Time `json:"expired_at"`
}
