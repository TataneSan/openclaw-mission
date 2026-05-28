# git-tag-list

List Git tags with details including date, type, author, and commit message. Supports multiple output formats and sorting options.

## Install

```bash
go install github.com/TataneSan/git-tag-list@latest
```

## Usage

```
git-tag-list [-d DIR] [-sort SORT] [-r] [-v] [-f FORMAT]
```

### Options

| Flag | Description | Default |
|------|-------------|---------|
| `-d` | Repository directory | `.` |
| `-sort` | Sort order: `date`, `name`, `version` | `date` |
| `-r` | Reverse sort order | off |
| `-v` | Verbose: show commit messages | off |
| `-f` | Output format: `table`, `list`, `json` | `table` |

## Examples

### List all tags

```bash
git-tag-list
```

Output:

```
NAME       DATE                   TYPE    COMMIT
----       ----                   ----    ------
v1.0.0     2024-01-15 10:30:00   tag     a1b2c3d4
v1.1.0     2024-02-20 14:15:00   tag     e5f6g7h8
v2.0.0     2024-03-10 09:00:00   tag     i9j0k1l2

3 tag(s)
```

### Verbose with commit messages

```bash
git-tag-list -v
```

### Sort by name, reversed

```bash
git-tag-list -sort name -r
```

### JSON output

```bash
git-tag-list -f json
```

### Specific repository

```bash
git-tag-list -d /path/to/repo
```

## License

MIT
