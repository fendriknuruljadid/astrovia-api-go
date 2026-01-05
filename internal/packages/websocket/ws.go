package ws

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Hub struct {
	mu      sync.RWMutex
	clients map[string]map[*websocket.Conn]bool // videoID -> conns
}

var ProgressHub = &Hub{
	clients: make(map[string]map[*websocket.Conn]bool),
}

func (h *Hub) Register(videoID string, conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.clients[videoID] == nil {
		h.clients[videoID] = make(map[*websocket.Conn]bool)
	}
	h.clients[videoID][conn] = true
}

func (h *Hub) Unregister(videoID string, conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()

	delete(h.clients[videoID], conn)
	if len(h.clients[videoID]) == 0 {
		delete(h.clients, videoID)
	}
}

func (h *Hub) Broadcast(videoID string, payload []byte) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	for conn := range h.clients[videoID] {
		_ = conn.WriteMessage(websocket.TextMessage, payload)
	}
}
