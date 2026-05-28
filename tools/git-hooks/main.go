// git-hooks installs pre-configured Git hooks into a repository.
//
// Supported hooks:
//   - pre-commit:    run linters / formatters before committing
//   - commit-msg:    enforce commit message conventions
//   - pre-push:      run tests before pushing
//   - pre-rebase:    warn before rebasing
//   - post-merge:    run commands after a successful merge
//
// Usage:
//   git-hooks install [--hooks pre-commit,commit-msg] [--dir .]
//   git-hooks list
//   git-hooks remove [--hooks pre-commit] [--dir .]
//   git-hooks show <hook>
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// Hook defines a git hook with its template.
type Hook struct {
	Name        string
	Description string
	Template    string
}

// pre-commit hook: runs configured commands before each commit.
// Reads optional config from .git-hooks.json in the repo root.
var preCommitTpl = `#!/usr/bin/env bash
# git-hooks: pre-commit
# Installed by git-hooks — edit .git-hooks.json to configure.

set -euo pipefail

HOOKS_CONFIG="$(git rev-parse --show-toplevel)/.git-hooks.json"

if [ ! -f "$HOOKS_CONFIG" ]; then
    echo "git-hooks: no .git-hooks.json found — pre-commit hook is inactive."
    exit 0
fi

FAIL=0

# Run pre-commit commands if defined
CMDS=$(python3 -c "
import json, sys
with open('$HOOKS_CONFIG') as f:
    cfg = json.load(f)
cmds = cfg.get('pre-commit', {}).get('commands', [])
print(' '.join(cmds))
" 2>/dev/null) || CMDS=""

if [ -n "$CMDS" ]; then
    echo "git-hooks: running pre-commit commands..."
    for CMD in $CMDS; do
        echo "  $ $CMD"
        if ! eval "$CMD"; then
            echo "  FAILED: $CMD"
            FAIL=1
        fi
    done
    echo ""
fi

# Check for trailing whitespace
if python3 -c "
import json
with open('$HOOKS_CONFIG') as f:
    cfg = json.load(f)
val = cfg.get('pre-commit', {}).get('trailing-whitespace', True)
sys.exit(0 if val else 1)
" 2>/dev/null; then
    TW_FILES=$(git diff --cached --diff-filter=ACM -z --name-only | \
        xargs -0 -I{} git diff --cached -p --word-diff-regex=.[[:space:]] {} | \
        grep -E '^\+[^+].*\x08' || true)
    if [ -n "$TW_FILES" ]; then
        echo "git-hooks: trailing whitespace detected:"
        echo "$TW_FILES"
        FAIL=1
    fi
fi

# Check for large files (>10MB)
if python3 -c "
import json
with open('$HOOKS_CONFIG') as f:
    cfg = json.load(f)
val = cfg.get('pre-commit', {}).get('large-files', True)
sys.exit(0 if val else 1)
" 2>/dev/null; then
    LARGE=$(git diff --cached --diff-filter=ACM -z --name-only | \
        xargs -0 -I{} sh -c 'SIZE=$(wc -c < "{}"); if [ "$SIZE" -gt 10485760 ]; then echo "{} ($(numfmt --to=iec "$SIZE"))"; fi' || true)
    if [ -n "$LARGE" ]; then
        echo "git-hooks: large files detected (>10MB):"
        echo "$LARGE"
        FAIL=1
    fi
fi

if [ $FAIL -ne 0 ]; then
    echo "git-hooks: pre-commit checks failed. Fix the issues above and try again."
    exit 1
fi

exit 0
`

// commit-msg hook: enforces conventional commit message format.
var commitMsgTpl = `#!/usr/bin/env bash
# git-hooks: commit-msg
# Enforces conventional commit message format.

set -euo pipefail

COMMIT_MSG_FILE="$1"
COMMIT_MSG=$(cat "$COMMIT_MSG_FILE")

# Skip merge commits
if echo "$COMMIT_MSG" | grep -qE "^Merge "; then
    exit 0
fi

# Conventional commit regex: type(optional-scope): description
if ! echo "$COMMIT_MSG" | grep -qE "^(feat|fix|docs|style|refactor|perf|test|build|ci|chore|revert)(\([a-z0-9\-]+\))?: .+"; then
    echo ""
    echo "ERROR: Commit message does not follow conventional commit format."
    echo ""
    echo "Expected: <type>(<scope>): <description>"
    echo ""
    echo "Types: feat, fix, docs, style, refactor, perf, test, build, ci, chore, revert"
    echo ""
    echo "Examples:"
    echo "  feat: add user authentication"
    echo "  fix(api): handle null response"
    echo "  docs: update installation guide"
    echo ""
    exit 1
fi

# Check subject line length (max 72 chars)
SUBJECT=$(echo "$COMMIT_MSG" | head -n1)
if [ ${#SUBJECT} -gt 72 ]; then
    echo ""
    echo "WARNING: Commit subject line is ${#SUBJECT} characters (max 72)."
fi

exit 0
`

// pre-push hook: runs tests before pushing.
var prePushTpl = `#!/usr/bin/env bash
# git-hooks: pre-push
# Runs tests before pushing to remote.

set -euo pipefail

HOOKS_CONFIG="$(git rev-parse --show-toplevel)/.git-hooks.json"

if [ ! -f "$HOOKS_CONFIG" ]; then
    echo "git-hooks: no .git-hooks.json found — pre-push hook is inactive."
    exit 0
fi

# Check if pre-push is enabled
ENABLED=$(python3 -c "
import json
with open('$HOOKS_CONFIG') as f:
    cfg = json.load(f)
val = cfg.get('pre-push', {}).get('enabled', True)
print('yes' if val else 'no')
" 2>/dev/null) || ENABLED="yes"

if [ "$ENABLED" = "no" ]; then
    exit 0
fi

FAIL=0

# Run pre-push commands if defined
CMDS=$(python3 -c "
import json
with open('$HOOKS_CONFIG') as f:
    cfg = json.load(f)
cmds = cfg.get('pre-push', {}).get('commands', [])
print(' '.join(cmds))
" 2>/dev/null) || CMDS=""

if [ -n "$CMDS" ]; then
    echo "git-hooks: running pre-push commands..."
    for CMD in $CMDS; do
        echo "  $ $CMD"
        if ! eval "$CMD"; then
            echo "  FAILED: $CMD"
            FAIL=1
        fi
    done
    echo ""
fi

if [ $FAIL -ne 0 ]; then
    echo "git-hooks: pre-push checks failed. Fix the issues above and try again."
    exit 1
fi

exit 0
`

// pre-rebase hook: warns before rebasing.
var preRebaseTpl = `#!/usr/bin/env bash
# git-hooks: pre-rebase
# Warns before rebasing shared branches.

set -euo pipefail

UPSTREAM="$1"
CURRENT="$2"

# Check if the branch has been pushed
if git rev-parse --verify "$UPSTREAM" >/dev/null 2>&1; then
    LOCAL=$(git rev-parse "$CURRENT")
    REMOTE=$(git rev-parse "$UPSTREAM")

    if [ "$LOCAL" != "$REMOTE" ]; then
        echo "WARNING: Your branch has diverged from $UPSTREAM."
        echo "Rebasing shared history can cause issues for collaborators."
        echo ""
        echo "Consider using 'git pull --rebase' instead."
    fi
fi

exit 0
`

// post-merge hook: runs commands after a successful merge.
var postMergeTpl = `#!/usr/bin/env bash
# git-hooks: post-merge
# Runs commands after a successful merge.

set -euo pipefail

MERGE_STATUS="$1"
HOOKS_CONFIG="$(git rev-parse --show-toplevel)/.git-hooks.json"

if [ ! -f "$HOOKS_CONFIG" ]; then
    exit 0
fi

if [ "$MERGE_STATUS" != "1" ]; then
    exit 0
fi

CMDS=$(python3 -c "
import json
with open('$HOOKS_CONFIG') as f:
    cfg = json.load(f)
cmds = cfg.get('post-merge', {}).get('commands', [])
print(' '.join(cmds))
" 2>/dev/null) || CMDS=""

if [ -n "$CMDS" ]; then
    echo "git-hooks: running post-merge commands..."
    for CMD in $CMDS; do
        echo "  $ $CMD"
        eval "$CMD" || true
    done
    echo ""
fi

exit 0
`

var hooks = map[string]Hook{
	"pre-commit": {"pre-commit", "Run linters, formatters, and checks before committing", preCommitTpl},
	"commit-msg": {"commit-msg", "Enforce conventional commit message format", commitMsgTpl},
	"pre-push":   {"pre-push", "Run tests before pushing to remote", prePushTpl},
	"pre-rebase": {"pre-rebase", "Warn before rebasing shared branches", preRebaseTpl},
	"post-merge": {"post-merge", "Run commands after a successful merge", postMergeTpl},
}

var defaultConfig = map[string]interface{}{
	"pre-commit": map[string]interface{}{
		"commands":           []string{},
		"trailing-whitespace": true,
		"large-files":        true,
	},
	"commit-msg": map[string]interface{}{
		"conventional": true,
	},
	"pre-push": map[string]interface{}{
		"enabled":  true,
		"commands": []string{},
	},
	"post-merge": map[string]interface{}{
		"commands": []string{},
	},
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "install":
		cmdInstall(os.Args[2:])
	case "list":
		cmdList()
	case "remove":
		cmdRemove(os.Args[2:])
	case "show":
		cmdShow(os.Args[2:])
	case "init":
		cmdInit(os.Args[2:])
	case "help", "--help", "-h":
		printUsage()
	default:
		fmt.Fprintf(os.Stderr, "git-hooks: unknown command %q\n", os.Args[1])
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println(`git-hooks — Install pre-configured Git hooks

Usage:
  git-hooks install [flags]    Install hooks into the current repository
  git-hooks list               List available hooks
  git-hooks remove [flags]     Remove installed hooks
  git-hooks show <hook>        Show the content of a hook
  git-hooks init [dir]         Create a default .git-hooks.json config

Flags for install/remove:
  -hooks    Comma-separated list of hooks (default: all)
  -dir      Repository root directory (default: current dir)

Examples:
  git-hooks install
  git-hooks install -hooks pre-commit,commit-msg
  git-hooks install -dir ../my-project
  git-hooks remove -hooks pre-commit
  git-hooks show pre-commit
  git-hooks init

Available hooks:
  pre-commit   Run checks before committing
  commit-msg   Enforce commit message conventions
  pre-push     Run tests before pushing
  pre-rebase   Warn before rebasing
  post-merge   Run commands after merging

Configure hooks by editing .git-hooks.json in your repo root.`)
}

func cmdInstall(args []string) {
	fs := flag.NewFlagSet("install", flag.ExitOnError)
	hooksFlag := fs.String("hooks", "", "Comma-separated list of hooks to install")
	dirFlag := fs.String("dir", ".", "Repository root directory")
	fs.Parse(args)

	repoRoot, err := findRepoRoot(*dirFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "git-hooks: %v\n", err)
		os.Exit(1)
	}

	hooksDir := filepath.Join(repoRoot, ".git", "hooks")

	// Determine which hooks to install
	var toInstall []string
	if *hooksFlag != "" {
		toInstall = strings.Split(*hooksFlag, ",")
	} else {
		for name := range hooks {
			toInstall = append(toInstall, name)
		}
	}

	// Create default config if it doesn't exist
	configPath := filepath.Join(repoRoot, ".git-hooks.json")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		if err := createDefaultConfig(configPath); err != nil {
			fmt.Fprintf(os.Stderr, "git-hooks: failed to create default config: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Created default config: %s\n", configPath)
	}

	installed := 0
	for _, name := range toInstall {
		name = strings.TrimSpace(name)
		hook, ok := hooks[name]
		if !ok {
			fmt.Fprintf(os.Stderr, "git-hooks: unknown hook %q\n", name)
			continue
		}

		hookPath := filepath.Join(hooksDir, name)
		if err := os.WriteFile(hookPath, []byte(hook.Template), 0755); err != nil {
			fmt.Fprintf(os.Stderr, "git-hooks: failed to write %s: %v\n", name, err)
			continue
		}
		fmt.Printf("Installed: %s\n", name)
		installed++
	}

	if installed == 0 {
		os.Exit(1)
	}
	fmt.Printf("\nInstalled %d hook(s) in %s\n", installed, hooksDir)
}

func cmdList() {
	fmt.Println("Available Git hooks:")
	fmt.Println()
	for _, name := range []string{"pre-commit", "commit-msg", "pre-push", "pre-rebase", "post-merge"} {
		h := hooks[name]
		fmt.Printf("  %-14s %s\n", h.Name+":", h.Description)
	}
	fmt.Println()
	fmt.Println("Configure by editing .git-hooks.json in your repo root.")
}

func cmdRemove(args []string) {
	fs := flag.NewFlagSet("remove", flag.ExitOnError)
	hooksFlag := fs.String("hooks", "", "Comma-separated list of hooks to remove")
	dirFlag := fs.String("dir", ".", "Repository root directory")
	fs.Parse(args)

	repoRoot, err := findRepoRoot(*dirFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "git-hooks: %v\n", err)
		os.Exit(1)
	}

	hooksDir := filepath.Join(repoRoot, ".git", "hooks")

	var toRemove []string
	if *hooksFlag != "" {
		toRemove = strings.Split(*hooksFlag, ",")
	} else {
		for name := range hooks {
			toRemove = append(toRemove, name)
		}
	}

	removed := 0
	for _, name := range toRemove {
		name = strings.TrimSpace(name)
		hookPath := filepath.Join(hooksDir, name)
		if _, ok := hooks[name]; !ok {
			fmt.Fprintf(os.Stderr, "git-hooks: unknown hook %q\n", name)
			continue
		}
		if err := os.Remove(hookPath); err != nil {
			if os.IsNotExist(err) {
				fmt.Printf("Not installed: %s\n", name)
				continue
			}
			fmt.Fprintf(os.Stderr, "git-hooks: failed to remove %s: %v\n", name, err)
			continue
		}
		fmt.Printf("Removed: %s\n", name)
		removed++
	}

	if removed == 0 {
		fmt.Println("No hooks were removed.")
	} else {
		fmt.Printf("\nRemoved %d hook(s) from %s\n", removed, hooksDir)
	}
}

func cmdShow(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "git-hooks: specify a hook name\n")
		fmt.Fprintf(os.Stderr, "Available: pre-commit, commit-msg, pre-push, pre-rebase, post-merge\n")
		os.Exit(1)
	}

	name := args[0]
	hook, ok := hooks[name]
	if !ok {
		fmt.Fprintf(os.Stderr, "git-hooks: unknown hook %q\n", name)
		fmt.Fprintf(os.Stderr, "Available: pre-commit, commit-msg, pre-push, pre-rebase, post-merge\n")
		os.Exit(1)
	}

	fmt.Printf("# %s: %s\n\n", hook.Name, hook.Description)
	fmt.Print(hook.Template)
}

func cmdInit(args []string) {
	dir := "."
	if len(args) > 0 {
		dir = args[0]
	}

	configPath := filepath.Join(dir, ".git-hooks.json")
	if _, err := os.Stat(configPath); err == nil {
		fmt.Fprintf(os.Stderr, "git-hooks: %s already exists\n", configPath)
		os.Exit(1)
	}

	if err := createDefaultConfig(configPath); err != nil {
		fmt.Fprintf(os.Stderr, "git-hooks: failed to create config: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created: %s\n", configPath)
	fmt.Println("\nEdit this file to configure your hooks, then run:")
	fmt.Println("  git-hooks install")
}

func createDefaultConfig(path string) error {
	data, err := json.MarshalIndent(defaultConfig, "", "  ")
	if err != nil {
		return err
	}
	// Add header comment as raw bytes before JSON
	var out bytes.Buffer
	out.WriteString("// git-hooks configuration\n")
	out.WriteString("// Edit this file to customize hook behavior.\n")
	out.WriteString("// See https://github.com/TataneSan/git-hooks for documentation.\n\n")
	out.Write(data)
	out.WriteString("\n")
	return os.WriteFile(path, out.Bytes(), 0644)
}

func findRepoRoot(dir string) (string, error) {
	abs, err := filepath.Abs(dir)
	if err != nil {
		return "", fmt.Errorf("failed to resolve path: %w", err)
	}

	for {
		gitDir := filepath.Join(abs, ".git")
		if _, err := os.Stat(gitDir); err == nil {
			return abs, nil
		}
		if abs == "/" || abs == "." {
			break
		}
		abs = filepath.Dir(abs)
	}
	return "", fmt.Errorf("not a git repository (or any parent up to mountpoint)")
}

// Ensure template is used to avoid unused import.
func init() {
	_ = template.FuncMap{}
}
