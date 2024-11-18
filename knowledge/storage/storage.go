package storage

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"sync"
)

// StorageProvider defines the interface for document storage
type StorageProvider interface {
	Save(id string, content string) error
	Load(id string) (string, error)
	Delete(id string) error
	List() ([]string, error)
}

// FileStorage implements StorageProvider using the local filesystem
type FileStorage struct {
	basePath string
	mutex    sync.RWMutex
}

// NewFileStorage creates a new FileStorage instance
func NewFileStorage(basePath string) (*FileStorage, error) {
	// Create base directory if it doesn't exist
	if err := os.MkdirAll(basePath, 0755); err != nil {
		return nil, err
	}

	return &FileStorage{
		basePath: basePath,
	}, nil
}

// Save stores a document to the filesystem
func (fs *FileStorage) Save(id string, content string) error {
	fs.mutex.Lock()
	defer fs.mutex.Unlock()

	filePath := filepath.Join(fs.basePath, id+".json")

	doc := struct {
		Content string `json:"content"`
	}{
		Content: content,
	}

	data, err := json.Marshal(doc)
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, data, 0644)
}

// Load retrieves a document from the filesystem
func (fs *FileStorage) Load(id string) (string, error) {
	fs.mutex.RLock()
	defer fs.mutex.RUnlock()

	filePath := filepath.Join(fs.basePath, id+".json")

	data, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return "", errors.New("document not found")
		}
		return "", err
	}

	var doc struct {
		Content string `json:"content"`
	}

	if err := json.Unmarshal(data, &doc); err != nil {
		return "", err
	}

	return doc.Content, nil
}

// Delete removes a document from the filesystem
func (fs *FileStorage) Delete(id string) error {
	fs.mutex.Lock()
	defer fs.mutex.Unlock()

	filePath := filepath.Join(fs.basePath, id+".json")

	if err := os.Remove(filePath); err != nil {
		if os.IsNotExist(err) {
			return errors.New("document not found")
		}
		return err
	}

	return nil
}

// List returns all document IDs in storage
func (fs *FileStorage) List() ([]string, error) {
	fs.mutex.RLock()
	defer fs.mutex.RUnlock()

	files, err := os.ReadDir(fs.basePath)
	if err != nil {
		return nil, err
	}

	var ids []string
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".json" {
			// Remove the .json extension to get the ID
			id := file.Name()[:len(file.Name())-5]
			ids = append(ids, id)
		}
	}

	return ids, nil
}
