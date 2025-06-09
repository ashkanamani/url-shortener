package routes

import (
	"fmt"
	"github.com/ashkanamani/url-shortener/internal/core"
	"github.com/ashkanamani/url-shortener/internal/db/redis"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// SetupRouter wires the public endpoints for this slice of the project.
func SetupRouter(app *core.App) {
	app.Router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	app.Router.POST("/shorten", newShortenURLHandler(app))
	app.Router.GET("/:redirect", newRedirectionHandler(app))
}

func newShortenURLHandler(app *core.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		type request struct {
			OriginalURL string `json:"original_url" binding:"required"`
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
			return
		}
		shorted := uuid.NewString()[:7]
		shortURL := fmt.Sprintf("http://localhost:%s/%s", app.Config.ServerPort, shorted)

		// TODO: Save to db: original_url -> shorted_url
		redis.Set(shorted, req.OriginalURL)
		c.JSON(http.StatusOK, gin.H{
			"original_url": req.OriginalURL,
			"short_url":    shortURL,
			"code":         shorted,
		})
	}
}

func newRedirectionHandler(app *core.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		code := c.Param("redirect")
		app.Logger.Println("Path to", code)

		originalURL, err := redis.Get(code)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
			return
		}
		c.Redirect(http.StatusMovedPermanently, originalURL)
	}
}
