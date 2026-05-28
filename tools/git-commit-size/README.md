# git-commit-size

Display commit sizes (files touched, lines of code) with visual bar charts.

## Features

- Per-commit statistics: files touched, insertions, deletions
- Visual bar charts for insertions and deletions
- Sortable by files, insertions, deletions, or total
- Per-file breakdown with `-files` flag
- JSON output support
- Colored terminal output
- Arbitrary revision ranges

## Install

```bash
go install github.com/TataneSan/git-commit-size@latest
```

Or build from source:

```bash
git clone https://github.com/TataneSan/git-commit-size.git
cd git-commit-size
go build -o git-commit-size .
```

## Usage

```
git-commit-size [options] [revision]
```

### Options

| Flag | Description | Default |
|------|-------------|---------|
| `-n` | Number of commits to show | `10` |
| `-all` | Show all commits in history | off |
| `-files` | Show per-file breakdown | off |
| `-sort` | Sort by: `files`, `insertions`, `deletions`, `total` | `files` |
| `-top` | Top N files per commit (0 = all) | `0` |
| `-bar-width` | Width of bar chart | `30` |
| `-no-color` | Disable colored output | off |
| `-json` | Output as JSON | off |

### Examples

Last 10 commits:

```bash
git-commit-size
```

Last 20 commits sorted by insertions:

```bash
git-commit-size -n 20 -sort insertions
```

All commits with per-file breakdown (top 5 files):

```bash
git-commit-size -all -files -top 5
```

Compare branches:

```bash
git-commit-size main..develop
```

JSON output:

```bash
git-commit-size -json -n 5 > commits.json
```

### Sample Output

```
HASH      FILES  +INS    -DEL     SUBJECT
------------------------------------------------------------------------------------------
a1b2c3d  15     342     88       feat: add user authentication
          +████████████████████████ -██████████████

e4f5g6h  3      120     5        fix: correct date formatting
          +██████████ -█

------------------------------------------------------------------------------------------
Total: 2 commits, 18 files, +462     ins, -93      del
```

## License

MIT
