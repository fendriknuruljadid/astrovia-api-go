package auth

import (
	"app/internal/middlewares"
	"app/internal/packages/db"
	"app/internal/services/v1/user/models"
	"net/http"

	"encoding/json"
	"fmt"
	"io"

	"app/internal/packages/response"

	"github.com/gin-gonic/gin"
)

type AuthRequest struct {
	Provider   string `json:"provider"`   // hanya: google
	OAuthToken string `json:"oauthToken"` // access_token dari Google Sign-In
	Email      string `json:"email"`      // optional (Google bisa kirim)
}

// Struktur data Google token validation response
type GoogleTokenInfo struct {
	Email         string `json:"email"`
	EmailVerified string `json:"email_verified"`
	Audience      string `json:"aud"`
}

func Auth(c *gin.Context) {
	var req AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Provider wajib Google
	if req.Provider != "google" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Only Google login allowed"})
		return
	}

	if req.OAuthToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "OAuth token required"})
		return
	}

	// === Validasi token Google ===
	googleUser, err := verifyGoogleToken(req.OAuthToken)
	if err != nil || googleUser.Email == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Google OAuth token"})
		return
	}

	// === Cek user dari database ===
	var user models.User
	err = db.DB.NewSelect().Model(&user).Where("email = ?", googleUser.Email).Scan(c)

	if err != nil { // user belum ada → auto register
		user = models.User{
			Email:    googleUser.Email,
			Name:     googleUser.Email[0 : len(googleUser.Email)-len("@gmail.com")],
			Provider: "google",
			Password: "", // kosong krn oauth
		}

		_, err = db.DB.NewInsert().Model(&user).Exec(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}
	}

	// === Generate JWT ===
	token, err := middlewares.GenerateToken(&user)
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
			"provider": "google",
		},
	}))

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
