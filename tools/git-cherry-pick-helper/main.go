// git-cherry-pick-helper lists commits from a range and lets you cherry-pick them interactively.
//
// Usage:
//
//	git-cherry-pick-helper main..develop
//	git-cherry-pick-helper abc123..def456
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Commit struct {
	Hash    string
	Subject string
}

func runGit(args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func listCommits(rangeStr string) ([]Commit, error) {
	out, err := runGit("log", rangeStr, "--format=%H|%s")
	if err != nil {
		return nil, fmt.Errorf("failed to list commits: %w", err)
	}

	var commits []Commit
	for _, line := range strings.Split(out, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, "|", 2)
		if len(parts) != 2 {
			continue
		}
		commits = append(commits, Commit{
			Hash:    parts[0],
			Subject: parts[1],
		})
	}
	return commits, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <commit-range>\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  %s main..develop\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s abc123..def456\n", os.Args[0])
		os.Exit(1)
	}

	rangeStr := os.Args[1]

	// Check we're in a git repo
	if _, err := runGit("rev-parse", "--git-dir"); err != nil {
		fmt.Fprintf(os.Stderr, "error: not a git repository\n")
		os.Exit(1)
	}

	commits, err := listCommits(rangeStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	if len(commits) == 0 {
		fmt.Println("No commits found in range.")
		os.Exit(0)
	}

	// Display commits
	fmt.Printf("Commits in %s:\n\n", rangeStr)
	for i, c := range commits {
		fmt.Printf("  [%d] %s %s\n", i+1, c.Hash[:7], c.Subject)
	}
	fmt.Println()

	// Ask for selection
	fmt.Print("Enter commit numbers to cherry-pick (e.g. 1,3,5 or 'all' or 'skip'): ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "skip" || input == "" {
		fmt.Println("Skipped.")
		os.Exit(0)
	}

	var selected []Commit
	if input == "all" {
		selected = commits
	} else {
		nums := strings.Split(input, ",")
		for _, n := range nums {
			n = strings.TrimSpace(n)
			idx, err := strconv.Atoi(n)
			if err != nil || idx < 1 || idx > len(commits) {
				fmt.Fprintf(os.Stderr, "Invalid number: %s\n", n)
				os.Exit(1)
			}
			selected = append(selected, commits[idx-1])
		}
	}

	if len(selected) == 0 {
		fmt.Println("No commits selected.")
		os.Exit(0)
	}

	// Cherry-pick in order
	var failed []string
	for _, c := range selected {
		fmt.Printf("Cherry-picking %s %s... ", c.Hash[:7], c.Subject)
		cmd := exec.Command("git", "cherry-pick", c.Hash)
		var stderr bytes.Buffer
		cmd.Stderr = &stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("FAILED")
			fmt.Fprintf(os.Stderr, "  %s\n", stderr.String())
			failed = append(failed, c.Hash[:7])
		} else {
			fmt.Println("OK")
		}
	}

	if len(failed) > 0 {
		fmt.Printf("\n%d commit(s) failed: %s\n", len(failed), strings.Join(failed, ", "))
		os.Exit(1)
	}

	fmt.Printf("\nAll %d commit(s) cherry-picked successfully.\n", len(selected))
}
