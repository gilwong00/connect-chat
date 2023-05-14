package ws

import (
	"log"

	"github.com/gorilla/websocket"
)

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

func (c *Client) writeMessage() {
	// close connection at the end
	defer func() {
		c.Conn.Close()
	}()

	for {
		// we get this message through the message channel
		message, ok := <-c.Message
		if !ok {
			return
		}
		// send message back to socket
		c.Conn.WriteJSON(message)
	}
}

func (c *Client) readMessage(hub *Hub) {
	defer func() {
		// unregister a client
		hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(
				err, websocket.CloseGoingAway,
				websocket.CloseAbnormalClosure,
			) {
				log.Panicf("error: %v", err)
			}
			break
		}
		msg := &Message{
			Content:  string(message),
			RoomID:   c.RoomID,
			Username: c.Username,
		}
		// broadcasting message to the hub
		hub.Broadcast <- msg
	}
}
