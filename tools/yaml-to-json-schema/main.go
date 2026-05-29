// yaml-to-json-schema generates a JSON Schema from a YAML file.
//
// It parses the YAML structure and generates a corresponding JSON Schema
// document with type inference, required fields, and format detection.
//
// Usage:
//
//	yaml-to-json-schema [flags] <file.yaml>
//
// Flags:
//   -t, --title        Title for the generated schema
//   -d, --description  Description for the generated schema
//   -o, --output       Output file (default: stdout)
//   -r, --required     Make all fields required (default: true)
//   -h, --help         Show help
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

// Schema represents a JSON Schema document.
type Schema struct {
	SchemaDraft    string                 `json:"$schema"`
	Type           string                 `json:"type,omitempty"`
	Title          string                 `json:"title,omitempty"`
	Description    string                 `json:"description,omitempty"`
	Properties     map[string]*Schema     `json:"properties,omitempty"`
	Required       []string               `json:"required,omitempty"`
	Items          *Schema                `json:"items,omitempty"`
	Additional     bool                   `json:"additionalProperties,omitempty"`
	Enum           []interface{}          `json:"enum,omitempty"`
	Format         string                 `json:"format,omitempty"`
	Minimum        *float64               `json:"minimum,omitempty"`
	Maximum        *float64               `json:"maximum,omitempty"`
	MinLength      *int                   `json:"minLength,omitempty"`
	MaxLength      *int                   `json:"maxLength,omitempty"`
	Pattern        string                 `json:"pattern,omitempty"`
	MinItems       *int                   `json:"minItems,omitempty"`
	MaxItems       *int                   `json:"maxItems,omitempty"`
	DefinitionName string                 `json:"-"`
	Definitions    map[string]*Schema     `json:"definitions,omitempty"`
	Ref            string                 `json:"$ref,omitempty"`
	AllOf          []*Schema              `json:"allOf,omitempty"`
	AnyOf          []*Schema             `json:"anyOf,omitempty"`
	OneOf          []*Schema             `json:"oneOf,omitempty"`
}

func inferSchemaFromNode(node *yaml.Node, allRequired bool, defs map[string]*Schema, path string) *Schema {
	switch node.Kind {
	case yaml.DocumentNode:
		if len(node.Content) > 0 {
			return inferSchemaFromNode(node.Content[0], allRequired, defs, path)
		}
		return &Schema{Type: "null"}

	case yaml.SequenceNode:
		schema := &Schema{Type: "array"}
		if len(node.Content) == 0 {
			schema.Items = &Schema{Type: "null"}
			return schema
		}

		// Infer item type from first element
		itemSchema := inferSchemaFromNode(node.Content[0], allRequired, defs, path+".[]")

		// Check if all items have the same type
		allSame := true
		for i := 1; i < len(node.Content); i++ {
			itemType := inferSchemaFromNode(node.Content[i], allRequired, defs, path+".["+strconv.Itoa(i)+"]")
			if itemType.Type != itemSchema.Type {
				allSame = false
				break
			}
		}

		if allSame {
			schema.Items = itemSchema
		} else {
			// Different types - use anyOf
			var anyOfTypes []*Schema
			seen := make(map[string]bool)
			for i := 0; i < len(node.Content); i++ {
				itemType := inferSchemaFromNode(node.Content[i], allRequired, defs, path+".["+strconv.Itoa(i)+"]")
				key := schemaKey(itemType)
				if !seen[key] {
					seen[key] = true
					anyOfTypes = append(anyOfTypes, itemType)
				}
			}
			schema.AnyOf = anyOfTypes
		}

		lenItems := len(node.Content)
		schema.MinItems = &lenItems

		return schema

	case yaml.MappingNode:
		schema := &Schema{Type: "object", Properties: make(map[string]*Schema)}

		for i := 0; i < len(node.Content); i += 2 {
			keyNode := node.Content[i]
			valNode := node.Content[i+1]

			key := keyNode.Value
			propSchema := inferSchemaFromNode(valNode, allRequired, defs, path+"."+key)
			schema.Properties[key] = propSchema

			if allRequired {
				schema.Required = append(schema.Required, key)
			}
		}

		schema.Additional = false
		return schema

	case yaml.ScalarNode:
		return inferSchemaFromScalar(node)

	default:
		return &Schema{Type: "null"}
	}
}

func inferSchemaFromScalar(node *yaml.Node) *Schema {
	value := node.Value
	tag := node.Tag

	switch tag {
	case "!!null", "!!nil":
		return &Schema{Type: "null"}

	case "!!bool":
		return &Schema{Type: "boolean"}

	case "!!int":
		schema := &Schema{Type: "integer"}
		if f, err := strconv.ParseFloat(value, 64); err == nil {
			schema.Minimum = &f
			schema.Maximum = &f
		}
		return schema

	case "!!float":
		schema := &Schema{Type: "number"}
		if f, err := strconv.ParseFloat(value, 64); err == nil {
			schema.Minimum = &f
			schema.Maximum = &f
		}
		return schema

	case "!!str", "!!timestamp":
		schema := &Schema{Type: "string"}

		// Detect formats
		if looksLikeEmail(value) {
			schema.Format = "email"
		} else if looksLikeDateTime(value) {
			schema.Format = "date-time"
		} else if looksLikeDate(value) {
			schema.Format = "date"
		} else if looksLikeUUID(value) {
			schema.Format = "uuid"
		} else if looksLikeURI(value) {
			schema.Format = "uri"
		}

		lenVal := len(value)
		schema.MinLength = &lenVal
		schema.MaxLength = &lenVal

		return schema

	default:
		// Try to infer from value
		if value == "null" || value == "~" || value == "" {
			return &Schema{Type: "null"}
		}
		if value == "true" || value == "false" {
			return &Schema{Type: "boolean"}
		}
		if _, err := strconv.ParseInt(value, 10, 64); err == nil {
			schema := &Schema{Type: "integer"}
			if f, _ := strconv.ParseFloat(value, 64); true {
				schema.Minimum = &f
				schema.Maximum = &f
			}
			return schema
		}
		if _, err := strconv.ParseFloat(value, 64); err == nil {
			schema := &Schema{Type: "number"}
			if f, _ := strconv.ParseFloat(value, 64); true {
				schema.Minimum = &f
				schema.Maximum = &f
			}
			return schema
		}
		return &Schema{Type: "string"}
	}
}

func looksLikeEmail(s string) bool {
	return strings.Contains(s, "@") && strings.Contains(s, ".")
}

func looksLikeDateTime(s string) bool {
	if len(s) >= 19 && s[4] == '-' && s[7] == '-' && s[10] == 'T' {
		return true
	}
	return false
}

func looksLikeDate(s string) bool {
	if len(s) == 10 && s[4] == '-' && s[7] == '-' {
		return true
	}
	return false
}

func looksLikeUUID(s string) bool {
	if len(s) == 36 && s[8] == '-' && s[13] == '-' && s[18] == '-' && s[23] == '-' {
		return true
	}
	return false
}

func looksLikeURI(s string) bool {
	return strings.HasPrefix(s, "http://") || strings.HasPrefix(s, "https://") || strings.HasPrefix(s, "ftp://")
}

func schemaKey(s *Schema) string {
	data, _ := json.Marshal(s)
	return string(data)
}

func printUsage() {
	fmt.Fprintf(os.Stderr, `yaml-to-json-schema - Generate a JSON Schema from a YAML file

Usage:
  yaml-to-json-schema [flags] <file.yaml>

Flags:
  -t, --title        Title for the generated schema
  -d, --description  Description for the generated schema
  -o, --output       Output file (default: stdout)
  -r, --required     Make all fields required (default: true)
  -h, --help         Show this help message

Examples:
  yaml-to-json-schema config.yaml                     Generate schema from YAML
  yaml-to-json-schema -t "Config" config.yaml        Set schema title
  yaml-to-json-schema -o schema.json config.yaml     Write to file
  yaml-to-json-schema --required=false config.yaml   No fields required

The tool parses the YAML structure and generates a corresponding JSON Schema
with type inference, format detection (email, date-time, uuid, uri), and
constraints (min/max for numbers, string lengths).

`)
}

func main() {
	title := flag.String("t", "", "Title for the generated schema")
	titleLong := flag.String("title", "", "Title for the generated schema")
	description := flag.String("d", "", "Description for the generated schema")
	descriptionLong := flag.String("description", "", "Description for the generated schema")
	output := flag.String("o", "", "Output file (default: stdout)")
	outputLong := flag.String("output", "", "Output file (default: stdout)")
	required := flag.Bool("r", true, "Make all fields required")
	requiredLong := flag.Bool("required", true, "Make all fields required")

	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Fprintln(os.Stderr, "Error: no file specified")
		printUsage()
		os.Exit(1)
	}

	// Merge short and long flags
	if *titleLong != "" {
		*title = *titleLong
	}
	if *descriptionLong != "" {
		*description = *descriptionLong
	}
	if *outputLong != "" {
		*output = *outputLong
	}
	if !*requiredLong {
		*required = false
	}

	filename := flag.Arg(0)

	// Read file
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Parse YAML
	var rootNode yaml.Node
	if err := yaml.Unmarshal(data, &rootNode); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing YAML: %v\n", err)
		os.Exit(1)
	}

	defs := make(map[string]*Schema)
	schema := inferSchemaFromNode(&rootNode, *required, defs, "")

	// Set metadata
	schema.SchemaDraft = "http://json-schema.org/draft-07/schema#"
	if *title != "" {
		schema.Title = *title
	}
	if *description != "" {
		schema.Description = *description
	}

	// Marshal to JSON
	jsonData, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	// Output
	if *output != "" {
		if err := os.WriteFile(*output, append(jsonData, '\n'), 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing output file: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Schema written to %s\n", *output)
	} else {
		fmt.Println(string(jsonData))
	}
}
