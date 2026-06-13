package handlers

import (
	"net/http"

	"collab-code-platform/internal/dto"
	"collab-code-platform/internal/services"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userService *services.UserService
}

func NewAuthHandler(
	userService *services.UserService,
) *AuthHandler {
	return &AuthHandler{
		userService: userService,
	}
}

func (h *AuthHandler) Signup(
	c *gin.Context,
) {

	var req dto.SignupRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "invalid request",
			},
		)

		return
	}

	err := h.userService.CreateUser(
		req.Email,
		req.Password,
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
		http.StatusCreated,
		gin.H{
			"message": "user created successfully",
		},
	)
}

func (h *AuthHandler) Login(
	c *gin.Context,
) {

	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "invalid request",
			},
		)

		return
	}

	token, err := h.userService.Login(
		req.Email,
		req.Password,
	)

	if err != nil {

		c.JSON(
			http.StatusUnauthorized,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"token": token,
			"email": req.Email,
		},
	)
}

func (h *AuthHandler) Me(
	c *gin.Context,
) {

	userID, _ := c.Get("user_id")

	email, _ := c.Get("email")

	c.JSON(
		http.StatusOK,
		gin.H{
			"user_id": userID,
			"email":   email,
		},
	)
}
