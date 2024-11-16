package mocks

import (
	"github.com/stretchr/testify/mock"
)

// MockConfig represents a mock implementation of Config
type MockConfig struct {
	mock.Mock
}

func NewMockConfig() *MockConfig {
	return &MockConfig{
		Mock: mock.Mock{},
	}
}

func (m *MockConfig) GetDatabaseConfig() map[string]interface{} {
	args := m.Called()
	return args.Get(0).(map[string]interface{})
}

func (m *MockConfig) GetServerConfig() map[string]interface{} {
	args := m.Called()
	return args.Get(0).(map[string]interface{})
}
