package repository

import (
	"time"
)

type AgentWithAccessRow struct {
	ID          string `bun:"id"`
	Name        string `bun:"name"`
	Description string `bun:"description"`
	Logo        string `bun:"logo"`
	URL         string `bun:"url"`

	Active    *bool      `bun:"active"` // pointer karena LEFT JOIN
	ExpiresAt *time.Time `bun:"expired"`
}
