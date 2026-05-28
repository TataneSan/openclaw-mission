# git-merge-base

Find the last common ancestor between two git branches. Displays the merge base commit with optional details and commit counts.

## Install

```bash
go install github.com/TataneSan/git-merge-base@latest
```

Or build from source:

```bash
git clone https://github.com/TataneSan/git-merge-base.git
cd git-merge-base
go build -o git-merge-base .
```

## Usage

```
git-merge-base [-v] [-c] <branch1> <branch2>
```

### Flags

| Flag | Description |
|------|-------------|
| `-v` | Show verbose commit details (hash, author, date, message) |
| `-c` | Show commit count from merge-base to each branch |

### Examples

Find merge base between two branches:

```bash
./git-merge-base main feature-login
```

With verbose details:

```bash
./git-merge-base -v main develop
```

With commit counts:

```bash
./git-merge-base -c main feature-login
```

Verbose and counts together:

```bash
./git-merge-base -v -c main develop
```

### Output

```
Merge base: a3f7b2c1

Hash:    a3f7b2c1
Author:  John Doe
Date:    2024-03-15
Message: refactor: restructure auth module

main: 12 commit(s) from merge-base
feature-login: 5 commit(s) from merge-base
```

## Features

- Finds the merge base (lowest common ancestor) between any two branches
- Verbose mode shows commit hash, author, date, and message
- Commit count mode shows how far each branch is from the merge base
- Works with branches, tags, and commit hashes

## License

MIT
