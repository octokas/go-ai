package mocks

import (
    "os"
    "github.com/stretchr/testify/mock"
)

type MockFileSystem struct {
    mock.Mock
}

func (m *MockFileSystem) MkdirAll(path string, perm os.FileMode) error {
    args := m.Called(path, perm)
    return args.Error(0)
}

func (m *MockFileSystem) WriteFile(filename string, data []byte, perm os.FileMode) error {
    args := m.Called(filename, data, perm)
    return args.Error(0)
} 