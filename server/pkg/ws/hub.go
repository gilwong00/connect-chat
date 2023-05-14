package ws

import "fmt"

type Room struct {
	ID      string             `json:"id"`
	Name    string             `json:"name"`
	Clients map[string]*Client `json:"clients"`
}

type Hub struct {
	Rooms      map[string]*Room
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[string]*Room),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message, 5),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			// check if room id exists
			if _, ok := h.Rooms[client.RoomID]; ok {
				fmt.Println(">>> has room id")
				r := h.Rooms[client.RoomID]
				// add client to room if it doesnt exist
				if _, ok := r.Clients[client.ID]; !ok {
					r.Clients[client.ID] = client
				}
			} else {
				fmt.Println(">> failed")
			}
		case client := <-h.Unregister:
			if _, ok := h.Rooms[client.RoomID]; ok {
				if _, ok := h.Rooms[client.ID]; ok {
					// broadcast message indicating client has left the room if there are still users
					if len(h.Rooms[client.RoomID].Clients) != 0 {
						h.Broadcast <- &Message{
							Content:  "user left the room",
							RoomID:   client.RoomID,
							Username: client.Username,
						}
					}
					// if the client exists delete it to unregister
					delete(h.Rooms[client.RoomID].Clients, client.ID)
					// close message channel for the client
					close(client.Message)
				}
			}
		case message := <-h.Broadcast:
			if _, ok := h.Rooms[message.RoomID]; ok {
				// iterate over all clients in the room to send message
				for _, client := range h.Rooms[message.RoomID].Clients {
					client.Message <- message
				}
			}
		}
	}
}
