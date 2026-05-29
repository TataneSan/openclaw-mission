// xml-to-ini converts XML files to INI format.
//
// XML elements become INI sections, attributes and child elements become key-value pairs.
//
// Usage:
//   xml-to-ini [options] <xml-file>
//   cat data.xml | xml-to-ini -
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// section maps section name to key-value pairs.
type section struct {
	name string
	kv   map[string]string
}

// doc holds all INI sections.
type doc struct {
	byName map[string]*section
	order  []string
}

func newDoc() *doc {
	return &doc{byName: make(map[string]*section)}
}

func (d *doc) getOrCreate(name string) *section {
	if s, ok := d.byName[name]; ok {
		return s
	}
	s := &section{name: name, kv: make(map[string]string)}
	d.byName[name] = s
	d.order = append(d.order, name)
	return s
}

// parseRoot reads the root XML element and builds the INI document.
func parseRoot(dec *xml.Decoder, d *doc, root xml.StartElement) error {
	rootSection := d.getOrCreate(sanitize(root.Name.Local))

	for _, a := range root.Attr {
		rootSection.kv[sanitize(a.Name.Local)] = a.Value
	}

	for {
		tok, err := nextToken(dec)
		if err != nil {
			return err
		}

		switch tk := tok.(type) {
		case xml.StartElement:
			if err := parseChild(dec, d, tk, rootSection, root.Name); err != nil {
				return err
			}
		case xml.EndElement:
			if tk.Name == root.Name {
				return nil
			}
		}
	}
}

// parseChild handles one child element.
// If the child has no sub-children, it becomes a key in currentSection.
// If it has sub-children, it becomes a new section.
func parseChild(dec *xml.Decoder, d *doc, elem xml.StartElement, currentSection *section, parent xml.Name) error {
	childName := elem.Name.Local

	// Peek: read the next token to decide if this is a simple or complex element.
	first, err := nextToken(dec)
	if err != nil {
		return err
	}

	switch ft := first.(type) {
	case xml.EndElement:
		if ft.Name == elem.Name {
			// Empty element
			currentSection.kv[sanitize(childName)] = ""
			return nil
		}
		// Mismatched end tag — fall through to handle as text

	case xml.CharData:
		// Read text content, then expect EndElement matching elem
		var buf strings.Builder
		buf.Write([]byte(ft))

		for {
			next, err := nextToken(dec)
			if err != nil {
				return err
			}

			switch n := next.(type) {
			case xml.CharData:
				buf.Write([]byte(n))
			case xml.EndElement:
				if n.Name == elem.Name {
					currentSection.kv[sanitize(childName)] = strings.TrimSpace(buf.String())
					return nil
				}
				// Mismatched — shouldn't happen in valid XML
				return fmt.Errorf("mismatched end tag: expected %s, got %s", elem.Name, n.Name)
			case xml.StartElement:
				// Text followed by a child element — treat as complex element
				// Discard the text we collected so far
				if err := handleComplex(dec, d, elem, n); err != nil {
					return err
				}
				return nil
			}
		}

	case xml.StartElement:
		// Complex element — has children
		if err := handleComplex(dec, d, elem, ft); err != nil {
			return err
		}
		return nil
	}

	return nil
}

// handleComplex processes an element that has child elements.
// firstChild is the already-consumed first child StartElement.
func handleComplex(dec *xml.Decoder, d *doc, elem xml.StartElement, firstChild xml.StartElement) error {
	childSection := d.getOrCreate(sanitize(elem.Name.Local))

	// Add the element's own attributes
	for _, a := range elem.Attr {
		childSection.kv[sanitize(a.Name.Local)] = a.Value
	}

	// Process the first child (already consumed)
	if err := parseGrandchild(dec, d, firstChild, childSection, elem.Name); err != nil {
		return err
	}

	// Continue reading siblings until we hit the EndElement matching elem
	for {
		tok, err := nextToken(dec)
		if err != nil {
			return err
		}

		switch tk := tok.(type) {
		case xml.StartElement:
			if err := parseGrandchild(dec, d, tk, childSection, elem.Name); err != nil {
				return err
			}
		case xml.EndElement:
			if tk.Name == elem.Name {
				return nil
			}
		}
	}
}

// parseGrandchild processes a child within a complex (section) element.
func parseGrandchild(dec *xml.Decoder, d *doc, elem xml.StartElement, section *section, parent xml.Name) error {
	name := elem.Name.Local

	first, err := nextToken(dec)
	if err != nil {
		return err
	}

	switch ft := first.(type) {
	case xml.EndElement:
		if ft.Name == elem.Name {
			section.kv[sanitize(name)] = ""
			return nil
		}

	case xml.CharData:
		var buf strings.Builder
		buf.Write([]byte(ft))

		for {
			next, err := nextToken(dec)
			if err != nil {
				return err
			}

			switch n := next.(type) {
			case xml.CharData:
				buf.Write([]byte(n))
			case xml.EndElement:
				if n.Name == elem.Name {
					section.kv[sanitize(name)] = strings.TrimSpace(buf.String())
					return nil
				}
				return fmt.Errorf("mismatched end tag: expected %s, got %s", elem.Name, n.Name)
			case xml.StartElement:
				// Text followed by nested element — flatten: store text, then recurse
				section.kv[sanitize(name)] = strings.TrimSpace(buf.String())
				if err := parseGrandchild(dec, d, n, section, parent); err != nil {
					return err
				}
				return nil
			}
		}

	case xml.StartElement:
		// Nested element with its own children — recurse
		if err := parseGrandchild(dec, d, ft, section, parent); err != nil {
			return err
		}
		return nil
	}

	return nil
}

func nextToken(dec *xml.Decoder) (xml.Token, error) {
	for {
		tok, err := dec.Token()
		if err != nil {
			return nil, err
		}
		// Skip comments, processing instructions, whitespace-only CharData
		switch t := tok.(type) {
		case xml.Comment, xml.ProcInst:
			continue
		case xml.CharData:
			if strings.TrimSpace(string(t)) == "" {
				continue
			}
		}
		return tok, nil
	}
}

func sanitize(name string) string {
	r := strings.ReplaceAll(name, "-", "_")
	r = strings.ReplaceAll(r, " ", "_")
	r = strings.ReplaceAll(r, ".", "_")
	if len(r) > 0 && r[0] != '_' && (r[0] < 'a' || r[0] > 'z') {
		r = "key_" + r
	}
	return r
}

func writeINI(w io.Writer, d *doc) {
	sort.Strings(d.order)

	first := true
	for _, name := range d.order {
		s := d.byName[name]
		if !first {
			fmt.Fprintln(w)
		}
		first = false

		fmt.Fprintf(w, "[%s]\n", s.name)

		keys := make([]string, 0, len(s.kv))
		for k := range s.kv {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, k := range keys {
			fmt.Fprintf(w, "%s = %s\n", k, s.kv[k])
		}
	}
}

func printUsage() {
	fmt.Println("xml-to-ini - Convert XML files to INI format")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  xml-to-ini [options] <xml-file>")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -o, --output <file>  Output file (default: same name with .ini extension)")
	fmt.Println("  -h, --help           Show this help message")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  xml-to-ini config.xml")
	fmt.Println("  xml-to-ini -o config.ini config.xml")
	fmt.Println("  cat data.xml | xml-to-ini -")
}

func main() {
	args := os.Args[1:]

	var inputFile, outputFile string

	i := 0
	for i < len(args) {
		switch args[i] {
		case "-h", "--help":
			printUsage()
			os.Exit(0)
		case "-o", "--output":
			if i+1 >= len(args) {
				fmt.Fprintln(os.Stderr, "error: -o requires an argument")
				os.Exit(1)
			}
			i++
			outputFile = args[i]
		default:
			inputFile = args[i]
		}
		i++
	}

	if inputFile == "" {
		inputFile = "-"
	}

	var reader io.Reader
	if inputFile == "-" {
		reader = os.Stdin
	} else {
		f, err := os.Open(inputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: cannot open file: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()
		reader = f
	}

	dec := xml.NewDecoder(reader)
	d := newDoc()

	found := false
	for {
		tok, err := dec.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Fprintf(os.Stderr, "error: invalid XML: %v\n", err)
			os.Exit(1)
		}

		if se, ok := tok.(xml.StartElement); ok {
			if !found {
				if err := parseRoot(dec, d, se); err != nil {
					fmt.Fprintf(os.Stderr, "error: failed to parse XML: %v\n", err)
					os.Exit(1)
				}
				found = true
			}
		}
	}

	if !found {
		fmt.Fprintln(os.Stderr, "error: no XML element found in input")
		os.Exit(1)
	}

	if outputFile == "" {
		if inputFile == "-" {
			outputFile = "-"
		} else {
			ext := filepath.Ext(inputFile)
			outputFile = strings.TrimSuffix(inputFile, ext) + ".ini"
		}
	}

	var writer io.Writer
	if outputFile == "-" {
		writer = os.Stdout
	} else {
		f, err := os.Create(outputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: cannot create output file: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()
		writer = f
	}

	writeINI(writer, d)

	if outputFile != "-" {
		fmt.Fprintf(os.Stderr, "converted: %s -> %s\n", inputFile, outputFile)
	}
}
