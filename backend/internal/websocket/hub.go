package websocket

import (
	"github.com/gorilla/websocket"
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
