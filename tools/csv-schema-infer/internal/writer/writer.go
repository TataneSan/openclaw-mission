package writer

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/TataneSan/csv-schema-infer/internal/inferrer"
)

// SchemaResult holds the full inferred schema for a CSV.
type SchemaResult struct {
	Columns []inferrer.ColumnSchema `json:"columns"`
}

// WriteJSON writes the schema as formatted JSON.
func WriteJSON(w io.Writer, schema *SchemaResult) error {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	return encoder.Encode(schema)
}

// WriteJSONFile writes the schema as formatted JSON to a file.
func WriteJSONFile(path string, schema *SchemaResult) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create file: %w", err)
	}
	defer f.Close()

	if err := WriteJSON(f, schema); err != nil {
		return fmt.Errorf("write json: %w", err)
	}
	return nil
}

// WriteTable writes the schema as a formatted text table.
func WriteTable(w io.Writer, schema *SchemaResult) error {
	if len(schema.Columns) == 0 {
		fmt.Fprintln(w, "No columns found.")
		return nil
	}

	// Calculate column widths
	nameWidth := len("Column")
	typeWidth := len("Type")
	confWidth := len("Confidence")
	nullWidth := len("Nullable")
	samplesWidth := len("Samples")

	for _, col := range schema.Columns {
		if len(col.Name) > nameWidth {
			nameWidth = len(col.Name)
		}
		if len(string(col.Type)) > typeWidth {
			typeWidth = len(string(col.Type))
		}
		confStr := fmt.Sprintf("%.0f%%", col.Confidence)
		if len(confStr) > confWidth {
			confWidth = len(confStr)
		}
		nullStr := fmt.Sprintf("%v", col.Nullable)
		if len(nullStr) > nullWidth {
			nullWidth = len(nullStr)
		}
		if col.NullCount > 0 {
			sampleText := fmt.Sprintf("[%d nulls]", col.NullCount)
			if len(sampleText) > samplesWidth {
				samplesWidth = len(sampleText)
			}
		}
	}

	// Header
	fmt.Fprintf(w, "%-*s  %-*s  %-*s  %-*s  %s\n",
		nameWidth, "Column", typeWidth, "Type", confWidth, "Confidence", nullWidth, "Nullable", "Samples")
	fmt.Fprintf(w, "%s\n", strings.Repeat("-", nameWidth+2+typeWidth+2+confWidth+2+nullWidth+2+samplesWidth))

	// Rows
	for _, col := range schema.Columns {
		confStr := fmt.Sprintf("%.0f%%", col.Confidence)
		nullStr := fmt.Sprintf("%v", col.Nullable)
		sampleStr := ""
		if col.NullCount > 0 {
			sampleStr = fmt.Sprintf("[%d nulls]", col.NullCount)
		}
		fmt.Fprintf(w, "%-*s  %-*s  %-*s  %-*s  %s\n",
			nameWidth, col.Name, typeWidth, col.Type, confWidth, confStr, nullWidth, nullStr, sampleStr)
	}

	return nil
}

// WriteSQL writes CREATE TABLE SQL from the schema.
func WriteSQL(w io.Writer, tableName string, schema *SchemaResult) error {
	fmt.Fprintf(w, "CREATE TABLE %s (\n", sanitizeSQL(tableName))

	for i, col := range schema.Columns {
		sqlType := toSQLType(col.Type)
		nullable := ""
		if !col.Nullable {
			nullable = " NOT NULL"
		}
		comma := ""
		if i < len(schema.Columns)-1 {
			comma = ","
		}
		fmt.Fprintf(w, "    %s %s%s%s\n",
			sanitizeSQL(col.Name), sqlType, nullable, comma)
	}

	fmt.Fprintln(w, ");")
	return nil
}

func toSQLType(t inferrer.DataType) string {
	switch t {
	case inferrer.TypeInteger:
		return "INTEGER"
	case inferrer.TypeFloat:
		return "REAL"
	case inferrer.TypeBoolean:
		return "BOOLEAN"
	case inferrer.TypeDate:
		return "DATE"
	case inferrer.TypeDateTime:
		return "TIMESTAMP"
	case inferrer.TypeEmail, inferrer.TypeURL, inferrer.TypeIP:
		return "TEXT"
	default:
		return "TEXT"
	}
}

func sanitizeSQL(s string) string {
	// Simple sanitization: replace spaces with underscores, ensure alphanumeric
	var result strings.Builder
	for i, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9' && i > 0) || r == '_' {
			result.WriteRune(r)
		} else {
			result.WriteRune('_')
		}
	}
	return result.String()
}
