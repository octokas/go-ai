package dirmap

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type FileNode struct {
	Name     string              `json:"name"`
	Path     string              `json:"path"`
	Type     string              `json:"type"`
	Language string              `json:"language,omitempty"`
	Size     int64               `json:"size"`
	Children map[string]FileNode `json:"children,omitempty"`
}

// Language detection map
var LanguageMap = map[string]string{
	".go":    "Go",
	".js":    "JavaScript",
	".ts":    "TypeScript",
	".jsx":   "React/JavaScript",
	".tsx":   "React/TypeScript",
	".py":    "Python",
	".rb":    "Ruby",
	".java":  "Java",
	".html":  "HTML",
	".css":   "CSS",
	".scss":  "SCSS",
	".sql":   "SQL",
	".md":    "Markdown",
	".json":  "JSON",
	".yaml":  "YAML",
	".yml":   "YAML",
	".sh":    "Shell",
	".bash":  "Bash",
	".php":   "PHP",
	".rs":    "Rust",
	".cpp":   "C++",
	".c":     "C",
	".swift": "Swift",
	".kt":    "Kotlin",
}

func ScanDirectory(root string, path string) (FileNode, error) {
	info, err := os.Stat(path)
	if err != nil {
		return FileNode{}, err
	}

	node := FileNode{
		Name: filepath.Base(path),
		Path: strings.TrimPrefix(path, root),
		Size: info.Size(),
	}

	if info.IsDir() {
		node.Type = "directory"
		node.Children = make(map[string]FileNode)

		entries, err := os.ReadDir(path)
		if err != nil {
			return node, err
		}

		for _, entry := range entries {
			// Skip .git directory and debug directory
			if entry.Name() == ".git" || entry.Name() == "debug" {
				continue
			}

			childPath := filepath.Join(path, entry.Name())
			childNode, err := ScanDirectory(root, childPath)
			if err != nil {
				fmt.Printf("Error scanning %s: %v\n", childPath, err)
				continue
			}

			node.Children[entry.Name()] = childNode
		}
	} else {
		node.Type = "file"
		ext := strings.ToLower(filepath.Ext(path))
		if lang, ok := LanguageMap[ext]; ok {
			node.Language = lang
		}
	}

	return node, nil
}

func GenerateDirectoryMap() {
	// Get the current working directory
	root, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting working directory: %v\n", err)
		return
	}

	// Create debug directory if it doesn't exist
	debugDir := filepath.Join(root, "debug")
	if err := os.MkdirAll(debugDir, 0755); err != nil {
		fmt.Printf("Error creating debug directory: %v\n", err)
		return
	}

	// Scan the directory
	fileTree, err := ScanDirectory(root, root)
	if err != nil {
		fmt.Printf("Error scanning directory: %v\n", err)
		return
	}

	// Create the output file
	outputPath := filepath.Join(debugDir, "directory_map.json")
	file, err := os.Create(outputPath)
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		return
	}
	defer file.Close()

	// Create JSON encoder with indentation
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	// Write the JSON
	if err := encoder.Encode(fileTree); err != nil {
		fmt.Printf("Error encoding JSON: %v\n", err)
		return
	}

	fmt.Printf("Directory map has been saved to %s\n", outputPath)
}
