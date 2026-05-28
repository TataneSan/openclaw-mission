# git-merge-driver

Configure and manage git merge drivers for specific file types.

## Install

```bash
pip install git-merge-driver
```

## Usage

```bash
# Register a builtin driver with file pattern
git-merge-driver builtin union -p "*.lock"
git-merge-driver builtin ours -p "*.json"

# Add a custom driver
git-merge-driver add concat-files -s 'cat "$1" > "$MERGED" && echo "" >> "$MERGED" && cat "$2" >> "$MERGED"' -p "*.txt"

# List drivers
git-merge-driver list

# Install .gitattributes into current repo
git-merge-driver install
```

## Builtin drivers

| Name    | Behavior                           |
|---------|------------------------------------|
| union   | Keep both versions (append)        |
| ours    | Always keep our version            |
| theirs  | Always keep their version          |
| concat  | Concatenate both versions          |
| empty   | Replace with empty file            |

## Examples

```bash
# Never conflict on lock files
git-merge-driver builtin union -p "*.lock"

# Always keep local config
git-merge-driver builtin ours -p "config.local.*"

# Install to current repo
git-merge-driver install
```

## License

MIT
