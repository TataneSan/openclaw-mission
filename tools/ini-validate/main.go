// ini-validate validates the syntax of INI configuration files.
//
// It checks for common issues like duplicate sections, duplicate keys,
// keys outside sections (if --strict), and malformed lines.
//
// Usage:
//   ini-validate config.ini
//   ini-validate -s config.ini    # strict mode: no keys outside sections
//   cat config.ini | ini-validate
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type issue struct {
	line    int
	message string
}

func validate(r io.Reader, strict bool) ([]issue, error) {
	var issues []issue
	scanner := bufio.NewScanner(r)

	sections := make(map[string]bool)
	currentSection := ""
	keyInSection := make(map[string]map[string]bool)

	lineNum := 0
	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, ";") {
			continue
		}

		// Section header
		if strings.HasPrefix(line, "[") {
			end := strings.Index(line, "]")
			if end == -1 {
				issues = append(issues, issue{lineNum, "missing closing bracket ']' for section"})
				continue
			}
			sectionName := strings.TrimSpace(line[1:end])
			if sectionName == "" {
				issues = append(issues, issue{lineNum, "empty section name"})
				continue
			}
			if sections[sectionName] {
				issues = append(issues, issue{lineNum, fmt.Sprintf("duplicate section [%s]", sectionName)})
			}
			sections[sectionName] = true
			currentSection = sectionName
			if _, ok := keyInSection[sectionName]; !ok {
				keyInSection[sectionName] = make(map[string]bool)
			}
			continue
		}

		// Key=Value pair
		eqIdx := strings.Index(line, "=")
		colonIdx := strings.Index(line, ":")

		var sepIdx int
		if eqIdx != -1 && (colonIdx == -1 || eqIdx < colonIdx) {
			sepIdx = eqIdx
		} else if colonIdx != -1 {
			sepIdx = colonIdx
		}

		if sepIdx == -1 {
			issues = append(issues, issue{lineNum, "malformed line: expected key=value or key: value"})
			continue
		}

		key := strings.TrimSpace(line[:sepIdx])
		if key == "" {
			issues = append(issues, issue{lineNum, "empty key"})
			continue
		}

		// Check for invalid key characters
		for _, c := range key {
			if c == '[' || c == ']' || c == '=' || c == '\n' || c == '\r' {
				issues = append(issues, issue{lineNum, fmt.Sprintf("invalid character '%c' in key %q", c, key)})
				break
			}
		}

		if strict && currentSection == "" {
			issues = append(issues, issue{lineNum, "key outside of any section (use --strict to allow)"})
		}

		section := currentSection
		if section == "" {
			section = "GLOBAL"
		}

		if keyInSection[section] == nil {
			keyInSection[section] = make(map[string]bool)
		}
		if keyInSection[section][key] {
			issues = append(issues, issue{lineNum, fmt.Sprintf("duplicate key %q in section [%s]", key, section)})
		}
		keyInSection[section][key] = true
	}

	return issues, scanner.Err()
}

func main() {
	args := os.Args[1:]
	var strictMode bool
	var inputPath string

	for _, arg := range args {
		switch arg {
		case "-s", "--strict":
			strictMode = true
		default:
			if strings.HasPrefix(arg, "-") {
				fmt.Fprintf(os.Stderr, "Unknown flag: %s\n", arg)
				fmt.Fprintf(os.Stderr, "Usage: ini-validate [-s|--strict] [file.ini]\n")
				os.Exit(1)
			}
			inputPath = arg
		}
	}

	var in io.Reader = os.Stdin
	if inputPath != "" {
		f, err := os.Open(inputPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()
		in = f
	}

	issues, err := validate(in, strictMode)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	if len(issues) == 0 {
		fmt.Println("OK: INI file is valid")
		return
	}

	fmt.Printf("Found %d issue(s):\n\n", len(issues))
	for _, iss := range issues {
		fmt.Printf("  Line %d: %s\n", iss.line, iss.message)
	}
	os.Exit(2)
}
