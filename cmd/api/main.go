package main

import (
	"context"
	"github.com/ashkanamani/url-shortener/internal/config"
	"github.com/ashkanamani/url-shortener/internal/repository"
	"github.com/ashkanamani/url-shortener/pkg/logger"
	"github.com/ashkanamani/url-shortener/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	Config *config.Config
	Router *gin.Engine
	Logger *logrus.Logger
	DB     *gorm.DB
}

// registerRoutes wires the public endpoints for this slice of the project.
func (a *App) registerRoutes() {
	a.Router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
}

// gracefulShutdown blocks until SIGINT/SIGTERM, then closes HTTP & DB.
func (a *App) gracefullyShutdown(srv *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	a.Logger.Info("shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		a.Logger.Fatalf("HTTP server shutdown error: %v", err)
	}

	sqlDB, err := a.DB.DB()
	if err == nil {
		_ = sqlDB.Close()
	}
	a.Logger.Info("graceful shutdown complete")
}

func main() {

	cfg := config.LoadConfig() // reads .env via godotenv
	log := logger.InitLogger(cfg.LogLevel)

	gin.SetMode(cfg.GinMode) // "debug" | "release"
	r := gin.New()
	r.Use(middleware.RequestLogger(log))
	r.Use(gin.Recovery())

	db := repository.InitDB(cfg.PostgresAddr, log)

	app := &App{
		Config: cfg,
		Router: r,
		Logger: log,
		DB:     db,
	}
	app.registerRoutes()

	app.Logger.Printf("🚀 Server starting on port %s...", app.Config.ServerPort)
	_ = app.Router.Run(":" + app.Config.ServerPort)
}
