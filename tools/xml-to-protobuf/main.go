package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type xmlStruct struct {
	XMLName xml.Name         `xml:"-"`
	Attrs   map[string]string `xml:"-"`
	Values  []xmlStruct       `xml:",any"`
	Text    string            `xml:",chardata"`
}

func toProtoType(value string) string {
	switch {
	case value == "true" || value == "false":
		return "bool"
	case looksLikeInt(value):
		return "int64"
	case looksLikeFloat(value):
		return "double"
	default:
		return "string"
	}
}

func looksLikeInt(s string) bool {
	if s == "" {
		return false
	}
	i := 0
	if s[0] == '-' || s[0] == '+' {
		i++
	}
	if i >= len(s) {
		return false
	}
	for _, c := range s[i:] {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func looksLikeFloat(s string) bool {
	if looksLikeInt(s) {
		return false
	}
	hasDot := false
	i := 0
	if s[0] == '-' || s[0] == '+' {
		i++
	}
	if i >= len(s) {
		return false
	}
	for _, c := range s[i:] {
		if c == '.' {
			if hasDot {
				return false
			}
			hasDot = true
		} else if c < '0' || c > '9' {
			return false
		}
	}
	return hasDot
}

func toProtoField(name string) string {
	var result strings.Builder
	for i, c := range name {
		if c == '-' || c == ' ' || c == '.' {
			if i > 0 && result.Len() > 0 {
				result.WriteRune('_')
			}
			continue
		}
		result.WriteRune(unicode.ToLower(c))
	}
	s := result.String()
	for strings.Contains(s, "__") {
		s = strings.ReplaceAll(s, "__", "_")
	}
	s = strings.Trim(s, "_")
	if s == "" || !unicode.IsLetter(rune(s[0])) {
		s = "field_" + s
	}
	return s
}

func toPascalCase(name string) string {
	var result strings.Builder
	nextUpper := true
	for _, c := range name {
		if c == '_' || c == '-' || c == ' ' || c == '.' {
			nextUpper = true
			continue
		}
		if nextUpper {
			result.WriteRune(unicode.ToUpper(c))
			nextUpper = false
		} else {
			result.WriteRune(c)
		}
	}
	s := result.String()
	if s == "" || !unicode.IsLetter(rune(s[0])) {
		s = "Field" + s
	}
	return s
}

func parseXML(content []byte) (*xmlStruct, error) {
	var root xmlStruct
	if err := xml.Unmarshal(content, &root); err != nil {
		return nil, err
	}
	return &root, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: xml-to-protobuf <file.xml> [-o output.proto]")
		fmt.Fprintln(os.Stderr, "\nGenerates a Protobuf message definition from an XML file.")
		os.Exit(1)
	}

	filename := os.Args[1]
	var outputPath string

	args := os.Args[2:]
	for len(args) > 0 {
		switch args[0] {
		case "-o":
			if len(args) < 2 {
				fmt.Fprintln(os.Stderr, "Error: -o requires a filename")
				os.Exit(1)
			}
			outputPath = args[1]
			args = args[2:]
		default:
			args = args[1:]
		}
	}

	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	root, err := parseXML(content)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing XML: %v\n", err)
		os.Exit(1)
	}

	var sb strings.Builder
	sb.WriteString("syntax = \"proto3\";\n\n")

	// Generate message from root element
	messageName := toPascalCase(root.XMLName.Local)
	sb.WriteString(fmt.Sprintf("message %s {\n", messageName))

	fieldNum := 1

	// Add attributes as fields
	if root.Attrs != nil {
		for k, v := range root.Attrs {
			fieldName := toProtoField(k)
			fieldType := toProtoType(v)
			sb.WriteString(fmt.Sprintf("  %s %s = %d;\n", fieldType, fieldName, fieldNum))
			fieldNum++
		}
	}

	// Add child elements as fields
	for _, child := range root.Values {
		fieldName := toProtoField(child.XMLName.Local)
		if len(child.Values) > 0 || child.Attrs != nil {
			// Nested element - create sub-message
			subMessageName := toPascalCase(child.XMLName.Local)
			sb.WriteString(fmt.Sprintf("  %s %s = %d;\n", subMessageName, fieldName, fieldNum))
			fieldNum++

			// Generate sub-message
			sb.WriteString(fmt.Sprintf("\nmessage %s {\n", subMessageName))
			for ak, av := range child.Attrs {
				subFieldName := toProtoField(ak)
				subFieldType := toProtoType(av)
				sb.WriteString(fmt.Sprintf("  %s %s = %d;\n", subFieldType, subFieldName, fieldNum))
				fieldNum++
			}
			if child.Text != "" && strings.TrimSpace(child.Text) != "" {
				sb.WriteString(fmt.Sprintf("  string value = %d;\n", fieldNum))
				fieldNum++
			}
			sb.WriteString("}\n")
		} else if child.Text != "" && strings.TrimSpace(child.Text) != "" {
			// Simple text element
			fieldType := toProtoType(strings.TrimSpace(child.Text))
			sb.WriteString(fmt.Sprintf("  %s %s = %d;\n", fieldType, fieldName, fieldNum))
			fieldNum++
		}
	}

	sb.WriteString("}\n")

	if outputPath != "" {
		if err := os.WriteFile(outputPath, []byte(sb.String()), 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing file: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Written to %s\n", outputPath)
	} else {
		fmt.Print(sb.String())
	}
}
