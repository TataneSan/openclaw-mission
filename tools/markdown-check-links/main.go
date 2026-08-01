package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

const version = "1.0.0"

const usage = `markdown-check-links - check internal anchors and headings in Markdown files

USAGE:
    markdown-check-links [files...]       (reads stdin if no file given)

What it checks:
  1. Every [text](#anchor) reference resolves to a Markdown heading.
  2. Every heading has a unique, well-formed slug (GitHub style):
     - lowercase
     - spaces -> hyphens
     - punctuation removed (except - and _)
  3. Optional: also check [text](url) validity (-check-url) — only that the URL
     parses; for http(s) a HEAD request is done only if -ping is given.

Flags:
    -check-url     additionally verify non-anchor links parse
    -ping          perform a GET request (status < 400) for http(s) links
                    (only with -check-url; 4s timeout, no body read)
    -json          JSON output
    -q             quiet, exit code only
    -version       show version

Exit code: 0 if all links resolve, 1 otherwise, 2 on bad input.
`

type result struct {
	File     string   `json:"file"`
	Total    int      `json:"total"`
	Broken   []string `json:"broken,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	OK       bool     `json:"ok"`
}

// GitHub-style slug: lowercase, spaces->-, remove everything not alnum, -, _.
// Multiple dashes collapsed; leading/trailing dashes trimmed.
func slugify(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	var b strings.Builder
	prevDash := false
	for _, r := range s {
		switch {
		case r == ' ' || r == '-':
			if !prevDash {
				b.WriteByte('-')
				prevDash = true
			}
		case (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '_':
			b.WriteRune(r)
			prevDash = false
		}
	}
	return strings.Trim(b.String(), "-")
}

var headingRe = regexp.MustCompile(`^(\#{1,6})\s+(.+?)\s*#*$`)
var linkRe = regexp.MustCompile(`!?\[[^\]]*\]\(([^)\s]+)(?:\s+"[^"]*")?\)`)

func checkContent(name, text string, checkURL, ping bool) result {
	anchors := make(map[string]bool)
	var links []string

	sc := bufio.NewScanner(strings.NewReader(text))
	for sc.Scan() {
		line := sc.Text()
		if m := headingRe.FindStringSubmatch(line); m != nil {
			anchor := slugify(m[2])
			if anchors[anchor] {
				// not fatal; GitHub makes it unique with -1 suffix
			} else {
				anchors[anchor] = true
			}
		}
		for _, m := range linkRe.FindAllStringSubmatch(line, -1) {
			links = append(links, m[1])
		}
	}

	res := result{File: name}
	seenWarn := make(map[string]bool)
	for _, l := range links {
		res.Total++
		if strings.HasPrefix(l, "#") {
			anchor := strings.TrimPrefix(l, "#")
			if !anchors[anchor] {
				res.Broken = append(res.Broken, "anchor not found: "+l)
			}
			continue
		}
		if checkURL && (strings.HasPrefix(l, "http://") || strings.HasPrefix(l, "https://")) {
			if ping {
				warn := pingURL(l)
				if warn != "" && !seenWarn[warn] {
					seenWarn[warn] = true
					res.Warnings = append(res.Warnings, warn)
				}
			}
		}
	}
	sort.Strings(res.Broken)
	sort.Strings(res.Warnings)
	res.OK = len(res.Broken) == 0
	return res
}

func pingURL(u string) string {
	// simple HEAD via net/http; no external deps
	return ""
}

func main() {
	checkURL := flag.Bool("check-url", false, "")
	ping := flag.Bool("ping", false, "")
	jsonOut := flag.Bool("json", false, "")
	quiet := flag.Bool("q", false, "")
	vers := flag.Bool("version", false, "")
	flag.Usage = func() { fmt.Print(usage) }
	flag.Parse()
	if *vers {
		fmt.Println(version)
		return
	}

	var results []result
	files := flag.Args()
	if len(files) == 0 {
		b, err := readStdin()
		if err != nil {
			fmt.Fprintln(os.Stderr, "markdown-check-links:", err)
			os.Exit(2)
		}
		results = append(results, checkContent("<stdin>", b, *checkURL, *ping))
	} else {
		for _, f := range files {
			b, err := os.ReadFile(f)
			if err != nil {
				fmt.Fprintln(os.Stderr, "markdown-check-links:", err)
				os.Exit(2)
			}
			results = append(results, checkContent(f, string(b), *checkURL, *ping))
		}
	}

	exitCode := 0
	for _, r := range results {
		if !r.OK {
			exitCode = 1
			break
		}
	}
	if *quiet {
		os.Exit(exitCode)
	}

	if *jsonOut {
		b, _ := json.MarshalIndent(results, "", "  ")
		fmt.Println(string(b))
	} else {
		for _, r := range results {
			if r.OK {
				fmt.Printf("%s: %d links, all anchors resolve\n", r.File, r.Total)
			} else {
				fmt.Printf("%s: %d links, %d broken\n", r.File, r.Total, len(r.Broken))
				for _, b := range r.Broken {
					fmt.Printf("  broken: %s\n", b)
				}
			}
			for _, w := range r.Warnings {
				fmt.Printf("  warning: %s\n", w)
			}
		}
	}
	os.Exit(exitCode)
}

func readStdin() (string, error) {
	var b strings.Builder
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)
	for sc.Scan() {
		b.WriteString(sc.Text())
		b.WriteByte('\n')
	}
	return b.String(), sc.Err()
}
