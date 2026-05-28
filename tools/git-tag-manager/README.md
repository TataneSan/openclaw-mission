# git-tag-manager

A CLI tool for managing Git tags — create, list, delete, move, and push tags with formatted output.

## Install

```bash
go install github.com/TataneSan/git-tag-manager@latest
```

Or build from source:

```bash
go build -o git-tag-manager .
```

## Usage

```
git-tag-manager <command> [options]

Commands:
  list [format]     List tags (table, json, verbose)
  create <name>     Create a new tag
  delete <name>     Delete a tag
  move <name> <ref> Move tag to a new commit
  push              Push all tags to origin
  latest            Show the latest tag

Options:
  -a, --annotated   Create annotated tag
  -m, --message     Tag message
  -s, --sign        GPG sign tag
  -f, --force       Force operation (overwrite existing tag)
  -h, --help        Show help
```

## Examples

### List tags

```bash
git-tag-manager list
```

```
TAG                       TYPE     DATE    SUBJECT
------------------------------------------------------------------------------------------
v1.2.0                    annotated 2026-05-28 Release v1.2.0
v1.1.0                    annotated 2026-05-27 Release v1.1.0
v1.0.0                    light     2026-05-20 Initial release

3 tag(s)
```

### List as JSON

```bash
git-tag-manager list json
```

### Verbose output

```bash
git-tag-manager list verbose
```

### Create a lightweight tag

```bash
git-tag-manager create v1.0.0
```

### Create an annotated tag

```bash
git-tag-manager create v1.0.0 -a -m "Initial release"
```

### Delete a tag

```bash
git-tag-manager delete v0.9.0
```

### Move a tag to a new commit

```bash
git-tag-manager move v1.0.0 abc1234
```

### Push tags to remote

```bash
git-tag-manager push
```

### Show latest tag

```bash
git-tag-manager latest
```

```
v1.2.0 (a1b2c3d4)
  Release v1.2.0
  2026-05-28
```

## License

MIT
