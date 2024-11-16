package mocks

import (
	"github.com/stretchr/testify/mock"
	"net/http"
)

type MockServer struct {
	mock.Mock
}

func (m *MockServer) Start() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockServer) Stop() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.Called(w, r)
} 