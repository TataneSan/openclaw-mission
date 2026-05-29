package reader

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

// Record represents a single row from a CSV file.
type Record []string

// Dataset represents the full CSV data including headers.
type Dataset struct {
	Headers []string
	Rows    [][]string
}

// ReadFile reads a CSV file and returns the parsed dataset.
func ReadFile(path string, delimiter rune) (*Dataset, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open file: %w", err)
	}
	defer f.Close()

	return ReadReader(f, delimiter)
}

// ReadReader reads CSV data from an io.Reader.
func ReadReader(r io.Reader, delimiter rune) (*Dataset, error) {
	reader := csv.NewReader(r)
	reader.LazyQuotes = true
	reader.TrimLeadingSpace = true
	reader.FieldsPerRecord = -1 // allow variable field count

	if delimiter != 0 {
		reader.Comma = delimiter
	}

	allRecords, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("parse csv: %w", err)
	}

	if len(allRecords) == 0 {
		return nil, fmt.Errorf("empty csv file")
	}

	headers := allRecords[0]
	rows := allRecords[1:]

	// Normalize headers
	for i, h := range headers {
		headers[i] = strings.TrimSpace(h)
	}

	return &Dataset{
		Headers: headers,
		Rows:    rows,
	}, nil
}

// ColumnValues extracts all values for a given column index.
func (d *Dataset) ColumnValues(colIdx int) []string {
	var values []string
	for _, row := range d.Rows {
		if colIdx < len(row) {
			values = append(values, strings.TrimSpace(row[colIdx]))
		}
	}
	return values
}

// ColumnNonEmptyValues extracts non-empty values for a given column index.
func (d *Dataset) ColumnNonEmptyValues(colIdx int) []string {
	var values []string
	for _, v := range d.ColumnValues(colIdx) {
		if v != "" {
			values = append(values, v)
		}
	}
	return values
}
