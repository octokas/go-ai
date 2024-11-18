package mongo

// import (
// 	"context"
// 	"fmt"
// 	"time"

// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// type MongoStore struct {
// 	client     *mongo.Client
// 	collection *mongo.Collection
// }

// type MongoDocument struct {
// 	ID        string                 `bson:"_id,omitempty"`
// 	Content   string                 `bson:"content"`
// 	Source    string                 `bson:"source"`
// 	Metadata  map[string]interface{} `bson:"metadata"`
// 	Embedding []float32              `bson:"embedding"`
// 	CreatedAt time.Time              `bson:"created_at"`
// }

// func NewMongoStore(uri, database string) (*MongoStore, error) {
// 	ctx := context.Background()
// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
// 	}

// 	store := &MongoStore{
// 		client:     client,
// 		collection: client.Database(database).Collection("documents"),
// 	}

// 	if err := store.createIndexes(); err != nil {
// 		client.Disconnect(context.Background())
// 		return nil, err
// 	}

// 	return store, nil
// }

// func (s *MongoStore) createIndexes() error {
// 	ctx := context.Background()
// 	_, err := s.collection.Indexes().CreateOne(ctx, mongo.IndexModel{
// 		Keys: map[string]interface{}{
// 			"embedding": "2dsphere",
// 		},
// 	})
// 	return err
// }

// // Implement Store interface
// func (s *MongoStore) Insert(docs []Document) error {
// 	// Implementation here
// 	return nil
// }

// func (s *MongoStore) Search(query string, limit int) ([]SearchResult, error) {
// 	// Implementation here
// 	return nil, nil
// }

// func (s *MongoStore) Close() error {
// 	return s.client.Disconnect(context.Background())
// }
