// git-branch-list lists git branches with their last commit info.
//
// Usage:
//
//	git-branch-list
//	git-branch-list --all
//	git-branch-list --sort commits
package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"
	"time"
)

type Branch struct {
	Name       string
	Hash       string
	Subject    string
	Author     string
	Date       time.Time
	IsCurrent  bool
	IsRemote   bool
	CommitsAgo int
}

func main() {
	showAll := false
	sortBy := "date"

	for _, arg := range os.Args[1:] {
		switch arg {
		case "-a", "--all":
			showAll = true
		case "--sort":
			if len(os.Args) > 2 {
				sortBy = os.Args[2]
			}
		}
	}

	branches, err := getBranches(showAll)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	if len(branches) == 0 {
		fmt.Println("no branches found")
		return
	}

	sortBranches(branches, sortBy)

	printHeader()
	for _, b := range branches {
		printBranch(b)
	}
	fmt.Printf("\nTotal: %d branch%s\n", len(branches), plural(len(branches)))
}

func getBranches(all bool) ([]Branch, error) {
	var args []string
	if all {
		args = []string{"branch", "-a", "-v", "--no-color"}
	} else {
		args = []string{"branch", "-v", "--no-color"}
	}

	cmd := exec.Command("git", args...)
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("git branch failed: %w", err)
	}

	var branches []Branch
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		if b := parseBranchLine(line); b != nil {
			branches = append(branches, *b)
		}
	}

	return branches, scanner.Err()
}

func parseBranchLine(line string) *Branch {
	isCurrent := false
	isRemote := false

	if strings.HasPrefix(line, "* ") {
		isCurrent = true
		line = line[2:]
	} else if strings.HasPrefix(line, "  ") {
		line = line[2:]
	} else if strings.HasPrefix(line, "*") {
		isCurrent = true
		line = strings.TrimPrefix(line, "*")
		line = strings.TrimSpace(line)
	}

	if strings.HasPrefix(line, "remotes/") {
		isRemote = true
	}

	parts := strings.SplitN(line, "  ", 2)
	if len(parts) < 2 {
		return nil
	}

	name := strings.TrimSpace(parts[0])
	rest := strings.TrimSpace(parts[1])

	// First 7 chars are the short hash
	hash := ""
	if len(rest) >= 7 {
		hash = rest[:7]
		rest = strings.TrimSpace(rest[7:])
	}

	subject := rest
	author := ""
	date := time.Time{}

	// Try to extract more info
	if !isRemote {
		commitInfo, err := getCommitInfo(name)
		if err == nil {
			author = commitInfo.author
			date = commitInfo.date
		}
	}

	return &Branch{
		Name:      name,
		Hash:      hash,
		Subject:   subject,
		Author:    author,
		Date:      date,
		IsCurrent: isCurrent,
		IsRemote:  isRemote,
	}
}

type commitInfo struct {
	author string
	date   time.Time
}

func getCommitInfo(branch string) (*commitInfo, error) {
	cmd := exec.Command("git", "log", "-1", "--format=%an|||%ai", branch)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	parts := strings.SplitN(strings.TrimSpace(string(output)), "|||", 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("failed to parse commit info")
	}

	date, err := time.Parse("2006-01-02 15:04:05 -0700", strings.TrimSpace(parts[1]))
	if err != nil {
		return nil, err
	}

	return &commitInfo{
		author: strings.TrimSpace(parts[0]),
		date:   date,
	}, nil
}

func sortBranches(branches []Branch, by string) {
	switch by {
	case "name":
		sort.Slice(branches, func(i, j int) bool {
			return branches[i].Name < branches[j].Name
		})
	case "author":
		sort.Slice(branches, func(i, j int) bool {
			return branches[i].Author < branches[j].Author
		})
	case "date":
		sort.Slice(branches, func(i, j int) bool {
			return branches[i].Date.After(branches[j].Date)
		})
	}
}

func printHeader() {
	fmt.Println("BRANCHES")
	fmt.Println(strings.Repeat("-", 80))
}

func printBranch(b Branch) {
	prefix := "  "
	if b.IsCurrent {
		prefix = "* "
	}

	dateStr := relativeTime(b.Date)
	authorStr := ""
	if b.Author != "" {
		authorStr = fmt.Sprintf(" (%s)", b.Author)
	}

	fmt.Printf("%s%s %s %s%s %s\n",
		prefix,
		pad(b.Name, 30),
		b.Hash,
		b.Subject,
		authorStr,
		dateStr)
}

func pad(s string, n int) string {
	if len(s) >= n {
		return s[:n]
	}
	return s + strings.Repeat(" ", n-len(s))
}

func relativeTime(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	delta := time.Since(t)
	if delta < time.Minute {
		return "just now"
	}
	if delta < time.Hour {
		return fmt.Sprintf("%dm ago", int(delta.Minutes()))
	}
	if delta < 24*time.Hour {
		return fmt.Sprintf("%dh ago", int(delta.Hours()))
	}
	if delta < 30*24*time.Hour {
		return fmt.Sprintf("%dd ago", int(delta.Hours()/24))
	}
	if delta < 365*24*time.Hour {
		return fmt.Sprintf("%dm ago", int(delta.Hours()/24/30))
	}
	return fmt.Sprintf("%dy ago", int(delta.Hours()/24/365))
}

func plural(n int) string {
	if n == 1 {
		return "e"
	}
	return "es"
}
