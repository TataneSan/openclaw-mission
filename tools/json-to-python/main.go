// json-to-python generates Python dataclasses from JSON input.
//
// Usage:
//
//	json-to-python data.json
//	json-to-python data.json -o models.py
//	json-to-python data.json -n User
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

// pythonKeyword checks if a string is a reserved Python keyword.
func pythonKeyword(s string) bool {
	keywords := map[string]bool{
		"False": true, "None": true, "True": true, "and": true, "as": true,
		"assert": true, "async": true, "await": true, "break": true, "class": true,
		"continue": true, "def": true, "del": true, "elif": true, "else": true,
		"except": true, "finally": true, "for": true, "from": true, "global": true,
		"if": true, "import": true, "in": true, "is": true, "lambda": true,
		"nonlocal": true, "not": true, "or": true, "pass": true, "raise": true,
		"return": true, "try": true, "while": true, "with": true, "yield": true,
	}
	return keywords[s]
}

// toFieldName converts a JSON key to a valid Python field name.
func toFieldName(key string) string {
	// Replace hyphens and spaces with underscores
	s := strings.ReplaceAll(key, "-", "_")
	s = strings.ReplaceAll(s, " ", "_")
	// If starts with digit, prefix with underscore
	if len(s) > 0 && s[0] >= '0' && s[0] <= '9' {
		s = "_" + s
	}
	// If it's a Python keyword, suffix with underscore
	if pythonKeyword(s) {
		s = s + "_"
	}
	if s == "" {
		s = "_field"
	}
	return s
}

// inferType determines the Python type hint from a JSON value.
func inferType(v interface{}) string {
	switch val := v.(type) {
	case nil:
		return "Any"
	case bool:
		return "bool"
	case float64:
		// Check if it's an integer value
		if val == float64(int64(val)) && val != 0 {
			return "int"
		}
		return "float"
	case string:
		return "str"
	case json.Number:
		return "int"
	case map[string]interface{}:
		return "Dict[str, Any]"
	case []interface{}:
		if len(val) == 0 {
			return "List[Any]"
		}
		// Infer from first element
		inner := inferType(val[0])
		return "List[" + inner + "]"
	default:
		return "Any"
	}
}

// FieldInfo holds a field name and its type hint.
type FieldInfo struct {
	Name string
	Type string
}

// extractFields extracts field names and types from a JSON object.
func extractFields(obj map[string]interface{}) []FieldInfo {
	fields := make([]FieldInfo, 0, len(obj))
	for key, val := range obj {
		fields = append(fields, FieldInfo{
			Name: toFieldName(key),
			Type: inferType(val),
		})
	}
	sort.Slice(fields, func(i, j int) bool {
		return fields[i].Name < fields[j].Name
	})
	return fields
}

// generateDataclass generates Python dataclass code from a JSON object.
func generateDataclass(name string, obj map[string]interface{}) string {
	fields := extractFields(obj)

	var sb strings.Builder
	sb.WriteString("from __future__ import annotations\n")
	sb.WriteString("from dataclasses import dataclass, field\n")
	sb.WriteString("from typing import Any, Dict, List\n\n\n")
	sb.WriteString("@dataclass\n")
	sb.WriteString("class ")
	sb.WriteString(name)
	sb.WriteString(":\n")

	for _, f := range fields {
		sb.WriteString("    ")
		sb.WriteString(f.Name)
		sb.WriteString(": ")
		sb.WriteString(f.Type)
		// For mutable types, use field(default_factory=...)
		if strings.HasPrefix(f.Type, "List[") {
			sb.WriteString(" = field(default_factory=list)")
		} else if strings.HasPrefix(f.Type, "Dict[") {
			sb.WriteString(" = field(default_factory=dict)")
		} else {
			sb.WriteString(" = None")
		}
		sb.WriteString("\n")
	}

	sb.WriteString("\n    @classmethod\n")
	sb.WriteString("    def from_dict(cls, data: dict) -> ")
	sb.WriteString(name)
	sb.WriteString(":\n")
	sb.WriteString("        return cls(**data)\n")

	return sb.String()
}

func printUsage() {
	fmt.Fprintf(os.Stderr, "json-to-python - Generate Python dataclasses from JSON input\n\n")
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "  json-to-python [flags] <input.json>\n\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "\nExamples:\n")
	fmt.Fprintf(os.Stderr, "  json-to-python user.json\n")
	fmt.Fprintf(os.Stderr, "  json-to-python user.json -n User -o models.py\n")
	fmt.Fprintf(os.Stderr, "  cat user.json | json-to-python -\n")
}

func main() {
	output := flag.String("o", "", "output file (default: stdout)")
	name := flag.String("n", "Model", "dataclass name")
	flag.Usage = printUsage
	flag.Parse()

	filePath := flag.Arg(0)
	var input io.Reader
	if filePath == "" || filePath == "-" {
		input = os.Stdin
	} else {
		f, err := os.Open(filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()
		input = f
	}

	data, err := io.ReadAll(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
		os.Exit(1)
	}

	var obj map[string]interface{}
	if err := json.Unmarshal(data, &obj); err != nil {
		fmt.Fprintf(os.Stderr, "error: input must be a JSON object: %v\n", err)
		os.Exit(1)
	}

	result := generateDataclass(*name, obj)

	if *output != "" {
		if err := os.WriteFile(*output, []byte(result), 0644); err != nil {
			fmt.Fprintf(os.Stderr, "error writing file: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Print(result)
	}
}
