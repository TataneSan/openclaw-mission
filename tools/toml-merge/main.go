// toml-merge merges multiple TOML files with configurable priority.
// Later files override earlier ones by default.
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	files := []string{}
	output := ""
	format := "toml"
	reverse := false

	args := os.Args[1:]
	i := 0
	for i < len(args) {
		switch args[i] {
		case "-o", "--output":
			if i+1 >= len(args) {
				die("flag %s requires a value", args[i])
			}
			i++
			output = args[i]
		case "-f", "--format":
			if i+1 >= len(args) {
				die("flag %s requires a value", args[i])
			}
			i++
			format = args[i]
		case "-r", "--reverse":
			reverse = true
		case "-h", "--help":
			usage()
			return
		default:
			files = append(files, args[i])
		}
		i++
	}

	if len(files) < 2 {
		die("at least 2 input files are required")
	}

	if format != "toml" && format != "json" {
		die("unsupported format %q, must be 'toml' or 'json'", format)
	}

	merged := map[string]any{}

	for _, f := range files {
		data, err := readFile(f)
		if err != nil {
			die("failed to read %s: %v", f, err)
		}
		parsed := map[string]any{}
		_, err = toml.Decode(string(data), &parsed)
		if err != nil {
			die("failed to parse %s: %v", f, err)
		}
		deepMerge(merged, parsed)
	}

	if reverse {
		// When reversing, earlier files take priority — re-process in reverse
		// Actually, reverse means first file wins, so we merge in reverse order
		merged = map[string]any{}
		for i := len(files) - 1; i >= 0; i-- {
			data, err := readFile(files[i])
			if err != nil {
				die("failed to read %s: %v", files[i], err)
			}
			parsed := map[string]any{}
			_, err = toml.Decode(string(data), &parsed)
			if err != nil {
				die("failed to parse %s: %v", files[i], err)
			}
			deepMerge(merged, parsed)
		}
	}

	out := getWriter(output)
	defer closeWriter(out)

	switch format {
	case "toml":
		enc := toml.NewEncoder(out)
		if err := enc.Encode(merged); err != nil {
			die("failed to encode TOML: %v", err)
		}
	case "json":
		enc := json.NewEncoder(out)
		enc.SetIndent("", "  ")
		if err := enc.Encode(merged); err != nil {
			die("failed to encode JSON: %v", err)
		}
	}
}

// deepMerge merges src into dst. Values in src override dst.
func deepMerge(dst, src map[string]any) {
	for k, v := range src {
		if dv, ok := dst[k]; ok {
			if dmap, dok := dv.(map[string]any); dok {
				if smap, sok := v.(map[string]any); sok {
					dst[k] = dmap
					deepMerge(dmap, smap)
					continue
				}
			}
		}
		dst[k] = v
	}
}

func readFile(path string) ([]byte, error) {
	if path == "-" {
		return io.ReadAll(os.Stdin)
	}
	return os.ReadFile(filepath.Clean(path))
}

func getWriter(path string) io.Writer {
	if path == "" || path == "-" {
		return os.Stdout
	}
	f, err := os.Create(filepath.Clean(path))
	if err != nil {
		die("failed to create output file: %v", err)
	}
	return f
}

func closeWriter(w io.Writer) {
	if c, ok := w.(io.Closer); ok {
		c.Close()
	}
}

func usage() {
	fmt.Println(`toml-merge — Merge multiple TOML files

Usage:
  toml-merge [OPTIONS] FILE...

Options:
  -o, --output FILE   Output file (default: stdout)
  -f, --format FMT    Output format: toml, json (default: toml)
  -r, --reverse       First file takes priority (default: last wins)
  -h, --help          Show this help message

Examples:
  toml-merge base.toml override.toml
  toml-merge -o merged.toml a.toml b.toml c.toml
  toml-merge -f json config.toml extra.toml
  toml-merge -r default.toml user.toml  # user values win

Files are merged left to right. Later files override earlier ones.
Nested tables are merged recursively.`)
}

func die(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "toml-merge: error: "+format+"\n", args...)
	os.Exit(1)
}
