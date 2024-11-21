package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Logging  LoggingConfig  `mapstructure:"logging"`
	Security SecurityConfig `mapstructure:"security"`
	API      APIConfig      `mapstructure:"api"`
}

type ServerConfig struct {
	Port         int           `mapstructure:"port"`
	Host         string        `mapstructure:"host"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
	Environment  string        `mapstructure:"environment"`
}

type DatabaseConfig struct {
	SQLite SQLiteConfig `mapstructure:"sqlite"`
	Mongo  MongoConfig  `mapstructure:"mongo"`
}

type SQLiteConfig struct {
	Path string `mapstructure:"path"`
}

type MongoConfig struct {
	URI      string `mapstructure:"uri"`
	Database string `mapstructure:"database"`
}

type LoggingConfig struct {
	Level      string `mapstructure:"level"`
	Format     string `mapstructure:"format"`
	OutputPath string `mapstructure:"output_path"`
}

type SecurityConfig struct {
	JWTSecret     string        `mapstructure:"jwt_secret"`
	TokenDuration time.Duration `mapstructure:"token_duration"`
}

type APIConfig struct {
	Version     string        `mapstructure:"version"`
	RateLimit   int           `mapstructure:"rate_limit"`
	RateWindow  time.Duration `mapstructure:"rate_window"`
	GraphQLPath string        `mapstructure:"graphql_path"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("reading config file: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unmarshaling config: %w", err)
	}

	return &config, nil
}
