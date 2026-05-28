# go-vet-check

Run `go vet` with formatted output and colored reports.

Wraps `go vet` to provide cleaner package grouping, optional color support, and JSON output for CI integration.

## Install

```bash
go install github.com/TataneSan/go-vet-check@latest
```

## Usage

```bash
go-vet-check [OPTIONS] [PACKAGES]
```

### Options

| Flag | Description |
|------|-------------|
| `-f, --format FORMAT` | Output format: `text` (default), `json` |
| `-n, --no-color` | Disable colored output |
| `-d, --dir DIRECTORY` | Run in specified directory (default: `.`) |
| `-h, --help` | Show help message |

### Exit Codes

| Code | Meaning |
|------|---------|
| `0` | No issues found |
| `1` | Issues found |
| `2` | Internal error |

## Examples

Run on current directory:

```bash
go-vet-check
# ok  no issues found
```

Run on all packages with JSON output:

```bash
go-vet-check -f json ./...
# [{"package": "myapp", "issues": ["./main.go:5:2: ..."]}]
```

Run in a specific directory:

```bash
go-vet-check -d ./myproject ./...
```

Disable colors:

```bash
go-vet-check -n
```

## CI Integration

GitHub Actions:

```yaml
- name: Go vet
  run: |
    go install github.com/TataneSan/go-vet-check@latest
    go-vet-check -n ./...
```

Makefile:

```makefile
vet:
	go-vet-check -n ./...
```

## Output Formats

### Text (default)

```
FAIL  2 issues in 1 package

github.com/user/project:
  ./main.go:5:2: fmt.Printf format %s has arg 42 of wrong type int
  ./util.go:10:1: unreachable code
```

### JSON

```json
[
  {
    "package": "github.com/user/project",
    "issues": [
      "./main.go:5:2: fmt.Printf format %s has arg 42 of wrong type int",
      "./util.go:10:1: unreachable code"
    ]
  }
]
```

## License

MIT
