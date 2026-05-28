package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/ini.v1"
	"gopkg.in/yaml.v3"
)

type ValidationResult struct {
	File     string   `json:"file"`
	Format   string   `json:"format"`
	Valid    bool     `json:"valid"`
	Errors   []string `json:"errors,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
}

type ValidateOptions struct {
	Recursive bool
	Strict    bool
	Output    string // "text" or "json"
	Extensions []string
}

func detectFormat(path string) string {
	ext := strings.ToLower(filepath.Ext(path))
	switch ext {
	case ".json":
		return "json"
	case ".yaml", ".yml":
		return "yaml"
	case ".toml":
		return "toml"
	case ".ini", ".cfg", ".conf":
		return "ini"
	default:
		return ""
	}
}

func validateJSON(path string, strict bool) *ValidationResult {
	result := &ValidationResult{
		File:   path,
		Format: "json",
		Valid:  true,
	}

	data, err := os.ReadFile(path)
	if err != nil {
		result.Valid = false
		result.Errors = append(result.Errors, fmt.Sprintf("cannot read file: %v", err))
		return result
	}

	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		result.Valid = false
		result.Errors = append(result.Errors, fmt.Sprintf("invalid JSON: %v", err))
		return result
	}

	if strict {
		checkJSONStrict(data, v, "", result)
	}

	return result
}

func checkJSONStrict(data []byte, v interface{}, path string, result *ValidationResult) {
	switch val := v.(type) {
	case map[string]interface{}:
		for k, v := range val {
			currentPath := path + "." + k
			if strings.Contains(k, " ") {
				result.Warnings = append(result.Warnings, fmt.Sprintf("%s: key contains spaces: %q", currentPath, k))
			}
			if strings.HasSuffix(k, " ") || strings.HasPrefix(k, " ") {
				result.Warnings = append(result.Warnings, fmt.Sprintf("%s: key has leading/trailing spaces: %q", currentPath, k))
			}
			checkJSONStrict(data, v, currentPath, result)
		}
	case []interface{}:
		for i, item := range val {
			checkJSONStrict(data, item, fmt.Sprintf("%s[%d]", path, i), result)
		}
	}

	// Check for trailing commas (common JSON error)
	content := string(data)
	if strings.Contains(content, ",\n]") || strings.Contains(content, ",\n}") {
		result.Warnings = append(result.Warnings, "trailing comma detected")
	}
}

func validateYAML(path string, strict bool) *ValidationResult {
	result := &ValidationResult{
		File:   path,
		Format: "yaml",
		Valid:  true,
	}

	data, err := os.ReadFile(path)
	if err != nil {
		result.Valid = false
		result.Errors = append(result.Errors, fmt.Sprintf("cannot read file: %v", err))
		return result
	}

	var v interface{}
	if err := yaml.Unmarshal(data, &v); err != nil {
		result.Valid = false
		result.Errors = append(result.Errors, fmt.Sprintf("invalid YAML: %v", err))
		return result
	}

	if strict {
		lines := strings.Split(string(data), "\n")
		for i, line := range lines {
			if strings.Contains(line, "\t") {
				result.Errors = append(result.Errors, fmt.Sprintf("line %d: tabs not allowed in YAML, use spaces", i+1))
			}
			trimmed := strings.TrimRight(line, " \t")
			if len(line) != len(trimmed) && line[len(trimmed)] == ' ' {
				result.Warnings = append(result.Warnings, fmt.Sprintf("line %d: trailing whitespace", i+1))
			}
		}
	}

	return result
}

func validateTOML(path string, strict bool) *ValidationResult {
	result := &ValidationResult{
		File:   path,
		Format: "toml",
		Valid:  true,
	}

	// Use Go's encoding/toml via ini package fallback
	// Since we don't have a TOML library, we'll do basic validation
	data, err := os.ReadFile(path)
	if err != nil {
		result.Valid = false
		result.Errors = append(result.Errors, fmt.Sprintf("cannot read file: %v", err))
		return result
	}

	// Basic TOML validation
	lines := strings.Split(string(data), "\n")
	sectionCount := make(map[string]int)

	for i, line := range lines {
		trimmed := strings.TrimSpace(line)

		// Skip empty lines and comments
		if trimmed == "" || strings.HasPrefix(trimmed, "#") {
			continue
		}

		// Section header
		if strings.HasPrefix(trimmed, "[") && strings.HasSuffix(trimmed, "]") {
			section := trimmed[1 : len(trimmed)-1]
			sectionCount[section]++
			if sectionCount[section] > 1 {
				result.Valid = false
				result.Errors = append(result.Errors, fmt.Sprintf("line %d: duplicate section: %s", i+1, section))
			}
			_ = section
			continue
		}

		// Key-value pair
		if !strings.Contains(trimmed, "=") {
			result.Valid = false
			result.Errors = append(result.Errors, fmt.Sprintf("line %d: invalid TOML syntax (expected key = value)", i+1))
		}
	}

	if strict {
		for i, line := range lines {
			trimmed := strings.TrimRight(line, " \t")
			if len(line) != len(trimmed) && line[len(trimmed)] == ' ' {
				result.Warnings = append(result.Warnings, fmt.Sprintf("line %d: trailing whitespace", i+1))
			}
		}
	}

	return result
}

func validateINI(path string, strict bool) *ValidationResult {
	result := &ValidationResult{
		File:   path,
		Format: "ini",
		Valid:  true,
	}

	_, err := ini.Load(path)
	if err != nil {
		result.Valid = false
		result.Errors = append(result.Errors, fmt.Sprintf("invalid INI: %v", err))
		return result
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return result
	}

	if strict {
		lines := strings.Split(string(data), "\n")
		for i, line := range lines {
			trimmed := strings.TrimSpace(line)
			if trimmed == "" || strings.HasPrefix(trimmed, ";") || strings.HasPrefix(trimmed, "#") {
				continue
			}
			if strings.HasPrefix(trimmed, "[") {
				if !strings.HasSuffix(trimmed, "]") {
					result.Errors = append(result.Errors, fmt.Sprintf("line %d: unclosed section header", i+1))
				}
			}
		}
	}

	return result
}

func validateFile(path string, strict bool) *ValidationResult {
	format := detectFormat(path)
	if format == "" {
		return &ValidationResult{
			File:   path,
			Format: "unknown",
			Valid:  false,
			Errors: []string{"unsupported file format"},
		}
	}

	switch format {
	case "json":
		return validateJSON(path, strict)
	case "yaml":
		return validateYAML(path, strict)
	case "toml":
		return validateTOML(path, strict)
	case "ini":
		return validateINI(path, strict)
	default:
		return &ValidationResult{
			File:   path,
			Format: format,
			Valid:  false,
			Errors: []string{"unsupported format"},
		}
	}
}

func findConfigFiles(dir string, extensions []string) ([]string, error) {
	var files []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			// Skip hidden directories and common non-config directories
			if strings.HasPrefix(info.Name(), ".") || info.Name() == "node_modules" || info.Name() == "vendor" {
				return filepath.SkipDir
			}
			return nil
		}

		ext := strings.ToLower(filepath.Ext(path))
		if len(extensions) == 0 {
			// Default extensions
			extensions = []string{".json", ".yaml", ".yml", ".toml", ".ini", ".cfg", ".conf"}
		}

		for _, e := range extensions {
			if ext == strings.ToLower(e) {
				files = append(files, path)
				break
			}
		}
		return nil
	})

	return files, err
}

func formatText(results []*ValidationResult) string {
	var sb strings.Builder

	validCount := 0
	invalidCount := 0

	for _, r := range results {
		if r.Valid {
			validCount++
		} else {
			invalidCount++
		}
	}

	sb.WriteString(fmt.Sprintf("Config Validator - %d files checked\n", len(results)))
	sb.WriteString(fmt.Sprintf("Valid: %d | Invalid: %d\n\n", validCount, invalidCount))

	for _, r := range results {
		status := "PASS"
		if !r.Valid {
			status = "FAIL"
		}

		sb.WriteString(fmt.Sprintf("[%s] %s (%s)\n", status, r.File, r.Format))

		for _, err := range r.Errors {
			sb.WriteString(fmt.Sprintf("  ERROR: %s\n", err))
		}
		for _, warn := range r.Warnings {
			sb.WriteString(fmt.Sprintf("  WARN:  %s\n", warn))
		}
	}

	return sb.String()
}

func formatJSONOutput(results []*ValidationResult) (string, error) {
	data, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func printUsage() {
	fmt.Println("config-validator - Validate configuration files (JSON, YAML, TOML, INI)")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  config-validator [options] <file|directory>")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -r, --recursive    Recursively search directories for config files")
	fmt.Println("  -s, --strict       Enable strict validation with additional checks")
	fmt.Println("  -o, --output       Output format: text (default) or json")
	fmt.Println("  -e, --ext          File extensions to check (comma-separated)")
	fmt.Println("                     Supported: .json, .yaml, .yml, .toml, .ini, .cfg, .conf")
	fmt.Println("  -h, --help         Show this help message")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  config-validator config.json")
	fmt.Println("  config-validator -s config.yaml")
	fmt.Println("  config-validator -r ./configs")
	fmt.Println("  config-validator -r -e .json,.yaml ./")
	fmt.Println("  config-validator -o json config.toml")
}

func main() {
	// Custom flag parsing to handle positional args
	fs := flag.NewFlagSet("config-validator", flag.ContinueOnError)
	fs.Bool("r", false, "Recursively search directories")
	strict := fs.Bool("s", false, "Strict validation mode")
	output := fs.String("o", "text", "Output format: text or json")
	extFlag := fs.String("e", "", "Comma-separated list of extensions")
	help := fs.Bool("h", false, "Show help")

	fs.Parse(os.Args[1:])

	if *help {
		printUsage()
		os.Exit(0)
	}

	if fs.NArg() == 0 {
		fmt.Fprintln(os.Stderr, "Error: no file or directory specified")
		fmt.Fprintln(os.Stderr)
		printUsage()
		os.Exit(1)
	}

	target := fs.Arg(0)
	var extensions []string
	if *extFlag != "" {
		extensions = strings.Split(*extFlag, ",")
		for i, e := range extensions {
			e = strings.TrimSpace(e)
			if !strings.HasPrefix(e, ".") {
				e = "." + e
			}
			extensions[i] = e
		}
	}

	info, err := os.Stat(target)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	var results []*ValidationResult

	if info.IsDir() {
		// Directory mode
		files, err := findConfigFiles(target, extensions)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error scanning directory: %v\n", err)
			os.Exit(1)
		}

		if len(files) == 0 {
			fmt.Fprintln(os.Stdout, "No configuration files found")
			os.Exit(0)
		}

		for _, file := range files {
			result := validateFile(file, *strict)
			results = append(results, result)
		}
	} else {
		// Single file mode
		result := validateFile(target, *strict)
		results = append(results, result)
	}

	// Output results
	switch *output {
	case "json":
		jsonOut, err := formatJSONOutput(results)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error formatting JSON output: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(jsonOut)
	default:
		fmt.Print(formatText(results))
	}

	// Exit code based on validation results
	for _, r := range results {
		if !r.Valid {
			os.Exit(1)
		}
	}
}
