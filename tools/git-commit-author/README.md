# git-commit-author

Display detailed author information for the last commit in a git repository.

## Install

```bash
go install github.com/TataneSan/git-commit-author@latest
```

Or build from source:

```bash
git clone https://github.com/TataneSan/git-commit-author.git
cd git-commit-author
go build -o git-commit-author .
```

## Usage

```
git-commit-author [options]
```

### Options

| Flag | Default | Description |
|------|---------|-------------|
| `-ref` | `HEAD` | Git ref to inspect (commit hash, branch, tag, etc.) |
| `-format` | `human` | Output format: `human`, `table`, `json`, `raw` |
| `-name` | | Show only author name |
| `-email` | | Show only author email |
| `-date` | | Show only author date (RFC3339) |
| `-hash` | | Show only full commit hash |
| `-short` | | Show only short commit hash |
| `-verbose` | | Show full hash and tree in human format |

## Examples

### Show last commit author info

```bash
git-commit-author
```

Output:
```
  a1b2c3d
  ───────

  Author:  John Doe
  Email:   john@example.com
  Date:    2026-05-29 14:30:00 UTC

  Subject: feat: add user authentication
```

### Show only author name

```bash
git-commit-author -name
# John Doe
```

### Show only author email

```bash
git-commit-author -email
# john@example.com
```

### JSON output for scripting

```bash
git-commit-author -format json
```

Output:
```json
{
  "hash": "a1b2c3d4e5f6...",
  "short_hash": "a1b2c3d",
  "author": {
    "name": "John Doe",
    "email": "john@example.com"
  },
  "date": "2026-05-29T14:30:00+00:00",
  "author_date": "2026-05-29T14:30:00+00:00",
  "subject": "feat: add user authentication",
  "body": "",
  "tree": "f6e5d4c3b2a1...",
  "parents": ["prev123"]
}
```

### Inspect a specific commit

```bash
git-commit-author -ref abc1234
```

### Table format

```bash
git-commit-author -format table
```

Output:
```
╔══════════════════════════════════════════════════════════════╗
║                    COMMIT AUTHOR INFO                       ║
╠══════════════════════════════════════════════════════════════╣
║  Hash:       a1b2c3d                              ║
║  Author:     John Doe                              ║
║  Email:      john@example.com                      ║
║  Date:       2026-05-29 14:30:00 UTC               ║
╠══════════════════════════════════════════════════════════════╣
║  Subject:    feat: add user authentication          ║
╚══════════════════════════════════════════════════════════════╝
```

### Raw format (for shell scripts)

```bash
git-commit-author -format raw
# John Doe <john@example.com> 2026-05-29T14:30:00+00:00 a1b2c3d
```

### Verbose output

```bash
git-commit-author -verbose
```

Shows full commit hash and tree hash in addition to the standard output.

## License

MIT
