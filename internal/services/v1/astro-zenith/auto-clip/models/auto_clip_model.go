package models

import (
	"app/internal/packages/utils"
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/uptrace/bun"
)

type Clip struct {
	bun.BaseModel `bun:"table:clips"`

	ID         string                 `bun:",pk"`
	ClipsURL   string                 `bun:"clips_url"`
	MetaData   map[string]interface{} `bun:"meta_data,type:json"`
	Title      string
	ViralScore float64 `bun:"viral_score"`
	Reason     string
	Duration   float64
	HookText   string `bun:"hook_text"`
	Status     string
	VideosID   string    `bun:"videos_id"`
	CreatedAt  time.Time `bun:"created_at"`
}

type Videos struct {
	bun.BaseModel `bun:"table:videos"`

	ID             string    `bun:"id,pk,notnull" json:"id"`
	DateUpload     time.Time `bun:"date_upload,nullzero,default:current_timestamp" json:"date_upload"`
	VideoUrl       string    `bun:"video_url" json:"video_url"`
	VideoName      string    `bun:"video_name" json:"video_name"`
	VideoTitle     string    `bun:"video_title" json:"video_title"`
	Thumbnail      string    `bun:"thumbnail" json:"thumbnail"`
	TranscriptFile string    `bun:"transcript_file" json:"transcript_file"`
	UsersID        string    `bun:"users_id" json:"users_id"`
	UserAgentsID   string    `bun:"user_agents_id" json:"user_agents_id"`
	TokenUsage     int64     `bun:"token_usage" json:"token_usage"`
	Done           bool      `bun:"done" json:"done"`
	Clips          []*Clip   `bun:"rel:has-many,join:id=videos_id"`
	AspectRatio    string    `bun:"aspect_ratio" json:"aspect_ratio"`
	ResizeMode     string    `bun:"resize_mode" json:"resize_mode"`
	OutputType     string    `bun:"output_type" json:"output_type"`
	CaptionPreset json.RawMessage `bun:"caption_preset,type:jsonb" json:"caption_preset"`
	CreatedAt      time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt      time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp" json:"updated_at"`
}

// Generate ID dan hash password otomatis sebelum insert
func (u *Videos) BeforeAppendModel(ctx context.Context, q bun.Query) error {
	if u.ID == "" {
		u.ID = "vid-" + strings.ToLower(utils.NewULID())
	}
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	return nil
}

// Update timestamp otomatis
func (u *Videos) BeforeUpdate(ctx context.Context, q bun.Query) error {
	u.UpdatedAt = time.Now()
	return nil
}
