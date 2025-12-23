package models

import (
	"context"
	"strings"
	"time"

	"app/internal/packages/utils"

	"github.com/uptrace/bun"
)

type Agent struct {
	bun.BaseModel `bun:"table:agents"`

	ID          string    `bun:"id,pk,notnull" json:"id"`
	Name        string    `bun:"name,notnull" json:"name"`
	Description string    `bun:"description" json:"description"`
	Logo        string    `bun:"logo" json:"logo"`
	URL         string    `bun:"url" json:"url"`
	CreatedAt   time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt   time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"updated_at"`
}

// Generate ID dan hash password otomatis sebelum insert
func (u *Agent) BeforeAppendModel(ctx context.Context, q bun.Query) error {
	if u.ID == "" {
		u.ID = "agn-" + strings.ToLower(utils.NewULID())
	}
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	return nil
}

// Update timestamp otomatis
func (u *Agent) BeforeUpdate(ctx context.Context, q bun.Query) error {
	u.UpdatedAt = time.Now()
	return nil
}
