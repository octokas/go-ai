package fix_yaml_comments

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFixYAMLComments(t *testing.T) {
	// Create a temporary directory for test files
	tmpDir, err := os.MkdirTemp("", "yaml-test")
	assert.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	testCases := []struct {
		name     string
		content  string
		expected string
		isYAML   bool
	}{
		{
			name: "single comment",
			content: `---
name: Test
# Single comment
key: value`,
			expected: `---
name: Test
## Single comment
key: value`,
			isYAML: true,
		},
		{
			name: "multiple comments",
			content: `name: Test
# First comment
key1: value1
# Second comment
key2: value2`,
			expected: `name: Test
## First comment
key1: value1
## Second comment
key2: value2`,
			isYAML: true,
		},
		{
			name: "already double comments",
			content: `name: Test
## Already correct
key: value`,
			expected: `name: Test
## Already correct
key: value`,
			isYAML: true,
		},
		{
			name: "not yaml file",
			content: `This is not a YAML file
# with some comments
regular text`,
			expected: `This is not a YAML file
# with some comments
regular text`,
			isYAML: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create test file
			filename := filepath.Join(tmpDir, "test.yaml")
			err := os.WriteFile(filename, []byte(tc.content), 0644)
			assert.NoError(t, err)

			// Run the function
			wasChanged, err := fixYAMLComments(filename)
			assert.NoError(t, err)

			// Read the result
			result, err := os.ReadFile(filename)
			assert.NoError(t, err)

			// Verify the result
			if tc.isYAML {
				assert.Equal(t, tc.expected, string(result))
				assert.True(t, wasChanged || tc.content == tc.expected)
			} else {
				assert.Equal(t, tc.content, string(result))
				assert.False(t, wasChanged)
			}
		})
	}
} 