package vectorstore

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
