package databases

import (
	"context"
	"time"

	"go-kas/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	MongoDB *mongo.Database
	SQLite  *SQLiteDB
}

func NewDatabase(cfg *config.Config) (*Database, error) {
	// MongoDB setup
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoDB.URI))
	if err != nil {
		return nil, err
	}

	// SQLite setup
	sqliteDB, err := NewSQLiteDB()
	if err != nil {
		return nil, err
	}

	return &Database{
		MongoDB: mongoClient.Database(cfg.MongoDB.DBName),
		SQLite:  sqliteDB,
	}, nil
}
