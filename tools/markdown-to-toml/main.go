// markdown-to-toml extracts structured data from Markdown tables and converts to TOML.
package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Table holds parsed data from a Markdown table.
type Table struct {
	Headers []string
	Rows    [][]string
}

// parseMarkdownTables extracts all tables from Markdown content.
func parseMarkdownTables(content string) []Table {
	var tables []Table
	lines := strings.Split(content, "\n")

	i := 0
	for i < len(lines) {
		line := lines[i]

		// Check if this line looks like a table separator (| --- | --- |)
		if isSeparator(line) && i > 0 {
			// Previous line should be the header
			headerLine := lines[i-1]
			headers := parseTableRow(headerLine)

			if len(headers) > 0 {
				i++ // Skip separator (header is at i-1, already processed)
				var rows [][]string
				for i < len(lines) {
					rowLine := lines[i]
					// Stop if we hit an empty line, another separator, or a heading
					if strings.TrimSpace(rowLine) == "" || isSeparator(rowLine) || isHeading(rowLine) {
						break
					}
					row := parseTableRow(rowLine)
					if len(row) == len(headers) {
						rows = append(rows, row)
					}
					i++
				}

				if len(rows) > 0 {
					tables = append(tables, Table{
						Headers: headers,
						Rows:    rows,
					})
				}
			}
		}
		i++
	}

	return tables
}

// isSeparator checks if a line is a Markdown table separator.
func isSeparator(line string) bool {
	trimmed := strings.TrimSpace(line)
	if !strings.HasPrefix(trimmed, "|") {
		return false
	}
	// Must contain at least one cell with dashes
	cellRe := regexp.MustCompile(`\|\s*[-:]+`)
	return cellRe.MatchString(trimmed)
}

// isHeading checks if a line is a Markdown heading.
func isHeading(line string) bool {
	trimmed := strings.TrimSpace(line)
	return strings.HasPrefix(trimmed, "#")
}

// parseTableRow parses a single Markdown table row into cells.
func parseTableRow(line string) []string {
	trimmed := strings.TrimSpace(line)
	// Remove leading/trailing pipes
	trimmed = strings.TrimPrefix(trimmed, "|")
	trimmed = strings.TrimSuffix(trimmed, "|")
	trimmed = strings.TrimSpace(trimmed)

	if trimmed == "" {
		return nil
	}

	cells := strings.Split(trimmed, "|")
	var result []string
	for _, cell := range cells {
		trimmed := strings.TrimSpace(cell)
		// Remove inline formatting (bold, italic)
		trimmed = strings.ReplaceAll(trimmed, "**", "")
		trimmed = strings.ReplaceAll(trimmed, "*", "")
		trimmed = strings.ReplaceAll(trimmed, "__", "")
		trimmed = strings.ReplaceAll(trimmed, "_", " ")
		trimmed = strings.TrimSpace(trimmed)
		result = append(result, trimmed)
	}

	return result
}

// toTOMLKey converts a string to a valid TOML key.
func toTOMLKey(s string) string {
	var result strings.Builder
	needsQuotes := false
	for i, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '_' || r == '-' {
			if i == 0 && (r >= '0' && r <= '9') {
				needsQuotes = true
			}
			result.WriteRune(r)
		} else {
			result.WriteRune('_')
		}
	}
	key := result.String()
	if key == "" {
		key = "field"
	}
	if needsQuotes {
		return fmt.Sprintf(`"%s"`, key)
	}
	return key
}

// toTOMLValue converts a string value to a TOML value.
func toTOMLValue(key, value string) string {
	// Try integer
	if isInteger(value) {
		return value
	}
	// Try float
	if isFloat(value) {
		return value
	}
	// Try boolean
	lower := strings.ToLower(value)
	if lower == "true" || lower == "false" {
		return lower
	}
	// Default: string
	return fmt.Sprintf("%q", value)
}

// isInteger checks if a string is a valid integer.
func isInteger(s string) bool {
	if s == "" {
		return false
	}
	start := 0
	if len(s) > 0 && (s[0] == '-' || s[0] == '+') {
		start = 1
	}
	if start >= len(s) {
		return false
	}
	for _, r := range s[start:] {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

// isFloat checks if a string is a valid float.
func isFloat(s string) bool {
	if s == "" {
		return false
	}
	start := 0
	if len(s) > 0 && (s[0] == '-' || s[0] == '+') {
		start = 1
	}
	if start >= len(s) {
		return false
	}
	dotCount := 0
	for _, r := range s[start:] {
		if r == '.' {
			dotCount++
			if dotCount > 1 {
				return false
			}
		} else if (r < '0' || r > '9') && r != 'e' && r != 'E' {
			return false
		}
	}
	return dotCount == 1
}

func main() {
	asArray := flag.Bool("a", false, "Output as TOML array of tables [[table]]")
	tableName := flag.String("t", "data", "Table name for array of tables mode")
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] <file.md>\n\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Extract structured data from Markdown tables and convert to TOML.")
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	filename := flag.Arg(0)
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading %s: %v\n", filename, err)
		os.Exit(1)
	}

	content := string(data)
	tables := parseMarkdownTables(content)

	if len(tables) == 0 {
		fmt.Fprintf(os.Stderr, "No Markdown tables found in %s\n", filename)
		os.Exit(1)
	}

	first := true
	for tIdx, table := range tables {
		if !first {
			fmt.Println()
		}
		first = false

		if *asArray {
			name := *tableName
			if len(tables) > 1 {
				name = fmt.Sprintf("%s_%d", *tableName, tIdx+1)
			}
			fmt.Printf("[[%s]]\n", name)
			for _, row := range table.Rows {
				for cIdx, cell := range row {
					key := toTOMLKey(table.Headers[cIdx])
					fmt.Printf("%s = %s\n", key, toTOMLValue(key, cell))
				}
				fmt.Println()
			}
		} else {
			// Simple key-value mode: first column is key, rest are values
			if len(table.Headers) >= 2 {
				keyCol := toTOMLKey(table.Headers[0])
				for _, row := range table.Rows {
					key := row[0]
					if len(row) == 1 {
						fmt.Printf("%s = true\n", toTOMLKey(key))
					} else if len(row) == 2 {
						fmt.Printf("%s = %s\n", toTOMLKey(key), toTOMLValue(keyCol, row[1]))
					} else {
						// Multiple value columns: create inline table
						fmt.Printf("%s = {", toTOMLKey(key))
						for vIdx, val := range row[1:] {
							if vIdx > 0 {
								fmt.Print(", ")
							}
							vKey := toTOMLKey(table.Headers[vIdx+1])
							fmt.Printf("%s = %s", vKey, toTOMLValue(vKey, val))
						}
						fmt.Println("}")
					}
				}
			} else {
				// Single column: just output values
				key := toTOMLKey(table.Headers[0])
				fmt.Printf("%s = [\n", key)
				for _, row := range table.Rows {
					fmt.Printf("  %s,\n", toTOMLValue(key, row[0]))
				}
				fmt.Println("]")
			}
		}
	}
}
