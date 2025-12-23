package dto

import "app/internal/services/v1/user-agent/models"

type CreateDTO struct {
	UsersID  string `json:"users_id" binding:"required"`
	AgentsID string `json:"agents_id" binding:"required"`
	Tokens   int64  `json:"tokens" binding:"required,min=0"`
	Active   bool   `json:"active"`
	Expired  bool   `json:"expired"`
}

type UpdateDTO struct {
	UsersID  *string `json:"users_id,omitempty"`
	AgentsID *string `json:"agents_id,omitempty"`
	Tokens   *int64  `json:"tokens,omitempty" binding:"omitempty,min=0"`
	Active   *bool   `json:"active,omitempty"`
	Expired  *bool   `json:"expired,omitempty"`
}

type ResponseDTO struct {
	ID        string `json:"id"`
	UsersID   string `json:"users_id"`
	AgentsID  string `json:"agents_id"`
	Tokens    int64  `json:"tokens"`
	Active    bool   `json:"active"`
	Expired   bool   `json:"expired"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func ToResponseDTO(ua *models.UserAgent) ResponseDTO {
	return ResponseDTO{
		ID:        ua.ID,
		UsersID:   ua.UsersID,
		AgentsID:  ua.AgentsID,
		Tokens:    ua.Tokens,
		Active:    ua.Active,
		Expired:   ua.Expired,
		CreatedAt: ua.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: ua.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func ToResponseDTOs(items []models.UserAgent) []ResponseDTO {
	res := make([]ResponseDTO, len(items))
	for i, ua := range items {
		res[i] = ToResponseDTO(&ua)
	}
	return res
}
