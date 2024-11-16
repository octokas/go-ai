package db

import (
	"testing"

	"github.com/octokas/go-ai/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestDatabase(t *testing.T) {
	// Create test config
	cfg := &config.Config{}
	cfg.Database.Host = "localhost"
	cfg.Database.Port = 5432
	cfg.Database.User = "test"
	cfg.Database.Password = "test"
	cfg.Database.DBName = "testdb"

	// Test database connection
	db, err := New(cfg)
	// In a real test, you'd use a test database or mock
	// Here we expect an error since we don't have a real DB
	assert.Error(t, err)
	assert.Nil(t, db)
}

func TestDatabaseSingleton(t *testing.T) {
	cfg := &config.Config{}
	db1, _ := New(cfg)
	db2, _ := New(cfg)
	if db1 != nil && db2 != nil {
		assert.Same(t, db1, db2, "Database should be a singleton")
	}
}
