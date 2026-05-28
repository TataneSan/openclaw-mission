package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type JSONSchema struct {
	SchemaID      string                 `json:"$schema"`
	Type          string                 `json:"type"`
	Title         string                 `json:"title,omitempty"`
	Properties    map[string]Property    `json:"properties"`
	Required      []string               `json:"required,omitempty"`
	AdditionalProps bool                 `json:"additionalProperties"`
}

type Property struct {
	Type        string `json:"type"`
	Format      string `json:"format,omitempty"`
	Example     string `json:"example,omitempty"`
	Description string `json:"description,omitempty"`
}

func detectType(value string) (string, string) {
	if value == "true" || value == "false" {
		return "boolean", ""
	}
	if strings.Contains(value, ".") {
		return "number", ""
	}
	if value != "" {
		allDigits := true
		for _, r := range value {
			if r < '0' || r > '9' {
				allDigits = false
				break
			}
		}
		if allDigits {
			return "integer", ""
		}
	}
	if len(value) >= 36 && !strings.Contains(value, " ") {
		return "string", "uuid"
	}
	if strings.Contains(value, "@") {
		return "string", "email"
	}
	if strings.HasPrefix(value, "http") {
		return "string", "uri"
	}
	return "string", ""
}

func parseEnv(filename string) (map[string]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	vars := make(map[string]string)
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		value = strings.Trim(value, "\"")
		value = strings.Trim(value, "'")
		vars[key] = value
	}
	return vars, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: env-to-jsonschema <file.env> [title]")
		os.Exit(1)
	}

	filename := os.Args[1]
	title := "Environment Configuration"
	if len(os.Args) > 2 {
		title = os.Args[2]
	}

	vars, err := parseEnv(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	properties := make(map[string]Property)
	var required []string

	for key, value := range vars {
		typ, format := detectType(value)
		prop := Property{
			Type: typ,
		}
		if format != "" {
			prop.Format = format
		}
		if value != "" {
			prop.Example = value
		}
		prop.Description = fmt.Sprintf("Environment variable: %s", key)
		properties[key] = prop
		required = append(required, key)
	}

	schema := JSONSchema{
		SchemaID:        "http://json-schema.org/draft-07/schema#",
		Type:            "object",
		Title:           title,
		Properties:      properties,
		Required:        required,
		AdditionalProps: false,
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	encoder.Encode(schema)
}
