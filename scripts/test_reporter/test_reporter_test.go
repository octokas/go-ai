package test_reporter

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockCommandExecutor struct {
	mock.Mock
}

func (m *MockCommandExecutor) Command(name string, args ...string) *exec.Cmd {
	args = append([]string{name}, args...)
	iargs := make([]interface{}, len(args))
	for i, v := range args {
		iargs[i] = v
	}
	ret := m.Called(iargs...)
	return ret.Get(0).(*exec.Cmd)
}

func TestReporter(t *testing.T) {
	// Create a new mock command executor
	mockExec := new(MockCommandExecutor)

	// Set up the mock expectation
	mockExec.On("Command", "go", "test", "-coverprofile", mock.Anything, "./...").Return(exec.Command(""), nil)

	// Create a new reporter instance
	reporter := NewReporter(mockExec)

	// Call the function you want to test
	err := reporter.RunTests()

	// Assert the expected behavior
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	// Verify the mock expectations
	mockExec.AssertExpectations(t)
}
