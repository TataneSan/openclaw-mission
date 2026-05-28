package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

type vetResult struct {
	PackagePath string   `json:"package"`
	Issues      []string `json:"issues"`
}

func parseVetOutput(output string) []vetResult {
	var results []vetResult
	var current *vetResult

	// go vet outputs:
	//   # package/path          (compilation info, skip)
	//   # [package/path]        (issues follow for this package)
	//   ./file.go:1:2: message  (issue line)
	issuePkgRe := regexp.MustCompile(`^\[\s*(\S+?)\s*\]$`)
	plainPkgRe := regexp.MustCompile(`^(\S+)$`)

	lines := strings.Split(strings.TrimSpace(output), "\n")
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			continue
		}

		// Check if this is a "# [pkg]" line (issue header)
		if strings.HasPrefix(trimmed, "#") {
			afterHash := strings.TrimSpace(trimmed[1:])

			if bracketMatch := issuePkgRe.FindStringSubmatch(afterHash); bracketMatch != nil {
				if current != nil && len(current.Issues) > 0 {
					results = append(results, *current)
				}
				current = &vetResult{
					PackagePath: bracketMatch[1],
					Issues:      []string{},
				}
				continue
			}

			// Plain "# pkg" line - just note the package if no bracket header seen yet
			if plainMatch := plainPkgRe.FindStringSubmatch(afterHash); plainMatch != nil {
				if current == nil {
					current = &vetResult{
						PackagePath: plainMatch[1],
						Issues:      []string{},
					}
				}
			}
			continue
		}

		if current != nil {
			current.Issues = append(current.Issues, trimmed)
		}
	}

	if current != nil && len(current.Issues) > 0 {
		results = append(results, *current)
	}

	return results
}

func runVet(dir string, packages []string) (string, error) {
	args := append([]string{"vet"}, packages...)
	cmd := exec.Command("go", args...)
	cmd.Dir = dir
	cmd.Env = append(os.Environ(), "GO111MODULE=on")

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	output := stdout.String()
	if stderr.Len() > 0 {
		if output != "" {
			output += "\n" + stderr.String()
		} else {
			output = stderr.String()
		}
	}

	if err != nil && output == "" {
		return "", fmt.Errorf("go vet failed: %w", err)
	}

	return output, nil
}

func printReport(results []vetResult, format string, color bool) {
	if len(results) == 0 {
		if color {
			fmt.Println("\033[32mok  no issues found\033[0m")
		} else {
			fmt.Println("ok  no issues found")
		}
		return
	}

	totalIssues := 0
	for _, r := range results {
		totalIssues += len(r.Issues)
	}

	if format == "json" {
		data, _ := json.MarshalIndent(results, "", "  ")
		fmt.Println(string(data))
		return
	}

	if color {
		fmt.Printf("\033[31mFAIL\033[0m  %d issue%s in %d package%s\n\n",
			totalIssues, plural(totalIssues), len(results), plural(len(results)))
	} else {
		fmt.Printf("FAIL  %d issue%s in %d package%s\n\n",
			totalIssues, plural(totalIssues), len(results), plural(len(results)))
	}

	for _, r := range results {
		if color {
			fmt.Printf("\033[1m%s\033[0m:\n", r.PackagePath)
		} else {
			fmt.Printf("%s:\n", r.PackagePath)
		}
		for _, issue := range r.Issues {
			if color {
				fmt.Printf("  \033[33m%s\033[0m\n", issue)
			} else {
				fmt.Printf("  %s\n", issue)
			}
		}
		fmt.Println()
	}
}

func plural(n int) string {
	if n == 1 {
		return ""
	}
	return "s"
}

func printHelp() {
	fmt.Println(`go-vet-check — run go vet and format the report

USAGE:
    go-vet-check [OPTIONS] [PACKAGES]

OPTIONS:
    -f, --format FORMAT    Output format: text (default), json
    -n, --no-color         Disable colored output
    -d, --dir DIRECTORY    Run in specified directory (default: .)
    -h, --help             Show this help message

EXIT CODES:
    0    No issues found
    1    Issues found
    2    Internal error (go vet failed to run)

EXAMPLES:
    go-vet-check
        Run go vet on current directory.

    go-vet-check ./...
        Run go vet on all packages.

    go-vet-check -f json
        Run go vet with JSON output.

    go-vet-check -d ./myproject ./...
        Run go vet in a specific directory.

NOTES:
    This tool wraps 'go vet' and provides formatted output with
    package grouping and optional color support. It is useful in
    CI pipelines for clearer vet reports.
`)
}

func main() {
	dir := "."
	format := "text"
	color := true
	packages := []string{"."}
	args := os.Args[1:]

	for i := 0; i < len(args); i++ {
		arg := args[i]
		switch arg {
		case "-h", "--help":
			printHelp()
			os.Exit(0)
		case "-f", "--format":
			if i+1 < len(args) {
				format = args[i+1]
				i++
			}
		case "-n", "--no-color":
			color = false
		case "-d", "--dir":
			if i+1 < len(args) {
				dir = args[i+1]
				i++
			}
		default:
			if !strings.HasPrefix(arg, "-") {
				packages = append(packages, arg)
			} else {
				fmt.Fprintf(os.Stderr, "unknown flag: %s\n", arg)
				os.Exit(2)
			}
		}
	}

	absDir, err := filepath.Abs(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(2)
	}

	output, err := runVet(absDir, packages)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(2)
	}

	if strings.TrimSpace(output) == "" {
		printReport(nil, format, color)
		os.Exit(0)
	}

	results := parseVetOutput(output)
	printReport(results, format, color)

	if len(results) > 0 {
		os.Exit(1)
	}
}
