package repository

import (
	"app/internal/packages/db"
	orderModels "app/internal/services/v1/payment/models"
	"app/internal/services/v1/pricing/models"
	userAgentModels "app/internal/services/v1/user-agent/models"
	userModels "app/internal/services/v1/user/models"
	"database/sql"
	"errors"

	"time"

	"context"
	"fmt"

	"strings"

	"app/internal/packages/utils"

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

func GetUserByID(id string) (*userModels.User, error) {
	var user userModels.User

	err := db.DB.NewSelect().
		Model(&user).
		Where("id = ? ", id).
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

func UpdateUser(user *userModels.User) error {
	_, err := db.DB.NewUpdate().
		Model(user).
		Where("id = ?", user.ID).
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
			existingUserAgent := new(userAgentModels.UserAgent)
			err := tx.NewSelect().
				Model(existingUserAgent).
				Where("users_id = ?", order.UsersID).
				Where("agents_id = ?", order.AgentsID).
				Limit(1).
				Scan(ctx)
			now := time.Now()
			oneMonth := now.AddDate(0, 1, 0)
			if err == nil {
				var newExpired time.Time
				if existingUserAgent.Expired.After(now) {
					newExpired = existingUserAgent.Expired.AddDate(0, 1, 0)
				} else {
					newExpired = oneMonth
				}

				existingUserAgent.Expired = newExpired
				existingUserAgent.Active = true

				if _, err := tx.NewUpdate().
					Model(existingUserAgent).
					Column("expired", "active").
					WherePK().
					Exec(ctx); err != nil {
					return err
				}

			} else if errors.Is(err, sql.ErrNoRows) {
				userAgent := &userAgentModels.UserAgent{
					ID:       "usr-agn-" + strings.ToLower(utils.NewULID()),
					UsersID:  order.UsersID,
					AgentsID: order.AgentsID,
					Active:   true,
					Expired:  oneMonth,
					// Tokens:   int64(pricing.TokenMonthly),
				}

				if _, err := tx.NewInsert().
					Model(userAgent).
					Exec(ctx); err != nil {
					return err
				}

			} else {
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
