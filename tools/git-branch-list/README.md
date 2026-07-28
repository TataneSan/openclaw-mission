# git-branch-list

Lists git branches with last commit information.

Shows branch name, short hash, commit subject, and marks the current branch with `*`.

## Install

```bash
go install github.com/TataneSan/git-branch-list@latest
```

Or build from source:

```bash
git clone https://github.com/TataneSan/git-branch-list.git
cd git-branch-list
go build -o git-branch-list .
```

## Usage

```
git-branch-list [OPTIONS]
```

### Options

| Flag | Description |
|------|-------------|
| `-s`, `--simple` | Show only branch names |
| `-h`, `--help` | Show help |

## Examples

```bash
# Show all branches with commit info
git-branch-list

# Show only branch names
git-branch-list -s
```

### Output

```
* main                           a1b2c3d  feat: add user authentication
  develop                        e4f5g6h  fix: resolve login redirect
  feature/oauth                  i7j8k9l  feat: add OAuth2 support
```

## License

MIT
