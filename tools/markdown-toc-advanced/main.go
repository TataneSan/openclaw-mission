// markdown-toc-advanced generates or updates a table of contents in Markdown
// files. Features: heading levels, GitHub slugs, exclusion zones, max depth,
// multiple files, JSON, in-place update.
package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type tocOptions struct {
	MinDepth  int    `json:"min_depth"`
	MaxDepth  int    `json:"max_depth"`
	Ordered   bool   `json:"ordered"`
	Link      bool   `json:"link"`
	Title     string `json:"title"`
	Insert    bool   `json:"insert"`
	StartMark string `json:"start_mark"`
	EndMark   string `json:"end_mark"`
}

type tocEntry struct {
	Level int    `json:"level"`
	Title string `json:"title"`
	Slug  string `json:"slug"`
	Line  int    `json:"line"`
}

var headingRE = regexp.MustCompile(`^(#{1,6})\s+(.+?)\s*#*\s*$`)

func slugify(s string) string {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "-")
	var b strings.Builder
	for _, r := range s {
		switch {
		case r >= 'a' && r <= 'z' || r >= '0' && r <= '9' || r == '-':
			b.WriteRune(r)
		case r == '_' || r == '/':
			b.WriteRune('-')
		}
	}
	// collapse multiple dashes
	s = b.String()
	for strings.Contains(s, "--") {
		s = strings.ReplaceAll(s, "--", "-")
	}
	return strings.Trim(s, "-")
}

func parse(r interface{ Read([]byte) (int, error) }, maxDepth int) ([]tocEntry, error) {
	var entries []tocEntry
	sc := bufio.NewScanner(r)
	lineNo := 0
	inCode := false
	for sc.Scan() {
		lineNo++
		line := sc.Text()
		if strings.HasPrefix(line, "```") {
			inCode = !inCode
			continue
		}
		if inCode {
			continue
		}
		m := headingRE.FindStringSubmatch(line)
		if m == nil {
			continue
		}
		level := len(m[1])
		if level > maxDepth {
			continue
		}
		title := strings.TrimSpace(m[2])
		entries = append(entries, tocEntry{
			Level: level,
			Title: title,
			Slug:  slugify(title),
			Line:  lineNo,
		})
	}
	return entries, sc.Err()
}

func generate(entries []tocEntry, minDepth int, ordered, link bool) string {
	var b strings.Builder
	for _, e := range entries {
		if e.Level < minDepth {
			continue
		}
		indent := strings.Repeat("  ", e.Level-minDepth)
		title := e.Title
		if link {
			title = fmt.Sprintf("[%s](#%s)", title, e.Slug)
		}
		if ordered {
			fmt.Fprintf(&b, "%s1. %s\n", indent, title)
		} else {
			fmt.Fprintf(&b, "%s- %s\n", indent, title)
		}
	}
	return strings.TrimRight(b.String(), "\n")
}

func main() {
	files := flag.Args()
	maxDepth := flag.Int("d", 6, "max heading depth")
	minDepth := flag.Int("min", 1, "min heading depth")
	ordered := flag.Bool("ordered", false, "ordered list style")
	link := flag.Bool("l", true, "create links")
	title := flag.String("t", "## Table of Contents", "TOC title")
	insert := flag.Bool("i", false, "insert TOC into file in-place")
	startMark := flag.String("start", "<!-- TOC -->", "TOC start marker")
	endMark := flag.String("end", "<!-- /TOC -->", "TOC end marker")
	jsonOut := flag.Bool("json", false, "output JSON")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "markdown-toc-advanced — generate TOC for Markdown files\n\n")
		fmt.Fprintf(os.Stderr, "Usage:\n")
		fmt.Fprintf(os.Stderr, "  markdown-toc-advanced file.md\n")
		fmt.Fprintf(os.Stderr, "  markdown-toc-advanced -i file.md    # insert in-place between markers\n")
		fmt.Fprintf(os.Stderr, "  markdown-toc-advanced -d 3 -ordered file.md\n\nFlags:\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	files = flag.Args()

	if len(files) == 0 {
		fmt.Fprintln(os.Stderr, "error: no files")
		os.Exit(2)
	}

	for _, path := range files {
		data, err := os.ReadFile(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error opening %s: %v\n", path, err)
			continue
		}
		contents := string(data)

		entries, err := parse(strings.NewReader(contents), *maxDepth)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error parsing %s: %v\n", path, err)
			continue
		}

		if *jsonOut {
			enc := json.NewEncoder(os.Stdout)
			enc.SetIndent("", "  ")
			enc.Encode(map[string]interface{}{
				"file":    path,
				"entries": entries,
			})
			continue
		}

		toc := generate(entries, *minDepth, *ordered, *link)
		if *insert {
			if contents == "" {
				data, _ := os.ReadFile(path)
				contents = string(data)
			}
			if !strings.Contains(contents, *startMark) {
				fmt.Fprintf(os.Stderr, "%s: start marker %q not found\n", path, *startMark)
				os.Exit(1)
			}
			if !strings.Contains(contents, *endMark) {
				fmt.Fprintf(os.Stderr, "%s: end marker %q not found\n", path, *endMark)
				os.Exit(1)
			}
			parts := strings.SplitN(contents, *startMark, 2)
			before := parts[0]
			rest := parts[1]
			after := ""
			if p2 := strings.SplitN(rest, *endMark, 2); len(p2) == 2 {
				after = p2[1]
			}
			newBody := before + *startMark + "\n" + *title + "\n\n" + toc + "\n" + *endMark + after
			if err := os.WriteFile(path, []byte(newBody), 0644); err != nil {
				fmt.Fprintf(os.Stderr, "error writing %s: %v\n", path, err)
				continue
			}
			fmt.Printf("%s updated\n", path)
			continue
		}
		fmt.Println(toc)
	}
}
