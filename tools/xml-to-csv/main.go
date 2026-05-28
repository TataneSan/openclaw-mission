package main

import (
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"sort"
)

func main() {
	args := os.Args[1:]
	if len(args) > 0 && (args[0] == "--help" || args[0] == "-h") {
		printUsage()
		return
	}

	var input string
	if len(args) > 0 {
		input = args[0]
	}

	var data []byte
	var err error

	if input == "" || input == "-" {
		data, err = io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
			os.Exit(1)
		}
	} else {
		data, err = os.ReadFile(input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	}

	var doc xml.Document
	if err := xml.Unmarshal(data, &doc); err != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid XML: %v\n", err)
		os.Exit(1)
	}

	elements, keys := extractElements(doc)
	if len(elements) == 0 {
		return
	}

	sort.Strings(keys)

	writer := csv.NewWriter(os.Stdout)
	defer writer.Flush()

	writer.Write(keys)

	for _, el := range elements {
		row := make([]string, len(keys))
		for i, k := range keys {
			row[i] = el[k]
		}
		writer.Write(row)
	}
}

func extractElements(doc xml.Document) ([]map[string]string, []string) {
	var elements []map[string]string
	keySet := make(map[string]bool)

	// Find the repeating element (children of root)
	if len(doc.Declarations) > 0 {
		_ = doc.Declarations
	}

	if doc.Root == nil {
		return elements, nil
	}

	root := doc.Root
	if root.Children == nil {
		// Root itself is the only element
		el := flattenXML(root)
		for k := range el {
			keySet[k] = true
		}
		elements = append(elements, el)
	} else {
		for _, child := range root.Children {
			el := flattenXML(child)
			for k := range el {
				keySet[k] = true
			}
			elements = append(elements, el)
		}
	}

	var keys []string
	for k := range keySet {
		keys = append(keys, k)
	}
	return elements, keys
}

func flattenXML(el *xml.Element) map[string]string {
	result := make(map[string]string)

	if el.Text != "" {
		result["value"] = el.Text
	}

	for _, attr := range el.Attributes {
		result[attr.Name.Local] = attr.Value
	}

	if el.Children != nil {
		for _, child := range el.Children {
			prefix := child.Name.Local
			if child.Text != "" && len(child.Children) == 0 {
				result[prefix] = child.Text
			} else {
				sub := flattenXML(child)
				for k, v := range sub {
					result[prefix+"."+k] = v
				}
			}
		}
	}

	return result
}

func printUsage() {
	fmt.Println(`Usage: xml-to-csv [FILE]

Convert XML files to CSV.

Reads from stdin if no file is provided.

Examples:
  xml-to-csv data.xml
  cat data.xml | xml-to-csv
  xml-to-csv < data.xml > output.csv`)
}
