package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	ServerPort   string
	PostgresAddr string
	JWTSecret    string
	RedisAddr    string
	LogLevel     string // "debug" | "info" | "warn" | "error"
	GinMode      string // "debug" | "release" | "test"
}

func LoadConfig() *Config {
	// Load .env if present. Ignore error when file doesn’t exist.
	_ = godotenv.Load()

	// helper that falls back to default
	getEnv := func(key, fallback string) string {
		if val, ok := os.LookupEnv(key); ok {
			return val
		}
		return fallback
	}
	cfg := &Config{
		ServerPort:   getEnv("SERVER_PORT", "8080"),
		GinMode:      getEnv("GIN_MODE", "debug"),
		LogLevel:     getEnv("LOG_LEVEL", "info"),
		PostgresAddr: getEnv("POSTGRES_ADDR", "postgres://localhost:5432/mydb?sslmode=disable"),
		RedisAddr:    getEnv("REDIS_ADDR", "localhost:6379"),
		JWTSecret:    os.Getenv("JWT_SECRET"),
	}

	if cfg.JWTSecret == "" {
		panic("config error: JWT_SECRET is required but missing")
	}
	return cfg
}
