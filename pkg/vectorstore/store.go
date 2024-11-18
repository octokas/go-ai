package vectorstore

import (
	"fmt"
)

// Common interfaces and types
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

type Store interface {
	Insert(docs []Document) error
	Search(query string, limit int) ([]SearchResult, error)
	Close() error
}

// Factory function to create the appropriate store
func NewStore(config VectorStoreConfig) (Store, error) {
	if err := config.VectorValidate(); err != nil {
		return nil, err
	}

	switch config.Type {
	case "mongodb":
		return NewMongoStore(config.MongoURI, config.MongoDatabase)
	case "postgres":
		return NewPostgresStore(config.PostgresURI, config.PostgresDatabase)
	default:
		return nil, fmt.Errorf("unsupported store type: %s", config.Type)
	}
}
