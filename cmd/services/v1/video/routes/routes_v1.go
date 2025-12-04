package routesV1

import (
	"app/internal/services/v1/video/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterPublicRoutes(r *gin.RouterGroup) {

}

func RegisterProtectedRoutes(r *gin.RouterGroup) {
	r.POST("/video", handlers.CreateVideo)
	r.GET("/video", handlers.GetVideo)
	r.GET("/video/:id", handlers.GetVideoByID)
	r.PUT("/video/:id", handlers.UpdateVideo)
	r.DELETE("/video/:id", handlers.DeleteVideo)
}
