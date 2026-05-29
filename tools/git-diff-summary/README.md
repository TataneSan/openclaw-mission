# git-diff-summary

Summarize changes between two git commits with file-level stats.

## Install

```bash
go install github.com/TataneSan/git-diff-summary@latest
```

Or build from source:

```bash
go build -o git-diff-summary .
```

## Usage

```
git-diff-summary [options]
```

### Options

| Flag | Default | Description |
|------|---------|-------------|
| `--from` | `HEAD~1` | Source commit (ref, SHA, tag) |
| `--to` | `HEAD` | Target commit (ref, SHA, tag) |

## Examples

Show changes between last commit and HEAD:

```bash
git-diff-summary
```

Compare two specific commits:

```bash
git-diff-summary --from abc1234 --to def5678
```

Compare two branches:

```bash
git-diff-summary --from main --to feature
```

Compare a tag with HEAD:

```bash
git-diff-summary --from v1.0.0 --to HEAD
```

## Output

```
Diff: HEAD~1..HEAD
============================================================
Files changed: 5
  Inserted:    2
  Deleted:     1
Lines added:   142
Lines removed: 38
Net change:    +104 lines

  [M] src/main.go      +85 -20
  [A] src/utils.go      +57 -0
  [D] old/legacy.go      +0 -18
  [M] README.md         +3 -0
  [R] config.yml         -7 -0
```

### Status codes

| Code | Meaning |
|------|---------|
| `M` | Modified |
| `A` | Added |
| `D` | Deleted |
| `R` | Renamed |

## Features

- File-level line change stats (added/removed)
- Automatic status detection (modified, added, deleted, renamed)
- Summary with total files, insertions, deletions, and net change
- Aligned output for easy reading
- Works with any git ref (branches, tags, SHAs)

## License

MIT
