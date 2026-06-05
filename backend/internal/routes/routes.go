package routes

import (
	"collab-code-platform/configs"
	"collab-code-platform/internal/handlers"
	"collab-code-platform/internal/repositories"
	"collab-code-platform/internal/services"
	"collab-code-platform/internal/middleware"
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

	r.POST(
		"/signup",
		authHandler.Signup,
	)
	r.POST(
		"/login",
		authHandler.Login,
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
	}
}
