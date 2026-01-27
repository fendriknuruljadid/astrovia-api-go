package models

import (
	"context"
	"strings"
	"time"

	"app/internal/packages/utils"

	"github.com/uptrace/bun"
)

type Pricing struct {
	bun.BaseModel `bun:"table:pricing"`

	ID           string    `bun:"id,pk,notnull" json:"id"`
	Duration     int       `bun:"duration,notnull" json:"duration"`
	AgentsID     string    `bun:"agents_id,notnull" json:"agents_id"`
	MonthlyPrice int64     `bun:"monthly_price,notnull" json:"monthly_price"`
	YearlyPrice  int64     `bun:"yearly_price,notnull" json:"yearly_price"`
	TokenMonthly int64     `bun:"token_monthly,notnull" json:"token_monthly"`
	CreatedAt    time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt    time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"updated_at"`
}

// Generate ID dan hash password otomatis sebelum insert
func (u *Pricing) BeforeAppendModel(ctx context.Context, q bun.Query) error {
	if u.ID == "" {
		u.ID = "prc-" + strings.ToLower(utils.NewULID())
	}
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	return nil
}

// Update timestamp otomatis
func (u *Pricing) BeforeUpdate(ctx context.Context, q bun.Query) error {
	u.UpdatedAt = time.Now()
	return nil
}
