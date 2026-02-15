package auth

import (
	"app/internal/middlewares"
	"app/internal/packages/db"
	auth_models "app/internal/services/v1/auth/models"
	"app/internal/services/v1/user/models"
	"net/http"
	"time"

	"encoding/json"
	"fmt"
	"io"

	"app/internal/packages/response"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// type AuthRequest struct {
// 	Provider   string `json:"provider"`   // hanya: google
// 	OAuthToken string `json:"oauthToken"` // access_token dari Google Sign-In
// 	Email      string `json:"email"`      // optional (Google bisa kirim)
// }

type AuthRequest struct {
	Provider   string `json:"provider"` // google | local
	OAuthToken string `json:"oauthToken,omitempty"`
	Email      string `json:"email"`
	Password   string `json:"password,omitempty"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

// Struktur data Google token validation response
type GoogleTokenInfo struct {
	Email         string `json:"email"`
	EmailVerified string `json:"email_verified"`
	Audience      string `json:"aud"`
}

// func Auth(c *gin.Context) {
// 	var req AuthRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
// 		return
// 	}

// 	// Provider wajib Google
// 	if req.Provider != "google" {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Only Google login allowed"})
// 		return
// 	}

// 	if req.OAuthToken == "" {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "OAuth token required"})
// 		return
// 	}

// 	// === Validasi token Google ===
// 	googleUser, err := verifyGoogleToken(req.OAuthToken)
// 	if err != nil || googleUser.Email == "" {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Google OAuth token"})
// 		return
// 	}

// 	// === Cek user dari database ===
// 	var user models.User
// 	err = db.DB.NewSelect().Model(&user).Where("email = ?", googleUser.Email).Scan(c)

// 	if err != nil { // user belum ada → auto register
// 		user = models.User{
// 			Email:      googleUser.Email,
// 			Name:       googleUser.Email[0 : len(googleUser.Email)-len("@gmail.com")],
// 			Provider:   "google",
// 			Password:   nil, // kosong krn oauth
// 			IsVerified: true,
// 		}

// 		_, err = db.DB.NewInsert().Model(&user).Exec(c)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
// 			return
// 		}
// 	}

// 	// === Generate JWT ===
// 	deviceId := c.GetHeader("X-DeviceId")
// 	token, err := middlewares.GenerateToken(&user, deviceId)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
// 		return
// 	}
// 	c.JSON(200, response.Success(200, "success", gin.H{
// 		"success":      true,
// 		"access_token": token,
// 		"token_type":   "Bearer",
// 		"expires_in":   3600,
// 		"user": gin.H{
// 			"id":       user.ID,
// 			"email":    user.Email,
// 			"name":     user.Name,
// 			"provider": "google",
// 		},
// 	}))

// }

func Auth(c *gin.Context) {
	var req AuthRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	switch req.Provider {

	case "google":
		handleGoogleLogin(c, req)

	case "local":
		handleLocalLogin(c, req)

	default:
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unsupported provider"})
	}
}

func handleGoogleLogin(c *gin.Context, req AuthRequest) {

	if req.OAuthToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "OAuth token required"})
		return
	}

	googleUser, err := verifyGoogleToken(req.OAuthToken)
	if err != nil || googleUser.Email == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Google OAuth token"})
		return
	}

	var user models.User
	err = db.DB.NewSelect().
		Model(&user).
		Where("email = ?", googleUser.Email).
		Scan(c)

	if err != nil { // auto register
		user = models.User{
			Email:      googleUser.Email,
			Name:       googleUser.Email,
			Provider:   "google",
			Password:   nil,
			IsVerified: true,
		}

		_, err = db.DB.NewInsert().Model(&user).Exec(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}
	}

	generateLoginResponse(c, &user)
}

func handleLocalLogin(c *gin.Context, req AuthRequest) {

	if req.Email == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password required"})
		return
	}

	var user models.User
	err := db.DB.NewSelect().
		Model(&user).
		Where("email = ?", req.Email).
		Scan(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if user.Password == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Account registered via OAuth"})
		return
	}

	// compare password
	err = bcrypt.CompareHashAndPassword(
		[]byte(*user.Password),
		[]byte(req.Password),
	)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if !user.IsVerified {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email not verified"})
		return
	}

	generateLoginResponse(c, &user)
}

func generateLoginResponse(c *gin.Context, user *models.User) {

	deviceId := c.GetHeader("X-DeviceId")

	token, err := middlewares.GenerateToken(user, deviceId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(200, response.Success(200, "success", gin.H{
		"success":      true,
		"access_token": token,
		"token_type":   "Bearer",
		"expires_in":   3600,
		"user": gin.H{
			"id":       user.ID,
			"email":    user.Email,
			"name":     user.Name,
			"provider": user.Provider,
		},
	}))
}

func RefreshToken(c *gin.Context) {
	var req RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, response.Error(400, "invalid request", nil))
		return
	}

	deviceId := c.GetHeader("X-DeviceId")
	if deviceId == "" {
		c.JSON(401, response.Error(401, "device id required", nil))
		return
	}

	tx, _ := db.DB.Begin()
	defer tx.Rollback()

	var rt auth_models.RefreshTokens
	err := tx.NewSelect().
		Model(&rt).
		Where("token = ?", req.RefreshToken).
		Where("device_id = ?", deviceId).
		Where("revoke = false").
		Scan(c)

	if err != nil || time.Now().After(rt.ExpiredAt) {
		c.JSON(401, response.Error(401, "invalid refresh token", nil))
		return
	}

	var user models.User
	_ = tx.NewSelect().
		Model(&user).
		Where("id = ?", rt.UserID).
		Scan(c)

	// revoke lama
	// _, err = tx.NewUpdate().
	// 	Model(&rt).
	// 	Set("revoke = true").
	// 	Where("id = ?", rt.ID).
	// 	Exec(c)
	_, err = tx.NewUpdate().
		Model((*auth_models.RefreshTokens)(nil)).
		Set("revoke = true").
		Where("id = ?", rt.ID).
		Exec(c)

	if err != nil {
		c.JSON(500, response.Error(500, "failed revoke token", nil))
		return
	}

	if err = tx.Commit(); err != nil {
		c.JSON(500, response.Error(500, "commit failed", nil))
		return
	}

	accessToken, err := middlewares.GenerateToken(&user, deviceId)
	if err != nil {
		c.JSON(500, response.Error(500, "failed generate token", nil))
		return
	}

	c.JSON(200, response.Success(200, "success", gin.H{
		"access_token": accessToken,
	}))
}

func Logout(c *gin.Context) {
	// user dari JWT middleware
	// user, exists := c.Get("user_id")
	// if !exists {
	// 	c.JSON(401, response.Error(401, "unauthorized", nil))
	// 	return
	// }

	// u := user.(*models.User)

	var req RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, response.Error(400, "invalid request", nil))
		return
	}

	deviceId := c.GetHeader("X-DeviceId")
	if deviceId == "" {
		c.JSON(400, response.Error(400, "device id required", nil))
		return
	}

	ctx := c.Request.Context()

	// revoke semua refresh token di device ini
	_, err := db.DB.NewUpdate().
		Model((*auth_models.RefreshTokens)(nil)).
		Set("revoke = true").
		Where("token = ?", req.RefreshToken).
		// Where("users_id = ?", user).
		Where("device_id = ?", deviceId).
		Where("revoke = false").
		Exec(ctx)

	if err != nil {
		c.JSON(500, response.Error(500, "failed logout", nil))
		return
	}

	c.JSON(200, response.Success(200, "logout success", nil))
}

func LogoutAll(c *gin.Context) {
	user := c.MustGet("user_id")
	ctx := c.Request.Context()

	_, err := db.DB.NewUpdate().
		Model((*auth_models.RefreshTokens)(nil)).
		Set("revoke = true").
		Where("users_id = ?", user).
		Where("revoke = false").
		Exec(ctx)

	if err != nil {
		c.JSON(500, response.Error(500, "failed logout all", nil))
		return
	}

	c.JSON(200, response.Success(200, "logout all success", nil))
}

// Verify Google Access Token
func verifyGoogleToken(accessToken string) (*GoogleTokenInfo, error) {
	resp, err := http.Get("https://www.googleapis.com/oauth2/v3/tokeninfo?access_token=" + accessToken)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var tokenInfo GoogleTokenInfo
	_ = json.Unmarshal(body, &tokenInfo)

	if tokenInfo.Email == "" {
		return nil, fmt.Errorf("invalid token")
	}

	return &tokenInfo, nil
}
