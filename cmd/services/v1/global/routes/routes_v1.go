package routesV1

import (
	agent "app/internal/services/v1/agent/handlers"
	payment "app/internal/services/v1/payment/handlers"
	pricing "app/internal/services/v1/pricing/handlers"
	userAgent "app/internal/services/v1/user-agent/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterPublicRoutes(r *gin.RouterGroup) {
	r.GET("/agents/public", agent.GetAgentsPublic)
	r.GET("/agents/public/:id", agent.GetAgentPublicByID)
	r.GET("/user-agents/public", userAgent.GetUserAgents)
	r.GET("/payment/payment-method/:id", payment.GetPaymentMethod)
	r.POST("/payment/order-public", payment.CreateOrderPublic)
	r.POST("/callback/duitku", payment.CallbackDuitku)
}

func RegisterProtectedRoutes(r *gin.RouterGroup) {
	r.GET("/agents", agent.GetAgents)
	r.GET("/agents/:id", agent.GetAgentByID)
	r.PUT("/agents/:id", agent.UpdateAgent)
	r.DELETE("/agents/:id", agent.DeleteAgent)
	r.POST("/agents", agent.CreateAgent)

	//pricing
	r.GET("/pricing", pricing.GetPricings)
	r.GET("/pricing/:id", pricing.GetPricingByID)
	r.PUT("/pricing/:id", pricing.UpdatePricing)
	r.DELETE("/pricing/:id", pricing.DeletePricing)
	r.POST("/pricing", pricing.CreatePricing)

	//user-agent
	r.GET("/user-agents", userAgent.GetUserAgents)
	r.GET("/user-agents/:id", userAgent.GetUserAgentByID)
	r.PUT("/user-agents/:id", userAgent.UpdateUserAgent)
	r.DELETE("/user-agents/:id", userAgent.DeleteUserAgent)
	r.POST("/user-agents", userAgent.CreateUserAgent)
}
