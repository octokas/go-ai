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

func (m *MockCommandExecutor) CombinedOutput() ([]byte, error) {
	ret := m.Called()
	return ret.Get(0).([]byte), ret.Error(1)
}

func (m *MockCommandExecutor) Run() error {
	ret := m.Called()
	return ret.Error(0)
}

func TestReporter(t *testing.T) {
	mockExec := new(MockCommandExecutor)

	// Set up mock expectations
	mockExec.On("Command", "go", "test", "-coverprofile", mock.Anything, "./...").Return(exec.Command(""))
	mockExec.On("CombinedOutput").Return([]byte("test output"), nil)
	mockExec.On("Run").Return(nil)

	reporter := &Reporter{
		commandExecutor: mockExec,
	}

	// Call the function you want to test
	err := reporter.SaveTestReports()

	// Assert the expected behavior
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	// Verify the mock expectations
	mockExec.AssertExpectations(t)
}
