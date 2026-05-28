// toml-to-sql converts TOML files to SQL INSERT statements.
//
// Each TOML table becomes a table name, and its key-value pairs
// become columns in INSERT statements. Sub-tables are flattened
// with dot notation by default.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"

	"github.com/BurntSushi/toml"
)

func escape(s string) string {
	return strings.ReplaceAll(s, "'", "''")
}

func toSQLValue(v any) string {
	switch val := v.(type) {
	case nil:
		return "NULL"
	case bool:
		if val {
			return "1"
		}
		return "0"
	case int64:
		return fmt.Sprintf("%d", val)
	case float64:
		return fmt.Sprintf("%g", val)
	case string:
		return "'" + escape(val) + "'"
	case []any:
		b, _ := json.Marshal(val)
		return "'" + escape(string(b)) + "'"
	case map[string]any:
		b, _ := json.Marshal(val)
		return "'" + escape(string(b)) + "'"
	default:
		return "'" + escape(fmt.Sprintf("%v", v)) + "'"
	}
}

func flattenMap(prefix string, m map[string]any) map[string]any {
	result := make(map[string]any)
	for k, v := range m {
		key := k
		if prefix != "" {
			key = prefix + "." + k
		}
		switch val := v.(type) {
		case map[string]any:
			sub := flattenMap(key, val)
			for sk, sv := range sub {
				result[sk] = sv
			}
		case map[any]any:
			flat := make(map[string]any)
			for mk, mv := range val {
				flat[fmt.Sprintf("%v", mk)] = mv
			}
			sub := flattenMap(key, flat)
			for sk, sv := range sub {
				result[sk] = sv
			}
		default:
			result[key] = v
		}
	}
	return result
}

// asMapStringAny converts a map[any]any to map[string]any
func asMapStringAny(v any) (map[string]any, bool) {
	if m, ok := v.(map[string]any); ok {
		return m, true
	}
	if m, ok := v.(map[any]any); ok {
		flat := make(map[string]any)
		for mk, mv := range m {
			flat[fmt.Sprintf("%v", mk)] = mv
		}
		return flat, true
	}
	return nil, false
}

// extractRows uses reflection to get rows from any slice type containing maps
func extractRows(v any) []map[string]any {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Slice {
		return nil
	}
	rows := make([]map[string]any, 0, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		item := rv.Index(i).Interface()
		if m, ok := asMapStringAny(item); ok {
			rows = append(rows, m)
		}
	}
	return rows
}

func columnNames(m map[string]any) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func generateInsert(table string, data map[string]any, sep string) string {
	keys := columnNames(data)
	cols := make([]string, 0, len(keys))
	vals := make([]string, 0, len(keys))
	for _, k := range keys {
		cols = append(cols, "`"+k+"`")
		vals = append(vals, toSQLValue(data[k]))
	}
	return fmt.Sprintf("INSERT INTO %s (%s)%sVALUES (%s);",
		tableName(table), strings.Join(cols, ", "), sep, strings.Join(vals, ", "))
}

func generateCreateTable(table string, data map[string]any, sep string) string {
	keys := columnNames(data)
	cols := make([]string, 0, len(keys))
	for _, k := range keys {
		colType := "TEXT"
		switch data[k].(type) {
		case int64:
			colType = "INTEGER"
		case float64:
			colType = "REAL"
		case bool:
			colType = "INTEGER"
		}
		cols = append(cols, "  `"+k+"` "+colType)
	}
	return fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s);%s",
		tableName(table), strings.Join(cols, ",\n"), sep)
}

func tableName(name string) string {
	return "`" + name + "`"
}

func processTOML(data map[string]any, output *strings.Builder, createTables bool) {
	tables := make([]string, 0, len(data))
	for k := range data {
		tables = append(tables, k)
	}
	sort.Strings(tables)

	sep := "\n"
	for _, table := range tables {
		val := data[table]

		// Check if it's a slice (array of tables like [[foo]])
		if rows := extractRows(val); len(rows) > 0 {
			for _, row := range rows {
				flat := flattenMap("", row)
				if createTables {
					output.WriteString(generateCreateTable(table, flat, sep))
				}
				output.WriteString(generateInsert(table, flat, sep))
			}
			continue
		}

		// Single map (regular table like [foo])
		if m, ok := asMapStringAny(val); ok {
			flat := flattenMap("", m)
			if createTables {
				output.WriteString(generateCreateTable(table, flat, sep))
			}
			output.WriteString(generateInsert(table, flat, sep))
		}
	}
}

func main() {
	createFlag := flag.Bool("create", false, "Generate CREATE TABLE statements before INSERTs")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: toml-to-sql [options] <file.toml>\n\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Convert TOML files to SQL INSERT statements.\n\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Options:\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	filename := flag.Arg(0)
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	var data map[string]any
	_, err = toml.Decode(string(content), &data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing TOML: %v\n", err)
		os.Exit(1)
	}

	var output strings.Builder
	processTOML(data, &output, *createFlag)

	fmt.Print(output.String())
}
