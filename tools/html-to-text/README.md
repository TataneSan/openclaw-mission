# html-to-text

Convert HTML files to plain text. Strips tags, decodes entities, preserves structure.

## Install

```bash
go install github.com/TataneSan/html-to-text@latest
```

Or build from source:

```bash
git clone https://github.com/TataneSan/html-to-text.git
cd html-to-text
go build -o html-to-text .
```

## Usage

```
html-to-text [options] [file]
```

### Options

| Flag | Description | Default |
|------|-------------|---------|
| `-w, --width N` | Output line width | 0 (no wrapping) |
| `-h, --help` | Show help | — |

## Examples

```bash
# Convert HTML to text
html-to-text page.html

# From stdin
curl -s https://example.com | html-to-text

# Pipe to pager
html-to-text page.html | less
```

## License

MIT
