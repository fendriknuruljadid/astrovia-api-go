package dto

import (
	"app/internal/services/v1/astro-zenith/auto-clip/models"
	"time"
)

type CaptionPresetDTO struct {
	PresetName string `json:"preset_name" binding:"required"`
	Position   string `json:"position" binding:"required"`
}

type CreateDTO struct {
	VideoUrl      string           `json:"video_url" binding:"required,url"`
	VideoTitle    string           `json:"video_title" binding:"required"`
	Thumbnail     string           `json:"thumbnail" binding:"required"`
	ResizeMode    string           `json:"resize_mode" binding:"required"`
	AspectRatio   string           `json:"aspect_ratio" binding:"required"`
	CaptionPreset CaptionPresetDTO `json:"caption_preset" binding:"required"`
	OutputType    string           `json:"output_type" binding:"required"`
}

type UpdateDTO struct {
	VideoName      *string `json:"video_name,omitempty"`
	TranscriptFile *string `json:"transcript_file,omitempty"`
	Done           *bool   `json:"done,omitempty"`
	TokenUsage     *int64  `json:"token_usage,omitempty"`
}

type ResponseDTO struct {
	ID             string            `json:"id"`
	DateUpload     string            `json:"date_upload"`
	VideoUrl       string            `json:"video_url"`
	VideoTitle     string            `json:"video_title"`
	Thumbnail      string            `json:"thumbnail"`
	VideoName      string            `json:"video_name"`
	TranscriptFile string            `json:"transcript_file"`
	UsersID        string            `json:"users_id"`
	UserAgentsID   string            `json:"user_agents_id"`
	TokenUsage     int64             `json:"token_usage"`
	Done           bool              `json:"done"`
	CreatedAt      string            `json:"created_at"`
	UpdatedAt      string            `json:"updated_at"`
	Clips          []ClipResponseDTO `json:"clips"` // <-- tambahkan ini
}

type ClipResponseDTO struct {
	ID         string  `json:"id"`
	ClipsURL   string  `json:"clips_url"`
	Title      string  `json:"title"`
	Reason     string  `json:"reason"`
	HookText   string  `json:"hook_text"`
	ViralScore float64 `json:"viral_score"`
	Duration   float64 `json:"duration"`
	Status     string  `json:"status"`
}

func ToResponseDTO(v *models.Videos) ResponseDTO {
	clips := make([]ClipResponseDTO, len(v.Clips))
	for i, c := range v.Clips {
		clips[i] = ClipResponseDTO{
			ID:         c.ID,
			ClipsURL:   c.ClipsURL,
			Title:      c.Title,
			Reason:     c.Reason,
			HookText:   c.HookText,
			ViralScore: c.ViralScore,
			Duration:   c.Duration,
			Status:     c.Status,
		}
	}
	return ResponseDTO{
		ID:             v.ID,
		DateUpload:     v.DateUpload.Format(time.DateTime),
		VideoUrl:       v.VideoUrl,
		VideoTitle:     v.VideoTitle,
		Thumbnail:      v.Thumbnail,
		VideoName:      v.VideoName,
		TranscriptFile: v.TranscriptFile,
		UsersID:        v.UsersID,
		UserAgentsID:   v.UserAgentsID,
		TokenUsage:     v.TokenUsage,
		Done:           v.Done,
		CreatedAt:      v.CreatedAt.Format(time.DateTime),
		UpdatedAt:      v.UpdatedAt.Format(time.DateTime),
		Clips:          clips,
	}
}

func ToResponseDTOs(videos []models.Videos) []ResponseDTO {
	res := make([]ResponseDTO, len(videos))
	for i, v := range videos {
		res[i] = ToResponseDTO(&v)
	}
	return res
}
