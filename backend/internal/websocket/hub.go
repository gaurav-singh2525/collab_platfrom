package websocket

import (
	"collab-code-platform/internal/models"
	"encoding/json"
	"github.com/gorilla/websocket"
)

type Hub struct {
	Rooms map[string]*Room
}

func NewHub() *Hub {

	return &Hub{
		Rooms: make(
			map[string]*Room,
		),
	}
}

func (h *Hub) BroadcastToRoom(
	roomID string,
	message []byte,
) {

	clients :=
		h.Rooms[roomID].Clients

	for client := range clients {

		client.Conn.WriteMessage(
			websocket.TextMessage,
			message,
		)
	}
}

func (h *Hub) AddClient(
	client *Client,
) {

	roomID := client.RoomID

	if _, exists := h.Rooms[roomID]; !exists {
		h.Rooms[roomID] = &Room{
			Clients:  make(map[*Client]bool),
			Code:     "",
			Language: "python",
		}
	}

	h.Rooms[roomID].Clients[client] = true
}

func (h *Hub) RemoveClient(
	client *Client,
) {

	roomID := client.RoomID

	delete(
		h.Rooms[roomID].Clients,
		client)

	if len(
		h.Rooms[roomID].Clients) == 0 {

		delete(
			h.Rooms,
			roomID,
		)
	}
}

func (h *Hub) BroadcastPresence(
	roomID string,
) {
	count := len(
		h.Rooms[roomID].Clients,
	)
	msg := models.WSMessage{
		Type:  "presence",
		Count: count,
	}
	data, _ :=
		json.Marshal(msg)
	h.BroadcastToRoom(
		roomID,
		data,
	)
}
