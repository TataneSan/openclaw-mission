# git-tag-list

List git tags with details: date, type (annotated/lightweight), author, message, and target commit.

## Install

```bash
go install github.com/TataneSan/git-tag-list@latest
```

## Usage

```
git-tag-list [options] [repo]
```

## Options

| Flag | Description |
|------|-------------|
| `-f, --format FORMAT` | Output format: `table` (default), `json`, `csv` |
| `-s, --sort FIELD` | Sort field: `name` (default), `date`, `author` |
| `-r, --reverse` | Reverse sort order |
| `-h, --help` | Show help message |

## Arguments

| Arg | Description |
|-----|-------------|
| `repo` | Path to git repository (default: current directory) |

## Examples

List tags in current repo:

```bash
git-tag-list
```

Output as JSON:

```bash
git-tag-list --format json
```

Sort by date, newest first:

```bash
git-tag-list --sort date --reverse
```

List tags in a specific repo:

```bash
git-tag-list ./path/to/repo
```

## Output

### Table (default)

```
NAME                      TYPE         DATE       AUTHOR               MESSAGE
----------------------------------------------------------------------------------------------------
v1.0.0                    lightweight  2026-05-28 Alice                Initial release
v2.0.0                    annotated    2026-05-29 Alice                Release v2

2 tag(s)
```

### JSON

```json
[
  {
    "name": "v1.0.0",
    "type": "lightweight",
    "target": "a1b2c3d4e5",
    "date": "2026-05-28T12:00:00Z",
    "author": "Alice",
    "message": "Initial release"
  }
]
```

### CSV

```csv
name,type,target,date,author,message
v1.0.0,lightweight,a1b2c3d4e5,2026-05-28,Alice,Initial release
```

## Features

- Detects annotated vs lightweight tags
- Shows tagger date and author
- Retrieves commit author for lightweight tags
- Sortable by name, date, or author
- Three output formats: table, JSON, CSV

## Exit Codes

| Code | Meaning |
|------|---------|
| 0 | Success |
| 1 | Error (not a git repo, etc.) |

## License

MIT
