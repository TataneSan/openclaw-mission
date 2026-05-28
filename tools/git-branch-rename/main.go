package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	action := os.Args[1]

	switch action {
	case "replace":
		if len(os.Args) < 4 {
			fmt.Println("Usage: git-branch-rename replace <old> <new> [branches...]")
			fmt.Println("  Replaces 'old' with 'new' in branch names")
			fmt.Println("  If no branches listed, applies to all local branches")
			os.Exit(1)
		}
		old := os.Args[2]
		new := os.Args[3]
		branches := os.Args[4:]
		doReplace(old, new, branches)
	case "prefix":
		if len(os.Args) < 4 {
			fmt.Println("Usage: git-branch-rename prefix <prefix> [branches...]")
			fmt.Println("  Adds a prefix to branch names")
			os.Exit(1)
		}
		prefix := os.Args[2]
		branches := os.Args[3:]
		doPrefix(prefix, branches)
	case "suffix":
		if len(os.Args) < 4 {
			fmt.Println("Usage: git-branch-rename suffix <suffix> [branches...]")
			fmt.Println("  Adds a suffix to branch names")
			os.Exit(1)
		}
		suffix := os.Args[2]
		branches := os.Args[3:]
		doSuffix(suffix, branches)
	case "strip-prefix":
		if len(os.Args) < 4 {
			fmt.Println("Usage: git-branch-rename strip-prefix <prefix> [branches...]")
			fmt.Println("  Removes a prefix from branch names")
			os.Exit(1)
		}
		prefix := os.Args[2]
		branches := os.Args[3:]
		doStripPrefix(prefix, branches)
	case "strip-suffix":
		if len(os.Args) < 4 {
			fmt.Println("Usage: git-branch-rename strip-suffix <suffix> [branches...]")
			fmt.Println("  Removes a suffix from branch names")
			os.Exit(1)
		}
		suffix := os.Args[2]
		branches := os.Args[3:]
		doStripSuffix(suffix, branches)
	case "list":
		listBranches()
	case "dry-run":
		if len(os.Args) < 4 {
			fmt.Println("Usage: git-branch-rename dry-run replace <old> <new> [branches...]")
			os.Exit(1)
		}
		subAction := os.Args[2]
		if subAction == "replace" && len(os.Args) >= 5 {
			old := os.Args[3]
			new := os.Args[4]
			branches := os.Args[5:]
			doReplace(old, new, branches, true)
		}
	case "-h", "--help", "help":
		printUsage()
	default:
		fmt.Fprintf(os.Stderr, "Unknown action: %s\n", action)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("git-branch-rename - Batch rename git branches")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  git-branch-rename replace <old> <new> [branches...]")
	fmt.Println("  git-branch-rename prefix <prefix> [branches...]")
	fmt.Println("  git-branch-rename suffix <suffix> [branches...]")
	fmt.Println("  git-branch-rename strip-prefix <prefix> [branches...]")
	fmt.Println("  git-branch-rename strip-suffix <suffix> [branches...]")
	fmt.Println("  git-branch-rename list")
	fmt.Println("  git-branch-rename dry-run replace <old> <new> [branches...]")
	fmt.Println()
	fmt.Println("If no branches are specified, all local branches are considered.")
	fmt.Println("Protected branches (main, master) are never renamed automatically.")
}

func doReplace(old, new string, branches []string, dryRun ...bool) {
	isDryRun := len(dryRun) > 0 && dryRun[0]

	if len(branches) == 0 {
		branches = getLocalBranches()
	}

	renamed := 0
	for _, branch := range branches {
		branch = strings.TrimSpace(branch)
		if isProtected(branch) {
			fmt.Fprintf(os.Stderr, "Skipping protected branch: %s\n", branch)
			continue
		}
		if !strings.Contains(branch, old) {
			continue
		}
		newName := strings.ReplaceAll(branch, old, new)
		if branch == newName {
			continue
		}
		if isDryRun {
			fmt.Printf("[DRY RUN] %s -> %s\n", branch, newName)
		} else {
			if err := renameBranch(branch, newName); err != nil {
				fmt.Fprintf(os.Stderr, "Error renaming %s: %v\n", branch, err)
				continue
			}
			fmt.Printf("Renamed: %s -> %s\n", branch, newName)
		}
		renamed++
	}

	if isDryRun {
		fmt.Printf("\n%d branches would be renamed\n", renamed)
	} else {
		fmt.Printf("\n%d branches renamed\n", renamed)
	}
}

func doPrefix(prefix string, branches []string) {
	if len(branches) == 0 {
		branches = getLocalBranches()
	}

	renamed := 0
	for _, branch := range branches {
		branch = strings.TrimSpace(branch)
		if isProtected(branch) {
			fmt.Fprintf(os.Stderr, "Skipping protected branch: %s\n", branch)
			continue
		}
		if strings.HasPrefix(branch, prefix) {
			continue
		}
		newName := prefix + branch
		if err := renameBranch(branch, newName); err != nil {
			fmt.Fprintf(os.Stderr, "Error renaming %s: %v\n", branch, err)
			continue
		}
		fmt.Printf("Renamed: %s -> %s\n", branch, newName)
		renamed++
	}
	fmt.Printf("\n%d branches renamed\n", renamed)
}

func doSuffix(suffix string, branches []string) {
	if len(branches) == 0 {
		branches = getLocalBranches()
	}

	renamed := 0
	for _, branch := range branches {
		branch = strings.TrimSpace(branch)
		if isProtected(branch) {
			fmt.Fprintf(os.Stderr, "Skipping protected branch: %s\n", branch)
			continue
		}
		if strings.HasSuffix(branch, suffix) {
			continue
		}
		newName := branch + suffix
		if err := renameBranch(branch, newName); err != nil {
			fmt.Fprintf(os.Stderr, "Error renaming %s: %v\n", branch, err)
			continue
		}
		fmt.Printf("Renamed: %s -> %s\n", branch, newName)
		renamed++
	}
	fmt.Printf("\n%d branches renamed\n", renamed)
}

func doStripPrefix(prefix string, branches []string) {
	if len(branches) == 0 {
		branches = getLocalBranches()
	}

	renamed := 0
	for _, branch := range branches {
		branch = strings.TrimSpace(branch)
		if isProtected(branch) {
			fmt.Fprintf(os.Stderr, "Skipping protected branch: %s\n", branch)
			continue
		}
		if !strings.HasPrefix(branch, prefix) {
			continue
		}
		newName := strings.TrimPrefix(branch, prefix)
		if newName == "" {
			fmt.Fprintf(os.Stderr, "Skipping %s: would result in empty name\n", branch)
			continue
		}
		if err := renameBranch(branch, newName); err != nil {
			fmt.Fprintf(os.Stderr, "Error renaming %s: %v\n", branch, err)
			continue
		}
		fmt.Printf("Renamed: %s -> %s\n", branch, newName)
		renamed++
	}
	fmt.Printf("\n%d branches renamed\n", renamed)
}

func doStripSuffix(suffix string, branches []string) {
	if len(branches) == 0 {
		branches = getLocalBranches()
	}

	renamed := 0
	for _, branch := range branches {
		branch = strings.TrimSpace(branch)
		if isProtected(branch) {
			fmt.Fprintf(os.Stderr, "Skipping protected branch: %s\n", branch)
			continue
		}
		if !strings.HasSuffix(branch, suffix) {
			continue
		}
		newName := strings.TrimSuffix(branch, suffix)
		if newName == "" {
			fmt.Fprintf(os.Stderr, "Skipping %s: would result in empty name\n", branch)
			continue
		}
		if err := renameBranch(branch, newName); err != nil {
			fmt.Fprintf(os.Stderr, "Error renaming %s: %v\n", branch, err)
			continue
		}
		fmt.Printf("Renamed: %s -> %s\n", branch, newName)
		renamed++
	}
	fmt.Printf("\n%d branches renamed\n", renamed)
}

func listBranches() {
	branches := getLocalBranches()
	if len(branches) == 0 {
		fmt.Println("No local branches found")
		return
	}
	fmt.Println("Local branches:")
	for _, b := range branches {
		b = strings.TrimSpace(b)
		marker := "  "
		if isProtected(b) {
			marker = "[protected] "
		}
		fmt.Printf("  %s%s\n", marker, b)
	}
}

func getLocalBranches() []string {
	cmd := exec.Command("git", "branch", "--format=%(refname:short)")
	output, err := cmd.Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error listing branches: %v\n", err)
		os.Exit(1)
	}

	var branches []string
	for _, line := range strings.Split(strings.TrimSpace(string(output)), "\n") {
		if line != "" {
			branches = append(branches, line)
		}
	}
	return branches
}

func renameBranch(oldName, newName string) error {
	cmd := exec.Command("git", "branch", "-m", oldName, newName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%v\n%s", err, string(output))
	}
	return nil
}

func isProtected(name string) bool {
	protected := []string{"main", "master", "HEAD"}
	for _, p := range protected {
		if name == p {
			return true
		}
	}
	// Check if currently checked out
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	current, err := cmd.Output()
	if err == nil && strings.TrimSpace(string(current)) == name {
		return true
	}
	return false
}

func init() {
	// Suppress unused import warning
	_ = bufio.ScanLines
}
