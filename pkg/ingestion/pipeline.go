package vectorstore

import (
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

func NewMongoStore(uri, dbName string) (*MongoStore, error) {
	client, err := mongo.Connect(nil, options.Client().ApplyURI(uri))
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

	_, err := s.collection.InsertMany(nil, mongoDocs)
	return err
}

func (s *MongoStore) Search(query string, limit int) ([]SearchResult, error) {
	// Note: 'query' here should be the embedding vector, not the raw text
	pipeline := mongo.Pipeline{
		{{
			"$search": bson.D{
				{"index", "vector_index"},
				{"knnBeta", bson.D{
					{"vector", query},
					{"path", "embedding"},
					{"k", limit},
				}},
			},
		}},
		{{
			"$project": bson.D{
				{"score", bson.D{{"$meta", "searchScore"}}},
				{"content", 1},
				{"source", 1},
				{"metadata", 1},
				{"embedding", 1},
			},
		}},
	}

	cursor, err := s.collection.Aggregate(nil, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(nil)

	var results []SearchResult
	for cursor.Next(nil) {
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
		Keys: bson.D{{"embedding", "vector"}},
		Options: options.Index().
			SetName("vector_index").
			SetUnique(false).
			SetWeights(bson.D{
				{"numDimensions", 1536}, // For OpenAI embeddings
				{"similarity", "cosine"},
			}),
	}

	_, err := collection.Indexes().CreateOne(nil, indexModel)
	return err
}
