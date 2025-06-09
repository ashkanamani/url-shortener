package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"time"
)

func RequestLogger(log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ensure every request has a UUID (client -> X-Request-ID wins)
		reqID := c.GetHeader("X-Request-ID")
		if reqID == "" {
			reqID = uuid.NewString()
		}
		c.Writer.Header().Set("X-Request-ID", reqID)

		start := time.Now()
		c.Set("req_id", reqID)

		c.Next()
		latency := time.Since(start)

		log.WithFields(logrus.Fields{
			"req_id":   reqID,
			"status":   c.Writer.Status(),
			"method":   c.Request.Method,
			"path":     c.Request.URL.Path,
			"latency":  latency.Milliseconds(),
			"clientIP": c.ClientIP(),
		}).Info("request completed")

	}
}
