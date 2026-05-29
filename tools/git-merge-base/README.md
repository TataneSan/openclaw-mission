# git-merge-base

Find the common ancestor (merge base) between two git branches or commits.

## Install

```bash
go install github.com/TataneSan/git-merge-base@latest
```

Or build from source:

```bash
git clone https://github.com/TataneSan/git-merge-base.git
cd git-merge-base
go build -o git-merge-base .
```

## Usage

```
git-merge-base [--verbose] <branch1> <branch2>
```

### Flags

| Flag | Description |
|------|-------------|
| `--verbose`, `-v` | Show commit message, author, date, and ahead/behind counts |

## Examples

### Basic

```bash
git-merge-base main feature-login
# Merge base between "main" and "feature-login":
#   a1b2c3d4e5f6g7h8i9j0
```

### Verbose

```bash
git-merge-base main develop --verbose
# Merge base between "main" and "develop":
#   a1b2c3d4e5f6g7h8i9j0
#   Message: feat: add authentication module
#   Author:  Jane Doe <jane@example.com>
#   Date:    2026-05-15 14:30:00 +0200
#   main is 12 commits ahead, 3 behind
#   develop is 8 commits ahead, 1 behind
```

### With commit SHAs

```bash
git-merge-base abc1234 def5678
```

## How it works

Uses `git merge-base` to find the best common ancestor between two commits. In verbose mode, displays the commit message, author, date, and how many commits each branch is ahead/behind the merge base.

## License

MIT
