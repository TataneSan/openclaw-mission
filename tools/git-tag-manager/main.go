package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"
)

// TagInfo holds metadata about a git tag.
type TagInfo struct {
	Name       string `json:"name"`
	Sha        string `json:"sha"`
	Subject    string `json:"subject,omitempty"`
	Message    string `json:"message,omitempty"`
	Tagger     string `json:"tagger,omitempty"`
	Date       string `json:"date,omitempty"`
	IsAnnotated bool  `json:"annotated"`
}

func runGit(args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	out, err := cmd.CombinedOutput()
	return strings.TrimSpace(string(out)), err
}

func listTags() ([]TagInfo, error) {
	raw, err := runGit("tag", "-l")
	if err != nil {
		return nil, fmt.Errorf("failed to list tags: %w", err)
	}

	lines := strings.Split(raw, "\n")
	var tags []TagInfo
	for _, line := range lines {
		name := strings.TrimSpace(line)
		if name == "" {
			continue
		}
		info, err := getTagInfo(name)
		if err != nil {
			return nil, err
		}
		tags = append(tags, *info)
	}

	sort.Slice(tags, func(i, j int) bool {
		return tags[i].Date > tags[j].Date
	})
	return tags, nil
}

func getTagInfo(name string) (*TagInfo, error) {
	info := &TagInfo{Name: name}

	// Get the commit SHA
	sha, err := runGit("rev-list", "-n", "1", name)
	if err != nil {
		return nil, fmt.Errorf("tag %s not found", name)
	}
	info.Sha = sha

	// Check if annotated
	fmtOut, _ := runGit("cat-file", "-t", name)
	isAnnotated := strings.TrimSpace(fmtOut) == "tag"
	info.IsAnnotated = isAnnotated

	if isAnnotated {
		// Parse annotated tag object
		tagObj, _ := runGit("cat-file", "-p", name)
		scanner := bufio.NewScanner(strings.NewReader(tagObj))
		for scanner.Scan() {
			line := scanner.Text()
			if strings.HasPrefix(line, "object ") {
				continue // already have SHA from rev-list
			}
			if strings.HasPrefix(line, "type ") {
				continue
			}
			if strings.HasPrefix(line, "tagger ") {
				info.Tagger = strings.TrimPrefix(line, "tagger ")
			}
			if strings.HasPrefix(line, "date ") {
				info.Date = strings.TrimPrefix(line, "date ")
			}
			// Subject line (first non-empty line after headers)
			if info.Subject == "" && !strings.HasPrefix(line, " ") && line != "" {
				info.Subject = line
			}
			// Message lines (indented after blank line)
			if line == "" || strings.HasPrefix(line, " ") {
				if line != "" {
					info.Message += strings.TrimPrefix(line, " ") + "\n"
				}
			}
		}
		info.Message = strings.TrimSuffix(info.Message, "\n")
	} else {
		// For lightweight tags, get commit date
		date, _ := runGit("log", "-1", "--format=%ai", name)
		info.Date = date
		subj, _ := runGit("log", "-1", "--format=%s", name)
		info.Subject = subj
	}

	return info, nil
}

func createTag(name, message string, annotated bool, sign bool) error {
	// Check if tag already exists
	_, err := runGit("rev-parse", name)
	if err == nil {
		return fmt.Errorf("tag '%s' already exists", name)
	}

	args := []string{"tag"}
	if annotated {
		args = append(args, "-a")
	}
	if sign {
		args = append(args, "-s")
	}
	if message != "" {
		args = append(args, "-m", message)
	}
	args = append(args, name)

	out, err := runGit(args...)
	if err != nil {
		return fmt.Errorf("failed to create tag: %v\n%s", err, out)
	}
	fmt.Printf("Tag '%s' created.\n", name)
	return nil
}

func deleteTag(name string) error {
	// Check if tag exists
	_, err := runGit("rev-parse", name)
	if err != nil {
		return fmt.Errorf("tag '%s' does not exist", name)
	}

	out, err := runGit("tag", "-d", name)
	if err != nil {
		return fmt.Errorf("failed to delete tag: %v\n%s", err, out)
	}
	fmt.Printf("Tag '%s' deleted.\n", strings.TrimSpace(out))
	return nil
}

func moveTag(name, newRef string) error {
	args := []string{"tag", "-f", "-a", name, newRef}
	out, err := runGit(args...)
	if err != nil {
		return fmt.Errorf("failed to move tag: %v\n%s", err, out)
	}
	fmt.Printf("Tag '%s' moved to %s.\n", name, newRef)
	return nil
}

func pushTags() error {
	out, err := runGit("push", "origin", "--tags")
	if err != nil {
		return fmt.Errorf("failed to push tags: %v\n%s", err, out)
	}
	fmt.Println("Tags pushed to origin.")
	return nil
}

func renderTable(tags []TagInfo) {
	if len(tags) == 0 {
		fmt.Println("No tags found.")
		return
	}

	fmt.Printf("%-25s %-8s %-7s %s\n", "TAG", "TYPE", "DATE", "SUBJECT")
	fmt.Println(strings.Repeat("-", 90))
	for _, t := range tags {
		tagType := "light"
		if t.IsAnnotated {
			tagType = "annotated"
		}
		date := t.Date
		if len(date) > 7 {
			date = date[:7]
		}
		subj := t.Subject
		if len(subj) > 40 {
			subj = subj[:38] + ".."
		}
		fmt.Printf("%-25s %-8s %-7s %s\n", t.Name, tagType, date, subj)
	}
	fmt.Printf("\n%d tag(s)\n", len(tags))
}

func renderJSON(tags []TagInfo) {
	data, err := json.MarshalIndent(tags, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(data))
}

func renderVerbose(tags []TagInfo) {
	if len(tags) == 0 {
		fmt.Println("No tags found.")
		return
	}

	for i, t := range tags {
		if i > 0 {
			fmt.Println(strings.Repeat("-", 60))
		}
		fmt.Printf("Tag: %s\n", t.Name)
		fmt.Printf("  SHA:       %s\n", t.Sha)
		if t.IsAnnotated {
			fmt.Printf("  Type:      annotated\n")
		} else {
			fmt.Printf("  Type:      lightweight\n")
		}
		if t.Date != "" {
			fmt.Printf("  Date:      %s\n", t.Date)
		}
		if t.Subject != "" {
			fmt.Printf("  Subject:   %s\n", t.Subject)
		}
		if t.Tagger != "" {
			fmt.Printf("  Tagger:    %s\n", t.Tagger)
		}
		if t.Message != "" {
			fmt.Printf("  Message:   %s\n", t.Message)
		}
	}
	fmt.Printf("\n%d tag(s)\n", len(tags))
}

func printUsage() {
	fmt.Println(`git-tag-manager - Git tag management CLI

Usage:
  git-tag-manager list [format]     List tags (table, json, verbose)
  git-tag-manager create <name>     Create a new tag
  git-tag-manager delete <name>     Delete a tag
  git-tag-manager move <name> <ref> Move tag to a new commit
  git-tag-manager push              Push all tags to origin
  git-tag-manager latest            Show the latest tag

Options:
  -a, --annotated   Create annotated tag (for create)
  -m, --message     Tag message (for create)
  -s, --sign        GPG sign tag (for create)
  -f, --force       Force operation (overwrite existing tag)
  -h, --help        Show this help`)
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	// Parse global flags
	annotated := false
	message := ""
	sign := false
	force := false
	for _, arg := range os.Args[2:] {
		switch arg {
		case "-a", "--annotated":
			annotated = true
		case "-s", "--sign":
			sign = true
		case "-f", "--force":
			force = true
		case "-m", "--message":
			if len(os.Args) > 2 {
				// message is next arg - handled below
			}
		}
	}

	// Check for -m flag and extract message
	for i, arg := range os.Args {
		if (arg == "-m" || arg == "--message") && i+1 < len(os.Args) {
			message = os.Args[i+1]
			break
		}
	}

	switch os.Args[1] {
	case "list", "ls":
		format := "table"
		if len(os.Args) >= 3 && os.Args[2] != "-h" {
			format = os.Args[2]
		}
		tags, err := listTags()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		switch format {
		case "json":
			renderJSON(tags)
		case "verbose", "v":
			renderVerbose(tags)
		default:
			renderTable(tags)
		}

	case "create", "add":
		if len(os.Args) < 3 {
			fmt.Fprintf(os.Stderr, "Error: tag name required\n")
			os.Exit(1)
		}
		name := os.Args[2]
		if err := createTag(name, message, annotated, sign); err != nil {
			if force {
				// Force: delete and recreate
				_ = deleteTag(name)
				if err := createTag(name, message, annotated, sign); err != nil {
					fmt.Fprintf(os.Stderr, "Error: %v\n", err)
					os.Exit(1)
				}
			} else {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
		}

	case "delete", "del", "rm":
		if len(os.Args) < 3 {
			fmt.Fprintf(os.Stderr, "Error: tag name required\n")
			os.Exit(1)
		}
		name := os.Args[2]
		if err := deleteTag(name); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

	case "move":
		if len(os.Args) < 4 {
			fmt.Fprintf(os.Stderr, "Error: git-tag-manager move <name> <ref>\n")
			os.Exit(1)
		}
		name := os.Args[2]
		newRef := os.Args[3]
		if err := moveTag(name, newRef); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

	case "push":
		if err := pushTags(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

	case "latest":
		tags, err := listTags()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		if len(tags) == 0 {
			fmt.Println("No tags found.")
			os.Exit(0)
		}
		t := tags[0] // already sorted by date descending
		fmt.Printf("%s (%s)\n", t.Name, t.Sha[:8])
		if t.Subject != "" {
			fmt.Printf("  %s\n", t.Subject)
		}
		if t.Date != "" {
			fmt.Printf("  %s\n", t.Date)
		}

	case "help", "-h", "--help":
		printUsage()

	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", os.Args[1])
		printUsage()
		os.Exit(1)
	}
}
