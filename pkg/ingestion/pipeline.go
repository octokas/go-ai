package vectorstore

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStore struct {
	client     *mongo.Client
	collection *mongo.Collection
}

type MongoDocument struct {
	ID        string                 `bson:"_id,omitempty"`
	Content   string                 `bson:"content"`
	Source    string                 `bson:"source"`
	Metadata  map[string]interface{} `bson:"metadata"`
	Embedding []float32              `bson:"embedding"`
	CreatedAt time.Time              `bson:"created_at"`
}

type Document struct {
	ID        string
	Content   string
	Source    string
	Metadata  map[string]interface{}
	Embedding []float32
}

type SearchResult struct {
	Document
	Score float32
}

func NewMongoStore(uri, dbName string) (*MongoStore, error) {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Create vector search index if it doesn't exist
	err = createVectorIndex(client.Database(dbName).Collection("documents"))
	if err != nil {
		return nil, err
	}

	return &MongoStore{
		client:     client,
		collection: client.Database(dbName).Collection("documents"),
	}, nil
}

func (s *MongoStore) Insert(docs []Document) error {
	var mongoDocs []interface{}
	for _, doc := range docs {
		mongoDoc := MongoDocument{
			Content:   doc.Content,
			Source:    doc.Source,
			Metadata:  doc.Metadata,
			Embedding: doc.Embedding,
			CreatedAt: time.Now(),
		}
		mongoDocs = append(mongoDocs, mongoDoc)
	}

	_, err := s.collection.InsertMany(context.Background(), mongoDocs)
	return err
}

func (s *MongoStore) Search(query string, limit int) ([]SearchResult, error) {
	// Note: 'query' here should be the embedding vector, not the raw text
	pipeline := mongo.Pipeline{
		{{Key: "$search", Value: bson.D{
			{Key: "index", Value: "vector_index"},
			{Key: "knnBeta", Value: bson.D{
				{Key: "vector", Value: query},
				{Key: "path", Value: "embedding"},
				{Key: "k", Value: limit},
			}},
		}}},
		{{Key: "$project", Value: bson.D{
			{Key: "score", Value: bson.D{{Key: "$meta", Value: "searchScore"}}},
			{Key: "content", Value: 1},
			{Key: "source", Value: 1},
			{Key: "metadata", Value: 1},
			{Key: "embedding", Value: 1},
		}}},
	}

	cursor, err := s.collection.Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var results []SearchResult
	for cursor.Next(context.Background()) {
		var mongoDoc struct {
			MongoDocument `bson:",inline"`
			Score         float32 `bson:"score"`
		}
		if err := cursor.Decode(&mongoDoc); err != nil {
			return nil, err
		}

		results = append(results, SearchResult{
			Document: Document{
				ID:        mongoDoc.ID,
				Content:   mongoDoc.Content,
				Source:    mongoDoc.Source,
				Metadata:  mongoDoc.Metadata,
				Embedding: mongoDoc.Embedding,
			},
			Score: mongoDoc.Score,
		})
	}

	return results, nil
}

func createVectorIndex(collection *mongo.Collection) error {
	indexModel := mongo.IndexModel{
		Keys: bson.D{{Key: "embedding", Value: "vector"}},
		Options: options.Index().
			SetName("vector_index").
			SetUnique(false).
			SetWeights(bson.D{
				{Key: "numDimensions", Value: 1536}, // For OpenAI embeddings
				{Key: "similarity", Value: "cosine"},
			}),
	}

	_, err := collection.Indexes().CreateOne(context.Background(), indexModel)
	return err
}
