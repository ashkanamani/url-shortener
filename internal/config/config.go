package config

import "os"

type Config struct {
	Port      string
	DBURL     string
	JWTSecret string
	RedisAddr string
}

func LoadConfig() *Config {
	return &Config{
		Port:      getEnv("APP_PORT", "8080"),
		DBURL:     getEnv("DATABASE_URL", "postgres://user:pass@localhost:5432/urlshortener?sslmode=disable"),
		JWTSecret: getEnv("JWT_SECRET", "supersecretkey"),
		RedisAddr: getEnv("REDIS_ADDR", "localhost:6379"),
	}
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}
