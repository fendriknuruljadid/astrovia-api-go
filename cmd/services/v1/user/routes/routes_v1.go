package routesV1

import (
	auth "app/internal/services/v1/auth/handlers"
	"app/internal/services/v1/user/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterPublicRoutes(r *gin.RouterGroup) {
	r.POST("/users", handlers.CreateUser)
	r.POST("/users/check", handlers.CheckUser)
	r.POST("/users/create-password", handlers.CreatePassword)
	r.POST("/users/verify-verification", handlers.VerifyOTP)
	r.POST("/users/resend-verification", handlers.ResendOTP)
	r.POST("/auth/generate-token", auth.Auth)
	r.POST("/auth/refresh-token", auth.RefreshToken)
	r.POST("/auth/logout", auth.Logout)
}

func RegisterProtectedRoutes(r *gin.RouterGroup) {
	r.GET("/users", handlers.GetUsers)
	r.GET("/users/:id", handlers.GetUserByID)
	r.PUT("/users/:id", handlers.UpdateUser)
	r.DELETE("/users/:id", handlers.DeleteUser)
	r.POST("/auth/logout-all", auth.LogoutAll)
}
