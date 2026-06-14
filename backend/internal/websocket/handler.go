package websocket

import (
	"net/http"

	"collab-code-platform/internal/models"
	"collab-code-platform/internal/repositories"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Handler struct {
	hub *Hub

	roomRepo *repositories.RoomRepository
}

func NewHandler(
	hub *Hub,
	roomRepo *repositories.RoomRepository,
) *Handler {

	return &Handler{
		hub:      hub,
		roomRepo: roomRepo,
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

	username :=
		c.Query(
			"username",
		)
	client := &Client{
		Conn:     conn,
		RoomID:   roomID,
		Username: username,
	}

	h.hub.AddClient(
		client,
	)

	savedRoom, err :=
		h.roomRepo.GetRoom(
			roomID,
		)

	if err == nil {

		if room,
			ok := h.hub.Rooms[roomID]; ok {

			room.Code =
				savedRoom.Code

			room.Language =
				savedRoom.Language
		}
	}

	room := h.hub.Rooms[roomID]

	joinMsg := models.WSMessage{
		Type: "user_joined",

		Username: client.Username,
	}
	data1, _ :=
		json.Marshal(
			joinMsg,
		)
	h.hub.BroadcastToRoom(
		roomID,
		data1,
	)

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

		leaveMsg := models.WSMessage{
			Type: "user_left",

			Username: client.Username,
		}
		data, _ :=
			json.Marshal(
				leaveMsg,
			)

		h.hub.BroadcastToRoom(
			roomID,
			data,
		)

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
			h.roomRepo.SaveRoom(
				roomID,
				room.Code,
				room.Language,
			)
		}

		if msg.Type ==
			"language_change" {

			room.Language =
				msg.Language

			h.roomRepo.SaveRoom(
				roomID,
				room.Code,
				room.Language,
			)
		}

		h.hub.BroadcastToRoom(
			roomID,
			message,
		)
	}
}
