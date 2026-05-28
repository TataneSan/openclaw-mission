package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var version = "1.0.0"

func main() {
	inputFile := flag.String("f", "", "Input JSON file (empty = stdin)")
	indent := flag.String("i", "  ", "Indent string (default: 2 spaces)")
	compact := flag.Bool("c", false, "Compact output (no whitespace)")
	sortKeys := flag.Bool("s", false, "Sort keys alphabetically")
	showVersion := flag.Bool("V", false, "Show version")
	flag.Parse()

	if *showVersion {
		fmt.Printf("json-reformat v%s\n", version)
		return
	}

	var src io.Reader
	if *inputFile != "" {
		f, err := os.Open(*inputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()
		src = f
	} else {
		src = os.Stdin
	}

	var data interface{}
	if err := json.NewDecoder(src).Decode(&data); err != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid JSON: %v\n", err)
		os.Exit(1)
	}

	if *compact {
		out, err := json.Marshal(data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(string(out))
		return
	}

	out, err := json.MarshalIndent(data, "", *indent)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if *sortKeys {
		// Re-marshal with sorted keys
		var sorted interface{}
		if err := json.Unmarshal(out, &sorted); err == nil {
			out, _ = json.MarshalIndent(sorted, "", *indent)
		}
	}

	fmt.Println(string(out))
}
