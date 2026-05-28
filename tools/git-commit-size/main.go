package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"math"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
)

// CommitStat holds statistics for a single commit.
type CommitStat struct {
	Hash       string
	ShortHash  string
	Subject    string
	Files      int
	Insertions int
	Deletions  int
	TotalLOC   int
}

// FileStat holds per-file statistics within a commit.
type FileStat struct {
	Insertions int
	Deletions  int
	FileName   string
}

// parseDiffStat parses the output of `git diff --stat` into FileStat entries.
func parseDiffStat(output string) []FileStat {
	var stats []FileStat
	scanner := bufio.NewScanner(strings.NewReader(output))
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		// Format: " path/to/file.go | 12 +++++++++---"
		// or:     " path/to/file.go | Bin 0 -> 1234 bytes"
		parts := strings.Split(line, "|")
		if len(parts) < 2 {
			continue
		}

		fileName := strings.TrimSpace(parts[0])
		changePart := strings.TrimSpace(parts[1])

		var ins, del int
		// Skip binary files
		if strings.HasPrefix(changePart, "Bin") {
			ins = 0
			del = 0
		} else {
			// Parse "12 +++++++++---" or just "12"
			fields := strings.Fields(changePart)
			if len(fields) >= 1 {
				total, err := strconv.Atoi(fields[0])
				if err == nil {
					// Count + and - characters
					rest := ""
					if len(fields) >= 2 {
						rest = fields[1]
					}
					plusCount := strings.Count(rest, "+")
					minusCount := strings.Count(rest, "-")
					if plusCount+minusCount > 0 {
						ins = plusCount
						del = minusCount
					} else {
						// No visual indicators, use total as insertions
						ins = total
						del = 0
					}
				}
			}
		}

		stats = append(stats, FileStat{
			FileName:   fileName,
			Insertions: ins,
			Deletions:  del,
		})
	}
	return stats
}

// getCommits runs git log and returns commit hashes and subjects.
func getCommits(rev string, limit int) ([]string, []string, error) {
	args := []string{"log", "--format=%H%n%s", "--no-merges"}
	if rev != "" {
		args = append(args, rev)
	}
	if limit > 0 {
		args = append(args, fmt.Sprintf("-n%d", limit))
	}

	cmd := exec.Command("git", args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return nil, nil, err
	}

	lines := strings.Split(strings.TrimSpace(out.String()), "\n")
	var hashes, subjects []string
	for i := 0; i+1 < len(lines); i += 2 {
		hashes = append(hashes, lines[i])
		subjects = append(subjects, lines[i+1])
	}
	return hashes, subjects, nil
}

// getCommitStats runs git diff --stat for a commit and returns file stats.
func getCommitStats(hash string) ([]FileStat, error) {
	var parentArg string
	// Check if commit has a parent
	cmd := exec.Command("git", "rev-parse", hash+"^")
	if err := cmd.Run(); err != nil {
		// No parent (root commit) — diff against empty tree
		parentArg = "4b825dc642cb6eb9a060e54bf8d69288fbee4904"
	} else {
		parentArg = hash + "^"
	}

	cmd = exec.Command("git", "diff", "--stat", parentArg+".."+hash)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return nil, err
	}

	return parseDiffStat(out.String()), nil
}

// formatNumber formats a number with thousand separators.
func formatNumber(n int) string {
	if n < 0 {
		return "-" + formatNumber(-n)
	}
	s := strconv.Itoa(n)
	if len(s) <= 3 {
		return s
	}

	var result strings.Builder
	result.Grow(len(s) + (len(s)-1)/3)
	for i, c := range s {
		if i > 0 && (len(s)-i)%3 == 0 {
			result.WriteByte(',')
		}
		result.WriteRune(c)
	}
	return result.String()
}

// bar creates a simple text bar.
func bar(count, maxCount, width int) string {
	if maxCount == 0 {
		return strings.Repeat(" ", width)
	}
	filled := int(math.Round(float64(count)/float64(maxCount) * float64(width)))
	if filled > width {
		filled = width
	}
	return strings.Repeat("█", filled) + strings.Repeat(" ", width-filled)
}

func printUsage() {
	fmt.Fprintf(os.Stderr, `git-commit-size - Show commit sizes (files touched, lines of code)

Usage:
  git-commit-size [options] [revision]

Options:
  -n int
        Number of commits to show (default: 10)
  -all
        Show all commits in history
  -files
        Show per-file breakdown for each commit
  -sort string
        Sort by: files, insertions, deletions, total (default: "files")
  -top int
        When -files is set, show only top N files per commit (default: 0 = all)
  -bar-width int
        Width of the bar chart (default: 30)
  -no-color
        Disable colored output
  -json
        Output as JSON

Examples:
  git-commit-size
  git-commit-size -n 20
  git-commit-size -all
  git-commit-size -files -top 5
  git-commit-size -sort insertions main..develop
  git-commit-size -json -n 5 > commits.json

Revisions:
  (none)        Last N commits on current branch
  main          Last N commits on main branch
  main..develop Commits in develop not in main
  v1.0..v2.0    Commits between two tags

`)
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--help" {
		printUsage()
		os.Exit(0)
	}

	fs := flag.NewFlagSet("git-commit-size", flag.ExitOnError)
	limit := fs.Int("n", 10, "Number of commits to show")
	showAll := fs.Bool("all", false, "Show all commits")
	showFiles := fs.Bool("files", false, "Show per-file breakdown")
	sortBy := fs.String("sort", "files", "Sort by: files, insertions, deletions, total")
	topFiles := fs.Int("top", 0, "Top N files per commit (0 = all)")
	barWidth := fs.Int("bar-width", 30, "Width of bar chart")
	noColor := fs.Bool("no-color", false, "Disable colors")
	outputJSON := fs.Bool("json", false, "Output as JSON")

	fs.Usage = printUsage
	fs.Parse(os.Args[1:])

	rev := ""
	if fs.NArg() > 0 {
		rev = fs.Arg(0)
	}

	if *showAll {
		*limit = 0
	}

	// Validate sort
	validSorts := map[string]bool{"files": true, "insertions": true, "deletions": true, "total": true}
	if !validSorts[*sortBy] {
		fmt.Fprintf(os.Stderr, "Error: invalid sort field %q, must be one of: files, insertions, deletions, total\n", *sortBy)
		os.Exit(1)
	}

	// Check we're in a git repo
	if err := exec.Command("git", "rev-parse", "--git-dir").Run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error: not a git repository")
		os.Exit(1)
	}

	hashes, subjects, err := getCommits(rev, *limit)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting commits: %v\n", err)
		os.Exit(1)
	}

	if len(hashes) == 0 {
		fmt.Println("No commits found.")
		return
	}

	// Gather stats
	stats := make([]CommitStat, len(hashes))
	for i, hash := range hashes {
		fileStats, err := getCommitStats(hash)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: could not get stats for %s: %v\n", hash[:7], err)
			continue
		}

		var totalIns, totalDel int
		for _, fs := range fileStats {
			totalIns += fs.Insertions
			totalDel += fs.Deletions
		}

		stats[i] = CommitStat{
			Hash:       hash,
			ShortHash:  hash[:7],
			Subject:    subjects[i],
			Files:      len(fileStats),
			Insertions: totalIns,
			Deletions:  totalDel,
			TotalLOC:   totalIns + totalDel,
		}
	}

	// Sort
	switch *sortBy {
	case "files":
		sort.Slice(stats, func(i, j int) bool { return stats[i].Files > stats[j].Files })
	case "insertions":
		sort.Slice(stats, func(i, j int) bool { return stats[i].Insertions > stats[j].Insertions })
	case "deletions":
		sort.Slice(stats, func(i, j int) bool { return stats[i].Deletions > stats[j].Deletions })
	case "total":
		sort.Slice(stats, func(i, j int) bool { return stats[i].TotalLOC > stats[j].TotalLOC })
	}

	// JSON output
	if *outputJSON {
		jsonOutput(stats)
		return
	}

	// Find max values for bar scaling
	maxFiles, maxIns, maxDel := 1, 1, 1
	for _, s := range stats {
		if s.Files > maxFiles {
			maxFiles = s.Files
		}
		if s.Insertions > maxIns {
			maxIns = s.Insertions
		}
		if s.Deletions > maxDel {
			maxDel = s.Deletions
		}
	}

	w := *barWidth

	// Colors
	var green, red, cyan, yellow, reset string
	if !*noColor {
		green = "\033[32m"
		red = "\033[31m"
		cyan = "\033[36m"
		yellow = "\033[33m"
		reset = "\033[0m"
	}

	// Print header
	fmt.Printf("%-8s  %-6s %-7s %-7s  %s\n", "HASH", "FILES", "+INS", "-DEL", "SUBJECT")
	fmt.Println(strings.Repeat("-", 90))

	for _, s := range stats {
		insBar := bar(s.Insertions, maxIns, w/3)
		delBar := bar(s.Deletions, maxDel, w/3)

		fmt.Printf("%-8s  %s%-6d%s %s%-7s%s %s%-7s%s  %s\n",
			cyan+s.ShortHash+reset,
			yellow, s.Files, reset,
			green, formatNumber(s.Insertions), reset,
			red, formatNumber(s.Deletions), reset,
			s.Subject)

		fmt.Printf("          %s+%-7s%s %s-%-7s%s\n",
			green, insBar, reset,
			red, delBar, reset)

		// Per-file breakdown
		if *showFiles {
			fileStats, _ := getCommitStats(s.Hash)
			// Sort files by total changes
			sort.Slice(fileStats, func(i, j int) bool {
				ti := fileStats[i].Insertions + fileStats[i].Deletions
				tj := fileStats[j].Insertions + fileStats[j].Deletions
				return ti > tj
			})

			limitFiles := *topFiles
			if limitFiles > 0 && len(fileStats) > limitFiles {
				fileStats = fileStats[:limitFiles]
			}

			for _, fs := range fileStats {
				insStr := green + "+" + formatNumber(fs.Insertions) + reset
				delStr := red + "-" + formatNumber(fs.Deletions) + reset
				fmt.Printf("      %-20s %s  %s\n", fs.FileName, insStr, delStr)
			}
		}

		fmt.Println()
	}

	// Summary
	var totalFiles, totalIns, totalDel int
	for _, s := range stats {
		totalFiles += s.Files
		totalIns += s.Insertions
		totalDel += s.Deletions
	}
	fmt.Println(strings.Repeat("-", 90))
	fmt.Printf("Total: %d commits, %s%d%s files, %s+%-7s%s ins, %s-%-7s%s del\n",
		len(stats),
		yellow, totalFiles, reset,
		green, formatNumber(totalIns), reset,
		red, formatNumber(totalDel), reset)
}

// jsonOutput outputs stats as JSON.
func jsonOutput(stats []CommitStat) {
	fmt.Println("[")
	for i, s := range stats {
		comma := ","
		if i == len(stats)-1 {
			comma = ""
		}
		fmt.Printf("  {\"hash\": \"%s\", \"subject\": %q, \"files\": %d, \"insertions\": %d, \"deletions\": %d, \"total\": %d}%s\n",
			s.Hash, s.Subject, s.Files, s.Insertions, s.Deletions, s.TotalLOC, comma)
	}
	fmt.Println("]")
}
