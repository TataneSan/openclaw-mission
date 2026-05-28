# git-hooks

Install pre-configured git hooks for better development workflow. Includes pre-commit checks, commit message validation, pre-push tests, and more.

## Features

- **pre-commit**: Linting, tests, trailing whitespace detection, large file warnings, TODO/FIXME detection
- **commit-msg**: Conventional Commits format validation with helpful error messages
- **pre-push**: Test and build verification before pushing to remote
- **prepare-commit-msg**: Automatic branch name metadata in commit messages
- **post-merge**: Auto-run package install after merging (npm, pip, go)
- **post-checkout**: Auto-update dependencies when switching branches

## Installation

```bash
# Clone and link globally
git clone https://github.com/TataneSan/git-hooks.git
cd git-hooks
npm link

# Or install from source
npm install -g path/to/git-hooks
```

## Usage

```bash
# Install all hooks
git-hooks install

# Install a specific hook
git-hooks install pre-commit

# Remove all hooks
git-hooks uninstall

# Remove a specific hook
git-hooks uninstall pre-commit

# List available hooks
git-hooks list

# Check installation status
git-hooks status
```

## Available Hooks

| Hook | Description |
|------|-------------|
| `pre-commit` | Run linting and tests before committing |
| `commit-msg` | Validate commit message format (Conventional Commits) |
| `pre-push` | Run checks before pushing to remote |
| `prepare-commit-msg` | Add commit metadata to commit message |
| `post-merge` | Run package install after merge |
| `post-checkout` | Update dependencies when switching branches |

## Examples

### Install all hooks in your project

```bash
cd your-project
git-hooks install
```

### Only use commit message validation

```bash
git-hooks install commit-msg
```

### Check what's installed

```bash
git-hooks status
```

```
Available git hooks:

  [✓] pre-commit             Run linting and tests before committing
  [✓] commit-msg             Validate commit message format
  [ ] pre-push               Run checks before pushing to remote
  [ ] prepare-commit-msg     Add commit metadata to commit message
  [ ] post-merge             Run package install after merge
  [ ] post-checkout          Update dependencies when switching branches

2/6 hooks installed.
```

## Conventional Commits

The `commit-msg` hook enforces [Conventional Commits](https://www.conventionalcommits.org/) format:

```
type(scope): description
```

**Types**: `feat`, `fix`, `docs`, `style`, `refactor`, `perf`, `test`, `build`, `ci`, `chore`, `revert`

**Examples**:
```
feat: add user authentication
fix(api): handle null response
docs: update README
refactor(core): simplify data processing
test(utils): add edge case tests
```

## Package Manager Support

The `post-merge` and `post-checkout` hooks auto-detect and support:

- **npm** (`package.json`)
- **pip** (`requirements.txt`)
- **Go** (`go.mod`)

## License

MIT
