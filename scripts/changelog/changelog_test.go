package changelog

import (
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateChangelog(t *testing.T) {
	// Setup: Create a temporary git repository with some test commits
	tmpDir, err := os.MkdirTemp("", "changelog-test")
	assert.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	// Initialize git repo
	err = os.Chdir(tmpDir)
	assert.NoError(t, err)
	
	commands := [][]string{
		{"git", "init"},
		{"git", "config", "user.email", "test@example.com"},
		{"git", "config", "user.name", "Test User"},
		{"git", "commit", "--allow-empty", "-m", "Initial commit"},
		{"git", "tag", "v0.1.0"},
		{"git", "commit", "--allow-empty", "-m", "feat: add new feature"},
		{"git", "commit", "--allow-empty", "-m", "fix: fix bug"},
		{"git", "commit", "--allow-empty", "-m", "chore: update docs"},
	}

	for _, cmd := range commands {
		out, err := exec.Command(cmd[0], cmd[1:]...).CombinedOutput()
		assert.NoError(t, err, "Command failed: %s\nOutput: %s", cmd, out)
	}

	// Test the changelog generation
	// Note: Since GenerateChangelog() prints to stdout, 
	// you might want to modify the function to return a string instead
	// for better testability
	GenerateChangelog() // This will print to stdout
} 