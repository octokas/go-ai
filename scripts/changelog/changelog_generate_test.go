package changelog

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestChangelogGenerate(t *testing.T) {
	// Setup test directory
	tmpDir := t.TempDir()
	changelogsDir := filepath.Join(tmpDir, "changelogs")
	err := os.MkdirAll(changelogsDir, 0755)
	assert.NoError(t, err)

	// Mock git environment
	mockCommitID := "abcd1234"
	mockAuthor := "testuser"
	timestamp := time.Now().Format("20060102_150405")
	expectedFilename := filepath.Join(changelogsDir, "changelog_"+timestamp+"_"+mockCommitID[:8]+".md")

	t.Run("generate_changelog_file", func(t *testing.T) {
		// Run changelog generation
		changelog, err := generateChangelog(mockCommitID, mockAuthor)
		assert.NoError(t, err)
		assert.NotEmpty(t, changelog)

		// Write changelog to file
		err = os.WriteFile(expectedFilename, []byte(changelog), 0644)
		assert.NoError(t, err)

		// Verify file exists and content
		content, err := os.ReadFile(expectedFilename)
		assert.NoError(t, err)
		assert.NotEmpty(t, content)

		// Verify content structure
		contentStr := string(content)
		assert.Contains(t, contentStr, mockCommitID)
		assert.Contains(t, contentStr, mockAuthor)
		assert.Contains(t, contentStr, "# Changelog")
	})

	t.Run("verify_changelog_structure", func(t *testing.T) {
		// Verify directory exists
		_, err := os.Stat(changelogsDir)
		assert.NoError(t, err)

		// Verify file naming convention
		files, err := os.ReadDir(changelogsDir)
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, len(files), 1)

		// Check filename pattern
		filename := files[0].Name()
		assert.Regexp(t, `^changelog_\d{8}_\d{6}_[a-f0-9]{8}\.md$`, filename)
	})
}
