package handlers

import (
	"collab-code-platform/internal/dto"
	"collab-code-platform/internal/models"
	"collab-code-platform/internal/services"
	"collab-code-platform/internal/websocket"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ExecutionHandler struct {
	service *services.ExecutionService
	hub     *websocket.Hub
}

func NewExecutionHandler(
	service *services.ExecutionService,
	hub *websocket.Hub,
) *ExecutionHandler {

	return &ExecutionHandler{
		service: service,
		hub:     hub,
	}
}

func (h *ExecutionHandler) Execute(
	c *gin.Context,
) {

	var req dto.ExecuteRequest

	if err := c.ShouldBindJSON(
		&req,
	); err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "invalid request",
			},
		)

		return
	}

	result, err :=
		h.service.Execute(
			req.Language,
			req.Code,
		)

	output := result.Stdout

	if result.Stderr != "" {
		output = result.Stderr
	}

	msg := models.WSMessage{
		Type:   "execution_output",
		Output: output,
	}

	data, _ := json.Marshal(msg)
	h.hub.BroadcastToRoom(
		req.RoomID,
		data,
	)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		result,
	)
}
