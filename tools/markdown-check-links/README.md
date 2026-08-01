# markdown-check-links

Check internal anchors in Markdown files: every `[text](#anchor)` must resolve to a heading in the same file (GitHub-style slug). Single Go binary, stdlib only.

## Install

```bash
git clone https://github.com/TataneSan/markdown-check-links
cd markdown-check-links
go build -o markdown-check-links .
```

## Usage

```bash
# from stdin
markdown-check-links < README.md

# files
markdown-check-links docs/*.md

# JSON output
markdown-check-links -json README.md

# exit code only (good for CI)
markdown-check-links -q README.md && echo ok

# also validate the syntax of http(s) URLs and optionally ping them
markdown-check-links -check-url -ping README.md
```

## Slug rules

Anchors follow GitHub's Markdown slug rules:

- lowercase
- spaces become hyphens
- punctuation is removed (except `-` and `_`)
- multiple consecutive hyphens collapse into one
- leading/trailing hyphens are trimmed

A heading `# Café & Beer` therefore has the anchor `#caf-beer`.

## Exit codes

| code | meaning |
|---|---|
| 0 | all anchors resolve |
| 1 | at least one anchor is broken |
| 2 | input read error |

## License

MIT
