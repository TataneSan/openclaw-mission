# git-branch-rename

Batch rename git branches with replace, prefix, suffix, and strip operations.

Protects `main`, `master`, and the currently checked-out branch from accidental renames.

## Install

```bash
go install github.com/TataneSan/git-branch-rename@latest
```

Or build from source:

```bash
git clone https://github.com/TataneSan/git-branch-rename.git
cd git-branch-rename
go build -o git-branch-rename .
```

## Usage

```bash
git-branch-rename <action> [options] [branches...]
```

### Actions

| Action | Description |
|--------|-------------|
| `replace <old> <new>` | Replace text in branch names |
| `prefix <text>` | Add a prefix to branch names |
| `suffix <text>` | Add a suffix to branch names |
| `strip-prefix <text>` | Remove a prefix from branch names |
| `strip-suffix <text>` | Remove a suffix from branch names |
| `list` | List all local branches |
| `dry-run replace <old> <new>` | Preview replacements without applying |

### Examples

Replace `feature/` with `feat/` in all branches:

```bash
./git-branch-rename replace feature feat
```

Replace only in specific branches:

```bash
./git-branch-rename replace old new feature/login feature/signup
```

Add a `dev/` prefix to all branches:

```bash
./git-branch-rename prefix dev/
```

Add a `-wip` suffix:

```bash
./git-branch-rename suffix -wip
```

Remove a `dev/` prefix:

```bash
./git-branch-rename strip-prefix dev/
```

Preview changes without applying:

```bash
./git-branch-rename dry-run replace feature feat
```

List all branches:

```bash
./git-branch-rename list
```

## Safety

- `main` and `master` branches are never renamed
- The currently checked-out branch is never renamed
- Protected branches are skipped with a warning

## License

MIT
