package repository

import (
	"app/internal/packages/db"
	"app/internal/services/v1/auth/models"
	"context"
	"fmt"
)

func CreateRefreshTokens(auth *models.RefreshTokens) error {
	ctx := context.Background()
	_, err := db.DB.NewInsert().Model(auth).Exec(ctx)
	if err != nil {
		fmt.Println("CreateRefreshTokens failed:", err)
	}
	return err
}

func GetRefreshTokensByID(id string) (*models.RefreshTokens, error) {
	ctx := context.Background()
	auth := new(models.RefreshTokens)
	err := db.DB.NewSelect().Model(auth).Where("id = ?", id).Scan(ctx)
	if err != nil {
		fmt.Println("GetRefreshTokensByID failed:", err)
	}
	return auth, err
}

func UpdateRefreshTokens(auth *models.RefreshTokens) error {
	ctx := context.Background()
	_, err := db.DB.NewUpdate().Model(auth).WherePK().Exec(ctx)
	if err != nil {
		fmt.Println("UpdateRefreshTokens failed:", err)
	}
	return err
}

func DeleteRefreshTokens(id string) error {
	ctx := context.Background()
	auth := &models.RefreshTokens{ID: id}
	_, err := db.DB.NewDelete().Model(auth).WherePK().Exec(ctx)
	if err != nil {
		fmt.Println("DeleteRefreshTokens failed:", err)
	}
	return err
}
