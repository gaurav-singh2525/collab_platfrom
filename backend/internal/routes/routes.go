package routes

import (
	"collab-code-platform/configs"
	"collab-code-platform/internal/handlers"
	"collab-code-platform/internal/middleware"
	"collab-code-platform/internal/repositories"
	"collab-code-platform/internal/services"
	"collab-code-platform/internal/websocket"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	r *gin.Engine,
	db *sql.DB,
	cfg *configs.Config,
) {
	userRepo := repositories.NewUserRepository(
		db,
	)

	userService := services.NewUserService(
		userRepo,
		cfg.JWTSecret,
	)

	authHandler := handlers.NewAuthHandler(
		userService,
	)

	executionService :=
		services.NewExecutionService()

	executionHandler :=
		handlers.NewExecutionHandler(
			executionService,
		)

	hub :=
		websocket.NewHub()

	wsHandler :=
		websocket.NewHandler(
			hub,
		)

	r.POST(
		"/signup",
		authHandler.Signup,
	)
	r.POST(
		"/login",
		authHandler.Login,
	)
	r.GET(
		"/ws/:roomId",
		wsHandler.Connect,
	)
	protected := r.Group("/")
	protected.Use(
		middleware.AuthMiddleware(cfg.JWTSecret),
	)
	{
		protected.GET(
			"/me",
			authHandler.Me,
		)
		protected.POST(
			"/execute",
			executionHandler.Execute,
		)
	}
}
