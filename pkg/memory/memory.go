package memory

import (
	"time"

	"github.com/google/uuid"
)

type MemoryStore struct {
	documents map[string]Document
}

type Document struct {
	ID      string
	Content string
	Source  string
	Created time.Time
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		documents: make(map[string]Document),
	}
}

func (m *MemoryStore) Upload(content, source string) (string, error) {
	id := uuid.New().String()
	m.documents[id] = Document{
		ID:      id,
		Content: content,
		Source:  source,
		Created: time.Now(),
	}
	return id, nil
}

func (m *MemoryStore) Get(id string) (Document, bool) {
	doc, exists := m.documents[id]
	return doc, exists
}
