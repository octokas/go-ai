package test_reporter

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

var (
	execCommand = exec.Command
	fs          = &realFileSystem{}
)

type TestReporter interface {
	SaveTestReports() error
}

type FileSystem interface {
	MkdirAll(path string, perm os.FileMode) error
	WriteFile(filename string, data []byte, perm os.FileMode) error
}

type realFileSystem struct{}

func (fs *realFileSystem) MkdirAll(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

func (fs *realFileSystem) WriteFile(filename string, data []byte, perm os.FileMode) error {
	return os.WriteFile(filename, data, perm)
}

// type TestReporter interface {
// 	SaveTestReports() error
// }

type CommandExecutor interface {
	CombinedOutput() ([]byte, error)
	Run() error
}

type Reporter struct {
	commandExecutor CommandExecutor
}

func (r *Reporter) SaveTestReports() error {
	// Create reports directory if it doesn't exist
	reportsDir := "reports"
	if err := fs.MkdirAll(reportsDir, 0755); err != nil {
		return fmt.Errorf("failed to create reports directory: %w", err)
	}

	// Generate timestamp for unique filenames
	timestamp := time.Now().Format("2006-01-02_15-04-05")

	// Run tests with coverage
	coverageFile := filepath.Join(reportsDir, fmt.Sprintf("coverage_%s.out", timestamp))
	coverageCmd := r.commandExecutor
	coverageOutput, err := coverageCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to run tests with coverage: %w", err)
	}

	// Save test output
	testOutputFile := filepath.Join(reportsDir, fmt.Sprintf("test_output_%s.txt", timestamp))
	if err := fs.WriteFile(testOutputFile, coverageOutput, 0644); err != nil {
		return fmt.Errorf("failed to save test output: %w", err)
	}

	// Generate HTML coverage report
	htmlFile := filepath.Join(reportsDir, fmt.Sprintf("coverage_%s.html", timestamp))
	htmlCmd := r.commandExecutor
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

// NewReporter creates a new Reporter instance with the specified test command
func NewReporter(testCommand string, args ...string) *Reporter {
	return &Reporter{
		commandExecutor: execCommand(testCommand, args...),
	}
}

// Example usage would be:
// reporter := NewReporter("go", "test", "-coverprofile=coverage.out", "./...")
