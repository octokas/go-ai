package yaml_comments

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func isYAMLFile(firstLine string) bool {
	yamlStarters := []string{"---", "name:", "on:"}
	for _, starter := range yamlStarters {
		if strings.HasPrefix(firstLine, starter) {
			return true
		}
	}
	return false
}

func fixYAMLComments(filename string) (bool, error) {
	// Skip markdown files
	if ext := filepath.Ext(filename); ext == ".md" || ext == ".markdown" {
		return false, nil
	}

	content, err := os.ReadFile(filename)
	if err != nil {
		return false, err
	}

	// Check if file looks like YAML
	scanner := bufio.NewScanner(bytes.NewReader(content))
	if !scanner.Scan() || !isYAMLFile(scanner.Text()) {
		return false, nil
	}

	// Regex to match single # comments but not ## or # within text
	re := regexp.MustCompile(`(?m)^([^#]*[^#])#(?!#)\s`)
	modified := re.ReplaceAll(content, []byte("${1}## "))

	if !bytes.Equal(content, modified) {
		err = os.WriteFile(filename, modified, 0644)
		if err != nil {
			return false, err
		}
		return true, nil
	}

	return false, nil
}

func FixYAMLComments(filenames ...string) error {
	if len(filenames) == 0 {
		return fmt.Errorf("no files provided")
	}

	changed := false
	for _, filename := range filenames {
		wasChanged, err := fixYAMLComments(filename)
		if err != nil {
			return fmt.Errorf("error processing %s: %v", filename, err)
		}
		if wasChanged {
			fmt.Printf("Fixed comments in %s\n", filename)
			changed = true
		}
	}

	if changed {
		return fmt.Errorf("files were modified")
	}
	return nil
}
