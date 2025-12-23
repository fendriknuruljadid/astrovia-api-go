package repository

import (
	"app/internal/packages/db"
	"app/internal/services/v1/user-agent/models"
	"context"
	"fmt"
)

func CreateUserAgent(user *models.UserAgent) error {
	ctx := context.Background()
	_, err := db.DB.NewInsert().Model(user).Exec(ctx)
	if err != nil {
		fmt.Println("CreateUserAgent failed:", err)
	}
	return err
}

func GetUserAgents() ([]models.UserAgent, error) {
	ctx := context.Background()
	var users []models.UserAgent
	err := db.DB.NewSelect().Model(&users).Scan(ctx)
	if err != nil {
		fmt.Println("GetUserAgents failed:", err)
	}
	return users, err
}

func GetUserAgentByID(id string) (*models.UserAgent, error) {
	ctx := context.Background()
	user := new(models.UserAgent)
	err := db.DB.NewSelect().Model(user).Where("id = ?", id).Scan(ctx)
	if err != nil {
		fmt.Println("GetUserAgentByID failed:", err)
	}
	return user, err
}

func UpdateUserAgent(user *models.UserAgent) error {
	ctx := context.Background()
	_, err := db.DB.NewUpdate().Model(user).WherePK().Exec(ctx)
	if err != nil {
		fmt.Println("UpdateUserAgent failed:", err)
	}
	return err
}

func DeleteUserAgent(id string) error {
	ctx := context.Background()
	user := &models.UserAgent{ID: id}
	_, err := db.DB.NewDelete().Model(user).WherePK().Exec(ctx)
	if err != nil {
		fmt.Println("DeleteUserAgent failed:", err)
	}
	return err
}
