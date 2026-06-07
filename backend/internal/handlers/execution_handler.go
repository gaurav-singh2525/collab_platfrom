package handlers

import (
	"net/http"

	"collab-code-platform/internal/dto"
	"collab-code-platform/internal/services"

	"github.com/gin-gonic/gin"
)

type ExecutionHandler struct {
	service *services.ExecutionService
}

func NewExecutionHandler(
	service *services.ExecutionService,
) *ExecutionHandler {

	return &ExecutionHandler{
		service: service,
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
