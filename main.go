package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: FileConcat <file1> <file2> ... or wildcard pattern like **/*.swift")
		return
	}

	files := os.Args[1:]
	var fileList []string

	// Handle wildcards
	for _, pattern := range files {
		matches, err := filepath.Glob(pattern)
		if err != nil {
			fmt.Printf("Error with pattern %s: %v\n", pattern, err)
			continue
		}
		fileList = append(fileList, matches...)
	}

	// Concatenate files while omitting comments
	var output strings.Builder
	commentRegex := regexp.MustCompile(`(?s)//.*?$|/\*.*?\*/`)

	for _, file := range fileList {
		content, err := os.ReadFile(file)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", file, err)
			continue
		}

		scanner := bufio.NewScanner(strings.NewReader(string(content)))
		for scanner.Scan() {
			line := scanner.Text()
			// Remove inline comments
			line = commentRegex.ReplaceAllString(line, "")
			// Only add lines that are not empty after comment removal
			if strings.TrimSpace(line) != "" {
				output.WriteString(line)
				output.WriteString("\n")
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Printf("Error scanning file %s: %v\n", file, err)
			continue
		}

		// Add a couple of newlines between files
		output.WriteString("\n\n")
	}

	// Print the concatenated result to STDOUT
	fmt.Print(output.String())
}
