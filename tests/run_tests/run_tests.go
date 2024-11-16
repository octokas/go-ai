package run_tests

import (
	"log"

	test_reporter "github.com/octokas/go-ai/tests/test_reporter"
)

var testReporter test_reporter.TestReporter = test_reporter.NewReporter("go", "test", "-coverprofile=coverage.out", "./...")

func RunTests() error {
	if err := testReporter.SaveTestReports(); err != nil {
		log.Printf("Failed to save test reports: %v", err)
		return err
	}
	return nil
}
