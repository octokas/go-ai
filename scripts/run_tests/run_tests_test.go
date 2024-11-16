package run_tests

import (
	"testing"

	"github.com/octokas/go-ai/internal/mocks"
	"github.com/stretchr/testify/assert"
)

func TestRunTests(t *testing.T) {
	mockReporter := new(mocks.MockTestReporter)

	// Set expectations
	mockReporter.On("SaveTestReports").Return(nil)

	// Override the test reporter
	originalReporter := testReporter
	defer func() {
		testReporter = originalReporter
		assert.Equal(t, originalReporter, testReporter, "testReporter should be restored after test")
	}()
	testReporter = mockReporter

	// Run tests
	err := RunTests()

	// Verify expectations and results
	assert.NoError(t, err, "RunTests should not return an error")
	assert.True(t, mockReporter.AssertNumberOfCalls(t, "SaveTestReports", 1), "SaveTestReports should be called exactly once")
	mockReporter.AssertExpectations(t)
}
