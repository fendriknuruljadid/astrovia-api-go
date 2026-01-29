package repository

import (
	"app/internal/packages/db"
	orderModels "app/internal/services/v1/payment/models"
	"app/internal/services/v1/pricing/models"
	userModels "app/internal/services/v1/user/models"
	"context"
	"fmt"
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
