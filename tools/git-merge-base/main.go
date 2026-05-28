// git-merge-base finds the last common ancestor between two git branches.
//
// It displays the merge base commit with details (hash, author, date, message)
// and optionally shows how many commits each branch is ahead.
package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	verbose := flag.Bool("v", false, "verbose output with commit details")
	count := flag.Bool("c", false, "show commit count between merge-base and each branch")
	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: git-merge-base [-v] [-c] <branch1> <branch2>\n")
		fmt.Fprintf(os.Stderr, "\nFinds the last common ancestor between two git branches.\n")
		fmt.Fprintf(os.Stderr, "\nOptions:\n")
		fmt.Fprintf(os.Stderr, "  -v    Show verbose commit details\n")
		fmt.Fprintf(os.Stderr, "  -c    Show commit count from merge-base to each branch\n")
		os.Exit(1)
	}

	branch1, branch2 := args[0], args[1]

	// Find merge base
	cmd := exec.Command("git", "merge-base", branch1, branch2)
	output, err := cmd.Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: could not find merge base between '%s' and '%s': %v\n", branch1, branch2, err)
		os.Exit(1)
	}

	mergeBase := strings.TrimSpace(string(output))
	fmt.Printf("Merge base: %s\n", mergeBase)

	if *verbose {
		fmt.Println()
		showCommitDetails(mergeBase)
	}

	if *count {
		fmt.Println()
		count1 := countCommits(mergeBase, branch1)
		count2 := countCommits(mergeBase, branch2)
		fmt.Printf("%s: %d commit(s) from merge-base\n", branch1, count1)
		fmt.Printf("%s: %d commit(s) from merge-base\n", branch2, count2)
	}
}

func showCommitDetails(commit string) {
	cmd := exec.Command("git", "log", "-1", "--format=%H%n%an%n%ad%n%s", "--date=short", commit)
	output, err := cmd.Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting commit details: %v\n", err)
		return
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	if len(lines) >= 4 {
		fmt.Printf("Hash:    %s\n", lines[0][:8])
		fmt.Printf("Author:  %s\n", lines[1])
		fmt.Printf("Date:    %s\n", lines[2])
		fmt.Printf("Message: %s\n", lines[3])
	}
}

func countCommits(base, branch string) int {
	cmd := exec.Command("git", "rev-list", "--count", base+".."+branch)
	var out strings.Builder
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return -1
	}

	count := 0
	fmt.Sscanf(out.String(), "%d", &count)
	return count
}
