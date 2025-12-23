package routes

import "github.com/gofiber/fiber/v2"

// =================== Dummy Agent model untuk Swagger ===================

type Agent struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// =================== Dummy Handlers untuk Swagger ===================

// @Summary Get all agent
// @Description Get list of agent
// @Tags Agents
// @Security BearerAuth
// @Security X-Signature
// @Security X-Timestamp
// @Produce json
// @Success 200 {array} routes.Agent
// @Router /v1/agents [get]
func GetAgentsHandler(c *fiber.Ctx) error { return nil }

// @Summary Get all agent
// @Description Get list of agent
// @Tags Agents
// @Security BearerAuth
// @Security X-Signature
// @Security X-Timestamp
// @Produce json
// @Success 200 {array} routes.Agent
// @Router /v1/agents/public [get]
func GetAgentsPublicHandler(c *fiber.Ctx) error { return nil }

// @Summary Get agent by ID
// @Description Get single agent
// @Tags Agents
// @Security BearerAuth
// @Security X-Signature
// @Security X-Timestamp
// @Produce json
// @Param id path string true "Agent ID"
// @Success 200 {object} routes.Agent
// @Router /v1/agents/{id} [get]
func GetAgentHandler(c *fiber.Ctx) error { return nil }

// @Summary Create agent
// @Description Create new agent
// @Tags Agents
// @Security X-Signature
// @Security X-Timestamp
// @Accept json
// @Produce json
// @Param agent body routes.Agent true "Agent info"
// @Success 201 {object} routes.Agent
// @Router /v1/agents [post]
func CreateAgentHandler(c *fiber.Ctx) error { return nil }

// @Summary Update agent
// @Description Update agent by ID
// @Tags Agents
// @Security BearerAuth
// @Security X-Signature
// @Security X-Timestamp
// @Accept json
// @Produce json
// @Param id path string true "Agent ID"
// @Param agent body routes.Agent true "Agent info"
// @Success 200 {object} routes.Agent
// @Router /v1/agents/{id} [put]
func UpdateAgentHandler(c *fiber.Ctx) error { return nil }

// @Summary Delete agent
// @Description Delete agent by ID
// @Tags Agents
// @Security BearerAuth
// @Security X-Signature
// @Security X-Timestamp
// @Param id path string true "Agent ID"
// @Success 200 {object} map[string]bool
// @Router /v1/agents/{id} [delete]
func DeleteAgentHandler(c *fiber.Ctx) error { return nil }
