// toml-to-graphql converts TOML data to GraphQL schema types.
//
// Usage:
//
//	toml-to-graphql config.toml
//	toml-to-graphql schema.toml --prefix App
//	cat data.toml | toml-to-graphql
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
)

var (
	version = "1.0.0"
)

func main() {
	fs := flag.NewFlagSet("toml-to-graphql", flag.ExitOnError)
	input := fs.String("i", "", "Input TOML file (default: stdin)")
	prefix := fs.String("prefix", "", "Prefix for generated type names")
	output := fs.String("o", "", "Output file (default: stdout)")
	showVersion := fs.Bool("version", false, "Print version")

	fs.Parse(os.Args[1:])

	if *showVersion {
		fmt.Println("toml-to-graphql version", version)
		return
	}

	var in io.Reader
	if *input != "" {
		f, err := os.Open(*input)
		if err != nil {
			fatal("open input: %v", err)
		}
		defer f.Close()
		in = f
	} else {
		in = os.Stdin
	}

	var data map[string]interface{}
	if _, err := toml.NewDecoder(in).Decode(&data); err != nil {
		fatal("parse TOML: %v", err)
	}

	schema := generateSchema(data, *prefix)

	out := strings.NewReader(schema)
	if *output != "" {
		if err := os.WriteFile(*output, []byte(schema), 0644); err != nil {
			fatal("write output: %v", err)
		}
		fmt.Fprintf(os.Stderr, "Written to %s\n", *output)
	} else {
		io.Copy(os.Stdout, out)
		fmt.Println()
	}
}

func generateSchema(data map[string]interface{}, prefix string) string {
	var sb strings.Builder
	types := extractTypes(data, "", prefix)

	for name, fields := range types {
		sb.WriteString(fmt.Sprintf("type %s {\n", name))
		for _, f := range fields {
			sb.WriteString(fmt.Sprintf("  %s: %s\n", f.Name, f.Type))
		}
		sb.WriteString("}\n\n")
	}

	return strings.TrimSuffix(sb.String(), "\n")
}

type Field struct {
	Name string
	Type string
}

func extractTypes(data map[string]interface{}, path, prefix string) map[string][]Field {
	types := make(map[string][]Field)

	// Collect all nested tables as types
	for key, value := range data {
		fullPath := key
		if path != "" {
			fullPath = path + "." + key
		}

		switch v := value.(type) {
		case map[string]interface{}:
			typeName := toTypeName(prefix, fullPath)
			fields := mapToFields(v)
			types[typeName] = fields

			// Recurse into nested tables
			nested := extractTypes(v, fullPath, prefix)
			for n, f := range nested {
				types[n] = f
			}
		case []interface{}:
			if len(v) > 0 {
				switch first := v[0].(type) {
				case map[string]interface{}:
					typeName := toTypeName(prefix, fullPath)
					fields := mapToFields(first)
					types[typeName] = fields

					nested := extractTypes(first, fullPath, prefix)
					for n, f := range nested {
						types[n] = f
					}
				}
			}
		}
	}

	return types
}

func mapToFields(m map[string]interface{}) []Field {
	fields := make([]Field, 0, len(m))
	for key, value := range m {
		field := Field{
			Name: toFieldName(key),
			Type: toGraphQLType(value),
		}
		fields = append(fields, field)
	}
	return fields
}

func toGraphQLType(v interface{}) string {
	switch v.(type) {
	case nil:
		return "String"
	case bool:
		return "Boolean"
	case int, int64, float64:
		return "Float"
	case string:
		return "String"
	case []interface{}:
		return "[String]"
	case map[string]interface{}:
		return "JSON"
	default:
		return "String"
	}
}

func toTypeName(prefix, path string) string {
	parts := strings.Split(path, ".")
	var nameParts []string

	if prefix != "" {
		nameParts = append(nameParts, prefix)
	}

	for _, part := range parts {
		nameParts = append(nameParts, capitalize(part))
	}

	name := strings.Join(nameParts, "")
	if !isValidGraphQLName(name) {
		name = "_" + name
	}
	return name
}

func toFieldName(s string) string {
	var result strings.Builder
	upperNext := false
	for _, c := range s {
		switch {
		case c >= 'A' && c <= 'Z':
			result.WriteRune(c)
		case c >= 'a' && c <= 'z':
			if upperNext && result.Len() > 0 {
				upperNext = false
				// Capitalize
				result.WriteRune(c - 'a' + 'A')
			} else {
				result.WriteRune(c)
			}
		case c >= '0' && c <= '9':
			result.WriteRune(c)
		case c == '_':
			upperNext = true
		case c == '-':
			upperNext = true
		default:
			// Skip invalid characters
		}
	}

	name := result.String()
	if name == "" || (name[0] >= '0' && name[0] <= '9') {
		name = "_" + name
	}
	return name
}

func capitalize(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	runes[0] = rune(runes[0] - 'a' + 'A')
	return string(runes)
}

func isValidGraphQLName(name string) bool {
	if name == "" {
		return false
	}
	if name[0] == '_' || (name[0] >= 'a' && name[0] <= 'z') || (name[0] >= 'A' && name[0] <= 'Z') {
		return true
	}
	return false
}

func fatal(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "toml-to-graphql: "+format+"\n", args...)
	os.Exit(1)
}

// toJSON converts a TOML value to a JSON-compatible representation.
func toJSON(v interface{}) interface{} {
	switch val := v.(type) {
	case map[string]interface{}:
		result := make(map[string]interface{})
		for k, v := range val {
			result[k] = toJSON(v)
		}
		return result
	case []interface{}:
		result := make([]interface{}, len(val))
		for i, v := range val {
			result[i] = toJSON(v)
		}
		return result
	default:
		return val
	}
}

// mustJSON marshals a value to JSON.
func mustJSON(v interface{}) string {
	data, err := json.Marshal(v)
	if err != nil {
		return "{}"
	}
	return string(data)
}
