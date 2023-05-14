package ws

import (
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
		// limit our origin if this was a prod app
		// origin := r.Header.Get("Origin")
		// return origin == "http://localhost:3000"
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
	if strings.Contains(r.URL.Path, "/room/join") {
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
	roomID := r.URL.Query().Get("roomId")
	clientId := r.URL.Query().Get("clientId")
	username := r.URL.Query().Get("username")
	client := &Client{
		Conn: conn,
		//buffer channel
		Message:  make(chan *Message, 10),
		ID:       clientId,
		RoomID:   roomID,
		Username: username,
	}
	// new message to broadcast a new user has joined
	message := &Message{
		Content:  "A new user has joined the room",
		RoomID:   roomID,
		Username: username,
	}
	// registry a new client through the register channel
	h.hub.Register <- client
	// Broadcast this message
	h.hub.Broadcast <- message
	// write message to socket
	// handle in a separate go routine
	go client.writeMessage()
	// reading incoming messages
	client.readMessage(h.hub)
}
