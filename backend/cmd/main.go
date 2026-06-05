package main

import (
	"collab-code-platform/configs"
	"collab-code-platform/internal/middleware"
	"collab-code-platform/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg := configs.LoadConfig()

	db := configs.ConnectDB(cfg)

	defer db.Close()

	r := gin.Default()

	r.Use(middleware.LoggerMiddleware())
	r.Use(middleware.CORSMiddleware())

	routes.SetupRoutes(
	r,
	db,
	cfg,
)

	r.Run(":" + cfg.Port)
}
