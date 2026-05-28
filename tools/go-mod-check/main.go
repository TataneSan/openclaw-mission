package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

type ModuleInfo struct {
	Path       string   `json:"path"`
	GoVersion  string   `json:"go_version"`
	Requires   []string `json:"requires"`
	Replaces   []string `json:"replaces"`
	HasSum     bool     `json:"has_sum"`
	HasVendor  bool     `json:"has_vendor"`
	Issues     []string `json:"issues"`
}

func checkGoMod(path string) (*ModuleInfo, error) {
	info := &ModuleInfo{}

	// Read go.mod
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	inRequire := false
	inReplace := false

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		if strings.HasPrefix(trimmed, "module ") {
			parts := strings.SplitN(trimmed, " ", 2)
			if len(parts) == 2 {
				info.Path = strings.TrimSpace(parts[1])
			}
		}

		if strings.HasPrefix(trimmed, "go ") {
			parts := strings.SplitN(trimmed, " ", 2)
			if len(parts) == 2 {
				info.GoVersion = strings.TrimSpace(parts[1])
			}
		}

		if strings.HasPrefix(trimmed, "require (") {
			inRequire = true
			continue
		}
		if inRequire && trimmed == ")" {
			inRequire = false
			continue
		}
		if inRequire && trimmed != "" {
			parts := strings.Fields(trimmed)
			if len(parts) >= 1 {
				info.Requires = append(info.Requires, parts[0])
			}
		}
		if strings.HasPrefix(trimmed, "require ") && !strings.HasSuffix(trimmed, "(") {
			parts := strings.Fields(trimmed)
			if len(parts) >= 2 {
				info.Requires = append(info.Requires, parts[1])
			}
		}

		if strings.HasPrefix(trimmed, "replace (") {
			inReplace = true
			continue
		}
		if inReplace && trimmed == ")" {
			inReplace = false
			continue
		}
		if inReplace && trimmed != "" {
			parts := strings.Fields(trimmed)
			if len(parts) >= 3 {
				info.Replaces = append(info.Replaces, fmt.Sprintf("%s -> %s", parts[0], parts[2]))
			}
		}
		if strings.HasPrefix(trimmed, "replace ") {
			parts := strings.Fields(trimmed)
			if len(parts) >= 3 {
				info.Replaces = append(info.Replaces, fmt.Sprintf("%s -> %s", parts[1], parts[3]))
			}
		}
	}

	// Check for go.sum
	dir := filepath.Dir(path)
	if _, err := os.Stat(filepath.Join(dir, "go.sum")); err == nil {
		info.HasSum = true
	}

	// Check for vendor directory
	if _, err := os.Stat(filepath.Join(dir, "vendor")); err == nil {
		info.HasVendor = true
	}

	// Check issues
	if !info.HasSum && len(info.Requires) > 0 {
		info.Issues = append(info.Issues, "missing go.sum file")
	}

	if info.Path == "" {
		info.Issues = append(info.Issues, "no module path declared")
	}

	if info.GoVersion == "" {
		info.Issues = append(info.Issues, "no go version declared")
	}

	// Check for unused replace directives
	if len(info.Replaces) > 0 {
		for _, replace := range info.Replaces {
			from := strings.Fields(replace)[0]
			found := false
			for _, req := range info.Requires {
				if req == from {
					found = true
					break
				}
			}
			if !found {
				info.Issues = append(info.Issues, fmt.Sprintf("replace for unused module: %s", from))
			}
		}
	}

	return info, nil
}

func checkGoFiles(dir string) []string {
	var issues []string
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".go") {
			fset := token.NewFileSet()
			_, err := parser.ParseFile(fset, path, nil, parser.ImportsOnly)
			if err != nil {
				issues = append(issues, fmt.Sprintf("parse error in %s: %v", filepath.Rel(dir, path), err))
			}
		}
		return nil
	})
	return issues
}

func printText(info *ModuleInfo, goIssues []string) {
	fmt.Printf("Module: %s\n", info.Path)
	fmt.Printf("Go version: %s\n", info.GoVersion)
	fmt.Printf("Requires: %d\n", len(info.Requires))
	fmt.Printf("Replaces: %d\n", len(info.Replaces))
	fmt.Printf("Has go.sum: %v\n", info.HasSum)
	fmt.Printf("Has vendor: %v\n", info.HasVendor)

	if len(info.Issues) > 0 || len(goIssues) > 0 {
		fmt.Printf("\nIssues:\n")
		for _, issue := range info.Issues {
			fmt.Printf("  - %s\n", issue)
		}
		for _, issue := range goIssues {
			fmt.Printf("  - %s\n", issue)
		}
	} else {
		fmt.Printf("\nNo issues found\n")
	}
}

func main() {
	fs := flag.NewFlagSet("go-mod-check", flag.ExitOnError)
	jsonFlag := fs.Bool("json", false, "output as JSON")
	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: go-mod-check [flags] [directory]\n\n")
		fmt.Fprintf(os.Stderr, "Checks go.mod for issues and validates Go source files.\n\n")
		fmt.Fprintf(os.Stderr, "Flags:\n")
		fs.PrintDefaults()
	}
	fs.Parse(os.Args[1:])
	args := fs.Args()

	dir := "."
	if len(args) > 0 {
		dir = args[0]
	}

	goModPath := filepath.Join(dir, "go.mod")
	info, err := checkGoMod(goModPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading go.mod: %v\n", err)
		os.Exit(1)
	}

	goIssues := checkGoFiles(dir)
	info.Issues = append(info.Issues, goIssues...)

	if *jsonFlag {
		data, _ := json.MarshalIndent(info, "", "  ")
		fmt.Println(string(data))
	} else {
		printText(info, goIssues)
	}

	if len(info.Issues) > 0 {
		os.Exit(1)
	}
}
