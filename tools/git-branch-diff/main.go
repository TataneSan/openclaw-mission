// git-branch-diff shows formatted diff between two git branches.
package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "help", "-h", "--help":
		printUsage()
	default:
		exitCmd(run(os.Args[1:]))
	}
}

func exitCmd(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println(`git-branch-diff - Show formatted diff between two git branches

Usage:
  git-branch-diff [options] <source> <target>

Options:
  -s, --short    Short output: only summary stats
  -f, --files    Show affected files list
  -c, --commits  Show commit list
  -v, --verbose  Show full diff stats (added/removed per file)
  -q, --quiet    Quiet: suppress non-error output (for scripts)

Arguments:
  source         Source branch/ref (e.g. feature/login)
  target         Target branch/ref (e.g. main, develop)

Examples:
  git-branch-diff feature/login main
  git-branch-diff -s feature/login main
  git-branch-diff -f feature/login main
  git-branch-diff -c HEAD main
  git-branch-diff -v develop release/v1.0`)
}

type diffResult struct {
	Source      string
	Target      string
	CommitsAhead int
	CommitsBehind int
	Files       []fileStat
	TotalAdded  int
	TotalRemoved int
}

type fileStat struct {
	Path     string
	Added    int
	Removed  int
	Status   string
}

func git(args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	out, err := cmd.Output()
	return strings.TrimSpace(string(out)), err
}

func parseDiff(args []string) (*diffResult, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("both source and target branches are required")
	}

	source := args[0]
	target := args[1]

	result := &diffResult{
		Source: source,
		Target: target,
	}

	// Get ahead/behind counts
	out, err := git("rev-list", "--left-right", "--count", source+"..."+target)
	if err != nil {
		return nil, fmt.Errorf("failed to compare branches: %w", err)
	}
	parts := strings.Fields(out)
	if len(parts) >= 2 {
		result.CommitsAhead, _ = strconv.Atoi(parts[0])
		result.CommitsBehind, _ = strconv.Atoi(parts[1])
	}

	// Get diff stats
	out, err = git("diff", "--numstat", target+"..."+source)
	if err != nil {
		return nil, fmt.Errorf("failed to get diff stats: %w", err)
	}

	if out != "" {
		scanner := bufio.NewScanner(strings.NewReader(out))
		for scanner.Scan() {
			line := scanner.Text()
			fields := strings.Fields(line)
			if len(fields) < 3 {
				continue
			}

			added, _ := strconv.Atoi(fields[0])
			removed, _ := strconv.Atoi(fields[1])
			path := fields[2]

			// Detect rename
			status := "M"
			if strings.HasPrefix(path, "R100\t") || strings.HasPrefix(path, "R") {
				status = "R"
				if len(fields) > 3 {
					path = fields[3]
				}
			} else if strings.HasPrefix(path, "A") {
				status = "A"
			}

			fs := fileStat{
				Path:    path,
				Added:   added,
				Removed: removed,
				Status:  status,
			}
			result.Files = append(result.Files, fs)
			result.TotalAdded += added
			result.TotalRemoved += removed
		}
	}

	return result, nil
}

func run(args []string) error {
	short := false
	showFiles := false
	showCommits := false
	verbose := false
	quiet := false

	i := 0
	branchArgs := []string{}
	for i < len(args) {
		switch args[i] {
		case "-s", "--short":
			short = true
		case "-f", "--files":
			showFiles = true
		case "-c", "--commits":
			showCommits = true
		case "-v", "--verbose":
			verbose = true
		case "-q", "--quiet":
			quiet = true
		default:
			branchArgs = append(branchArgs, args[i])
		}
		i++
	}

	if len(branchArgs) < 2 {
		return fmt.Errorf("both source and target branches are required")
	}

	result, err := parseDiff(branchArgs)
	if err != nil {
		return err
	}

	if quiet {
		if len(result.Files) > 0 {
			fmt.Printf("%d\n", len(result.Files))
		} else {
			fmt.Print("0\n")
		}
		return nil
	}

	// Header
	fmt.Printf("\x1b[1mBranch diff: %s -> %s\x1b[0m\n\n", result.Source, result.Target)

	// Commits ahead/behind
	fmt.Printf("  Commits ahead:  \x1b[32m%d\x1b[0m\n", result.CommitsAhead)
	fmt.Printf("  Commits behind: \x1b[31m%d\x1b[0m\n", result.CommitsBehind)
	fmt.Println()

	if short {
		printSummary(result)
		return nil
	}

	// Files summary
	if len(result.Files) > 0 {
		fmt.Printf("  Files changed: \x1b[1m%d\x1b[0m\n", len(result.Files))
		fmt.Printf("  Insertions:    \x1b[32m+%d\x1b[0m\n", result.TotalAdded)
		fmt.Printf("  Deletions:     \x1b[31m-%d\x1b[0m\n", result.TotalRemoved)
		fmt.Println()
	} else {
		fmt.Println("  No file changes.")
		return nil
	}

	// Show commits list
	if showCommits {
		printCommits(result.Source, result.Target)
		fmt.Println()
	}

	// Show files
	if showFiles || verbose {
		printFiles(result, verbose)
	}

	return nil
}

func printSummary(result *diffResult) {
	fmt.Printf("  Files changed: \x1b[1m%d\x1b[0m\n", len(result.Files))
	fmt.Printf("  Insertions:    \x1b[32m+%d\x1b[0m\n", result.TotalAdded)
	fmt.Printf("  Deletions:     \x1b[31m-%d\x1b[0m\n", result.TotalRemoved)
}

func printCommits(source, target string) {
	fmt.Println("  \x1b[1mCommits:\x1b[0m")

	out, err := git("log", "--oneline", target+".."+source)
	if err != nil {
		fmt.Fprintf(os.Stderr, "  (could not fetch commits)\n")
		return
	}

	if out == "" {
		fmt.Println("    (no commits)")
		return
	}

	scanner := bufio.NewScanner(strings.NewReader(out))
	for scanner.Scan() {
		fmt.Printf("    %s\n", scanner.Text())
	}
}

func printFiles(result *diffResult, verbose bool) {
	fmt.Println("  \x1b[1mFiles:\x1b[0m")

	for _, f := range result.Files {
		statusColor := "\x1b[36m" // default: modified (cyan)
		status := "M"
		switch f.Status {
		case "A":
			statusColor = "\x1b[32m" // added (green)
			status = "A"
		case "R":
			statusColor = "\x1b[35m" // renamed (magenta)
			status = "R"
		}

		if verbose {
			addedStr := ""
			if f.Added > 0 {
				addedStr = fmt.Sprintf(" \x1b[32m+%d\x1b[0m", f.Added)
			}
			removedStr := ""
			if f.Removed > 0 {
				removedStr = fmt.Sprintf(" \x1b[31m-%d\x1b[0m", f.Removed)
			}
			fmt.Printf("    \x1b[%s%s\x1b[0m%s%s %s\n", statusColor, status, addedStr, removedStr, f.Path)
		} else {
			fmt.Printf("    \x1b[%s%s\x1b[0m %s\n", statusColor, status, f.Path)
		}
	}
}
