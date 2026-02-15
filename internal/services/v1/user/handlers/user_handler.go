package handlers

import (
	"app/internal/packages/errors"
	"app/internal/packages/mailer"
	"app/internal/packages/response"
	"app/internal/packages/utils"
	"app/internal/services/v1/user/dto"
	"app/internal/services/v1/user/models"
	"app/internal/services/v1/user/repository"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context) {
	var req dto.CreateDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		// Deteksi apakah ini validation error dari validator bawaan Gin
		if ve, ok := err.(validator.ValidationErrors); ok {
			c.JSON(402, response.Error(402, "validation failed", utils.ValidationErrors(ve)))
			return
		}

		// Kalau bukan validasi, baru fallback ke invalid JSON
		c.JSON(400, response.Error(400, "invalid request payload", err.Error()))
		return
	}
	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := repository.CreateUser(&user); err != nil {
		c.Error(errors.NewInternal(utils.ParseDBError(err)))
		return
	}

	res := dto.ToResponseDTO(&user)

	c.JSON(201, response.Success(201, "success", res))
}

func CheckUser(c *gin.Context) {
	var req dto.CheckDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		// Deteksi apakah ini validation error dari validator bawaan Gin
		if ve, ok := err.(validator.ValidationErrors); ok {
			c.JSON(402, response.Error(402, "validation failed", utils.ValidationErrors(ve)))
			return
		}

		// Kalau bukan validasi, baru fallback ke invalid JSON
		c.JSON(400, response.Error(400, "invalid request payload", err.Error()))
		return
	}
	user, err := repository.CheckOrCreateUser(req.Email)
	if err != nil {
		c.JSON(500, response.Error(500, "Internal Server Error", err.Error()))
		return
	}

	c.JSON(200, response.Success(200, "success", gin.H{
		"has_password":   user.Password != nil,
		"is_verified":    user.IsVerified,
		"otp_expired_at": user.OTPExpiredAt,
	}))
}

func CreatePassword(c *gin.Context) {
	var req dto.CreatePasswordDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		// Deteksi apakah ini validation error dari validator bawaan Gin
		if ve, ok := err.(validator.ValidationErrors); ok {
			c.JSON(402, response.Error(402, "validation failed", utils.ValidationErrors(ve)))
			return
		}

		// Kalau bukan validasi, baru fallback ke invalid JSON
		c.JSON(400, response.Error(400, "invalid request payload", err.Error()))
		return
	}

	user, err := repository.GetUserByEmail(req.Email)
	if err != nil {
		c.JSON(404, response.Error(404, "User Not Found", err.Error()))
		return
	}

	if user.IsVerified {
		c.JSON(400, response.Error(400, "user already verified", nil))
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		c.JSON(500, response.Error(500, "hash failed", err.Error()))
		return
	}

	hashStr := string(hashedPassword)
	user.Password = &hashStr

	// Generate OTP
	otp := utils.GenerateOTP(6)

	hashedOTP, err := bcrypt.GenerateFromPassword(
		[]byte(otp),
		bcrypt.DefaultCost,
	)

	if err != nil {
		c.JSON(500, response.Error(500, "otp hash failed", err.Error()))
		return
	}

	otpHashStr := string(hashedOTP)
	expired := time.Now().Add(10 * time.Minute)

	user.OTPHash = &otpHashStr
	user.OTPExpiredAt = &expired

	// Save to DB
	if err := repository.UpdateUser(user); err != nil {
		c.JSON(500, response.Error(500, "update failed", err.Error()))
		return
	}

	// Kirim email OTP
	go mailer.SendOTPEmail(user.Email, otp)

	// c.JSON(200, gin.H{
	// 	"message": "password created, verification email sent",
	// })

	c.JSON(200, response.Success(200, "success", gin.H{
		"email":          user.Email,
		"otp_expired_at": user.OTPExpiredAt,
		"is_verified":    user.IsVerified,
	}))
}

func VerifyOTP(c *gin.Context) {
	var req dto.VerifyOTPDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			c.JSON(402, response.Error(402, "validation failed", utils.ValidationErrors(ve)))
			return
		}

		c.JSON(400, response.Error(400, "invalid request payload", nil))
		return
	}

	user, err := repository.GetUserByEmail(req.Email)
	if err != nil {
		c.JSON(404, response.Error(404, "user not found", nil))
		return
	}

	// Sudah verified?
	if user.IsVerified {
		c.JSON(400, response.Error(400, "user already verified", nil))
		return
	}

	// Tidak ada OTP
	if user.OTPHash == nil || user.OTPExpiredAt == nil {
		c.JSON(400, response.Error(400, "otp not found", nil))
		return
	}

	// OTP expired?
	if time.Now().After(*user.OTPExpiredAt) {
		c.JSON(400, response.Error(400, "otp expired", nil))
		return
	}
	if strings.TrimSpace(req.OTP) == "" {
		c.JSON(400, response.Error(400, "otp required", nil))
		return
	}

	// Compare OTP
	err = bcrypt.CompareHashAndPassword(
		[]byte(*user.OTPHash),
		[]byte(req.OTP),
	)

	if err != nil {
		user.OTPAttempt += 1
		if user.OTPAttempt >= 5 {
			user.OTPHash = nil
			user.OTPExpiredAt = nil
			user.OTPAttempt = 0

			_ = repository.ClearOTP(user)

			c.JSON(400, response.Error(400,
				"too many otp attempts, otp blocked, please request new otp",
				nil,
			))
			return
		}
		_ = repository.UpdateUser(user)

		c.JSON(400, response.Error(400, "invalid otp", gin.H{
			"attempt_left": 5 - user.OTPAttempt,
		}))
		return
	}

	// SUCCESS
	user.IsVerified = true
	user.OTPHash = nil
	user.OTPExpiredAt = nil
	user.OTPAttempt = 0

	if err := repository.UpdateUser(user); err != nil {
		c.JSON(500, response.Error(500, "verification failed", nil))
		return
	}

	res := dto.VerifyOTPResponse{
		Email:      user.Email,
		IsVerified: user.IsVerified,
	}

	c.JSON(200, response.Success(200, "verification success", res))
}

func ResendOTP(c *gin.Context) {
	var req dto.ResendOTPDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			c.JSON(402, response.Error(402, "validation failed", utils.ValidationErrors(ve)))
			return
		}
		c.JSON(400, response.Error(400, "invalid request payload", nil))
		return
	}

	user, err := repository.GetUserByEmail(req.Email)
	if err != nil {
		c.JSON(404, response.Error(404, "user not found", nil))
		return
	}

	// Sudah verified?
	if user.IsVerified {
		c.JSON(400, response.Error(400, "user already verified", nil))
		return
	}

	// Cooldown check (60 detik)
	if user.LastOTPSentAt != nil {
		elapsed := time.Since(*user.LastOTPSentAt)
		if elapsed < 60*time.Second {
			remaining := 60 - int(elapsed.Seconds())

			c.JSON(429, response.Error(429,
				"please wait before requesting new otp",
				gin.H{
					"retry_after_seconds": remaining,
				},
			))
			return
		}
	}

	//Generate OTP baru
	otp := utils.GenerateOTP(6)

	hashedOTP, err := bcrypt.GenerateFromPassword(
		[]byte(otp),
		bcrypt.DefaultCost,
	)
	if err != nil {
		c.JSON(500, response.Error(500, "otp hash failed", nil))
		return
	}

	hashStr := string(hashedOTP)
	expired := time.Now().Add(10 * time.Minute)
	now := time.Now()

	user.OTPHash = &hashStr
	user.OTPExpiredAt = &expired
	user.OTPAttempt = 0
	user.LastOTPSentAt = &now

	if err := repository.UpdateUser(user); err != nil {
		c.JSON(500, response.Error(500, "update failed", nil))
		return
	}

	go mailer.SendOTPEmail(user.Email, otp)

	// res := dto.ResendOTPResponse{
	// 	Email:      user.Email,
	// 	IsVerified: user.IsVerified,
	// }
	c.JSON(200, response.Success(200, "success", gin.H{
		"email":          user.Email,
		"otp_expired_at": user.OTPExpiredAt,
		"is_verified":    user.IsVerified,
	}))

	// c.JSON(200, response.Success(200, "new otp sent", res))
}

func GetUsers(c *gin.Context) {
	users, err := repository.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(500, "Server internal error", err.Error()))
		return
	}

	res := dto.ToResponseDTOs(users)
	c.JSON(200, response.Success(200, "success", res))
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := repository.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	res := dto.ToResponseDTO(user)
	c.JSON(200, response.Success(200, "success", res))
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var req dto.UpdateDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		// Deteksi apakah ini validation error dari validator bawaan Gin
		if ve, ok := err.(validator.ValidationErrors); ok {
			c.JSON(402, response.Error(402, "validation failed", utils.ValidationErrors(ve)))
			return
		}

		// Kalau bukan validasi, baru fallback ke invalid JSON
		c.JSON(400, response.Error(400, "invalid request payload", err.Error()))
		return
	}

	user, err := repository.GetUserByID(id)
	if err != nil {
		c.Error(errors.NewNotFound("user not found"))
		return
	}

	// Apply updates safely
	if req.Name != nil {
		user.Name = *req.Name
	}
	if req.Email != nil {
		user.Email = *req.Email
	}
	if req.Password != nil {
		user.Password = req.Password
	}

	if err := repository.UpdateUser(user); err != nil {
		c.Error(errors.NewInternal(utils.ParseDBError(err)))
		return
	}

	c.JSON(200, response.Success(200, "updated", dto.ToResponseDTO(user)))
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	// Cek dulu apakah user ada
	user, err := repository.GetUserByID(id)
	if err != nil || user == nil {
		c.Error(errors.NewNotFound("user not found"))
		return
	}

	// Jika ada, lanjut hapus
	if err := repository.DeleteUser(id); err != nil {
		c.Error(errors.NewInternal(utils.ParseDBError(err)))
		return
	}

	c.JSON(200, response.Success(200, "deleted successfully", gin.H{
		"deleted": true,
		"id":      id,
	}))
}
