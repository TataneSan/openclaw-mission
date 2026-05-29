// git-tag-list lists git tags with details (date, type, author, message).
//
// Usage:
//
//	git-tag-list [options] [repo]
//
// Examples:
//
//	git-tag-list
//	git-tag-list --format table
//	git-tag-list --sort date ./my-repo
package main

import (
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"
	"time"
)

type TagInfo struct {
	Name    string
	Target  string // commit SHA
	Type    string // "annotated" or "lightweight"
	Date    time.Time
	Author  string
	Message string
}

func printUsage() {
	fmt.Fprintf(os.Stderr, `git-tag-list - List git tags with details

Usage:
  git-tag-list [options] [repo]

Options:
  -f, --format FORMAT   Output format: table (default), json, csv
  -s, --sort FIELD      Sort field: name (default), date, author
  -r, --reverse         Reverse sort order
  -h, --help            Show this help message

Arguments:
  repo                  Path to git repository (default: current directory)

Examples:
  git-tag-list
  git-tag-list --format json
  git-tag-list --sort date
  git-tag-list --sort date --reverse
  git-tag-list ./path/to/repo

Exit codes:
  0  Success
  1  Error (not a git repo, etc.)
`)
}

func runGit(dir string, args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	cmd.Dir = dir
	out, err := cmd.Output()
	return strings.TrimSpace(string(out)), err
}

func collectTags(dir string) ([]TagInfo, error) {
	// Get all tags
	out, err := runGit(dir, "tag", "-l")
	if err != nil {
		return nil, fmt.Errorf("not a git repository: %w", err)
	}

	tagNames := strings.Split(strings.TrimSpace(out), "\n")
	if len(tagNames) == 1 && tagNames[0] == "" {
		return nil, nil
	}

	var tags []TagInfo
	for _, name := range tagNames {
		name = strings.TrimSpace(name)
		if name == "" {
			continue
		}

		info := TagInfo{Name: name}

		// Get tag details via for-each-ref
		fmtStr := "%(taggerdate:iso8601)|%(taggername)|%(subject)|%(objecttype)"
		detailOut, err := runGit(dir, "for-each-ref", "--format="+fmtStr, "refs/tags/"+name)
		if err != nil {
			continue
		}

		parts := strings.SplitN(detailOut, "|", 4)
		if len(parts) >= 1 && parts[0] != "" {
			if t, err := time.Parse("2006-01-02 15:04:05 -0700", parts[0]); err == nil {
				info.Date = t
			}
		}
		if len(parts) >= 2 && parts[1] != "" {
			info.Author = parts[1]
		}
		if len(parts) >= 3 {
			info.Message = parts[2]
		}
		if len(parts) >= 4 {
			if parts[3] == "tag" {
				info.Type = "annotated"
			} else {
				info.Type = "lightweight"
			}
		}

		// For lightweight tags, get author from the commit
		if info.Type == "lightweight" && info.Author == "" {
			authorOut, err := runGit(dir, "log", "-1", "--format=%an", name)
			if err == nil {
				info.Author = authorOut
			}
		}

		// Get commit SHA
		sha, err := runGit(dir, "rev-list", "-n", "1", name)
		if err == nil {
			info.Target = sha[:10]
		}

		tags = append(tags, info)
	}

	return tags, nil
}

func sortTags(tags []TagInfo, field string, reverse bool) {
	sortFunc := func(i, j int) bool {
		switch field {
		case "date":
			return tags[i].Date.Before(tags[j].Date)
		case "author":
			return tags[i].Author < tags[j].Author
		default: // name
			return tags[i].Name < tags[j].Name
		}
	}
	sort.Slice(tags, func(i, j int) bool {
		result := sortFunc(i, j)
		if reverse {
			return !result
		}
		return result
	})
}

func printTable(tags []TagInfo) {
	if len(tags) == 0 {
		fmt.Println("No tags found.")
		return
	}

	// Headers
	fmt.Printf("%-25s %-12s %-10s %-20s %s\n", "NAME", "TYPE", "DATE", "AUTHOR", "MESSAGE")
	fmt.Println(strings.Repeat("-", 100))

	for _, t := range tags {
		dateStr := ""
		if !t.Date.IsZero() {
			dateStr = t.Date.Format("2006-01-02")
		}
		msg := t.Message
		if len(msg) > 35 {
			msg = msg[:32] + "..."
		}
		fmt.Printf("%-25s %-12s %-10s %-20s %s\n", t.Name, t.Type, dateStr, t.Author, msg)
	}

	fmt.Printf("\n%d tag(s)\n", len(tags))
}

func printJSON(tags []TagInfo) {
	fmt.Println("[")
	for i, t := range tags {
		comma := ","
		if i == len(tags)-1 {
			comma = ""
		}
		dateStr := ""
		if !t.Date.IsZero() {
			dateStr = t.Date.Format(time.RFC3339)
		}
		fmt.Printf("  {\n")
		fmt.Printf("    \"name\": %q,\n", t.Name)
		fmt.Printf("    \"type\": %q,\n", t.Type)
		fmt.Printf("    \"target\": %q,\n", t.Target)
		fmt.Printf("    \"date\": %q,\n", dateStr)
		fmt.Printf("    \"author\": %q,\n", t.Author)
		fmt.Printf("    \"message\": %q\n", t.Message)
		fmt.Printf("  }%s\n", comma)
	}
	fmt.Println("]")
}

func printCSV(tags []TagInfo) {
	if len(tags) == 0 {
		fmt.Println("name,type,target,date,author,message")
		return
	}

	fmt.Println("name,type,target,date,author,message")
	for _, t := range tags {
		dateStr := ""
		if !t.Date.IsZero() {
			dateStr = t.Date.Format("2006-01-02")
		}
		msg := strings.ReplaceAll(t.Message, "\"", "\"\"")
		if strings.Contains(msg, ",") || strings.Contains(msg, "\"") {
			msg = "\"" + msg + "\""
		}
		fmt.Printf("%s,%s,%s,%s,%s,%s\n", t.Name, t.Type, t.Target, dateStr, t.Author, msg)
	}
}

func main() {
	args := os.Args[1:]

	format := "table"
	sortField := "name"
	reverse := false
	repo := ""

	for len(args) > 0 {
		switch args[0] {
		case "-h", "--help":
			printUsage()
			os.Exit(0)
		case "-f", "--format":
			if len(args) < 2 {
				fmt.Fprintf(os.Stderr, "error: --format requires an argument (table, json, csv)\n")
				os.Exit(1)
			}
			format = args[1]
			args = args[2:]
		case "-s", "--sort":
			if len(args) < 2 {
				fmt.Fprintf(os.Stderr, "error: --sort requires an argument (name, date, author)\n")
				os.Exit(1)
			}
			sortField = args[1]
			args = args[2:]
		case "-r", "--reverse":
			reverse = true
			args = args[1:]
		default:
			repo = args[0]
			args = args[1:]
		}
	}

	if repo == "" {
		repo = "."
	}

	tags, err := collectTags(repo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	sortTags(tags, sortField, reverse)

	switch format {
	case "table":
		printTable(tags)
	case "json":
		printJSON(tags)
	case "csv":
		printCSV(tags)
	default:
		fmt.Fprintf(os.Stderr, "error: unknown format %q (valid: table, json, csv)\n", format)
		os.Exit(1)
	}
}
