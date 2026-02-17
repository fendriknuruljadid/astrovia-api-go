package handlers

import (
	"app/internal/packages/response"
	"crypto/md5"
	"fmt"
	"time"

	"os"

	"app/internal/services/v1/payment/dto"
	orderModels "app/internal/services/v1/payment/models"
	"app/internal/services/v1/payment/repository"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	merchantCode := os.Getenv("MERCHANT_CODE")
	apiKey := os.Getenv("PG_API_KEY")
	userID := c.GetString("user_id")
	var req dto.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, response.Error(400, "invalid request", err.Error()))
		return
	}
	pricing, err := repository.GetPricingByID(req.PricingId)
	if err != nil {
		c.JSON(404, response.Error(404, "Package not found", err.Error()))
		return
	}
	user, err := repository.GetUserByID(userID)
	if err != nil {
		c.JSON(404, response.Error(404, "user not found", err.Error()))
		return
	}
	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.Phone = req.Phone

	if err := repository.UpdateUser(user); err != nil {
		c.JSON(500, response.Error(500, "failed update user", err.Error()))
		return
	}

	order := &orderModels.Order{
		UsersID:       user.ID,
		PricingID:     pricing.ID,
		AgentsID:      pricing.AgentsID,
		PricingName:   pricing.Name,
		AgentName:     pricing.Agent.Name,
		Amount:        float64(pricing.MonthlyPrice),
		PaymentMethod: req.PaymentMethod,
		Status:        "PENDING",
		ExpiryPeriod:  10,
	}

	if err := repository.CreateOrder(order); err != nil {
		c.JSON(500, response.Error(500, "failed create order", err.Error()))
		return
	}

	// Signature: md5(merchantCode + merchantOrderId + amount + apiKey)
	signatureRaw := fmt.Sprintf("%s%s%d%s",
		merchantCode,
		order.ID,
		pricing.MonthlyPrice,
		apiKey,
	)
	signature := fmt.Sprintf("%x", md5.Sum([]byte(signatureRaw)))

	reqData := dto.DuitkuInquiryRequest{
		MerchantCode:    merchantCode,
		PaymentAmount:   pricing.MonthlyPrice,
		PaymentMethod:   req.PaymentMethod, // atau dari request user
		MerchantOrderID: order.ID,
		ProductDetails:  pricing.Name,
		CustomerVaName:  fmt.Sprintf("%s %s", user.FirstName, user.LastName),
		Email:           user.Email,
		CustomerDetail: dto.DuitkuCustomerDetail{
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			Email:       user.Email,
			PhoneNumber: user.Phone,
		},
		CallbackURL:  os.Getenv("DUITKU_CALLBACK_URL"),
		ReturnURL:    os.Getenv("DUITKU_RETURN_URL"),
		Signature:    signature,
		ExpiryPeriod: 10,
	}

	result, err := requestDuitkuInquiry(reqData)
	if err != nil {
		c.JSON(502, response.Error(502, err.Error(), nil))
		return
	}

	order.Reference = result.Reference
	order.VANumber = result.VANumber
	order.PaymentURL = result.PaymentURL
	order.QRCode = result.QRString
	expiredAt := time.Now().Add(10 * time.Minute)
	order.ExpiredAt = &expiredAt
	_ = repository.UpdateOrder(order)

	c.JSON(200, response.Success(200, "success", order))
}
