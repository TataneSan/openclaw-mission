# git-hooks

Install pre-configured Git hooks into any repository with a single command.

## Features

- **pre-commit**: Run linters, formatters, and checks before each commit
  - Trailing whitespace detection
  - Large file detection (>10MB)
  - Custom commands from config
- **commit-msg**: Enforce [Conventional Commits](https://www.conventionalcommits.org/) format
- **pre-push**: Run tests before pushing to remote
- **pre-rebase**: Warn before rebasing shared branches
- **post-merge**: Run commands after a successful merge (e.g., `npm install`)

All hooks are configured via a single `.git-hooks.json` file in your repo root.

## Install

```bash
go install github.com/TataneSan/git-hooks@latest
```

Or download a release from [GitHub Releases](https://github.com/TataneSan/git-hooks/releases).

## Usage

```bash
# Install all hooks into the current repo
git-hooks install

# Install specific hooks
git-hooks install -hooks pre-commit,commit-msg

# Install into a different directory
git-hooks install -dir ../my-project

# List available hooks
git-hooks list

# Show a hook's content
git-hooks show pre-commit

# Remove hooks
git-hooks remove
git-hooks remove -hooks pre-commit

# Create default config
git-hooks init
```

## Configuration

Run `git-hooks init` to create a default `.git-hooks.json`:

```json
{
  "pre-commit": {
    "commands": ["npm run lint", "npm run format"],
    "trailing-whitespace": true,
    "large-files": true
  },
  "commit-msg": {
    "conventional": true
  },
  "pre-push": {
    "enabled": true,
    "commands": ["npm test"]
  },
  "post-merge": {
    "commands": ["npm install"]
  }
}
```

### pre-commit

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `commands` | `string[]` | `[]` | Shell commands to run before commit |
| `trailing-whitespace` | `bool` | `true` | Check for trailing whitespace |
| `large-files` | `bool` | `true` | Check for files >10MB |

### commit-msg

Enforces the format: `<type>(<scope>): <description>`

Valid types: `feat`, `fix`, `docs`, `style`, `refactor`, `perf`, `test`, `build`, `ci`, `chore`, `revert`

### pre-push

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `enabled` | `bool` | `true` | Enable/disable the hook |
| `commands` | `string[]` | `[]` | Shell commands to run before push |

### post-merge

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `commands` | `string[]` | `[]` | Shell commands to run after merge |

## Examples

### JavaScript project

```json
{
  "pre-commit": {
    "commands": ["npx eslint .", "npx prettier --check ."],
    "trailing-whitespace": true,
    "large-files": true
  },
  "commit-msg": { "conventional": true },
  "pre-push": {
    "enabled": true,
    "commands": ["npm test"]
  },
  "post-merge": {
    "commands": ["npm install"]
  }
}
```

### Go project

```json
{
  "pre-commit": {
    "commands": ["go fmt ./...", "go vet ./..."],
    "trailing-whitespace": true,
    "large-files": true
  },
  "commit-msg": { "conventional": true },
  "pre-push": {
    "enabled": true,
    "commands": ["go test ./..."]
  },
  "post-merge": { "commands": [] }
}
```

### Python project

```json
{
  "pre-commit": {
    "commands": ["flake8", "black --check ."],
    "trailing-whitespace": true,
    "large-files": true
  },
  "commit-msg": { "conventional": true },
  "pre-push": {
    "enabled": true,
    "commands": ["pytest"]
  },
  "post-merge": { "commands": ["pip install -r requirements.txt"] }
}
```

## License

MIT
