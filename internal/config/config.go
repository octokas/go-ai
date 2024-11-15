package config

import (
	"encoding/json"
	"os"
	"sync"
)

type Config struct {
	Server struct {
		Port    int    `json:"port"`
		Host    string `json:"host"`
		Timeout int    `json:"timeout"`
	} `json:"server"`

	Database struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		DBName   string `json:"dbname"`
	} `json:"database"`

	AI struct {
		ModelPath    string  `json:"model_path"`
		Threshold    float64 `json:"threshold"`
		MaxBatchSize int     `json:"max_batch_size"`
	} `json:"ai"`
}

var (
	config *Config
	once   sync.Once
)

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
		config.Server.Port = 8080 // Convert string to int in real implementation
	}
	if host := os.Getenv("SERVER_HOST"); host != "" {
		config.Server.Host = host
	}
	// Add more environment variables as needed
} 