# git-tag-list

List and analyze git tags: versions, dates, pre-release flags, and version gaps.

## Features

- Lists tags with date and annotated/lightweight marker
- Semver-aware sorting (`v1.9.0` < `v1.10.0`)
- Pre-release detection (`rc`, `beta`, `alpha`...)
- Version gap detection (skipped minor/major numbers)
- Regex filtering, `--semver`-only mode, limit / reverse
- `--json` machine-readable output
- Exit code 2 when no tags match (CI-friendly)

## Install

```bash
pip install .
# or
pip install git+https://github.com/TataneSan/git-tag-list.git
```

## Usage

```bash
# inside a repo
git-tag-list

# tags of another repo, sorted by version (default)
git-tag-list /path/to/repo

# latest 5 releases, only strict semver
git-tag-list --semver --limit 5

# newest first with gap detection
git-tag-list --reverse --gaps

# JSON for scripting
git-tag-list --json | jq '.tags[-1].name'

# filter by pattern
git-tag-list --grep '^v2\.'
```

Sample output:

```
v1.0.0                   2026-01-12  [annotated]
v1.1.0-rc.1              2026-02-03  [annotated, pre:rc.1]
v1.1.0                   2026-02-10  [annotated]
v2.0.0                   2026-05-01  [annotated]
```

## Exit codes

| Code | Meaning |
|------|---------|
| 0 | Success |
| 1 | I/O or CLI error (not a git repo, git missing) |
| 2 | No tags found |

## License

MIT
