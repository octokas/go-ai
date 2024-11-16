package mocks

import "github.com/stretchr/testify/mock"

type MockLogger struct {
    mock.Mock
}

func (m *MockLogger) Info(args ...interface{}) {
    m.Called(args...)
}

func (m *MockLogger) Error(args ...interface{}) {
    m.Called(args...)
}

func (m *MockLogger) Fatal(args ...interface{}) {
    m.Called(args...)
}

func (m *MockLogger) Debug(args ...interface{}) {
    m.Called(args...)
}

func (m *MockLogger) Warn(args ...interface{}) {
    m.Called(args...)
} 