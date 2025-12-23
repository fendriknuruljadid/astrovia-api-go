package repository

import (
	"app/internal/packages/db"
	"app/internal/services/v1/agent/models"
	"context"
	"fmt"
)

func CreateAgent(agent *models.Agent) error {
	ctx := context.Background()
	_, err := db.DB.NewInsert().Model(agent).Exec(ctx)
	if err != nil {
		fmt.Println("CreateAgent failed:", err)
	}
	return err
}

func GetAgents() ([]models.Agent, error) {
	ctx := context.Background()
	var agents []models.Agent
	err := db.DB.NewSelect().Model(&agents).Scan(ctx)
	if err != nil {
		fmt.Println("GetAgents failed:", err)
	}
	return agents, err
}

func GetAgentByID(id string) (*models.Agent, error) {
	ctx := context.Background()
	agent := new(models.Agent)
	err := db.DB.NewSelect().Model(agent).Where("id = ?", id).Scan(ctx)
	if err != nil {
		fmt.Println("GetAgentByID failed:", err)
	}
	return agent, err
}

func UpdateAgent(agent *models.Agent) error {
	ctx := context.Background()
	_, err := db.DB.NewUpdate().Model(agent).WherePK().Exec(ctx)
	if err != nil {
		fmt.Println("UpdateAgent failed:", err)
	}
	return err
}

func DeleteAgent(id string) error {
	ctx := context.Background()
	agent := &models.Agent{ID: id}
	_, err := db.DB.NewDelete().Model(agent).WherePK().Exec(ctx)
	if err != nil {
		fmt.Println("DeleteAgent failed:", err)
	}
	return err
}
