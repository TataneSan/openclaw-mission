# git-ignored

Lists files that would be ignored by a `.gitignore` file. Useful for auditing what's being excluded from version control.

## Install

```bash
go install github.com/TataneSan/git-ignored@latest
```

Or build from source:

```bash
git clone https://github.com/TataneSan/git-ignored.git
cd git-ignored
go build -o git-ignored .
```

## Usage

```
git-ignored [flags] [directory]
```

### Flags

| Flag | Description |
|------|-------------|
| `-json` | Output as JSON |
| `-gitignore` | Path to .gitignore file (default: `.gitignore`) |

### Examples

Check current directory:
```bash
git-ignored
```

Check a specific directory:
```bash
git-ignored ./src
```

JSON output:
```bash
git-ignored -json
```

Custom gitignore file:
```bash
git-ignored -gitignore .gitignore-global ./project
```

### Output

Text mode:
```
build/output.log (matched by '*.log')
tmp/cache.tmp (matched by '*.tmp')

2 file(s) ignored
```

JSON mode:
```json
[
  {
    "file": "build/output.log",
    "ignored": true,
    "pattern": "*.log"
  }
]
```

## License

MIT
