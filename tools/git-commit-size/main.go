// git-commit-size shows the size of commits (files added, modified, deleted)
// in a git repository. Supports single commit, range, and summary modes.
package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

// FileChange represents a changed file in a commit.
type FileChange struct {
	Status rune // 'A', 'M', 'D'
	Path   string
}

// CommitSize holds aggregated stats for a single commit.
type CommitSize struct {
	Hash   string
	Short  string
	Author string
	Date   string
	Msg    string
	Added  int
	Mod    int
	Deleted int
	Total  int
	Files  []FileChange
}

func runGit(args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("git %s: %w", strings.Join(args, " "), err)
	}
	return strings.TrimSpace(string(out)), nil
}

func parseDiffNameStatus(output string) []FileChange {
	var changes []FileChange
	for _, line := range strings.Split(output, "\n") {
		if line == "" {
			continue
		}
		// Format from diff-tree: ":mode mode hash hash STATUS\tpath"
		// e.g. ":100644 000000 hash 0000000... D\tpath"
		// Find the last tab-separated field and the status before it
		lastTab := strings.LastIndex(line, "\t")
		if lastTab == -1 {
			continue
		}
		beforeTab := line[:lastTab]
		path := line[lastTab+1:]
		// Status is the last field before the tab
		parts := strings.Fields(beforeTab)
		if len(parts) < 5 {
			continue
		}
		statusStr := parts[len(parts)-1]
		var s rune = 'M'
		switch statusStr {
		case "A":
			s = 'A'
		case "M":
			s = 'M'
		case "D":
			s = 'D'
		case "C":
			s = 'M' // copy counts as modify
		}
		changes = append(changes, FileChange{Status: s, Path: path})
	}
	return changes
}

func getCommitSize(commitRef string) (*CommitSize, error) {
	cs := &CommitSize{Hash: commitRef}

	hash, err := runGit("rev-parse", commitRef)
	if err != nil {
		return nil, err
	}
	cs.Hash = hash
	cs.Short = hash[:7]

	short, _ := runGit("log", "-1", "--format=%h", commitRef)
	if short != "" {
		cs.Short = short
	}

	author, _ := runGit("log", "-1", "--format=%an", commitRef)
	cs.Author = author

	date, _ := runGit("log", "-1", "--format=%ad", "--date=short", commitRef)
	cs.Date = date

	msg, _ := runGit("log", "-1", "--format=%s", commitRef)
	cs.Msg = msg

	diffOut, err := runGit("diff-tree", "--root", "--no-commit-id", "-r", "--diff-filter=ACDM", "--no-renames", commitRef)
	if err != nil {
		return nil, err
	}

	changes := parseDiffNameStatus(diffOut)
	cs.Files = changes
	for _, f := range changes {
		switch f.Status {
		case 'A':
			cs.Added++
		case 'M':
			cs.Mod++
		case 'D':
			cs.Deleted++
		}
	}
	cs.Total = cs.Added + cs.Mod + cs.Deleted
	return cs, nil
}

func formatSingle(cs *CommitSize) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Commit: %s\n", cs.Short))
	sb.WriteString(fmt.Sprintf("Author: %s\n", cs.Author))
	sb.WriteString(fmt.Sprintf("Date:   %s\n", cs.Date))
	sb.WriteString(fmt.Sprintf("Message: %s\n", cs.Msg))
	sb.WriteString(fmt.Sprintf("\nFiles: %d added, %d modified, %d deleted (total: %d)\n",
		cs.Added, cs.Mod, cs.Deleted, cs.Total))

	if len(cs.Files) > 0 {
		sb.WriteString("\nChanged files:\n")
		for _, f := range cs.Files {
			var marker string
			switch f.Status {
			case 'A':
				marker = "A"
			case 'M':
				marker = "M"
			case 'D':
				marker = "D"
			}
			sb.WriteString(fmt.Sprintf("  [%s] %s\n", marker, f.Path))
		}
	}
	return sb.String()
}

func formatTable(commits []*CommitSize) string {
	if len(commits) == 0 {
		return "No commits found."
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%-10s %-8s %-8s %-8s %-6s  %s\n",
		"Commit", "Added", "Modified", "Deleted", "Total", "Message"))
	sb.WriteString(strings.Repeat("-", 90) + "\n")

	for _, cs := range commits {
		msg := cs.Msg
		if len(msg) > 40 {
			msg = msg[:37] + "..."
		}
		sb.WriteString(fmt.Sprintf("%-10s %-8d %-8d %-8d %-6d  %s\n",
			cs.Short, cs.Added, cs.Mod, cs.Deleted, cs.Total, msg))
	}

	// Summary row
	var totalAdded, totalMod, totalDel int
	for _, cs := range commits {
		totalAdded += cs.Added
		totalMod += cs.Mod
		totalDel += cs.Deleted
	}
	sb.WriteString(strings.Repeat("-", 90) + "\n")
	sb.WriteString(fmt.Sprintf("TOTAL      %-8d %-8d %-8d %-6d  %d commits\n",
		totalAdded, totalMod, totalDel, totalAdded+totalMod+totalDel, len(commits)))

	return sb.String()
}

func formatSummary(commits []*CommitSize) string {
	if len(commits) == 0 {
		return "No commits found."
	}

	type authorStats struct {
		Author    string
		Commits   int
		Added     int
		Modified  int
		Deleted   int
	}

	authors := make(map[string]*authorStats)
	var totalAdded, totalMod, totalDel int

	for _, cs := range commits {
		totalAdded += cs.Added
		totalMod += cs.Mod
		totalDel += cs.Deleted
		if _, ok := authors[cs.Author]; !ok {
			authors[cs.Author] = &authorStats{Author: cs.Author}
		}
		a := authors[cs.Author]
		a.Commits++
		a.Added += cs.Added
		a.Modified += cs.Mod
		a.Deleted += cs.Deleted
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Summary: %d commits\n\n", len(commits)))
	sb.WriteString(fmt.Sprintf("Total files: %d added, %d modified, %d deleted\n\n",
		totalAdded, totalMod, totalDel))

	// Sort authors by total changes
	var authorList []*authorStats
	for _, a := range authors {
		authorList = append(authorList, a)
	}
	sort.Slice(authorList, func(i, j int) bool {
		ti := authorList[i].Added + authorList[i].Modified + authorList[i].Deleted
		tj := authorList[j].Added + authorList[j].Modified + authorList[j].Deleted
		return ti > tj
	})

	sb.WriteString(fmt.Sprintf("%-25s %-8s %-7s %-8s %-7s %-6s\n",
		"Author", "Commits", "Added", "Modified", "Deleted", "Total"))
	sb.WriteString(strings.Repeat("-", 75) + "\n")
	for _, a := range authorList {
		total := a.Added + a.Modified + a.Deleted
		sb.WriteString(fmt.Sprintf("%-25s %-8d %-7d %-8d %-7d %-6d\n",
			a.Author, a.Commits, a.Added, a.Modified, a.Deleted, total))
	}

	return sb.String()
}

func getCommitsFromRange(ref string) ([]*CommitSize, error) {
	logOut, err := runGit("log", "--format=%H", ref)
	if err != nil {
		return nil, err
	}

	hashes := strings.Fields(logOut)
	if len(hashes) == 0 {
		return nil, fmt.Errorf("no commits found for ref %q", ref)
	}

	var commits []*CommitSize
	for _, h := range hashes {
		cs, err := getCommitSize(h)
		if err != nil {
			return nil, fmt.Errorf("commit %s: %w", h[:7], err)
		}
		commits = append(commits, cs)
	}
	return commits, nil
}

func humanSize(n int64) string {
	if n < 1024 {
		return fmt.Sprintf("%d B", n)
	}
	if n < 1024*1024 {
		return fmt.Sprintf("%.1f KB", float64(n)/1024)
	}
	return fmt.Sprintf("%.1f MB", float64(n)/(1024*1024))
}

func main() {
	var ref, output, repo string
	var summary, allFiles bool

	flag.StringVar(&ref, "ref", "HEAD", "Commit ref or range (e.g. HEAD, main..develop, HEAD~5..HEAD)")
	flag.StringVar(&output, "output", "table", "Output format: table, summary, json")
	flag.StringVar(&repo, "repo", ".", "Path to git repository")
	flag.BoolVar(&summary, "summary", false, "Show summary by author (shorthand for --output=summary)")
	flag.BoolVar(&allFiles, "files", false, "Show all changed files per commit")
	flag.Parse()

	if summary {
		output = "summary"
	}

	// Change to repo directory
	absRepo, err := filepath.Abs(repo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	if err := os.Chdir(absRepo); err != nil {
		fmt.Fprintf(os.Stderr, "Error changing to %s: %v\n", absRepo, err)
		os.Exit(1)
	}

	// Check it's a git repo
	if _, err := runGit("rev-parse", "--git-dir"); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s is not a git repository\n", absRepo)
		os.Exit(1)
	}

	// Single commit mode: if ref doesn't contain '..' or '...' and no range
	isRange := strings.Contains(ref, "..")
	var commits []*CommitSize

	if !isRange {
		// Could be a single commit or a branch name
		// If it looks like a single hash or HEAD~N, treat as single
		// Otherwise get recent commits
		if strings.Contains(ref, "~") || strings.HasPrefix(ref, "HEAD") || len(ref) >= 7 {
			// Check if it's a single commit
			cs, err := getCommitSize(ref)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			commits = []*CommitSize{cs}
		} else {
			// Branch/tag - get last 10 commits
			commits, err = getCommitsFromRange(ref + "~0")
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			if len(commits) > 10 {
				commits = commits[:10]
			}
		}
	} else {
		commits, err = getCommitsFromRange(ref)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	}

	switch output {
	case "json":
		fmt.Println(jsonOutput(commits))
	case "summary":
		fmt.Println(formatSummary(commits))
	case "table":
		if len(commits) == 1 && !isRange {
			fmt.Println(formatSingle(commits[0]))
		} else {
			fmt.Println(formatTable(commits))
			if allFiles {
				for _, cs := range commits {
					if len(cs.Files) > 0 {
						fmt.Printf("\n%s files:\n", cs.Short)
						for _, f := range cs.Files {
							var marker string
							switch f.Status {
							case 'A':
								marker = "A"
							case 'M':
								marker = "M"
							case 'D':
								marker = "D"
							}
							fmt.Printf("  [%s] %s\n", marker, f.Path)
						}
					}
				}
			}
		}
	default:
		fmt.Fprintf(os.Stderr, "Unknown output format: %s\n", output)
		os.Exit(1)
	}
}

func jsonOutput(commits []*CommitSize) string {
	var sb strings.Builder
	sb.WriteString("[\n")
	for i, cs := range commits {
		if i > 0 {
			sb.WriteString(",\n")
		}
		sb.WriteString("  {\n")
		sb.WriteString(fmt.Sprintf("    \"commit\": %q,\n", cs.Short))
		sb.WriteString(fmt.Sprintf("    \"author\": %q,\n", cs.Author))
		sb.WriteString(fmt.Sprintf("    \"date\": %q,\n", cs.Date))
		sb.WriteString(fmt.Sprintf("    \"message\": %q,\n", cs.Msg))
		sb.WriteString(fmt.Sprintf("    \"added\": %d,\n", cs.Added))
		sb.WriteString(fmt.Sprintf("    \"modified\": %d,\n", cs.Mod))
		sb.WriteString(fmt.Sprintf("    \"deleted\": %d,\n", cs.Deleted))
		sb.WriteString(fmt.Sprintf("    \"total\": %d\n", cs.Total))
		sb.WriteString("  }")
	}
	sb.WriteString("\n]\n")
	return sb.String()
}

// unused helper kept for potential future use
var _ = humanSize
var _ = strconv.Atoi
