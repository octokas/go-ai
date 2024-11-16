package mocks

import "github.com/stretchr/testify/mock"

type MockConfig struct {
    mock.Mock
}

func (m *MockConfig) GetServerConfig() interface{} {
    args := m.Called()
    return args.Get(0)
}

func (m *MockConfig) GetDatabaseConfig() interface{} {
    args := m.Called()
    return args.Get(0)
} 