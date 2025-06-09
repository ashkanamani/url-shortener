package main

import (
	"github.com/ashkanamani/url-shortener/internal/config"
	"github.com/ashkanamani/url-shortener/internal/core"
	"github.com/ashkanamani/url-shortener/internal/db/postgres"
	"github.com/ashkanamani/url-shortener/internal/http/middleware"
	"github.com/ashkanamani/url-shortener/internal/http/routes"
	"github.com/ashkanamani/url-shortener/pkg/logger"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.LoadConfig() // reads .env via godotenv
	log := logger.InitLogger(cfg.LogLevel)

	gin.SetMode(cfg.GinMode) // "debug" | "release"
	router := gin.New()
	router.Use(middleware.RequestLogger(log))
	router.Use(gin.Recovery())

	db := postgres.InitDB(cfg.PostgresAddr, log)

	app := &core.App{
		Config: cfg,
		Router: router,
		Logger: log,
		DB:     db,
	}
	routes.SetupRouter(app)

	app.Logger.Printf("🚀 Server starting on port %s...", app.Config.ServerPort)
	_ = app.Router.Run(":" + app.Config.ServerPort)
}
