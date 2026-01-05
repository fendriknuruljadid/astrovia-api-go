package handlers

import (
	"net/http"

	ws "app/internal/packages/websocket"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // nanti bisa diketatkan
	},
}

func ProgressWS(c *gin.Context) {
	videoID := c.Query("video_id")
	if videoID == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "video_id required"})
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	ws.ProgressHub.Register(videoID, conn)

	defer func() {
		ws.ProgressHub.Unregister(videoID, conn)
		conn.Close()
	}()

	// keep alive
	for {
		if _, _, err := conn.ReadMessage(); err != nil {
			break
		}
	}
}
