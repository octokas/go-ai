package config

import (
	"encoding/json"
	"os"
	"strconv"
	"sync"
)

type Config struct {
	Port     int
	Env      string
	LogLevel string
	Server   ServerConfig
	AI       AIConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Port    int    `json:"port"`
	Host    string `json:"host"`
	Timeout int    `json:"timeout"`
}

type AIConfig struct {
	ModelPath    string
	Threshold    float64
	MaxBatchSize int
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

var (
	config *Config
	once   sync.Once
)

func LoadConfig() (Config, error) {
	// Read environment variables and populate config
	port := os.Getenv("PORT")
	portNum, err := strconv.Atoi(port)
	if err != nil {
		return Config{}, err
	}

	return Config{
		Port:     portNum,
		Env:      os.Getenv("ENV"),
		LogLevel: os.Getenv("LOG_LEVEL"),
	}, nil
}

func Load() (*Config, error) {
	var err error
	once.Do(func() {
		config = &Config{}
		err = loadFromFile("config.json")
		if err != nil {
			return
		}
		loadFromEnv()
	})
	return config, err
}

func loadFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(config)
}

func loadFromEnv() {
	if port := os.Getenv("SERVER_PORT"); port != "" {
		if portNum, err := strconv.Atoi(port); err == nil {
			config.Server.Port = portNum
		}
	}
	if host := os.Getenv("SERVER_HOST"); host != "" {
		config.Server.Host = host
	}
	// Add more environment variables as needed
}
