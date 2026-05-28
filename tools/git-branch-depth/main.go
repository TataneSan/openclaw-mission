package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
)

type branchDepth struct {
	Name   string
	Depth  int
	IsHead bool
}

func runGit(args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func getHEAD() (string, error) {
	out, err := runGit("rev-parse", "--abbrev-ref", "HEAD")
	return out, err
}

func getBranchDepth(name string) (int, error) {
	// git rev-list --count <branch>
	out, err := runGit("rev-list", "--count", name)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(out)
}

func main() {
	// Check we're in a git repo
	if _, err := runGit("rev-parse", "--git-dir"); err != nil {
		fmt.Fprintf(os.Stderr, "Error: not a git repository\n")
		os.Exit(1)
	}

	head, err := getHEAD()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting HEAD: %v\n", err)
		os.Exit(1)
	}

	// Get all local branches
	branchesOut, err := runGit("for-each-ref", "--format=%(refname:short)", "refs/heads/")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error listing branches: %v\n", err)
		os.Exit(1)
	}

	var branches []string
	scanner := bufio.NewScanner(strings.NewReader(branchesOut))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			branches = append(branches, line)
		}
	}

	var results []branchDepth
	for _, name := range branches {
		depth, err := getBranchDepth(name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: could not get depth for '%s': %v\n", name, err)
			continue
		}
		results = append(results, branchDepth{
			Name:   name,
			Depth:  depth,
			IsHead: name == head,
		})
	}

	// Sort by depth descending
	sort.Slice(results, func(i, j int) bool {
		return results[i].Depth > results[j].Depth
	})

	if len(results) == 0 {
		fmt.Println("No branches found.")
		return
	}

	// Print header
	fmt.Printf("%-40s %8s  %s\n", "BRANCH", "COMMITS", "")
	fmt.Println(strings.Repeat("-", 55))

	for _, b := range results {
		marker := ""
		if b.IsHead {
			marker = "<-"
		}
		fmt.Printf("%-40s %8d  %s\n", b.Name, b.Depth, marker)
	}

	fmt.Printf("\nTotal: %d branch(es)\n", len(results))
}
