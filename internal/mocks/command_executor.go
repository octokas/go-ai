package mocks

import "github.com/stretchr/testify/mock"

type MockCommandExecutor struct {
	mock.Mock
}

func (m *MockCommandExecutor) Command(name string, args ...string) ([]byte, error) {
	callArgs := append([]interface{}{name}, args)
	returnArgs := m.Called(callArgs...)
	return returnArgs.Get(0).([]byte), returnArgs.Error(1)
}
