package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	Environment string
	MongoDB     MongoDBConfig
	JWT         JWTConfig
	API         APIConfig
}

type MongoDBConfig struct {
	URI    string
	DBName string
}

type JWTConfig struct {
	Secret string
	Expiry string
}

type APIConfig struct {
	Version string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	return &Config{
		Port:        GetEnv("PORT", "8080"),
		Environment: GetEnv("ENV", "development"),
		MongoDB: MongoDBConfig{
			URI:    GetEnv("MONGODB_URI", "mongodb://localhost:27017"),
			DBName: GetEnv("DB_NAME", "go-kas"),
		},
		JWT: JWTConfig{
			Secret: GetEnv("JWT_SECRET", "your-secret-key-here"),
			Expiry: GetEnv("JWT_EXPIRY", "24h"),
		},
		API: APIConfig{
			Version: GetEnv("API_VERSION", "v1"),
		},
	}
}

func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
