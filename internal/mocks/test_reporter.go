package mocks

import "github.com/stretchr/testify/mock"

type MockTestReporter struct {
	mock.Mock
}

func (m *MockTestReporter) SaveTestReports() error {
	args := m.Called()
	return args.Error(0)
}
