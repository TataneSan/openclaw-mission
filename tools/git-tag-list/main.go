package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"
)

type TagInfo struct {
	Name        string
	Target      string
	Type        string // "commit", "tree", "blob"
	Date        string
	Author      string
	Message     string
	CommitMsg   string
}

func run() int {
	dir := flag.String("d", ".", "repository directory")
	sortFlag := flag.String("sort", "date", "sort order: date, name, version")
	reverse := flag.Bool("r", false, "reverse sort order")
	verbose := flag.Bool("v", false, "show verbose output with commit message")
	format := flag.String("f", "table", "output format: table, list, json")
	flag.Parse()

	tags, err := listTags(*dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return 1
	}

	sortTags(tags, *sortFlag, *reverse)

	switch *format {
	case "table":
		printTable(tags, *verbose)
	case "list":
		printList(tags, *verbose)
	case "json":
		printJSON(tags)
	default:
		fmt.Fprintf(os.Stderr, "error: unknown format %q\n", *format)
		return 1
	}

	return 0
}

func listTags(dir string) ([]TagInfo, error) {
	out, err := exec.Command("git", "-C", dir, "tag", "-l").Output()
	if err != nil {
		return nil, fmt.Errorf("not a git repository: %w", err)
	}

	var tags []TagInfo
	scanner := bufio.NewScanner(strings.NewReader(string(out)))
	for scanner.Scan() {
		name := strings.TrimSpace(scanner.Text())
		if name == "" {
			continue
		}
		info, err := tagInfo(dir, name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "warning: could not read tag %q: %v\n", name, err)
			continue
		}
		tags = append(tags, info)
	}

	return tags, scanner.Err()
}

func tagInfo(dir, name string) (TagInfo, error) {
	info := TagInfo{Name: name}

	// Get tag type and target via for-each-ref
	out, err := exec.Command("git", "-C", dir, "for-each-ref",
		fmt.Sprintf("refs/tags/%s", name),
		"--format=%(objecttype)|%(objectname)|%(creatordate:iso8601)|%(authorname)|%(contents)",
	).Output()
	if err != nil {
		return info, err
	}

	parts := strings.SplitN(string(out), "|", 5)
	if len(parts) >= 1 {
		info.Type = parts[0]
	}
	if len(parts) >= 2 {
		info.Target = parts[1][:8]
	}
	if len(parts) >= 3 {
		info.Date = parts[2]
	}
	if len(parts) >= 4 {
		info.Author = parts[3]
	}
	if len(parts) >= 5 {
		info.Message = strings.TrimSpace(parts[4])
	}

	// Get commit message for the target
	commitOut, _ := exec.Command("git", "-C", dir, "log", "-1", "--format=%s", name).Output()
	info.CommitMsg = strings.TrimSpace(string(commitOut))

	return info, nil
}

func sortTags(tags []TagInfo, by string, reverse bool) {
	sort.Slice(tags, func(i, j int) bool {
		var a, b string
		switch by {
		case "name":
			a, b = tags[i].Name, tags[j].Name
		case "version":
			a, b = tags[i].Name, tags[j].Name
		default: // date
			a, b = tags[i].Date, tags[j].Date
		}
		if reverse {
			return a > b
		}
		return a < b
	})
}

func printTable(tags []TagInfo, verbose bool) {
	if len(tags) == 0 {
		fmt.Println("no tags found")
		return
	}

	// Calculate column widths
	nameW, dateW, typeW := 5, 4, 4
	for _, t := range tags {
		if l := len(t.Name); l > nameW {
			nameW = l
		}
		if l := len(t.Date); l > dateW {
			dateW = l
		}
		if l := len(t.Type); l > typeW {
			typeW = l
		}
	}

	// Header
	fmt.Printf("%-*s  %-*s  %-*s  %s\n", nameW, "NAME", dateW, "DATE", typeW, "TYPE", "COMMIT")
	fmt.Printf("%-*s  %-*s  %-*s  %s\n", nameW, strings.Repeat("-", nameW), dateW, strings.Repeat("-", dateW), typeW, strings.Repeat("-", typeW), "------")

	for _, t := range tags {
		msg := ""
		if verbose && t.CommitMsg != "" {
			msg = "  " + t.CommitMsg
		}
		fmt.Printf("%-*s  %-*s  %-*s  %s%s\n", nameW, t.Name, dateW, t.Date, typeW, t.Type, t.Target, msg)
	}

	fmt.Printf("\n%d tag(s)\n", len(tags))
}

func printList(tags []TagInfo, verbose bool) {
	for _, t := range tags {
		fmt.Printf("%s (%s, %s)\n", t.Name, t.Date, t.Type)
		if verbose && t.CommitMsg != "" {
			fmt.Printf("  %s\n", t.CommitMsg)
		}
	}
}

func printJSON(tags []TagInfo) {
	fmt.Println("[")
	for i, t := range tags {
		comma := ","
		if i == len(tags)-1 {
			comma = ""
		}
		fmt.Printf("  {\"name\": %q, \"date\": %q, \"type\": %q, \"target\": %q, \"message\": %q, \"commit_message\": %q}%s\n",
			t.Name, t.Date, t.Type, t.Target, t.Message, t.CommitMsg, comma)
	}
	fmt.Println("]")
}

func main() {
	os.Exit(run())
}
