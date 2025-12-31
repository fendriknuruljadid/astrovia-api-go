package models

import (
	"context"
	"strings"
	"time"

	"app/internal/packages/utils"

	"github.com/uptrace/bun"
)

type RefreshTokens struct {
	bun.BaseModel `bun:"table:refresh_tokens"`

	ID        string    `bun:"id,pk,notnull" json:"id"`
	UserID    string    `bun:"users_id,notnull" json:"users_id"`
	Revoke    bool      `bun:"revoke,notnull,default:false" json:"revoke"`
	ExpiredAt time.Time `bun:"expired_at,notnull" json:"expired_at"`
	Token     string    `bun:"token,notnull" json:"token"`
	DeviceId  string    `bun:"device_id,notnull" json:"device_id"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"created_at"`
}

// Generate ID dan hash password otomatis sebelum insert
func (u *RefreshTokens) BeforeAppendModel(ctx context.Context, q bun.Query) error {
	if u.ID == "" {
		u.ID = "ref-" + strings.ToLower(utils.NewULID())
	}
	u.CreatedAt = time.Now()

	return nil
}
