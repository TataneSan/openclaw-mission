# git-commit-size

CLI tool that shows the size of git commits — files added, modified, and deleted.

## Install

```bash
go install github.com/TataneSan/git-commit-size@latest
```

Or build from source:

```bash
go build -o git-commit-size .
```

## Usage

```
git-commit-size [flags]
```

### Flags

| Flag | Default | Description |
|------|---------|-------------|
| `--ref` | `HEAD` | Commit ref or range (e.g. `HEAD`, `main..develop`, `HEAD~5..HEAD`) |
| `--output` | `table` | Output format: `table`, `summary`, `json` |
| `--repo` | `.` | Path to git repository |
| `--summary` | `false` | Show summary by author (shorthand for `--output=summary`) |
| `--files` | `false` | Show all changed files per commit |

### Examples

Show the current commit:

```bash
git-commit-size
```

```
Commit: a1b2c3d
Author: Jane Doe
Date:   2026-05-29
Message: feat: add user authentication

Files: 3 added, 1 modified, 0 deleted (total: 4)

Changed files:
  [A] auth/login.go
  [A] auth/token.go
  [A] auth/middleware.go
  [M] main.go
```

Show the last 10 commits in table format:

```bash
git-commit-size --ref main
```

```
Commit     Added    Modified Deleted  Total   Message
------------------------------------------------------------------------------------------
a1b2c3d    3        1        0        4       feat: add user authentication
e4f5g6h    0        2        1        3       fix: correct validation logic
i7j8k9l    1        0        0        1       docs: update README
------------------------------------------------------------------------------------------
TOTAL      4        3        1        8       3 commits
```

Show a commit range:

```bash
git-commit-size --ref main..develop
```

Show summary by author:

```bash
git-commit-size --ref main --summary
```

```
Summary: 15 commits

Total files: 42 added, 28 modified, 10 deleted

Author                    Commits  Added   Modified Deleted Total
---------------------------------------------------------------------------
Jane Doe                  10       30      15       5       50
John Smith                5        12      13       5       30
```

Output as JSON:

```bash
git-commit-size --ref HEAD~5..HEAD --output json
```

Show all changed files for each commit:

```bash
git-commit-size --ref main --files
```

Use a specific repository:

```bash
git-commit-size --repo /path/to/repo
```

## License

MIT
