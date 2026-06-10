package websocket

import (
	"net/http"

	"collab-code-platform/internal/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Handler struct {
	hub *Hub
}

func NewHandler(
	hub *Hub,
) *Handler {

	return &Handler{
		hub: hub,
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(
		r *http.Request,
	) bool {

		return true
	},
}

func (h *Handler) Connect(
	c *gin.Context,
) {
	conn, err :=
		upgrader.Upgrade(
			c.Writer,
			c.Request,
			nil,
		)

	if err != nil {
		return
	}

	roomID :=
		c.Param(
			"roomId",
		)

	client := &Client{
		Conn:   conn,
		RoomID: roomID,
	}

	h.hub.AddClient(
		client,
	)

	room := h.hub.Rooms[roomID]

	codeMsg := models.WSMessage{
		Type: "code_change",
		Code: room.Code,
	}

	data, err := json.Marshal(codeMsg)

	if err == nil {
		conn.WriteMessage(
			websocket.TextMessage,
			data,
		)
	}

	langMsg := models.WSMessage{
		Type:     "language_change",
		Language: room.Language,
	}

	data, err = json.Marshal(langMsg)

	if err == nil {
		conn.WriteMessage(
			websocket.TextMessage,
			data,
		)
	}

	h.hub.BroadcastPresence(
		roomID,
	)
	defer func() {

		h.hub.RemoveClient(
			client,
		)

		h.hub.BroadcastPresence(
			roomID,
		)

		conn.Close()

	}()
	for {

		_, message, err :=
			conn.ReadMessage()

		if err != nil {
			break
		}

		var msg models.WSMessage

		err = json.Unmarshal(
			message,
			&msg,
		)

		if err != nil {
			continue
		}

		room := h.hub.Rooms[roomID]

		if msg.Type == "code_change" {
			room.Code = msg.Code
		}

		if msg.Type == "language_change" {
			room.Language = msg.Language
		}

		h.hub.BroadcastToRoom(
			roomID,
			message,
		)
	}
}
