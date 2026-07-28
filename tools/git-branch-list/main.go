// git-branch-list lists git branches with last commit info.
package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Branch struct {
	Name      string
	Hash      string
	Subject   string
	Author    string
	Date      string
	IsCurrent bool
}

func runGit(args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func getBranches() ([]Branch, error) {
	out, err := runGit("for-each-ref", "--sort=-committerdate",
		"refs/heads/",
		"--format=%(refname:short)|%(objectname:short)|%(subject)|%(authorname)|%(committerdate:short)")
	if err != nil {
		return nil, fmt.Errorf("list branches: %w", err)
	}

	current, _ := runGit("rev-parse", "--abbrev-ref", "HEAD")

	var branches []Branch
	for _, line := range strings.Split(out, "\n") {
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, "|", 5)
		if len(parts) < 5 {
			continue
		}
		b := Branch{
			Name:      parts[0],
			Hash:      parts[1],
			Subject:   parts[2],
			Author:    parts[3],
			Date:      parts[4],
			IsCurrent: parts[0] == current,
		}
		branches = append(branches, b)
	}
	return branches, nil
}

func printTable(branches []Branch) {
	if len(branches) == 0 {
		fmt.Println("No branches found.")
		return
	}

	for _, b := range branches {
		marker := "  "
		if b.IsCurrent {
			marker = "* "
		}
		fmt.Printf("%s%-30s %-8s %s\n", marker, b.Name, b.Hash, b.Subject)
	}
}

func printSimple(branches []Branch) {
	for _, b := range branches {
		if b.IsCurrent {
			fmt.Println(b.Name)
			return
		}
	}
}

func printUsage() {
	fmt.Fprintf(os.Stderr, "Usage: git-branch-list [OPTIONS]\n\n")
	fmt.Fprintf(os.Stderr, "Lists git branches with last commit information.\n\n")
	fmt.Fprintf(os.Stderr, "Options:\n")
	fmt.Fprintf(os.Stderr, "  -s, --simple   Show only branch names\n")
	fmt.Fprintf(os.Stderr, "  -h, --help     Show this help\n")
}

func main() {
	simple := false
	for _, arg := range os.Args[1:] {
		switch arg {
		case "-s", "--simple":
			simple = true
		case "-h", "--help":
			printUsage()
			os.Exit(0)
		default:
			fmt.Fprintf(os.Stderr, "unknown option: %s\n", arg)
			os.Exit(1)
		}
	}

	branches, err := getBranches()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	if simple {
		printSimple(branches)
	} else {
		printTable(branches)
	}
}
