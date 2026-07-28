package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func runGit(args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	out, err := cmd.CombinedOutput()
	return strings.TrimSpace(string(out)), err
}

func printUsage() {
	fmt.Fprintf(os.Stderr, `git-commit-diff - Affiche le diff d'un commit spécifique

USAGE:
    git-commit-diff [OPTIONS] <commit>

OPTIONS:
    -s, --stat         Affiche uniquement les statistiques de diff
    -n, --name-only    Affiche uniquement les noms de fichiers modifiés
    -N, --name-status  Affiche les noms et statut (A/M/D) des fichiers
    -p, --patch        Affiche le diff complet (défaut)
    -U, --context N    Nombre de lignes de contexte (défaut: 3)
    -c, --color        Force la coloration du diff
    -f, --files        Affiche la liste des fichiers modifiés
    -h, --help         Affiche l'aide

EXEMPLES:
    git-commit-diff abc1234                     # Diff complet du commit
    git-commit-diff HEAD                        # Diff du commit actuel
    git-commit-diff HEAD~1                      # Diff du parent de HEAD
    git-commit-diff abc1234 --stat              # Statistiques du commit
    git-commit-diff abc1234 --name-only         # Fichiers modifiés
    git-commit-diff abc1234 --name-status       # Fichiers avec statut
    git-commit-diff abc1234 --context 5         # 5 lignes de contexte
    git-commit-diff abc1234 --color             # Diff coloré

DESCRIPTION:
    Affiche le diff d'un commit spécifique de manière formatée.
    Supporte les mêmes références que git (SHA, branch, tag, HEAD~N).
`)
}

func main() {
	statOnly := false
	nameOnly := false
	nameStatus := false
	filesOnly := false
	forceColor := false
	contextLines := 3
	commit := ""

	args := os.Args[1:]
	for len(args) > 0 {
		switch args[0] {
		case "-s", "--stat":
			statOnly = true
			args = args[1:]
		case "-n", "--name-only":
			nameOnly = true
			args = args[1:]
		case "-N", "--name-status":
			nameStatus = true
			args = args[1:]
		case "-p", "--patch":
			args = args[1:]
		case "-U", "--context":
			if len(args) < 2 {
				fmt.Fprintf(os.Stderr, "Erreur: --context nécessite un nombre\n")
				os.Exit(1)
			}
			fmt.Sscanf(args[1], "%d", &contextLines)
			args = args[2:]
		case "-c", "--color":
			forceColor = true
			args = args[1:]
		case "-f", "--files":
			filesOnly = true
			args = args[1:]
		case "-h", "--help":
			printUsage()
			os.Exit(0)
		default:
			if strings.HasPrefix(args[0], "-") && args[0] != "-" {
				fmt.Fprintf(os.Stderr, "Option inconnue: %s\n", args[0])
				printUsage()
				os.Exit(1)
			}
			commit = args[0]
			args = args[1:]
		}
	}

	if commit == "" {
		fmt.Fprintf(os.Stderr, "Erreur: aucun commit spécifié\n")
		printUsage()
		os.Exit(1)
	}

	// Vérifier qu'on est dans un repo git
	_, err := runGit("rev-parse", "--git-dir")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erreur: aucun dépôt git trouvé\n")
		os.Exit(1)
	}

	// Vérifier que le commit existe
	sha, err := runGit("rev-parse", commit)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erreur: commit '%s' introuvable\n", commit)
		os.Exit(1)
	}

	// Afficher les infos du commit
	info, _ := runGit("log", "-1", "--format:%H %an <%ae> %aI %s", commit)
	parts := strings.SplitN(info, " ", 4)
	if len(parts) >= 4 {
		fmt.Printf("Commit: %s\n", parts[0][:8])
		fmt.Printf("Auteur: %s\n", parts[2])
		date := strings.Split(parts[3], " ")[0]
		fmt.Printf("Date:   %s\n", date)
		// Extract subject (everything after date)
		dateEnd := strings.Index(info, parts[3]) + len(parts[3])
		if dateEnd < len(info) {
			subject := strings.TrimSpace(info[dateEnd:])
			fmt.Printf("Sujet:  %s\n", subject)
		}
		fmt.Println(strings.Repeat("-", 60))
	}

	// Force color if requested
	if forceColor {
		os.Setenv("GIT_CONFIG_GLOBAL", "/dev/null")
		runGit("config", "--local", "color.diff", "always")
		runGit("config", "--local", "color.ui", "always")
	}

	// Execute the requested diff mode
	var output string
	if nameOnly {
		output, err = runGit("diff-tree", "--no-commit-id", "--name-only", "-r", sha)
	} else if nameStatus {
		output, err = runGit("diff-tree", "--no-commit-id", "--name-status", "-r", sha)
	} else if filesOnly {
		output, err = runGit("diff-tree", "--no-commit-id", "--name-only", "-r", sha)
	} else if statOnly {
		output, err = runGit("diff", "--stat", sha+"^.."+sha)
	} else {
		output, err = runGit("diff", fmt.Sprintf("-U%d", contextLines), sha+"^.."+sha)
	}

	if err != nil {
		// Root commit has no parent, try without ^
		if strings.Contains(err.Error(), "unknown revision") {
			if nameOnly || nameStatus || filesOnly {
				output, err = runGit("diff-tree", "--root", "--no-commit-id", "--name-only", "-r", sha)
			} else if statOnly {
				output, err = runGit("diff", "--stat", "--root", sha)
			} else {
				output, err = runGit("diff", fmt.Sprintf("-U%d", contextLines), "--root", sha)
			}
			if err != nil {
				fmt.Fprintf(os.Stderr, "Erreur lors du diff: %v\n", err)
				os.Exit(1)
			}
			fmt.Println("(commit racine)")
		} else {
			fmt.Fprintf(os.Stderr, "Erreur lors du diff: %v\n", err)
			os.Exit(1)
		}
	}

	if output != "" {
		fmt.Println(output)
	} else {
		fmt.Println("(aucun changement)")
	}
}
