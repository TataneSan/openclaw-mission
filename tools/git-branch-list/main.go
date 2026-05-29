package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"
	"time"
)

type Branch struct {
	Name       string
	Hash       string
	Subject    string
	Author     string
	Date       time.Time
	IsCurrent  bool
	IsRemote   bool
	Ahead      int
	Behind     int
}

func runGit(args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	cmd.Dir = repoDir
	out, err := cmd.Output()
	return strings.TrimSpace(string(out)), err
}

func parseGitForEachRef(output string, currentBranch string) []Branch {
	var branches []Branch
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, "\t", 5)
		if len(parts) < 5 {
			continue
		}

		fields := strings.Fields(parts[0])
		if len(fields) < 2 {
			continue
		}

		hash := fields[1]
		ref := fields[2]
		isRemote := strings.HasPrefix(ref, "refs/remotes/")
		name := ref
		if strings.HasPrefix(name, "refs/heads/") {
			name = strings.TrimPrefix(name, "refs/heads/")
		} else if strings.HasPrefix(name, "refs/remotes/") {
			name = strings.TrimPrefix(name, "refs/remotes/")
		}

		subject := parts[1]
		author := parts[2]
		dateStr := parts[3]
		aheadBehind := parts[4]

		date, _ := time.Parse("2006-01-02 15:04 -0700", dateStr)

		var ahead, behind int
		abParts := strings.Fields(aheadBehind)
		for len(abParts) > 0 {
			switch {
			case abParts[0] == "ahead" && len(abParts) > 1:
				fmt.Sscanf(abParts[1], "%d", &ahead)
				abParts = abParts[2:]
			case abParts[0] == "behind" && len(abParts) > 1:
				fmt.Sscanf(abParts[1], "%d", &behind)
				abParts = abParts[2:]
			default:
				break
			}
		}

		branch := Branch{
			Name:      name,
			Hash:      hash[:8],
			Subject:   subject,
			Author:    author,
			Date:      date,
			IsCurrent: name == currentBranch,
			IsRemote:  isRemote,
			Ahead:     ahead,
			Behind:    behind,
		}
		branches = append(branches, branch)
	}
	return branches
}

func formatRelativeTime(t time.Time) string {
	diff := time.Since(t)
	if diff < minute {
		return "il y a quelques secondes"
	} else if diff < hour {
		m := int(diff.Minutes())
		if m == 1 {
			return "il y a 1 minute"
		}
		return fmt.Sprintf("il y a %d minutes", m)
	} else if diff < 24*hour {
		h := int(diff.Hours())
		if h == 1 {
			return "il y a 1 heure"
		}
		return fmt.Sprintf("il y a %d heures", h)
	} else if diff < 30*24*hour {
		d := int(diff.Hours() / 24)
		if d == 1 {
			return "il y a 1 jour"
		}
		return fmt.Sprintf("il y a %d jours", d)
	} else if diff < 365*24*hour {
		m := int(diff.Hours() / 24 / 30)
		if m == 1 {
			return "il y a 1 mois"
		}
		return fmt.Sprintf("il y a %d mois", m)
	}
	y := int(diff.Hours() / 24 / 365)
	if y == 1 {
		return "il y a 1 an"
	}
	return fmt.Sprintf("il y a %d ans", y)
}

func printTable(branches []Branch) {
	if len(branches) == 0 {
		fmt.Println("Aucune branche trouvee.")
		return
	}

	localBranches := []Branch{}
	remoteBranches := []Branch{}
	for _, b := range branches {
		if b.IsRemote {
			remoteBranches = append(remoteBranches, b)
		} else {
			localBranches = append(localBranches, b)
		}
	}

	sort.Slice(localBranches, func(i, j int) bool {
		return localBranches[i].Date.After(localBranches[j].Date)
	})
	sort.Slice(remoteBranches, func(i, j int) bool {
		return remoteBranches[i].Date.After(remoteBranches[j].Date)
	})

	if len(localBranches) > 0 {
		fmt.Println("\033[1mBranches locales\033[0m")
		fmt.Println(strings.Repeat("-", 80))
		for _, b := range localBranches {
			marker := "  "
			if b.IsCurrent {
				marker = "* "
			}
			dateStr := formatRelativeTime(b.Date)
			ab := ""
			if b.Ahead > 0 || b.Behind > 0 {
				ab = fmt.Sprintf("  ahead %d, behind %d", b.Ahead, b.Behind)
			}
			if b.IsCurrent {
				fmt.Printf("\033[32m%s\033[0m %-25s %s  %s  %s%s\n",
					marker, b.Name, b.Hash, b.Author, dateStr, ab)
			} else {
				fmt.Printf("%s %-25s %s  %s  %s%s\n",
					marker, b.Name, b.Hash, b.Author, dateStr, ab)
			}
		}
	}

	if len(remoteBranches) > 0 {
		fmt.Println("\n\033[1mBranches distantes\033[0m")
		fmt.Println(strings.Repeat("-", 80))
		for _, b := range remoteBranches {
			dateStr := formatRelativeTime(b.Date)
			fmt.Printf("  %-30s %s  %s  %s\n",
				b.Name, b.Hash, b.Author, dateStr)
		}
	}

	fmt.Printf("\nTotal: %d locale(s), %d distante(s)\n", len(localBranches), len(remoteBranches))
}

func printJSON(branches []Branch) {
	for i, b := range branches {
		if i > 0 {
			fmt.Print(",\n")
		}
		ab := ""
		if b.Ahead > 0 || b.Behind > 0 {
			ab = fmt.Sprintf(", \"ahead\": %d, \"behind\": %d", b.Ahead, b.Behind)
		}
		fmt.Printf(`{"name": "%s", "hash": "%s", "subject": "%s", "author": "%s", "date": "%s", "is_current": %t, "is_remote": %t%s}`,
			b.Name, b.Hash, b.Subject, b.Author, b.Date.Format(time.RFC3339), b.IsCurrent, b.IsRemote, ab)
	}
	fmt.Println()
}

var repoDir string
var minute = time.Minute
var hour = time.Hour

func main() {
	format := flag.String("f", "table", "Format de sortie: table, json")
	repoFlag := flag.String("repo", "", "Chemin du repertoire git (defaut: repertoire courant)")
	flag.Parse()

	if *repoFlag != "" {
		repoDir = *repoFlag
	} else {
		repoDir = "."
	}

	_, err := exec.Command("git", "-C", repoDir, "rev-parse").Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erreur: pas de repo git dans %s\n", repoDir)
		os.Exit(1)
	}

	currentBranch, err := runGit("rev-parse", "--abbrev-ref", "HEAD")
	if err != nil {
		currentBranch = "detached"
	}

	output, err := runGit("for-each-ref",
		"--sort=-committerdate",
		"--format=%(objectname)%09%(subject)%09%(authorname)%09%(committerdate:iso8601)%09%(upstream:track)",
		"refs/heads/", "refs/remotes/")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erreur git: %v\n", err)
		os.Exit(1)
	}

	branches := parseGitForEachRef(output, currentBranch)

	switch *format {
	case "json":
		fmt.Print("[\n")
		printJSON(branches)
		fmt.Print("]\n")
	default:
		printTable(branches)
	}
}
