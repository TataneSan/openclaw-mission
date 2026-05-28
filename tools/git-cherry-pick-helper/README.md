# git-cherry-pick-helper

Interactive cherry-pick assistant for Git repositories.

## Features

- Lists commits in a range with numbered output
- Select commits by number, "all", or "skip"
- Cherry-picks in order with per-commit status
- Reports failures with error details

## Install

```bash
go install github.com/TataneSan/git-cherry-pick-helper@latest
```

## Usage

```bash
git-cherry-pick-helper <commit-range>
```

### Examples

```bash
# Cherry-pick from develop to main
git-cherry-pick-helper main..develop

# Cherry-pick between specific commits
git-cherry-pick-helper abc123..def456
```

### Interactive flow

```
$ git-cherry-pick-helper main..develop

Commits in main..develop:

  [1] a1b2c3d Fix login bug
  [2] e4f5g6h Add user profile page
  [3] i7j8k9l Update dependencies

Enter commit numbers to cherry-pick (e.g. 1,3,5 or 'all' or 'skip'): 1,3
Cherry-picking a1b2c3d Fix login bug... OK
Cherry-picking i7j8k9l Update dependencies... OK

All 2 commit(s) cherry-picked successfully.
```

## License

MIT
