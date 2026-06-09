package websocket

import (
	"net/http"

	"fmt"
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

		fmt.Println(
			"Received:",
			string(message),
		)
		h.hub.BroadcastToRoom(
			roomID,
			message,
		)
	}
}
