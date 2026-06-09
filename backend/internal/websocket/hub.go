package websocket

import (
	"collab-code-platform/internal/models"
	"github.com/gorilla/websocket"
	"encoding/json"
)

type Hub struct {
	Rooms map[string]map[*Client]bool
}

func NewHub() *Hub {

	return &Hub{
		Rooms: make(
			map[string]map[*Client]bool,
		),
	}
}

func (h *Hub) BroadcastToRoom(
	roomID string,
	message []byte,
) {

	clients :=
		h.Rooms[roomID]

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

	if _, exists :=
		h.Rooms[roomID]; !exists {

		h.Rooms[roomID] =
			make(
				map[*Client]bool,
			)
	}

	h.Rooms[roomID][client] =
		true
}

func (h *Hub) RemoveClient(
	client *Client,
) {

	roomID := client.RoomID

	delete(
		h.Rooms[roomID],
		client,
	)

	if len(
		h.Rooms[roomID],
	) == 0 {

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
		h.Rooms[roomID],
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
