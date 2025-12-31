package routesV1

import (
	auth "app/internal/services/v1/auth/handlers"
	"app/internal/services/v1/user/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterPublicRoutes(r *gin.RouterGroup) {
	r.POST("/users", handlers.CreateUser)
	r.POST("/auth/generate-token", auth.Auth)
}

func RegisterProtectedRoutes(r *gin.RouterGroup) {
	r.GET("/users", handlers.GetUsers)
	r.GET("/users/:id", handlers.GetUserByID)
	r.PUT("/users/:id", handlers.UpdateUser)
	r.DELETE("/users/:id", handlers.DeleteUser)
	r.POST("/auth/refresh-token", auth.RefreshToken)
	r.POST("/auth/logout", auth.Logout)
	r.POST("/auth/logout-all", auth.LogoutAll)
}
