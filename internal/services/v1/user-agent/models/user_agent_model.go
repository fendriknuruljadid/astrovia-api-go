package models

import (
	"context"
	"strings"
	"time"

	"app/internal/packages/utils"

	"github.com/uptrace/bun"
)

type UserAgent struct {
	bun.BaseModel `bun:"table:user_agents"`

	ID        string    `bun:"id,pk,notnull" json:"id"`
	UsersID   string    `bun:"users_id,notnull" json:"users_id"`
	AgentsID  string    `bun:"agents_id,notnull" json:"agents_id"`
	Active    bool      `bun:"active,notnull,default:false" json:"active"`
	Expired   bool      `bun:"expired,notnull,default:false" json:"expired"`
	Tokens    int64     `bun:"tokens,notnull,default:0" json:"tokens"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"updated_at"`
}

// Generate ID dan hash password otomatis sebelum insert
func (u *UserAgent) BeforeAppendModel(ctx context.Context, q bun.Query) error {
	if u.ID == "" {
		u.ID = "usr-agn-" + strings.ToLower(utils.NewULID())
	}
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	return nil
}

// Update timestamp otomatis
func (u *UserAgent) BeforeUpdate(ctx context.Context, q bun.Query) error {
	u.UpdatedAt = time.Now()
	return nil
}
