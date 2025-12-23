package models

import (
	"context"
	"strings"
	"time"

	"app/internal/packages/utils"

	"github.com/uptrace/bun"
)

type Videos struct {
	bun.BaseModel `bun:"table:video"`

	ID             string `bun:"id,pk,notnull" json:"id"`
	DateUpload     string `bun:"date_upload" json:"date_upload"`
	VideoURL       string `bun:"video_url" json:"video_url"`
	VideoName      string `bun:"video_name" json:"video_name"`
	TranscriptFile string `bun:"transcript_file" json:"transcript_file"`
	UsersID        string `bun:"users_id" json:"users_id"`
	UserAgentsID   string `bun:"user_agents_id" json:"user_agents_id"`
	TokenUsage     int64  `bun:"token_usage" json:"token_usage"`
	Done           bool   `bun:"done" json:"done"`

	CreatedAt time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp" json:"updated_at"`
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
