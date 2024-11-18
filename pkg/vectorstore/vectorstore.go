package vectorstore

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

// type VectorStoreConfig struct {
// 	Type string // "mongodb" or "postgres"

// 	// MongoDB specific
// 	MongoURI      string
// 	MongoDatabase string

// 	// Postgres specific
// 	PostgresURI      string
// 	PostgresDatabase string
// }

// func LoadVectorStoreConfig() VectorStoreConfig {
// 	return VectorStoreConfig{
// 		Type: getVectorStoreType(),

// 		// MongoDB config
// 		MongoURI:      os.Getenv("MONGO_URI"),
// 		MongoDatabase: os.Getenv("MONGO_DATABASE"),

// 		// Postgres config
// 		PostgresURI:      os.Getenv("POSTGRES_URI"),
// 		PostgresDatabase: os.Getenv("POSTGRES_DATABASE"),
// 	}
// }

// func (c VectorStoreConfig) Validate() error {
// 	switch c.Type {
// 	case "mongodb":
// 		if c.MongoURI == "" {
// 			return fmt.Errorf("MONGO_URI is required when STORE_TYPE is mongodb")
// 		}
// 		if c.MongoDatabase == "" {
// 			return fmt.Errorf("MONGO_DATABASE is required when STORE_TYPE is mongodb")
// 		}
// 	case "postgres":
// 		if c.PostgresURI == "" {
// 			return fmt.Errorf("POSTGRES_URI is required when STORE_TYPE is postgres")
// 		}
// 		if c.PostgresDatabase == "" {
// 			return fmt.Errorf("POSTGRES_DATABASE is required when STORE_TYPE is postgres")
// 		}
// 	default:
// 		return fmt.Errorf("unsupported store type: %s", c.Type)
// 	}
// 	return nil
// }

// func (c VectorStoreConfig) String() string {
// 	return fmt.Sprintf(
// 		"Vector Store Configuration:\n"+
// 			"  Type: %s\n"+
// 			"  MongoDB URI: %s\n"+
// 			"  MongoDB Database: %s\n"+
// 			"  Postgres URI: %s\n"+
// 			"  Postgres Database: %s",
// 		c.Type,
// 		maskURI(c.MongoURI),
// 		c.MongoDatabase,
// 		maskURI(c.PostgresURI),
// 		c.PostgresDatabase,
// 	)
// }

// func getVectorStoreType() string {
// 	storeType := os.Getenv("STORE_TYPE")
// 	if storeType == "" {
// 		return "memory" // default value
// 	}
// 	return strings.ToLower(storeType)
// }

// func maskURI(uri string) string {
// 	if uri == "" {
// 		return ""
// 	}
// 	return "[MASKED]"
// }

// MemoryStore provides a simple in-memory implementation of Store
type MemoryStore struct {
	documents []Document
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		documents: make([]Document, 0),
	}
}

func (m *MemoryStore) Insert(docs []Document) error {
	m.documents = append(m.documents, docs...)
	return nil
}

func (m *MemoryStore) Search(query string, limit int) ([]SearchResult, error) {
	results := make([]SearchResult, 0, len(m.documents))
	for _, doc := range m.documents {
		results = append(results, SearchResult{
			Document: doc,
			Score:    1.0, // Simple implementation returns all docs with same score
		})
	}
	if len(results) > limit {
		results = results[:limit]
	}
	return results, nil
}

func (m *MemoryStore) Close() error {
	m.documents = nil
	return nil
}
