# git-aliases

Manage git aliases from the command line: add, list, remove, and export.

## Features

- List all git aliases in a formatted table
- Add new aliases with built-in conflict detection
- Remove aliases safely
- Export aliases in multiple formats: JSON, shell functions, gitconfig
- Validates against built-in git command names
- Single binary, no dependencies

## Install

```bash
go install github.com/TataneSan/git-aliases@latest
```

Or build from source:

```bash
git clone https://github.com/TataneSan/git-aliases.git
cd git-aliases
go build -o git-aliases .
cp git-aliases /usr/local/bin/
```

## Usage

```
git-aliases <command> [arguments]
```

### Commands

**List aliases:**
```bash
git-aliases list
```

**Add an alias:**
```bash
git-aliases add co checkout
git-aliases add br branch
git-aliases add st 'status -sb'
git-aliases add last 'log -1 HEAD'
```

**Remove an alias:**
```bash
git-aliases remove co
```

**Export aliases:**
```bash
git-aliases export --format json
git-aliases export --format shell
git-aliases export --format gitconfig
```

### Export Formats

**JSON:**
```json
[
  {"Name": "co", "Command": "checkout"},
  {"Name": "br", "Command": "branch"}
]
```

**Shell functions:**
```bash
git-co() {
    git checkout "$@"
}
```

**Git config:**
```ini
[alias]
    co = checkout
    br = branch
```

## Output Example

```
$ git-aliases list
  ALIAS                COMMAND
  ────────────────────────────────────────────────────
  br                   branch
  co                   checkout
  last                 log -1 HEAD
  st                   status -sb

  4 alias(es)
```

## Requirements

- Go 1.21+
- Git

## License

MIT
