package config

import (
	"os"
	"testing"
)

func TestLoadConfigDefaults(t *testing.T) {
	os.Clearenv()
	// required var missing -> should fatal; use defer to catch it
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("expected fatal due to missing JWT_SECRET")
		}
	}()
	LoadConfig()
}

func TestLoadConfigSuccess(t *testing.T) {
	t.Setenv("POSTGRES_ADDR", "postgres://test@localhost:5432/db")
	t.Setenv("SERVER_PORT", "9999")
	t.Setenv("GIN_MODE", "release")
	t.Setenv("LOG_LEVEL", "debug")
	t.Setenv("JWT_SECRET", "secret")
	t.Setenv("REDIS_ADDR", "redis://test@localhost:6379")

	cfg := LoadConfig()

	if cfg.ServerPort != "9999" {
		t.Errorf("want port 9999, got %s", cfg.ServerPort)
	}
	if cfg.GinMode != "release" {
		t.Errorf("gin mode not loaded")
	}
	if cfg.LogLevel != "debug" {
		t.Errorf("log level not loaded")
	}
	if cfg.PostgresAddr == "" {
		t.Errorf("postgres addr missing")
	}
	if cfg.JWTSecret == "" {
		t.Errorf("jwt secret missing")
	}
	if cfg.RedisAddr != "redis://test@localhost:6379" {
		t.Errorf("redis addr missing")
	}

}
