package fix_yaml_comments

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

	// Fix: Simpler regex that matches single # at start of line
	re := regexp.MustCompile(`(?m)^#([^#])`)
	modified := re.ReplaceAllString(string(content), "##$1")

	if string(content) != modified {
		err = os.WriteFile(filename, []byte(modified), 0644)
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
