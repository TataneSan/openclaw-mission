// csv-schema-infer infers column data types from CSV files.
//
// It analyzes each column and determines the most likely data type
// (integer, float, boolean, date, datetime, email, url, ip, string)
// with confidence scores and null detection.
package main

import (
	"flag"
	"fmt"
	"os"
	"unicode/utf8"

	"github.com/TataneSan/csv-schema-infer/internal/inferrer"
	"github.com/TataneSan/csv-schema-infer/internal/reader"
	"github.com/TataneSan/csv-schema-infer/internal/writer"
)

func main() {
	cmdInfer := flag.NewFlagSet("infer", flag.ExitOnError)
	inferFile := cmdInfer.String("f", "", "CSV file to analyze")
	inferDelim := cmdInfer.String("d", ",", "CSV delimiter (comma, tab, pipe, semicolon)")
	inferFormat := cmdInfer.String("format", "table", "Output format: table, json, sql")
	inferTable := cmdInfer.String("table", "table", "Table name for SQL output")

	cmdSamples := flag.NewFlagSet("samples", flag.ExitOnError)
	samplesFile := cmdSamples.String("f", "", "CSV file to show samples")
	samplesDelim := cmdSamples.String("d", ",", "CSV delimiter")
	samplesCount := cmdSamples.Int("n", 3, "Number of sample values per column")

	cmdNulls := flag.NewFlagSet("nulls", flag.ExitOnError)
	nullsFile := cmdNulls.String("f", "", "CSV file to check for nulls")
	nullsDelim := cmdNulls.String("d", ",", "CSV delimiter")

	cmdColumns := flag.NewFlagSet("columns", flag.ExitOnError)
	columnsFile := cmdColumns.String("f", "", "CSV file to list columns")
	columnsDelim := cmdColumns.String("d", ",", "CSV delimiter")

	switch {
	case len(os.Args) < 2:
		printUsage()
		os.Exit(1)

	case os.Args[1] == "infer":
		cmdInfer.Parse(os.Args[2:])
		if err := runInfer(*inferFile, *inferDelim, *inferFormat, *inferTable); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

	case os.Args[1] == "samples":
		cmdSamples.Parse(os.Args[2:])
		if err := runSamples(*samplesFile, *samplesDelim, *samplesCount); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

	case os.Args[1] == "nulls":
		cmdNulls.Parse(os.Args[2:])
		if err := runNulls(*nullsFile, *nullsDelim); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

	case os.Args[1] == "columns":
		cmdColumns.Parse(os.Args[2:])
		if err := runColumns(*columnsFile, *columnsDelim); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

	case os.Args[1] == "--help" || os.Args[1] == "-h" || os.Args[1] == "help":
		printUsage()

	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", os.Args[1])
		printUsage()
		os.Exit(1)
	}
}

func parseDelim(delim string) rune {
	switch delim {
	case "tab", "\\t", "t":
		return '\t'
	case "pipe", "|", "p":
		return '|'
	case "semicolon", ";", "s":
		return ';'
	default:
		r, _ := utf8.DecodeRuneInString(delim)
		return r
	}
}

func runInfer(file, delimStr, format, tableName string) error {
	if file == "" {
		return fmt.Errorf("file is required, use -f")
	}

	dataset, err := reader.ReadFile(file, parseDelim(delimStr))
	if err != nil {
		return err
	}

	var columns []inferrer.ColumnSchema
	for i, header := range dataset.Headers {
		values := dataset.ColumnValues(i)
		schema := inferrer.Infer(header, values)
		columns = append(columns, *schema)
	}

	result := &writer.SchemaResult{Columns: columns}

	switch format {
	case "json":
		return writer.WriteJSON(os.Stdout, result)
	case "sql":
		return writer.WriteSQL(os.Stdout, tableName, result)
	case "table":
		return writer.WriteTable(os.Stdout, result)
	default:
		return fmt.Errorf("unknown format: %s (valid: table, json, sql)", format)
	}
}

func runSamples(file, delimStr string, count int) error {
	if file == "" {
		return fmt.Errorf("file is required, use -f")
	}

	dataset, err := reader.ReadFile(file, parseDelim(delimStr))
	if err != nil {
		return err
	}

	fmt.Printf("Samples from %d columns:\n\n", len(dataset.Headers))

	for i, header := range dataset.Headers {
		values := dataset.ColumnValues(i)
		fmt.Printf("[%s]\n", header)
		limit := count
		if len(values) < limit {
			limit = len(values)
		}
		for j := 0; j < limit; j++ {
			if values[j] == "" {
				fmt.Printf("  %d: (null)\n", j+1)
			} else {
				fmt.Printf("  %d: %s\n", j+1, values[j])
			}
		}
		fmt.Println()
	}

	return nil
}

func runNulls(file, delimStr string) error {
	if file == "" {
		return fmt.Errorf("file is required, use -f")
	}

	dataset, err := reader.ReadFile(file, parseDelim(delimStr))
	if err != nil {
		return err
	}

	totalRows := len(dataset.Rows)
	fmt.Printf("Null analysis (%d rows):\n\n", totalRows)

	for i, header := range dataset.Headers {
		values := dataset.ColumnValues(i)
		nullCount := 0
		for _, v := range values {
			if v == "" {
				nullCount++
			}
		}
		pct := 0.0
		if totalRows > 0 {
			pct = float64(nullCount) / float64(totalRows) * 100
		}
		status := "OK"
		if pct > 0 {
			status = "HAS NULLS"
		}
		fmt.Printf("%-20s %4d/%-4d (%5.1f%%) %s\n", header, nullCount, totalRows, pct, status)
	}

	return nil
}

func runColumns(file, delimStr string) error {
	if file == "" {
		return fmt.Errorf("file is required, use -f")
	}

	dataset, err := reader.ReadFile(file, parseDelim(delimStr))
	if err != nil {
		return err
	}

	fmt.Printf("Columns (%d):\n", len(dataset.Headers))
	for i, header := range dataset.Headers {
		fmt.Printf("  %3d. %s\n", i+1, header)
	}

	return nil
}

func printUsage() {
	fmt.Println(`csv-schema-infer - Infer column data types from CSV files

Usage:
  csv-schema-infer <command> [options]

Commands:
  infer     Infer schema from CSV file
  samples   Show sample values per column
  nulls     Analyze null values per column
  columns   List column names

Infer Options:
  -f FILE       CSV file to analyze (required)
  -d DELIM      Delimiter: , (default), tab, pipe, semicolon
  -format FMT   Output: table (default), json, sql
  -table NAME   Table name for SQL output (default: "table")

Samples Options:
  -f FILE       CSV file (required)
  -d DELIM      Delimiter
  -n COUNT      Number of samples per column (default: 3)

Nulls Options:
  -f FILE       CSV file (required)
  -d DELIM      Delimiter

Columns Options:
  -f FILE       CSV file (required)
  -d DELIM      Delimiter

Examples:
  csv-schema-infer infer -f data.csv
  csv-schema-infer infer -f data.csv -format json
  csv-schema-infer infer -f data.csv -format sql -table users
  csv-schema-infer samples -f data.csv -n 5
  csv-schema-infer nulls -f data.csv
  csv-schema-infer columns -f data.csv
  csv-schema-infer infer -f data.tsv -d tab
  csv-schema-infer infer -f data.csv -d pipe -format json

Detected Types:
  integer, float, boolean, date, datetime, email, url, ip, string`)
}
