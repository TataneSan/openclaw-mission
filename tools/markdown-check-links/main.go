package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var linkRe = regexp.MustCompile(`\[([^\]]*)\]\(([^)\s]+)(?:\s+"[^"]*")?\)`)
var autoLinkRe = regexp.MustCompile(`<([a-zA-Z][a-zA-Z0-9+.\-]*://[^>]+)>`)

type LinkCheck struct {
	Text  string
	URL   string
	OK    bool
	Code  int
	Error string
}

func extractURLs(file string) []string {
	f, err := os.Open(file)
	if err != nil {
		return nil
	}
	defer f.Close()

	var urls []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		for _, m := range linkRe.FindAllStringSubmatch(line, -1) {
			urls = append(urls, m[2])
		}
		for _, m := range autoLinkRe.FindAllStringSubmatch(line, -1) {
			urls = append(urls, m[1])
		}
	}
	return urls
}

func checkURL(url string) LinkCheck {
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return LinkCheck{URL: url, OK: false, Error: err.Error()}
	}
	req.Header.Set("User-Agent", "markdown-check-links/1.0")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return LinkCheck{URL: url, OK: false, Error: err.Error()}
	}
	defer resp.Body.Close()

	ok := resp.StatusCode >= 200 && resp.StatusCode < 400
	return LinkCheck{URL: url, OK: ok, Code: resp.StatusCode}
}

func printUsage() {
	fmt.Fprintf(os.Stderr, `markdown-check-links - Check links in Markdown files for broken URLs

Usage:
  markdown-check-links [options] [file]

Options:
  -q, --quiet    Only show broken links
  -h, --help     Show this help

Examples:
  markdown-check-links README.md
  markdown-check-links --quiet README.md
`)
}

func main() {
	quiet := false
	file := ""

	args := os.Args[1:]
	i := 0
	for i < len(args) {
		switch args[i] {
		case "-q", "--quiet":
			quiet = true
		case "-h", "--help":
			printUsage()
			return
		default:
			if strings.HasPrefix(args[i], "-") {
				fmt.Fprintf(os.Stderr, "error: unknown flag %s\n", args[i])
				os.Exit(1)
			}
			file = args[i]
		}
		i++
	}

	if file == "" {
		printUsage()
		os.Exit(1)
	}

	urls := extractURLs(file)
	if len(urls) == 0 {
		fmt.Println("no links found")
		return
	}

	broken := 0
	for _, url := range urls {
		result := checkURL(url)
		if !result.OK {
			broken++
			if quiet {
				fmt.Printf("BROKEN: %s (%s)\n", result.URL, result.Error)
			} else {
				status := fmt.Sprintf("HTTP %d", result.Code)
				if result.Error != "" {
					status = result.Error
				}
				fmt.Printf("BROKEN: %s (%s)\n", result.URL, status)
			}
		} else if !quiet {
			fmt.Printf("OK:     %s (HTTP %d)\n", result.URL, result.Code)
		}
	}

	if quiet && broken == 0 {
		fmt.Println("all links OK")
	} else if !quiet {
		fmt.Printf("\n%d/%d links broken\n", broken, len(urls))
	}

	if broken > 0 {
		os.Exit(1)
	}
}
