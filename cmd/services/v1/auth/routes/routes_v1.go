package routesV1

import (
	"github.com/gin-gonic/gin"
    "astrovia-api-go/internal/services/v1/user/handlers"	
)

func RegisterPublicRoutes(r *gin.RouterGroup) {
    r.POST("/users", handlers.CreateUser)
}

func RegisterProtectedRoutes(r *gin.RouterGroup) {
    r.GET("/users", handlers.GetUsers)
    r.GET("/users/:id", handlers.GetUserByID)
    r.PUT("/users/:id", handlers.UpdateUser)
    r.DELETE("/users/:id", handlers.DeleteUser)
}
