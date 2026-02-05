package models

import (
	"context"
	"strings"
	"time"

	"app/internal/packages/utils"
	"app/internal/services/v1/agent/models"

	"github.com/uptrace/bun"
)

type Order struct {
	bun.BaseModel `bun:"table:orders"`

	ID               string        `bun:"id,pk,notnull" json:"id"`
	OrderDate        *time.Time    `bun:"order_date,nullzero" json:"order_date"`
	Amount           float64       `bun:"amount,nullzero" json:"amount"`
	PricingName      string        `bun:"pricing_name,nullzero" json:"pricing_name"`
	AgentName        string        `bun:"agent_name,nullzero" json:"agent_name"`
	PaymentMethod    string        `bun:"payment_method,nullzero" json:"payment_method"`
	Discount         float64       `bun:"discount,nullzero" json:"discount"`
	Status           string        `bun:"status,nullzero" json:"status"`
	ExpiryPeriod     int           `bun:"expiry_period,nullzero" json:"expiry_period"`
	ExpiredAt        *time.Time    `bun:"expired_at,nullzero" json:"expired_at"`
	VANumber         string        `bun:"va_number,nullzero" json:"va_number"`
	Reference        string        `bun:"reference,nullzero" json:"reference"`
	QRCode           string        `bun:"qr_code,nullzero" json:"qr_code"`
	PaymentURL       string        `bun:"payment_url,nullzero" json:"payment_url"`
	PublisherOrderID string        `bun:"publisher_order_id,nullzero" json:"publisher_order_id"`
	IssuerCode       string        `bun:"issuer_code,nullzero" json:"issuer_code"`
	InvoiceNumber    string        `bun:"invoice_number,nullzero" json:"invoice_number"`
	UsersID          string        `bun:"users_id,nullzero" json:"users_id"`
	PricingID        string        `bun:"pricing_id,nullzero" json:"pricing_id"`
	AgentsID         string        `bun:"agents_id,nullzero" json:"agents_id"`
	Agent            *models.Agent `bun:"rel:belongs-to,join:agents_id=id"`
	CreatedAt        *time.Time    `bun:"created_at,nullzero" json:"created_at"`
	UpdatedAt        *time.Time    `bun:"updated_at,nullzero" json:"updated_at"`
}

// Generate ID & set timestamp sebelum insert
func (o *Order) BeforeAppendModel(ctx context.Context, q bun.Query) error {
	if o.ID == "" {
		ulid := strings.ToLower(utils.NewULID())

		o.ID = "order-" + ulid
		o.InvoiceNumber = "order-" + ulid[len(ulid)-6:]
	}

	now := time.Now()
	if o.CreatedAt == nil {
		o.CreatedAt = &now
	}
	o.UpdatedAt = &now

	if o.OrderDate == nil {
		o.OrderDate = &now
	}

	return nil
}

// Update timestamp otomatis
func (o *Order) BeforeUpdate(ctx context.Context, q bun.Query) error {
	now := time.Now()
	o.UpdatedAt = &now
	return nil
}
