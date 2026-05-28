// toml-to-sql converts TOML files to SQL INSERT statements.
//
// It reads TOML data from a file or stdin and generates CREATE TABLE
// and INSERT statements for PostgreSQL or MySQL.
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/BurntSushi/toml"
)

type config struct {
	file    string
	table   string
	format  string
	dialect string
}

func main() {
	cfg := parseFlags()

	data := loadData(cfg.file)

	var tables map[string]interface{}
	if _, err := toml.Decode(string(data), &tables); err != nil {
		fmt.Fprintf(os.Stderr, "error: failed to parse TOML: %v\n", err)
		os.Exit(1)
	}

	if len(tables) == 0 {
		fmt.Fprintf(os.Stderr, "error: TOML file is empty or contains no tables\n")
		os.Exit(1)
	}

	for tableName, tableData := range tables {
		if cfg.table != "" {
			tableName = cfg.table
		}
		rows, err := extractRows(tableData)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: failed to extract rows from table '%s': %v\n", tableName, err)
			continue
		}
		if len(rows) == 0 {
			continue
		}
		printTable(tableName, rows, cfg.dialect)
	}
}

func parseFlags() config {
	cfg := config{table: "", format: "text", dialect: "postgres"}
	args := os.Args[1:]
	i := 0
	for i < len(args) {
		switch args[i] {
		case "--table", "-t":
			if i+1 < len(args) {
				cfg.table = args[i+1]
				i += 2
			} else {
				fmt.Fprintf(os.Stderr, "error: --table requires a value\n")
				os.Exit(1)
			}
		case "--format", "-f":
			if i+1 < len(args) {
				cfg.format = args[i+1]
				i += 2
			} else {
				fmt.Fprintf(os.Stderr, "error: --format requires a value\n")
				os.Exit(1)
			}
		case "--dialect", "-d":
			if i+1 < len(args) {
				cfg.dialect = args[i+1]
				i += 2
			} else {
				fmt.Fprintf(os.Stderr, "error: --dialect requires a value\n")
				os.Exit(1)
			}
		case "--file", "--":
			if i+1 < len(args) {
				cfg.file = args[i+1]
				i += 2
			} else {
				fmt.Fprintf(os.Stderr, "error: --file requires a value\n")
				os.Exit(1)
			}
		case "--help", "-h":
			printUsage()
			os.Exit(0)
		default:
			if args[i][0] != '-' {
				cfg.file = args[i]
			}
			i++
		}
	}
	return cfg
}

func printUsage() {
	fmt.Println("Usage: toml-to-sql [OPTIONS] [FILE]")
	fmt.Println()
	fmt.Println("Converts TOML files to SQL INSERT statements.")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -f, --format FORMAT   Output format: text, json (default: text)")
	fmt.Println("  -t, --table NAME      Table name (default: TOML table name)")
	fmt.Println("  -d, --dialect DIALECT SQL dialect: postgres, mysql (default: postgres)")
	fmt.Println("      --file FILE       Input file (default: stdin)")
	fmt.Println("  -h, --help            Show this help message")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  toml-to-sql data.toml")
	fmt.Println("  cat data.toml | toml-to-sql")
	fmt.Println("  toml-to-sql --dialect mysql --table users data.toml")
}

func loadData(file string) []byte {
	var reader io.Reader
	if file != "" {
		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: failed to open file '%s': %v\n", file, err)
			os.Exit(1)
		}
		defer f.Close()
		reader = f
	} else {
		reader = os.Stdin
	}
	buf := bufio.NewReader(reader)
	data, err := io.ReadAll(buf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: failed to read input: %v\n", err)
		os.Exit(1)
	}
	return data
}

func extractRows(data interface{}) ([]map[string]interface{}, error) {
	switch v := data.(type) {
	case map[string]interface{}:
		return []map[string]interface{}{v}, nil
	case []interface{}:
		var rows []map[string]interface{}
		for _, item := range v {
			m, ok := item.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("array items must be tables")
			}
			rows = append(rows, m)
		}
		return rows, nil
	case []map[string]interface{}:
		return v, nil
	default:
		return nil, fmt.Errorf("unsupported type: %T", data)
	}
}

func printTable(tableName string, rows []map[string]interface{}, dialect string) {
	if len(rows) == 0 {
		return
	}

	keys := getAllKeys(rows)
	sort.Strings(keys)
	types := inferTypes(rows, keys)

	escapedName := escapeIdentifier(tableName, dialect)

	fmt.Printf("-- Table: %s\n", tableName)
	fmt.Printf("CREATE TABLE %s (\n", escapedName)
	for i, k := range keys {
		escaped := escapeIdentifier(k, dialect)
		sqlType := goTypeToSQL(types[i], dialect)
		comma := ","
		if i == len(keys)-1 {
			comma = ""
		}
		fmt.Printf("    %s %s%s\n", escaped, sqlType, comma)
	}
	fmt.Println(");")
	fmt.Println()

	for _, row := range rows {
		values := make([]string, len(keys))
		for i, k := range keys {
			values[i] = formatValue(row[k], types[i], dialect)
		}
		escapedKeys := make([]string, len(keys))
		for i, k := range keys {
			escapedKeys[i] = escapeIdentifier(k, dialect)
		}
		fmt.Printf("INSERT INTO %s (%s) VALUES (%s);\n",
			escapedName,
			strings.Join(escapedKeys, ", "),
			strings.Join(values, ", "))
	}
	fmt.Println()
}

func getAllKeys(rows []map[string]interface{}) []string {
	keySet := make(map[string]bool)
	for _, row := range rows {
		for k := range row {
			keySet[k] = true
		}
	}
	keys := make([]string, 0, len(keySet))
	for k := range keySet {
		keys = append(keys, k)
	}
	return keys
}

func inferTypes(rows []map[string]interface{}, keys []string) []string {
	types := make([]string, len(keys))
	for i, k := range keys {
		types[i] = "text"
		allBool := true
		allInt := true
		allFloat := true
		for _, row := range rows {
			v, ok := row[k]
			if !ok || v == nil {
				continue
			}
			switch val := v.(type) {
			case bool:
				allInt = false
				allFloat = false
			case int64, int32, int, uint64, uint32, uint:
				allBool = false
			case float64:
				allBool = false
				if float64(int64(val)) != val {
					allInt = false
				}
			case string:
				allBool = false
				allInt = false
				allFloat = false
			default:
				allBool = false
				allInt = false
				allFloat = false
			}
		}
		if allBool {
			types[i] = "bool"
		} else if allInt {
			types[i] = "int"
		} else if allFloat {
			types[i] = "float"
		}
	}
	return types
}

func goTypeToSQL(gotype string, dialect string) string {
	switch gotype {
	case "bool":
		if dialect == "mysql" {
			return "TINYINT(1)"
		}
		return "BOOLEAN"
	case "int":
		return "INTEGER"
	case "float":
		return "REAL"
	default:
		return "TEXT"
	}
}

func formatValue(v interface{}, gotype string, dialect string) string {
	if v == nil {
		return "NULL"
	}
	switch val := v.(type) {
	case bool:
		if dialect == "mysql" {
			if val {
				return "1"
			}
			return "0"
		}
		if val {
			return "TRUE"
		}
		return "FALSE"
	case int64, int32, int, uint64, uint32, uint:
		return fmt.Sprintf("%d", val)
	case float64:
		if float64(int64(val)) == val {
			return fmt.Sprintf("%d", int64(val))
		}
		return fmt.Sprintf("%g", val)
	case string:
		escaped := strings.ReplaceAll(val, "'", "''")
		return fmt.Sprintf("'%s'", escaped)
	default:
		data, _ := json.Marshal(val)
		escaped := strings.ReplaceAll(string(data), "'", "''")
		return fmt.Sprintf("'%s'", escaped)
	}
}

func escapeIdentifier(name string, dialect string) string {
	if dialect == "mysql" {
		escaped := strings.ReplaceAll(name, "`", "``")
		return fmt.Sprintf("`%s`", escaped)
	}
	escaped := strings.ReplaceAll(name, "\"", "\"\"")
	return fmt.Sprintf("\"%s\"", escaped)
}
