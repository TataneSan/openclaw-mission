# git-merge-base

Find the common ancestor (merge base) between two git branches.

A simple CLI tool that wraps `git merge-base` with a cleaner interface, branch listing, and verbose commit details.

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
git-merge-base <branch1> <branch2> [options]
```

### Options

| Flag | Description |
|------|-------------|
| `-v, --verbose` | Show detailed commit info for the merge base |
| `-l, --list` | List all branches (use with `--all` for remote branches) |
| `-a, --all` | Include remote branches in listing |
| `-h, --help` | Show help message |

## Examples

### Find merge base between two branches

```bash
$ git-merge-base main feature-x
a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6q7r8s9t0
```

### Verbose output

```bash
$ git-merge-base main feature-x -v
Merge base between 'main' and 'feature-x':

  Commit: a1b2c3d4e5f6
  John Doe <john@example.com> 2026-05-28 Add user authentication
```

### List branches

```bash
$ git-merge-base --list
main
feature-x
develop
```

### Use with any git ref

```bash
$ git-merge-base HEAD~5 develop
$ git-merge-base v1.0.0 v2.0.0
```

## Requirements

- Git installed and available in PATH
- Must be run inside a git repository

## License

MIT
