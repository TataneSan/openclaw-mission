package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type fileStat struct {
	Path     string
	Added    int
	Deleted  int
	Status   string // M, A, D, R
	Renamed  string
}

type diffSummary struct {
	From    string
	To      string
	Files   []fileStat
	TotalAdded  int
	TotalDeleted int
	FilesChanged int
	FilesInserted int
	FilesDeleted int
}

func runGit(args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("git %s: %w\n%s", strings.Join(args, " "), err, out.String())
	}
	return strings.TrimSpace(out.String()), nil
}

func parseDiffStat(from, to string) (*diffSummary, error) {
	output, err := runGit("diff", "--stat", "--numstat", from+".."+to)
	if err != nil {
		return nil, err
	}

	summary := &diffSummary{
		From: from,
		To:   to,
	}

	scanner := bufio.NewScanner(strings.NewReader(output))
	skipBlank := false
	for scanner.Scan() {
		line := scanner.Text()

		if skipBlank {
			if line == "" {
				skipBlank = false
			}
			continue
		}

		// Skip the summary line (e.g., " 3 files changed, 42 insertions(+), 8 deletions(-)")
		if strings.Contains(line, "files changed") {
			skipBlank = true
			continue
		}

		// Parse numstat lines: <added>\t<deleted>\t<file>
		parts := strings.SplitN(line, "\t", 3)
		if len(parts) != 3 {
			continue
		}

		added, _ := strconv.Atoi(parts[0])
		deleted, _ := strconv.Atoi(parts[1])
		path := parts[2]

		if parts[0] == "-" {
			added = 0
		}
		if parts[1] == "-" {
			deleted = 0
		}

		fs := fileStat{
			Path:    path,
			Added:   added,
			Deleted: deleted,
		}

		// Detect rename
		if strings.HasPrefix(path, "rename from") || strings.HasPrefix(path, "copy from") {
			continue
		}
		if strings.Contains(path, " => ") {
			parts := strings.Split(path, " => ")
			if len(parts) == 2 {
				pctStr := strings.TrimSuffix(strings.TrimPrefix(parts[1], "("), ")")
				pctStr = strings.TrimSpace(strings.TrimSuffix(pctStr, "%"))
				fs.Renamed = parts[0]
				fs.Path = strings.TrimSpace(strings.Split(parts[1], "(")[0])
				fs.Status = "R"
			}
		}

		if added > 0 && deleted == 0 && fs.Status == "" {
			fs.Status = "A"
			summary.FilesInserted++
		} else if added == 0 && deleted > 0 && fs.Status == "" {
			fs.Status = "D"
			summary.FilesDeleted++
		} else if fs.Status == "" {
			fs.Status = "M"
		}

		summary.Files = append(summary.Files, fs)
		summary.TotalAdded += added
		summary.TotalDeleted += deleted
		summary.FilesChanged++
	}

	return summary, nil
}

func formatSummary(s *diffSummary) {
	fmt.Printf("Diff: %s..%s\n", s.From, s.To)
	fmt.Println(strings.Repeat("=", 60))
	fmt.Printf("Files changed: %d\n", s.FilesChanged)
	if s.FilesInserted > 0 {
		fmt.Printf("  Inserted:    %d\n", s.FilesInserted)
	}
	if s.FilesDeleted > 0 {
		fmt.Printf("  Deleted:     %d\n", s.FilesDeleted)
	}
	fmt.Printf("Lines added:   %d\n", s.TotalAdded)
	fmt.Printf("Lines removed: %d\n", s.TotalDeleted)
	net := s.TotalAdded - s.TotalDeleted
	sign := "+"
	if net < 0 {
		sign = ""
		net = -net
	}
	fmt.Printf("Net change:    %s%d lines\n", sign, net)
	fmt.Println()

	if len(s.Files) == 0 {
		fmt.Println("(no changes)")
		return
	}

	// Find max path length for alignment
	maxLen := 0
	for _, f := range s.Files {
		if len(f.Path) > maxLen {
			maxLen = len(f.Path)
		}
	}

	for _, f := range s.Files {
		status := f.Status
		pad := strings.Repeat(" ", maxLen-len(f.Path))
		fmt.Printf("  [%s] %s%s  +%d -%d\n", status, f.Path, pad, f.Added, f.Deleted)
	}
}

func run() error {
	from := flag.String("from", "HEAD~1", "Source commit (ref, SHA, tag)")
	to := flag.String("to", "HEAD", "Target commit (ref, SHA, tag)")
	flag.Parse()

	// Check we're in a git repo
	if _, err := runGit("rev-parse", "--git-dir"); err != nil {
		return fmt.Errorf("not a git repository")
	}

	summary, err := parseDiffStat(*from, *to)
	if err != nil {
		return err
	}

	formatSummary(summary)
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "git-diff-summary: %v\n", err)
		os.Exit(1)
	}
}
