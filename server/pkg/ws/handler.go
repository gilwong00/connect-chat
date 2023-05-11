package ws

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

type HubHandler struct {
	hub *Hub
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewHubHandler(h *Hub) *HubHandler {
	return &HubHandler{
		hub: h,
	}
}

func (h *HubHandler) AppendNewRoom(roomID string, roomName string) {
	h.hub.Rooms[roomID] = &Room{
		ID:      roomID,
		Name:    roomName,
		Clients: make(map[string]*Client),
	}
}

func (h *HubHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.URL.Path, "/room/join/") {
		h.joinRoom(w, r)
		return
	}
}

func (h *HubHandler) joinRoom(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "failed to creat websocket conection", http.StatusInternalServerError)
		return
	}
	fmt.Println("conn", conn)
}
