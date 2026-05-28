package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func runGit(args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func isInGitRepo() bool {
	_, err := runGit("rev-parse", "--git-dir")
	return err == nil
}

func listBranches(all bool) ([]string, error) {
	args := []string{"branch"}
	if all {
		args = append(args, "-a")
	}
	out, err := runGit(args...)
	if err != nil {
		return nil, err
	}
	var branches []string
	for _, line := range strings.Split(out, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		line = strings.TrimPrefix(line, "* ")
		line = strings.TrimPrefix(line, "  ")
		if strings.HasPrefix(line, "remotes/") {
			line = strings.TrimPrefix(line, "remotes/")
		}
		branches = append(branches, line)
	}
	return branches, nil
}

func getMergeBase(branch1, branch2 string) (string, error) {
	return runGit("merge-base", branch1, branch2)
}

func printCommit(sha string) {
	fmt.Printf("  Commit: %s\n", sha[:12])
	oneline, _ := runGit("log", "-1", "--format=%an <%ae> %as %s", sha)
	fmt.Printf("  %s\n", oneline)
}

func printUsage() {
	fmt.Fprintf(os.Stderr, `Usage: git-merge-base <branch1> <branch2> [options]

Find the common ancestor (merge base) between two git branches.

Options:
  -v, --verbose    Show detailed commit info for the merge base
  -l, --list       List all branches (use with --all for remote branches)
  -a, --all        Include remote branches in listing
  -h, --help       Show this help message

Examples:
  git-merge-base main feature-x
  git-merge-base main feature-x -v
  git-merge-base HEAD~5 develop
  git-merge-base --list
`)
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	verbose := false
	listMode := false
	allBranches := false
	var branch1, branch2 string

	i := 1
	for i < len(os.Args) {
		arg := os.Args[i]
		switch arg {
		case "-h", "--help":
			printUsage()
			os.Exit(0)
		case "-v", "--verbose":
			verbose = true
		case "-l", "--list":
			listMode = true
		case "-a", "--all":
			allBranches = true
		default:
			if branch1 == "" {
				branch1 = arg
			} else {
				branch2 = arg
			}
		}
		i++
	}

	if !isInGitRepo() {
		fmt.Fprintf(os.Stderr, "Error: not a git repository\n")
		os.Exit(1)
	}

	if listMode {
		branches, err := listBranches(allBranches)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error listing branches: %v\n", err)
			os.Exit(1)
		}
		for _, b := range branches {
			fmt.Println(b)
		}
		return
	}

	if branch1 == "" || branch2 == "" {
		fmt.Fprintf(os.Stderr, "Error: two branches are required\n\n")
		printUsage()
		os.Exit(1)
	}

	mergeBase, err := getMergeBase(branch1, branch2)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error finding merge base: %v\n", err)
		os.Exit(1)
	}

	if verbose {
		fmt.Printf("Merge base between '%s' and '%s':\n\n", branch1, branch2)
		printCommit(mergeBase)
	} else {
		fmt.Println(mergeBase)
	}
}
