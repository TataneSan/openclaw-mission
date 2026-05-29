// git-merge-base finds the common ancestor between two git branches or commits.
//
// Usage:
//   git-merge-base main feature-login
//   git-merge-base main feature-login --verbose
//   git-merge-base abc123 def456
package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	verbose := false
	args := []string{}
	for _, a := range os.Args[1:] {
		if a == "--verbose" || a == "-v" {
			verbose = true
		} else {
			args = append(args, a)
		}
	}

	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: git-merge-base [--verbose] <branch1> <branch2>\n")
		os.Exit(1)
	}

	branch1 := args[0]
	branch2 := args[1]

	if !isGitRepo() {
		fmt.Fprintf(os.Stderr, "error: not a git repository\n")
		os.Exit(1)
	}

	if !branchExists(branch1) {
		fmt.Fprintf(os.Stderr, "error: branch or commit %q not found\n", branch1)
		os.Exit(1)
	}
	if !branchExists(branch2) {
		fmt.Fprintf(os.Stderr, "error: branch or commit %q not found\n", branch2)
		os.Exit(1)
	}

	mergeBase := runGit("merge-base", branch1, branch2)
	if mergeBase == "" {
		fmt.Println("No common ancestor found.")
		os.Exit(1)
	}

	sha := strings.TrimSpace(mergeBase)
	fmt.Printf("Merge base between %q and %q:\n", branch1, branch2)
	fmt.Printf("  %s\n", sha)

	if verbose {
		msg := runGit("log", "-1", "--format=%s", sha)
		author := runGit("log", "-1", "--format=%an <%ae>", sha)
		date := runGit("log", "-1", "--format=%ai", sha)
		fmt.Printf("  Message: %s\n", strings.TrimSpace(msg))
		fmt.Printf("  Author:  %s\n", strings.TrimSpace(author))
		fmt.Printf("  Date:    %s\n", strings.TrimSpace(date))

		ahead1, behind1 := getAheadBehind(branch1, sha)
		ahead2, behind2 := getAheadBehind(branch2, sha)
		fmt.Printf("  %s is %d commits ahead, %d behind\n", branch1, ahead1, behind1)
		fmt.Printf("  %s is %d commits ahead, %d behind\n", branch2, ahead2, behind2)
	}
}

func isGitRepo() bool {
	cmd := exec.Command("git", "rev-parse", "--git-dir")
	return cmd.Run() == nil
}

func branchExists(ref string) bool {
	cmd := exec.Command("git", "rev-parse", "--verify", ref)
	return cmd.Run() == nil
}

func runGit(args ...string) string {
	cmd := exec.Command("git", args...)
	out, err := cmd.Output()
	if err != nil {
		return ""
	}
	return string(out)
}

func getAheadBehind(branch, base string) (int, int) {
	out := runGit("rev-list", "--count", base+".."+branch)
	ahead := 0
	fmt.Sscanf(strings.TrimSpace(out), "%d", &ahead)

	out = runGit("rev-list", "--count", branch+".."+base)
	behind := 0
	fmt.Sscanf(strings.TrimSpace(out), "%d", &behind)

	return ahead, behind
}
