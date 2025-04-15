package main

import (
	"github.com/ashkanamani/url-shortener/internal/config"
	"github.com/ashkanamani/url-shortener/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type App struct {
	Config *config.Config
	Router *gin.Engine
	Logger *logrus.Logger
}

func (a *App) registerRoutes() {
	a.Router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
}

func main() {
	app := &App{
		Config: config.LoadConfig(),
		Router: gin.Default(),
		Logger: utils.InitLogger(),
	}
	app.registerRoutes()

	app.Logger.Printf("🚀 Server starting on port %s...", app.Config.Port)
	_ = app.Router.Run(":" + app.Config.Port)
}
