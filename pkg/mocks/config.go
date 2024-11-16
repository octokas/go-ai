package mocks

import (
	"io"

	"github.com/octokas/go-ai/pkg/config"
	"github.com/stretchr/testify/mock"
)

type ConfigMock struct {
	mock.Mock
}

func NewConfigMock() *ConfigMock {
	return &ConfigMock{}
}

// Implement all methods from the Configer interface
func (m *ConfigMock) Load() (*config.Config, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*config.Config), args.Error(1)
}

func (m *ConfigMock) LoadFromReader(reader io.Reader) (*config.Config, error) {
	args := m.Called(reader)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*config.Config), args.Error(1)
}

func (m *ConfigMock) LoadFromString(configStr string) error {
	args := m.Called(configStr)
	return args.Error(0)
}

func (m *ConfigMock) GetGitHubToken() string {
	args := m.Called()
	return args.String(0)
}

func (m *ConfigMock) Reset() {
	m.Called()
}
