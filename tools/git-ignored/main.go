package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type IgnoreResult struct {
	File     string `json:"file"`
	Ignored  bool   `json:"ignored"`
	Pattern  string `json:"pattern,omitempty"`
}

func isIgnored(path string, patterns []string) (bool, string) {
	cleanPath := filepath.ToSlash(path)

	for _, pattern := range patterns {
		pattern = strings.TrimSpace(pattern)
		if pattern == "" || strings.HasPrefix(pattern, "#") {
			continue
		}

		// Negation pattern
		if strings.HasPrefix(pattern, "!") {
			continue
		}

		// Directory-only pattern
		if strings.HasSuffix(pattern, "/") {
			dirPattern := strings.TrimSuffix(pattern, "/")
			if matchPattern(cleanPath, dirPattern) {
				return true, pattern
			}
			continue
		}

		// Glob pattern
		if matchPattern(cleanPath, pattern) {
			return true, pattern
		}
	}

	return false, ""
}

func matchPattern(path, pattern string) bool {
	// Handle leading slash (root-relative)
	if strings.HasPrefix(pattern, "/") {
		pattern = strings.TrimPrefix(pattern, "/")
		return matchGlob(path, pattern) || matchGlob(filepath.Base(path), pattern)
	}

	// Handle ** (match any directories)
	if strings.Contains(pattern, "**") {
		return matchDoubleStar(path, pattern)
	}

	// Pattern with slash - match from root
	if strings.Contains(pattern, "/") {
		return matchGlob(path, pattern)
	}

	// No slash - match any part of the path
	return matchGlob(filepath.Base(path), pattern) || matchGlob(path, pattern)
}

func matchGlob(path, pattern string) bool {
	matched, _ := filepath.Match(pattern, path)
	return matched
}

func matchDoubleStar(path, pattern string) bool {
	parts := strings.Split(pattern, "**")
	if len(parts) == 2 {
		prefix := strings.TrimSuffix(parts[0], "/")
		suffix := strings.TrimPrefix(parts[1], "/")

		if prefix == "" && suffix == "" {
			return true
		}
		if prefix == "" {
			return matchGlob(path, "*"+suffix) || matchGlob(filepath.Base(path), suffix)
		}
		if suffix == "" {
			return strings.HasPrefix(path, prefix+"/") || path == prefix
		}

		// Check if path starts with prefix
		if strings.HasPrefix(path, prefix+"/") {
			remaining := strings.TrimPrefix(path, prefix+"/")
			return matchGlob(remaining, "*"+suffix) || matchGlob(filepath.Base(remaining), suffix)
		}
	}
	return false
}

func parseGitignore(content string) []string {
	var patterns []string
	for _, line := range strings.Split(content, "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		patterns = append(patterns, line)
	}
	return patterns
}

func findFiles(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			rel, err := filepath.Rel(root, path)
			if err != nil {
				return err
			}
			files = append(files, rel)
		}
		return nil
	})
	return files, err
}

func main() {
	fs := flag.NewFlagSet("git-ignored", flag.ExitOnError)
	jsonFlag := fs.Bool("json", false, "output as JSON")
	gitignoreFlag := fs.String("gitignore", ".gitignore", "path to .gitignore file")
	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: git-ignored [flags] [directory]\n\n")
		fmt.Fprintf(os.Stderr, "Lists files that would be ignored by .gitignore\n\n")
		fmt.Fprintf(os.Stderr, "Flags:\n")
		fs.PrintDefaults()
	}
	fs.Parse(os.Args[1:])
	args := fs.Args()

	dir := "."
	if len(args) > 0 {
		dir = args[0]
	}

	gitignorePath := filepath.Join(dir, *gitignoreFlag)
	content, err := os.ReadFile(gitignorePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading %s: %v\n", gitignorePath, err)
		os.Exit(1)
	}

	patterns := parseGitignore(string(content))
	files, err := findFiles(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error walking directory: %v\n", err)
		os.Exit(1)
	}

	var results []IgnoreResult
	for _, f := range files {
		ignored, pattern := isIgnored(f, patterns)
		if ignored {
			results = append(results, IgnoreResult{
				File:    f,
				Ignored: true,
				Pattern: pattern,
			})
		}
	}

	if *jsonFlag {
		data, _ := json.MarshalIndent(results, "", "  ")
		fmt.Println(string(data))
	} else {
		if len(results) == 0 {
			fmt.Println("No ignored files found")
			return
		}
		for _, r := range results {
			fmt.Printf("%s (matched by '%s')\n", r.File, r.Pattern)
		}
		fmt.Printf("\n%d file(s) ignored\n", len(results))
	}
}
