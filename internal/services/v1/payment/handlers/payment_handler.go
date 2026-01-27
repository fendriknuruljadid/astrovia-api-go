package handlers

import (
	"app/internal/packages/response"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type PaymentRequest struct {
	MerchantCode string `json:"merchantcode"`
	Amount       int    `json:"amount"`
	Datetime     string `json:"datetime"`
	Signature    string `json:"signature"`
}

type PaymentResponse struct {
	StatusCode int                    `json:"status_code"`
	Message    string                 `json:"message"`
	Data       map[string]interface{} `json:"data,omitempty"`
}

type DuitkuPaymentMethodResponse struct {
	PaymentFee      []PaymentFee `json:"paymentFee"`
	ResponseCode    string       `json:"responseCode"`
	ResponseMessage string       `json:"responseMessage"`
}

type PaymentFee struct {
	PaymentImage  string `json:"paymentImage"`
	PaymentMethod string `json:"paymentMethod"`
	PaymentName   string `json:"paymentName"`
	TotalFee      string `json:"totalFee"`
}

// Helper function untuk request ke Duitku
func requestDuitkuPaymentMethod(
	merchantCode, apiKey string,
	amount int,
) ([]PaymentFee, error) {
	// Generate datetime
	datetime := time.Now().Format("2006-01-02 15:04:05")

	// Generate signature
	signatureData := fmt.Sprintf("%s%d%s%s", merchantCode, amount, datetime, apiKey)
	hash := sha256.Sum256([]byte(signatureData))
	signature := hex.EncodeToString(hash[:])

	payload := PaymentRequest{
		MerchantCode: merchantCode,
		Amount:       amount,
		Datetime:     datetime,
		Signature:    signature,
	}

	payloadBytes, _ := json.Marshal(payload)

	url := "https://sandbox.duitku.com/webapi/api/merchant/paymentmethod/getpaymentmethod"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("RAW DUITKU RESPONSE:")
	fmt.Println(string(body))
	if resp.StatusCode == 200 {
		var result DuitkuPaymentMethodResponse

		if err := json.Unmarshal(body, &result); err != nil {
			return nil, err
		}

		if result.ResponseCode != "00" {
			return nil, fmt.Errorf("duitku error: %s", result.ResponseMessage)
		}

		if result.PaymentFee == nil {
			return []PaymentFee{}, nil
		}

		return result.PaymentFee, nil
	} else {
		var errResp map[string]interface{}
		_ = json.Unmarshal(body, &errResp)
		msg := ""
		if m, ok := errResp["Message"].(string); ok {
			msg = m
		}
		return nil, fmt.Errorf("duitku error %d: %s", resp.StatusCode, msg)
	}
}

func GetPaymentMethod(c *gin.Context) {
	merchantCode := "DS27467"
	apiKey := "c2376b872e008811acdc47158992f2a2"
	amount := 89000

	data, err := requestDuitkuPaymentMethod(merchantCode, apiKey, amount)
	if err != nil {
		c.JSON(http.StatusBadGateway, response.Error(
			502,
			err.Error(),
			nil,
		))
		return
	}

	c.JSON(http.StatusOK, response.Success(
		200, "success", data,
	))
}
