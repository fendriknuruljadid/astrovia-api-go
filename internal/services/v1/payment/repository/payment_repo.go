package repository

import (
	"app/internal/packages/db"
	orderModels "app/internal/services/v1/payment/models"
	"app/internal/services/v1/pricing/models"
	userAgentModels "app/internal/services/v1/user-agent/models"
	userModels "app/internal/services/v1/user/models"
	"time"

	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func GetPricingByID(id string) (*models.Pricing, error) {
	ctx := context.Background()
	pricing := new(models.Pricing)
	err := db.DB.NewSelect().Model(pricing).Where("pricing.id = ?", id).Relation("Agent").Scan(ctx)
	if err != nil {
		fmt.Println("GetPricingByID failed:", err)
	}
	return pricing, err
}

func GetUserByEmailOrPhone(email, phone string) (*userModels.User, error) {
	var user userModels.User

	err := db.DB.NewSelect().
		Model(&user).
		Where("email = ? OR phone = ?", email, phone).
		Limit(1).
		Scan(context.Background())

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func CreateUser(user *userModels.User) error {
	_, err := db.DB.NewInsert().
		Model(user).
		Exec(context.Background())
	return err
}

func CreateOrder(order *orderModels.Order) error {
	_, err := db.DB.NewInsert().
		Model(order).
		Exec(context.Background())
	return err
}

func UpdateOrder(order *orderModels.Order) error {
	_, err := db.DB.NewUpdate().
		Model(order).
		Where("id = ?", order.ID).
		Exec(context.Background())

	return err
}
func CreatePayment(orderId string, publisherOrderId string, issuerCode string, resultCode string) (*orderModels.Order, error) {
	ctx := context.Background()

	var order orderModels.Order
	var pricing models.Pricing

	err := db.DB.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		// 1. Ambil order
		if err := tx.NewSelect().
			Model(&order).
			Where("id = ?", orderId).
			Scan(ctx); err != nil {
			return err
		}

		// 2. Cegah double payment
		if order.Status == "PAID" {
			return fmt.Errorf("order already paid")
		}

		// 3. Ambil pricing
		if err := tx.NewSelect().
			Model(&pricing).
			Where("id = ?", order.PricingID).
			Scan(ctx); err != nil {
			return err
		}

		now := time.Now()
		expiredAt := now.AddDate(0, 1, 0)

		// 4. Update order → PAID
		if resultCode == "00" {
			order.Status = "PAID"
			order.PublisherOrderID = publisherOrderId
			order.IssuerCode = issuerCode
			order.UpdatedAt = &now

			if _, err := tx.NewUpdate().
				Model(&order).
				Column("status",
					"publisher_order_id",
					"issuer_code", "updated_at").
				Where("id = ?", order.ID).
				Exec(ctx); err != nil {
				return err
			}

			// 5. Create user agent
			userAgent := &userAgentModels.UserAgent{
				ID:       uuid.NewString(),
				UsersID:  order.UsersID,
				AgentsID: order.AgentsID,
				Active:   true,
				Expired:  expiredAt,
				Tokens:   int64(pricing.TokenMonthly),
				// ExpiredAt: ,
			}

			if _, err := tx.NewInsert().
				Model(userAgent).
				Exec(ctx); err != nil {
				return err
			}
		} else {
			order.Status = "FAILED"
			order.UpdatedAt = &now

			if _, err := tx.NewUpdate().
				Model(&order).
				Column("status", "updated_at").
				Where("id = ?", order.ID).
				Exec(ctx); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &order, nil
}
