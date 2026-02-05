package handlers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"os"

	"app/internal/services/v1/payment/repository"

	"github.com/gin-gonic/gin"
)

func CallbackDuitku(c *gin.Context) {
	if err := c.Request.ParseForm(); err != nil {
		c.JSON(400, gin.H{"message": "invalid form"})
		return
	}

	fmt.Println("FORM CALLBACK:")
	fmt.Println(c.Request.Form)

	apiKey := os.Getenv("PG_API_KEY")
	merchantCode := c.PostForm("merchantCode")
	amount := c.PostForm("amount")
	merchantOrderId := c.PostForm("merchantOrderId")
	productDetail := c.PostForm("productDetail")
	additionalParam := c.PostForm("additionalParam")
	paymentMethod := c.PostForm("paymentCode")
	resultCode := c.PostForm("resultCode")
	merchantUserId := c.PostForm("merchantUserId")
	reference := c.PostForm("reference")
	signature := c.PostForm("signature")
	publisherOrderId := c.PostForm("publisherOrderId")
	spUserHash := c.PostForm("spUserHash")
	settlementDate := c.PostForm("settlementDate")
	issuerCode := c.PostForm("issuerCode")

	_ = productDetail
	_ = additionalParam
	_ = paymentMethod
	_ = resultCode
	_ = merchantUserId
	_ = reference
	_ = publisherOrderId
	_ = spUserHash
	_ = settlementDate
	_ = issuerCode

	// Validasi parameter wajib
	if merchantCode == "" || amount == "" || merchantOrderId == "" || signature == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Bad Parameter",
		})
		return
	}

	params := merchantCode + amount + merchantOrderId + apiKey
	hash := md5.Sum([]byte(params))
	calcSignature := hex.EncodeToString(hash[:])

	if signature != calcSignature {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "Bad Signature",
		})
		return
	}
	repository.CreatePayment(merchantOrderId, publisherOrderId, issuerCode, resultCode)
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Callback validated",
	})

}
