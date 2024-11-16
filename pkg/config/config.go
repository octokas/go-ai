package config

import (
	"encoding/json"
	"io"
	"os"
	"strconv"
	"strings"
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
	mutex  sync.Mutex
)

func Reset() {
	mutex.Lock()
	defer mutex.Unlock()
	config = nil
	once = sync.Once{}
}

func LoadFromReader(reader io.Reader) (*Config, error) {
	cfg := &Config{}
	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

func Load() (*Config, error) {
	var err error
	once.Do(func() {
		var cfg *Config
		cfg, err = loadConfig()
		if err == nil {
			config = cfg
		}
	})
	return config, err
}

func loadConfig() (*Config, error) {
	cfg := &Config{}
	if file, err := os.Open("config.json"); err == nil {
		defer file.Close()
		if err := json.NewDecoder(file).Decode(cfg); err != nil {
			return nil, err
		}
	}

	if err := loadFromEnv(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func loadFromEnv(cfg *Config) error {
	if host := os.Getenv("SERVER_HOST"); host != "" {
		cfg.Server.Host = host
	}
	if port := os.Getenv("SERVER_PORT"); port != "" {
		portNum, err := strconv.Atoi(port)
		if err != nil {
			return err
		}
		cfg.Server.Port = portNum
	}
	// Add more environment variables as needed
	return nil
}

func GetGitHubToken() string {
	return os.Getenv("GITHUB_TOKEN")
}

// Add this function to support testing
func LoadFromString(configStr string) error {
	_, err := LoadFromReader(strings.NewReader(configStr))
	return err
}
