package middlewares

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"

	"app/internal/packages/response"
	"app/internal/packages/utils"
	auth_models "app/internal/services/v1/auth/models"
	"app/internal/services/v1/auth/repository"
	"app/internal/services/v1/user/models"
)

func init() {
	godotenv.Load()
}

func getJWTSecret() []byte {
	secret := os.Getenv("JWT_SECRET")
	return []byte(secret)
}

type Claims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

type AuthTokens struct {
	AccessToken  string
	RefreshToken string
	ExpiresIn    int64
}

// Generate Token
func GenerateToken(u *models.User, deviceId string) (*AuthTokens, error) {
	expirationTime := time.Now().Add(2 * time.Hour)
	// expirationTime := time.Now().Add(30 * time.Second)

	claims := &Claims{
		ID:       u.ID,
		Username: u.Name,
		Email:    u.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "astrovia-auth-service",
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(getJWTSecret())
	if err != nil {
		return nil, err
	}
	refreshToken := utils.GenerateRefreshToken()
	ref := auth_models.RefreshTokens{
		UserID:    u.ID,
		Token:     refreshToken,
		DeviceId:  deviceId,
		ExpiredAt: time.Now().Add(30 * 24 * time.Hour),
		Revoke:    false,
	}
	repository.CreateRefreshTokens(&ref)

	return &AuthTokens{
		AccessToken:  token,
		RefreshToken: refreshToken,
		ExpiresIn:    int64(2 * 60 * 60),
		// ExpiresIn: int64(30),
	}, nil
	// return token.SignedString(getJWTSecret())
}

// JWT AUTH Middleware (fully fixed)
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		secret := getJWTSecret()
		if len(secret) == 0 {
			c.JSON(http.StatusInternalServerError, response.Error(500, "JWT secret not set", nil))
			c.Abort()
			return
		}

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, response.Error(401, "Missing Authorization header", nil))
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, response.Error(401, "Invalid Authorization header", nil))
			c.Abort()
			return
		}

		tokenStr := parts[1]
		claims := &Claims{}

		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, response.Error(401, "Invalid or expired token", nil))
			c.Abort()
			return
		}

		c.Set("user_id", claims.ID)
		c.Set("username", claims.Username)
		c.Set("email", claims.Email)

		c.Next()
	}
}
