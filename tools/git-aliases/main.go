// git-aliases manages git aliases: add, list, remove, and export.
//
// Usage:
//
//	git-aliases list
//	git-aliases add <name> <command>
//	git-aliases remove <name>
//	git-aliases export [--format json|shell|gitconfig]
//
// Examples:
//
//	git-aliases add co git checkout
//	git-aliases add br git branch
//	git-aliases add st git status -sb
//	git-aliases list
//	git-aliases export --format shell
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"
)

type Alias struct {
	Name    string
	Command string
}

func runGit(args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("git %s: %w\n%s", strings.Join(args, " "), err, out)
	}
	return strings.TrimSpace(string(out)), nil
}

func listAliases() ([]Alias, error) {
	output, err := runGit("config", "--get-all", "alias.*")
	if err != nil {
		// git config --get-all returns error if no aliases exist
		if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() == 1 {
			return nil, nil
		}
		return nil, err
	}

	if output == "" {
		return nil, nil
	}

	var aliases []Alias
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// Format: alias.name = value
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		name := strings.TrimPrefix(strings.TrimSpace(parts[0]), "alias.")
		command := strings.TrimSpace(parts[1])

		aliases = append(aliases, Alias{
			Name:    name,
			Command: command,
		})
	}

	sort.Slice(aliases, func(i, j int) bool {
		return aliases[i].Name < aliases[j].Name
	})

	return aliases, nil
}

func addAlias(name, command string) error {
	// Validate name: should not conflict with built-in git commands
	builtins := map[string]bool{
		"add": true, "branch": true, "checkout": true, "clone": true,
		"commit": true, "diff": true, "fetch": true, "init": true,
		"log": true, "merge": true, "pull": true, "push": true,
		"remote": true, "reset": true, "revert": true, "show": true,
		"status": true, "tag": true, "stash": true, "submodule": true,
	}

	if builtins[name] {
		return fmt.Errorf("alias name '%s' conflicts with a built-in git command", name)
	}

	// Check if alias already exists
	_, err := runGit("config", "--get", "alias."+name)
	if err == nil {
		fmt.Printf("Alias '%s' already exists. Use 'git-aliases remove %s' first, or it will be overwritten.\n", name, name)
	}

	_, err = runGit("config", "alias."+name, command)
	return err
}

func removeAlias(name string) error {
	_, err := runGit("config", "--get", "alias."+name)
	if err != nil {
		return fmt.Errorf("alias '%s' does not exist", name)
	}
	_, err = runGit("config", "--unset", "alias."+name)
	return err
}

func exportAliases(aliases []Alias, format string) error {
	switch format {
	case "json":
		return exportJSON(aliases)
	case "shell":
		return exportShell(aliases)
	case "gitconfig":
		return exportGitConfig(aliases)
	default:
		return fmt.Errorf("unknown format: %s (supported: json, shell, gitconfig)", format)
	}
}

func exportJSON(aliases []Alias) error {
	data, err := json.MarshalIndent(aliases, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}

func exportShell(aliases []Alias) error {
	fmt.Println("# Git aliases as shell functions")
	fmt.Println()
	for _, a := range aliases {
		// Escape single quotes in command
		cmd := strings.ReplaceAll(a.Command, "'", "'\\''")
		fmt.Printf("git-%s() {\n", a.Name)
		fmt.Printf("    git %s \"$@\"\n", cmd)
		fmt.Printf("}\n\n")
	}
	return nil
}

func exportGitConfig(aliases []Alias) error {
	fmt.Println("[alias]")
	for _, a := range aliases {
		fmt.Printf("    %s = %s\n", a.Name, a.Command)
	}
	return nil
}

func printHelp() {
	fmt.Println("git-aliases — Manage git aliases from the command line")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  git-aliases <command> [arguments]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  list                    List all git aliases")
	fmt.Println("  add <name> <command>    Add a new alias")
	fmt.Println("  remove <name>           Remove an alias")
	fmt.Println("  export [--format fmt]   Export aliases (json, shell, gitconfig)")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  git-aliases add co checkout")
	fmt.Println("  git-aliases add br branch")
	fmt.Println("  git-aliases add st 'status -sb'")
	fmt.Println("  git-aliases add last 'log -1 HEAD'")
	fmt.Println("  git-aliases list")
	fmt.Println("  git-aliases export --format shell")
}

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(1)
	}

	if os.Args[1] == "-h" || os.Args[1] == "--help" {
		printHelp()
		os.Exit(0)
	}

	// Check we're in a git repo
	if _, err := runGit("rev-parse", "--git-dir"); err != nil {
		fmt.Fprintf(os.Stderr, "Error: not a git repository\n")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "list":
		aliases, err := listAliases()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		if len(aliases) == 0 {
			fmt.Println("No aliases found.")
			return
		}

		fmt.Printf("  %-20s %s\n", "ALIAS", "COMMAND")
		fmt.Println("  " + strings.Repeat("─", 60))
		for _, a := range aliases {
			fmt.Printf("  %-20s %s\n", a.Name, a.Command)
		}
		fmt.Printf("\n  %d alias(es)\n", len(aliases))

	case "add":
		if len(os.Args) < 4 {
			fmt.Fprintf(os.Stderr, "Usage: git-aliases add <name> <command>\n")
			os.Exit(1)
		}
		name := os.Args[2]
		cmdStr := strings.Join(os.Args[3:], " ")

		if err := addAlias(name, cmdStr); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Added alias: %s = %s\n", name, cmdStr)

	case "remove":
		if len(os.Args) < 3 {
			fmt.Fprintf(os.Stderr, "Usage: git-aliases remove <name>\n")
			os.Exit(1)
		}
		name := os.Args[2]

		if err := removeAlias(name); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Removed alias: %s\n", name)

	case "export":
		format := "json"
		for i := 2; i < len(os.Args); i++ {
			if os.Args[i] == "--format" && i+1 < len(os.Args) {
				format = os.Args[i+1]
				break
			}
		}

		aliases, err := listAliases()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		if err := exportAliases(aliases, format); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		fmt.Fprintf(os.Stderr, "Try 'git-aliases --help' for more information.\n")
		os.Exit(1)
	}
}
