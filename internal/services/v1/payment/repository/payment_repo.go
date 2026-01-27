package repository

import (
	"app/internal/packages/db"
	"app/internal/services/v1/pricing/models"
	"context"
	"fmt"
)

func CreatePricing(pricing *models.Pricing) error {
	ctx := context.Background()
	_, err := db.DB.NewInsert().Model(pricing).Exec(ctx)
	if err != nil {
		fmt.Println("CreatePricing failed:", err)
	}
	return err
}

func GetPricings() ([]models.Pricing, error) {
	ctx := context.Background()
	var pricings []models.Pricing
	err := db.DB.NewSelect().Model(&pricings).Scan(ctx)
	if err != nil {
		fmt.Println("GetPricings failed:", err)
	}
	return pricings, err
}

func GetPricingByID(id string) (*models.Pricing, error) {
	ctx := context.Background()
	pricing := new(models.Pricing)
	err := db.DB.NewSelect().Model(pricing).Where("id = ?", id).Scan(ctx)
	if err != nil {
		fmt.Println("GetPricingByID failed:", err)
	}
	return pricing, err
}

func UpdatePricing(pricing *models.Pricing) error {
	ctx := context.Background()
	_, err := db.DB.NewUpdate().Model(pricing).WherePK().Exec(ctx)
	if err != nil {
		fmt.Println("UpdatePricing failed:", err)
	}
	return err
}

func DeletePricing(id string) error {
	ctx := context.Background()
	pricing := &models.Pricing{ID: id}
	_, err := db.DB.NewDelete().Model(pricing).WherePK().Exec(ctx)
	if err != nil {
		fmt.Println("DeletePricing failed:", err)
	}
	return err
}
