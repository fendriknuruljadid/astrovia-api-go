package repository

import (
	database "app/internal/packages/db"
	"context"
)

func GetAgentsWithUserAccess(userID string) ([]AgentWithAccessRow, error) {
	var rows []AgentWithAccessRow

	err := database.DB.NewSelect().
		Table("agents").
		Column("agents.id", "agents.name", "agents.description", "agents.logo", "agents.url").
		ColumnExpr("ua.active").
		ColumnExpr("ua.expired").
		Join(`
			LEFT JOIN user_agents ua
			ON ua.agents_id = agents.id
			AND ua.users_id = ?
		`, userID).
		Scan(context.Background(), &rows)

	return rows, err
}
