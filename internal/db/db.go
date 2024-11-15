package db

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/octokas/go-ai/internal/config"
	"github.com/octokas/go-ai/internal/logger"
)

type Database struct {
	db     *sql.DB
	logger *logger.Logger
}

var (
	instance *Database
	once     sync.Once
)

func New(cfg *config.Config) (*Database, error) {
	var err error
	once.Do(func() {
		log := logger.New()

		dsn := fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			cfg.Database.Host,
			cfg.Database.Port,
			cfg.Database.User,
			cfg.Database.Password,
			cfg.Database.DBName,
		)

		db, err := sql.Open("postgres", dsn)
		if err != nil {
			return
		}

		err = db.Ping()
		if err != nil {
			return
		}

		instance = &Database{
			db:     db,
			logger: log,
		}
	})

	return instance, err
}

func (d *Database) Close() error {
	return d.db.Close()
}

// Example query method
func (d *Database) Query(query string, args ...interface{}) (*sql.Rows, error) {
	d.logger.Debug("Executing query:", query)
	return d.db.Query(query, args...)
}

// Example exec method
func (d *Database) Exec(query string, args ...interface{}) (sql.Result, error) {
	d.logger.Debug("Executing statement:", query)
	return d.db.Exec(query, args...)
}
