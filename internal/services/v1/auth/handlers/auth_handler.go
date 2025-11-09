package auth

import (
	"net/http"
	"app/internal/middlewares"
	"app/internal/services/v1/user/models"
	"app/internal/packages/db"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Auth(c *gin.Context) {
	var req AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	//  Ambil user dari database
	var user models.User
	err := db.DB.NewSelect().Model(&user).Where("email = ?", req.Email).Scan(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email not found"})
		return
	}

	// Check bcrypt password
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	// ✅ Buat token
	token, err := middlewares.GenerateToken(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":       token,
		"expires_in":  "2h",
		"user_id":     user.ID,
		"name":        user.Name,
		"email":       user.Email,
	})
}
