package routes

import "github.com/gofiber/fiber/v2"

// =================== Dummy Handlers untuk Swagger ===================
type CreateOrderRequest struct {
	PricingId     string `json:"pricing_id" binding:"required"`
	PaymentMethod string `json:"payment_method" binding:"required"`
	FirstName     string `json:"first_name" binding:"required"`
	LastName      string `json:"last_name" binding:"required"`
	Phone         string `json:"phone" binding:"required"`
	Email         string `json:"email" binding:"required,email"`
}

// @Summary Create order for subscription
// @Description Create Order
// @Tags Payments
// @Security X-Signature
// @Security X-Timestamp
// @Produce json
// @Accept json
// @Param order body routes.CreateOrderRequest true "Order info"
// @Success 201 {object} routes.CreateOrderRequest
// @Router /v1/payment/order-public [post]
func CreateOrderMethodHandler(c *fiber.Ctx) error { return nil }
