package test_reporter

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func SaveTestReports() error {
	// Create reports directory if it doesn't exist
	reportsDir := "reports"
	if err := os.MkdirAll(reportsDir, 0755); err != nil {
		return fmt.Errorf("failed to create reports directory: %w", err)
	}

	// Generate timestamp for unique filenames
	timestamp := time.Now().Format("2006-01-02_15-04-05")

	// Run tests with coverage
	coverageFile := filepath.Join(reportsDir, fmt.Sprintf("coverage_%s.out", timestamp))
	coverageCmd := exec.Command("go", "test", "-coverprofile", coverageFile, "./...")
	coverageOutput, err := coverageCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to run tests with coverage: %w", err)
	}

	// Save test output
	testOutputFile := filepath.Join(reportsDir, fmt.Sprintf("test_output_%s.txt", timestamp))
	if err := os.WriteFile(testOutputFile, coverageOutput, 0644); err != nil {
		return fmt.Errorf("failed to save test output: %w", err)
	}

	// Generate HTML coverage report
	htmlFile := filepath.Join(reportsDir, fmt.Sprintf("coverage_%s.html", timestamp))
	htmlCmd := exec.Command("go", "tool", "cover", "-html", coverageFile, "-o", htmlFile)
	if err := htmlCmd.Run(); err != nil {
		return fmt.Errorf("failed to generate HTML coverage report: %w", err)
	}

	fmt.Printf("Reports generated:\n"+
		"- Test output: %s\n"+
		"- Coverage data: %s\n"+
		"- Coverage HTML: %s\n",
		testOutputFile, coverageFile, htmlFile)

	return nil
}
