package routesV1

import (
	autoCaption "app/internal/services/v1/astro-zenith/auto-caption/handlers"
	autoClip "app/internal/services/v1/astro-zenith/auto-caption/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterPublicRoutes(r *gin.RouterGroup) {
	r.POST("/auto-caption", autoCaption.CreateVideo)
	r.GET("/auto-caption", autoCaption.GetVideo)
	r.GET("/auto-caption/:id", autoCaption.GetVideoByID)
	r.PUT("/auto-caption/:id", autoCaption.UpdateVideo)
	r.DELETE("/auto-caption/:id", autoCaption.DeleteVideo)

	r.POST("/auto-clip", autoClip.CreateVideo)
	r.GET("/auto-clip", autoClip.GetVideo)
	r.GET("/auto-clip/:id", autoClip.GetVideoByID)
	r.PUT("/auto-clip/:id", autoClip.UpdateVideo)
	r.DELETE("/auto-clip/:id", autoClip.DeleteVideo)
}

func RegisterProtectedRoutes(r *gin.RouterGroup) {

}
