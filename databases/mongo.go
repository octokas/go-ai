package databases

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	UsersCollection    = "users"
	TasksCollection    = "tasks"
	AssetsCollection   = "assets"
	CalendarCollection = "calendar"
	ReportsCollection  = "reports"
)

func (db *Database) Collection(name string) *mongo.Collection {
	return db.MongoDB.Collection(name)
}

func (db *Database) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return db.MongoDB.Client().Ping(ctx, nil)
}
