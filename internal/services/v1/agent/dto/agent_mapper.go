package dto

import (
	"app/internal/services/v1/agent/repository"
	"time"
)

func ToAgentResponseDTOs(rows []repository.AgentWithAccessRow) []AgentResponseDTO {
	res := make([]AgentResponseDTO, 0, len(rows))

	now := time.Now()

	for _, r := range rows {
		hasAccess := false
		expired := false

		if r.Active != nil && *r.Active {
			if r.ExpiresAt == nil || r.ExpiresAt.Before(now) {
				expired = true
			}
			hasAccess = true
		}

		res = append(res, AgentResponseDTO{
			ID:          r.ID,
			Name:        r.Name,
			Description: r.Description,
			Logo:        r.Logo,
			URL:         r.URL,
			HasAccess:   hasAccess,
			Expired:     expired,
			ExpiredAt:   r.ExpiresAt,
		})
	}

	return res
}
