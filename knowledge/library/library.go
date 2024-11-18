package library

import (
	"strings"
	"sync"

	"github.com/octokas/go-ai/knowledge/storage"
)

type Library struct {
	documents map[string]*Document
	storage   storage.StorageProvider
	mutex     sync.RWMutex
}

func NewLibrary(storage storage.StorageProvider) *Library {
	return &Library{
		documents: make(map[string]*Document),
		storage:   storage,
	}
}

func (l *Library) AddDocument(id string, content string) error {
	doc := NewDocument(content)

	l.mutex.Lock()
	defer l.mutex.Unlock()

	if err := l.storage.Save(id, content); err != nil {
		return err
	}

	l.documents[id] = doc
	return nil
}

func (l *Library) LoadAllDocuments() error {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	ids, err := l.storage.List()
	if err != nil {
		return err
	}

	for _, id := range ids {
		content, err := l.storage.Load(id)
		if err != nil {
			return err
		}

		l.documents[id] = NewDocument(content)
	}

	return nil
}

func (l *Library) Search(query string) []string {
	l.mutex.RLock()
	defer l.mutex.RUnlock()

	results := []string{}
	queryTerms := strings.Fields(strings.ToLower(query))

	for _, doc := range l.documents {
		if doc.MatchesQuery(queryTerms) {
			results = append(results, doc.GetRelevantExcerpts(queryTerms)...)
		}
	}

	return results
}
