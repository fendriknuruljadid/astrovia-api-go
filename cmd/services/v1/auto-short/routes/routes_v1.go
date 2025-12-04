package routesV1

import (
	"app/internal/services/v1/auto-short/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterPublicRoutes(r *gin.RouterGroup) {

}

func RegisterProtectedRoutes(r *gin.RouterGroup) {
	r.POST("/auto-short", handlers.CreateVideo)
	r.GET("/auto-short", handlers.GetVideo)
	r.GET("/auto-short/:id", handlers.GetVideoByID)
	r.PUT("/auto-short/:id", handlers.UpdateVideo)
	r.DELETE("/auto-short/:id", handlers.DeleteVideo)
}
