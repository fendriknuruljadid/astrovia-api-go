package handlers

import (
	"app/internal/packages/response"
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"os"

	"app/internal/services/v1/payment/dto"
	orderModels "app/internal/services/v1/payment/models"
	"app/internal/services/v1/payment/repository"
	userModels "app/internal/services/v1/user/models"

	"github.com/gin-gonic/gin"
)

func requestDuitkuInquiry(reqData dto.DuitkuInquiryRequest) (*dto.DuitkuInquiryResponse, error) {
	url := pasymentGatewayURL + "v2/inquiry"
	// sandbox: https://sandbox.duitku.com/webapi/api/merchant/v2/inquiry

	payload, _ := json.Marshal(reqData)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("RAW DUITKU RESPONSE:")
	fmt.Println(string(body))

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("duitku http error %d", resp.StatusCode)
	}

	var result dto.DuitkuInquiryResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	if result.StatusCode != "00" {
		return nil, fmt.Errorf("duitku error: %s", result.StatusMessage)
	}

	return &result, nil
}

func CreateOrderPublic(c *gin.Context) {
	merchantCode := os.Getenv("MERCHANT_CODE")
	apiKey := os.Getenv("PG_API_KEY")

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

	user, err := repository.GetUserByEmailOrPhone(req.Email, req.Phone)
	if err != nil {
		user = &userModels.User{
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Email:     req.Email,
			Phone:     req.Phone,
		}
		if err := repository.CreateUser(user); err != nil {
			c.JSON(500, response.Error(500, "failed create user", err.Error()))
			return
		}
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

	// loc, err := time.LoadLocation("Asia/Jakarta")
	// if err != nil {
	// 	// fallback aman
	// 	loc = time.FixedZone("GMT+7", 7*60*60)
	// }
	// expiredLocal := order.ExpiredAt.In(loc)
	// mailer.SendInvoiceEmailAsync(req.Email, map[string]any{
	// 	"Name":    req.FirstName,
	// 	"Invoice": order.InvoiceNumber,
	// 	"Amount":  order.Amount,
	// 	"Status":  "PENDING",
	// 	"Expired": expiredLocal.Format("02 Jan 2006 15:04"),
	// })

	c.JSON(200, response.Success(200, "success", order))
}
