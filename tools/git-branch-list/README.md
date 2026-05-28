# git-branch-list

List git branches with their last commit info, author, and relative date.

## Install

```bash
go install github.com/TataneSan/git-branch-list@latest
```

Or build from source:

```bash
go build -o git-branch-list .
```

## Usage

```
git-branch-list [options]
```

### Options

| Flag       | Description                     |
|------------|---------------------------------|
| `-a`, `-all`  | show remote branches too        |
| `--sort`      | sort by: `date`, `name`, `author` |

## Examples

### List local branches

```bash
git-branch-list
```

### List all branches including remotes

```bash
git-branch-list --all
```

### Sort by name

```bash
git-branch-list --sort name
```

## Output

Shows for each branch:
- Branch name (current branch marked with `*`)
- Short commit hash
- Last commit subject
- Author name
- Relative time (e.g. "2d ago", "3h ago")

## License

MIT
