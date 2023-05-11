package ws

import "github.com/gorilla/websocket"

// client for ws connection
type Client struct {
	Conn     *websocket.Conn
	Message  chan *Message
	ID       string `json:"id"`
	RoomID   string `json:"roomId"`
	Username string `json:"username"`
}

// incoming message from chat
type Message struct {
	Content  string `json:"content"`
	RoomID   string `json:"roomId"`
	Username string `json:"username"`
}
