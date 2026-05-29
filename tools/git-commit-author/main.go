package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

type CommitAuthor struct {
	Name        string
	Email       string
	Date        time.Time
	AuthorDate  time.Time
	Hash        string
	ShortHash   string
	Subject     string
	Body        string
	Parents     []string
	Tree        string
}

func runGit(args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("git %s failed: %w", strings.Join(args, " "), err)
	}
	return strings.TrimSpace(string(output)), nil
}

func parseCommit(ref string) (*CommitAuthor, error) {
	author := &CommitAuthor{}

	// Get commit hash
	hash, err := runGit("rev-parse", ref)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve ref %q: %w", ref, err)
	}
	author.Hash = hash
	author.ShortHash = hash[:7]

	// Get tree
	tree, err := runGit("rev-parse", ref+"^{tree}")
	if err != nil {
		return nil, fmt.Errorf("cannot get tree: %w", err)
	}
	author.Tree = tree

	// Get author name
	name, err := runGit("log", "-1", "--format=%an", ref)
	if err != nil {
		return nil, err
	}
	author.Name = name

	// Get author email
	email, err := runGit("log", "-1", "--format=%ae", ref)
	if err != nil {
		return nil, err
	}
	author.Email = email

	// Get author date
	authorDateStr, err := runGit("log", "-1", "--format=%ai", ref)
	if err != nil {
		return nil, err
	}
	authorDate, err := time.Parse("2006-01-02 15:04:05 -0700", authorDateStr)
	if err != nil {
		return nil, fmt.Errorf("cannot parse author date: %w", err)
	}
	author.AuthorDate = authorDate

	// Get committer date
	committerDateStr, err := runGit("log", "-1", "--format=%ci", ref)
	if err != nil {
		return nil, err
	}
	committerDate, err := time.Parse("2006-01-02 15:04:05 -0700", committerDateStr)
	if err != nil {
		return nil, fmt.Errorf("cannot parse committer date: %w", err)
	}
	author.Date = committerDate

	// Get subject
	subject, err := runGit("log", "-1", "--format=%s", ref)
	if err != nil {
		return nil, err
	}
	author.Subject = subject

	// Get body (everything after blank line)
	body, err := runGit("log", "-1", "--format=%b", ref)
	if err != nil {
		return nil, err
	}
	author.Body = strings.TrimSpace(body)

	// Get parents
	parentsStr, err := runGit("log", "-1", "--format=%P", ref)
	if err == nil && parentsStr != "" {
		author.Parents = strings.Fields(parentsStr)
	}

	return author, nil
}

func formatOutput(author *CommitAuthor, format string) {
	switch format {
	case "json":
		fmt.Printf("{\n")
		fmt.Printf("  \"hash\": \"%s\",\n", author.Hash)
		fmt.Printf("  \"short_hash\": \"%s\",\n", author.ShortHash)
		fmt.Printf("  \"author\": {\n")
		fmt.Printf("    \"name\": \"%s\",\n", escapeJSON(author.Name))
		fmt.Printf("    \"email\": \"%s\"\n", escapeJSON(author.Email))
		fmt.Printf("  },\n")
		fmt.Printf("  \"date\": \"%s\",\n", author.Date.Format(time.RFC3339))
		fmt.Printf("  \"author_date\": \"%s\",\n", author.AuthorDate.Format(time.RFC3339))
		fmt.Printf("  \"subject\": \"%s\",\n", escapeJSON(author.Subject))
		if author.Body != "" {
			fmt.Printf("  \"body\": \"%s\",\n", escapeJSON(author.Body))
		} else {
			fmt.Printf("  \"body\": \"\",\n")
		}
		fmt.Printf("  \"tree\": \"%s\",\n", author.Tree)
		if len(author.Parents) > 0 {
			fmt.Printf("  \"parents\": [%s]\n", strings.Join(func() []string {
				var p []string
				for _, parent := range author.Parents {
					p = append(p, fmt.Sprintf("\"%s\"", parent))
				}
				return p
			}(), ", "))
		} else {
			fmt.Printf("  \"parents\": []\n")
		}
		fmt.Printf("}\n")

	case "table":
		fmt.Println("╔══════════════════════════════════════════════════════════════╗")
		fmt.Println("║                    COMMIT AUTHOR INFO                       ║")
		fmt.Println("╠══════════════════════════════════════════════════════════════╣")
		fmt.Printf("║  Hash:       %s                              ║\n", author.ShortHash)
		fmt.Printf("║  Author:     %s                              ║\n", truncateRight(author.Name, 46))
		fmt.Printf("║  Email:      %s                              ║\n", truncateRight(author.Email, 46))
		fmt.Printf("║  Date:       %s                              ║\n", truncateRight(author.AuthorDate.Format("2006-01-02 15:04:05 MST"), 46))
		fmt.Println("╠══════════════════════════════════════════════════════════════╣")
		fmt.Printf("║  Subject:    %s                              ║\n", truncateRight(author.Subject, 46))
		if author.Body != "" {
			lines := strings.Split(author.Body, "\n")
			for _, line := range lines {
				fmt.Printf("║  %s                                        ║\n", truncateRight(" "+line, 46))
			}
		}
		if len(author.Parents) > 0 {
			fmt.Println("╠══════════════════════════════════════════════════════════════╣")
			for i, parent := range author.Parents {
				fmt.Printf("║  Parent %d:   %s                              ║\n", i+1, parent[:7])
			}
		}
		fmt.Println("╚══════════════════════════════════════════════════════════════╝")

	case "raw":
		fmt.Printf("%s %s <%s> %s\n", author.Name, author.Email, author.AuthorDate.Format(time.RFC3339), author.ShortHash)

	default: // human-readable
		fmt.Printf("\n  %s\n", author.ShortHash)
		fmt.Printf("  %s\n", strings.Repeat("─", min(len(author.ShortHash), 40)))
		fmt.Printf("\n")
		fmt.Printf("  Author:  %s\n", author.Name)
		fmt.Printf("  Email:   %s\n", author.Email)
		fmt.Printf("  Date:    %s\n", author.AuthorDate.Format("2006-01-02 15:04:05 MST"))
		fmt.Printf("\n")
		fmt.Printf("  Subject: %s\n", author.Subject)
		if author.Body != "" {
			fmt.Printf("\n")
			for _, line := range strings.Split(author.Body, "\n") {
				fmt.Printf("  %s\n", line)
			}
		}
		if len(author.Parents) > 0 {
			fmt.Printf("\n")
			fmt.Printf("  Parents:\n")
			for i, parent := range author.Parents {
				fmt.Printf("    %d. %s\n", i+1, parent[:7])
			}
		}
		fmt.Printf("\n")
	}
}

func escapeJSON(s string) string {
	s = strings.ReplaceAll(s, "\\", "\\\\")
	s = strings.ReplaceAll(s, "\"", "\\\"")
	s = strings.ReplaceAll(s, "\n", "\\n")
	s = strings.ReplaceAll(s, "\r", "\\r")
	s = strings.ReplaceAll(s, "\t", "\\t")
	return s
}

func truncateRight(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-2] + ".."
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	ref := flag.String("ref", "HEAD", "Git ref to inspect (commit hash, branch, tag, etc.)")
	format := flag.String("format", "human", "Output format: human, table, json, raw")
	nameOnly := flag.Bool("name", false, "Show only author name")
	emailOnly := flag.Bool("email", false, "Show only author email")
	dateOnly := flag.Bool("date", false, "Show only author date")
	hashOnly := flag.Bool("hash", false, "Show only commit hash")
	shortHashOnly := flag.Bool("short", false, "Show only short commit hash")
	verbose := flag.Bool("verbose", false, "Show full hash and tree")
	flag.Parse()

	// Check if we're in a git repo
	if _, err := runGit("rev-parse", "--git-dir"); err != nil {
		fmt.Fprintln(os.Stderr, "Error: not a git repository")
		os.Exit(1)
	}

	author, err := parseCommit(*ref)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Handle single-field flags
	switch {
	case *nameOnly:
		fmt.Println(author.Name)
	case *emailOnly:
		fmt.Println(author.Email)
	case *dateOnly:
		fmt.Println(author.AuthorDate.Format(time.RFC3339))
	case *hashOnly:
		fmt.Println(author.Hash)
	case *shortHashOnly:
		fmt.Println(author.ShortHash)
	default:
		formatOutput(author, *format)
		if *verbose {
			fmt.Printf("\n  Full hash: %s\n", author.Hash)
			fmt.Printf("  Tree:      %s\n", author.Tree)
		}
	}
}
