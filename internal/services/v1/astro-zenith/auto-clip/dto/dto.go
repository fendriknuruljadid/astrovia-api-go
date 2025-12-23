package dto

import (
	"app/internal/services/v1/astro-zenith/auto-clip/models"
	"time"
)

type CreateDTO struct {
	VideoURL string `json:"video_url" binding:"required,url"`
}

type UpdateDTO struct {
	VideoName      *string `json:"video_name,omitempty"`
	TranscriptFile *string `json:"transcript_file,omitempty"`
	Done           *bool   `json:"done,omitempty"`
	TokenUsage     *int64  `json:"token_usage,omitempty"`
}

type ResponseDTO struct {
	ID             string `json:"id"`
	DateUpload     string `json:"date_upload"`
	VideoURL       string `json:"video_url"`
	VideoName      string `json:"video_name"`
	TranscriptFile string `json:"transcript_file"`
	UsersID        string `json:"users_id"`
	UserAgentsID   string `json:"user_agents_id"`
	TokenUsage     int64  `json:"token_usage"`
	Done           bool   `json:"done"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

func ToResponseDTO(v *models.Videos) ResponseDTO {
	return ResponseDTO{
		ID:             v.ID,
		DateUpload:     v.DateUpload,
		VideoURL:       v.VideoURL,
		VideoName:      v.VideoName,
		TranscriptFile: v.TranscriptFile,
		UsersID:        v.UsersID,
		UserAgentsID:   v.UserAgentsID,
		TokenUsage:     v.TokenUsage,
		Done:           v.Done,
		CreatedAt:      v.CreatedAt.Format(time.DateTime),
		UpdatedAt:      v.UpdatedAt.Format(time.DateTime),
	}
}

func ToResponseDTOs(videos []models.Videos) []ResponseDTO {
	res := make([]ResponseDTO, len(videos))
	for i, v := range videos {
		res[i] = ToResponseDTO(&v)
	}
	return res
}
