package routes

import "github.com/gofiber/fiber/v2"

// =================== Dummy Handlers untuk Swagger ===================
type Payment struct {
	MerchantCode string `json:"merchantcode"`
	Amount       int    `json:"amount"`
	Datetime     string `json:"datetime"`
	Signature    string `json:"signature"`
}

// @Summary Get all payment method
// @Description Get list of payment method
// @Tags Payments
// @Security X-Signature
// @Security X-Timestamp
// @Produce json
// @Success 200 {array} routes.Payment
// @Param id path string true "Pricing ID"
// @Router /v1/payment/payment-method/{id} [get]
func GetPaymentMethodHandler(c *fiber.Ctx) error { return nil }
