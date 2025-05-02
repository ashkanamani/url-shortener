package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestLoggerAddsRequestID(t *testing.T) {
	// capture log output
	var buf bytes.Buffer
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(&buf)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(RequestLogger(log))
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	r.ServeHTTP(w, req)

	// 1) header should exist
	if rid := w.Header().Get("X-Request-ID"); rid == "" {
		t.Fatalf("X-Request-ID header missing")
	}

	// 2) log line should contain same req_id
	var entry map[string]any
	if err := json.Unmarshal(buf.Bytes(), &entry); err != nil {
		t.Fatalf("unmarshal log: %v", err)
	}

	if entry["req_id"] == "" {
		t.Fatalf("log entry missing req_id")
	}
	if entry["status"] != float64(200) { // json numbers -> float64
		t.Errorf("want status 200, got %v", entry["status"])
	}

}
