package config

import (
	"os"
)

type Config struct {
	ServerPort string
	MongoURI   string
	JWTSecret  string
}

func LoadConfig() *Config {
	return &Config{
		ServerPort: getEnv("SERVER_PORT", "8080"),
		MongoURI:   getEnv("MONGO_URI", "mongodb://localhost:27017"),
		JWTSecret:  getEnv("JWT_SECRET", "your-secret-key"),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
