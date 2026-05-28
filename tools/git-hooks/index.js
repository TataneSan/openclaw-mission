#!/usr/bin/env node
import { execSync } from 'node:child_process';
import { readFileSync, writeFileSync, mkdirSync, existsSync, readdirSync, unlinkSync } from 'node:fs';
import { join, dirname } from 'node:path';
import { fileURLToPath } from 'node:url';

const __dirname = dirname(fileURLToPath(import.meta.url));
const HOOKS_DIR = join(__dirname, 'hooks');

const HOOKS = {
  'pre-commit': {
    description: 'Run linting and tests before committing',
    template: `#!/usr/bin/env bash
# pre-commit hook
# Runs checks before each commit

set -e

echo "Running pre-commit checks..."

# Check for TODO/FIXME in staged files
STAGED_FILES=\$(git diff --cached --name-only --diff-filter=ACM)

if [ -n "\$STAGED_FILES" ]; then
  TODOS=\$(echo "\$STAGED_FILES" | xargs grep -l -i "TODO\\|FIXME" 2>/dev/null || true)
  if [ -n "\$TODOS" ]; then
    echo ""
    echo "Warning: Found TODO/FIXME comments in staged files:"
    echo "\$TODOS"
    echo ""
  fi
fi

# Check for trailing whitespace
WHITESPACE_FILES=\$(echo "\$STAGED_FILES" | xargs grep -l '[[:space:]]\$' 2>/dev/null || true)
if [ -n "\$WHITESPACE_FILES" ]; then
  echo "Warning: Trailing whitespace found in:"
  echo "\$WHITESPACE_FILES"
  echo ""
fi

# Check for large files (>5MB)
for file in \$STAGED_FILES; do
  if [ -f "\$file" ]; then
    SIZE=\$(wc -c < "\$file" | tr -d ' ')
    if [ "\$SIZE" -gt 5242880 ]; then
      echo "Warning: Large file detected (>5MB): \$file (\$(( SIZE / 1024 / 1024 ))MB)"
    fi
  fi
done

# Run linter if available
if command -v npm &> /dev/null && [ -f "package.json" ]; then
  if npm run lint &> /dev/null 2>&1; then
    npm run lint || {
      echo "Error: Linting failed. Fix issues before committing."
      exit 1
    }
  fi
fi

# Run tests if available (fast mode)
if command -v npm &> /dev/null && [ -f "package.json" ]; then
  if npm run test &> /dev/null 2>&1; then
    npm run test -- --passWithNoTests 2>/dev/null || {
      echo "Warning: Some tests failed. Consider fixing before committing."
    }
  fi
fi

echo "Pre-commit checks passed."
exit 0
`
  },

  'commit-msg': {
    description: 'Validate commit message format (Conventional Commits)',
    template: `#!/usr/bin/env bash
# commit-msg hook
# Validates commit messages follow Conventional Commits format

COMMIT_MSG_FILE=\$1
COMMIT_MSG=\$(cat "\$COMMIT_MSG_FILE")

# Skip merge commits and fixup/squash commits
if echo "\$COMMIT_MSG" | grep -qE "^(Merge |fixup! |squash! )"; then
  exit 0
fi

# Conventional Commits pattern: type(scope): description
if ! echo "\$COMMIT_MSG" | grep -qE "^(feat|fix|docs|style|refactor|perf|test|build|ci|chore|revert)(\\(.+\\))?!?: .+"; then
  echo ""
  echo "Error: Commit message does not follow Conventional Commits format."
  echo ""
  echo "Expected format: type(scope): description"
  echo ""
  echo "Types: feat, fix, docs, style, refactor, perf, test, build, ci, chore, revert"
  echo ""
  echo "Examples:"
  echo "  feat: add user authentication"
  echo "  fix(api): handle null response"
  echo "  docs: update README"
  echo "  refactor(core): simplify data processing"
  echo ""
  exit 1
fi

# Check subject line length (max 72 chars)
SUBJECT=\$(echo "\$COMMIT_MSG" | head -n 1)
LENGTH=\${#SUBJECT}
if [ "\$LENGTH" -gt 72 ]; then
  echo "Warning: Commit subject is \$LENGTH characters (max recommended: 72)"
fi

exit 0
`
  },

  'pre-push': {
    description: 'Run checks before pushing to remote',
    template: `#!/usr/bin/env bash
# pre-push hook
# Runs checks before pushing to remote

echo "Running pre-push checks..."

# Prevent pushing to main/master directly (optional)
REMOTE=\$1
URL=\$2

# Check if pushing to protected branches
for ref in \$3; do
  BRANCH=\$(echo "\$ref" | cut -d' ' -f2 | sed 's|refs/heads/||')

  if [ "\$BRANCH" = "main" ] || [ "\$BRANCH" = "master" ]; then
    echo "Warning: Pushing directly to '\$BRANCH' branch."
    echo "Consider using a feature branch and pull request instead."
  fi
done

# Run tests before pushing
if command -v npm &> /dev/null && [ -f "package.json" ]; then
  if npm run test &> /dev/null 2>&1; then
    npm test 2>/dev/null || {
      echo "Error: Tests failed. Fix issues before pushing."
      exit 1
    }
  fi
fi

# Check build if available
if command -v npm &> /dev/null && [ -f "package.json" ]; then
  if npm run build &> /dev/null 2>&1; then
    npm run build 2>/dev/null || {
      echo "Error: Build failed. Fix issues before pushing."
      exit 1
    }
  fi
fi

echo "Pre-push checks passed."
exit 0
`
  },

  'prepare-commit-msg': {
    description: 'Add commit metadata to commit message',
    template: `#!/usr/bin/env bash
# prepare-commit-msg hook
# Adds useful metadata to commit messages

COMMIT_MSG_FILE=\$1
COMMIT_TYPE=\$2
COMMIT_SHA=\$3

# Only modify regular commits (not merges, amends, etc.)
if [ "\$COMMIT_TYPE" = "message" ] || [ -z "\$COMMIT_TYPE" ]; then
  # Add branch name as reference if not a default branch
  BRANCH=\$(git symbolic-ref --short HEAD 2>/dev/null || echo "")

  if [ -n "\$BRANCH" ] && [ "\$BRANCH" != "main" ] && [ "\$BRANCH" != "master" ]; then
    # Check if branch reference already exists
    if ! grep -q "\$BRANCH" "\$COMMIT_MSG_FILE" 2>/dev/null; then
      echo "" >> "\$COMMIT_MSG_FILE"
      echo "Branch: \$BRANCH" >> "\$COMMIT_MSG_FILE"
    fi
  fi
fi

exit 0
`
  },

  'post-merge': {
    description: 'Run package install after merge',
    template: `#!/usr/bin/env bash
# post-merge hook
# Runs package manager install after merging

echo "Post-merge: checking for dependency updates..."

# Detect package manager and run install
if [ -f "package.json" ] && command -v npm &> /dev/null; then
  if [ -f "package-lock.json" ]; then
    DIFF=\$(git diff HEAD@{1} HEAD -- package-lock.json 2>/dev/null || true)
    if [ -n "\$DIFF" ]; then
      echo "package-lock.json changed, running npm install..."
      npm install
    fi
  fi
elif [ -f "requirements.txt" ] && command -v pip &> /dev/null; then
  DIFF=\$(git diff HEAD@{1} HEAD -- requirements.txt 2>/dev/null || true)
  if [ -n "\$DIFF" ]; then
    echo "requirements.txt changed, running pip install..."
    pip install -r requirements.txt
  fi
elif [ -f "go.mod" ] && command -v go &> /dev/null; then
  DIFF=\$(git diff HEAD@{1} HEAD -- go.mod 2>/dev/null || true)
  if [ -n "\$DIFF" ]; then
    echo "go.mod changed, running go mod tidy..."
    go mod tidy
  fi
fi

echo "Post-merge complete."
exit 0
`
  },

  'post-checkout': {
    description: 'Update dependencies when switching branches',
    template: `#!/usr/bin/env bash
# post-checkout hook
# Updates dependencies when switching branches

PREV_BRANCH=\$1
NEW_BRANCH=\$2
BRANCH_FLAG=\$3

# Only run on branch switches (not file checkouts)
if [ "\$BRANCH_FLAG" = "1" ]; then
  echo "Switched to branch: \$(git symbolic-ref --short HEAD 2>/dev/null || echo \$NEW_BRANCH)"

  # Run package install based on detected package manager
  if [ -f "package.json" ] && command -v npm &> /dev/null; then
    npm install --silent 2>/dev/null && echo "Dependencies installed."
  elif [ -f "requirements.txt" ] && command -v pip &> /dev/null; then
    pip install -r requirements.txt --quiet 2>/dev/null && echo "Dependencies installed."
  elif [ -f "go.mod" ] && command -v go &> /dev/null; then
    go mod tidy 2>/dev/null && echo "Dependencies tidied."
  fi
fi

exit 0
`
  }
};

function getGitDir() {
  try {
    return execSync('git rev-parse --git-dir', { encoding: 'utf-8' }).trim();
  } catch {
    return null;
  }
}

function printHooks() {
  console.log('');
  console.log('Available git hooks:');
  console.log('');
  for (const [name, hook] of Object.entries(HOOKS)) {
    const checked = isHookInstalled(name) ? '✓' : ' ';
    console.log(`  [${checked}] ${name.padEnd(24)} ${hook.description}`);
  }
  console.log('');
}

function isHookInstalled(name) {
  const gitDir = getGitDir();
  if (!gitDir) return false;
  return existsSync(join(gitDir, 'hooks', name));
}

function installHook(name) {
  const hook = HOOKS[name];
  if (!hook) {
    console.error(`Error: Unknown hook "${name}"`);
    process.exit(1);
  }

  const gitDir = getGitDir();
  if (!gitDir) {
    console.error('Error: Not a git repository. Run `git init` first.');
    process.exit(1);
  }

  const hooksPath = join(gitDir, 'hooks');
  if (!existsSync(hooksPath)) {
    mkdirSync(hooksPath, { recursive: true });
  }

  const hookPath = join(hooksPath, name);
  writeFileSync(hookPath, hook.template);

  // Make executable
  try {
    execSync(`chmod +x "${hookPath}"`);
  } catch {
    // chmod not available on Windows
  }

  console.log(`Installed: ${name}`);
}

function uninstallHook(name) {
  const gitDir = getGitDir();
  if (!gitDir) {
    console.error('Error: Not a git repository.');
    process.exit(1);
  }

  const hookPath = join(gitDir, 'hooks', name);
  if (existsSync(hookPath)) {
    unlinkSync(hookPath);
    console.log(`Removed: ${name}`);
  } else {
    console.log(`Not installed: ${name}`);
  }
}

function installAll() {
  for (const name of Object.keys(HOOKS)) {
    installHook(name);
  }
}

function uninstallAll() {
  for (const name of Object.keys(HOOKS)) {
    uninstallHook(name);
  }
}

function showUsage() {
  console.log(`
git-hooks - Install pre-configured git hooks

USAGE:
  git-hooks install [hook-name]   Install hook(s)
  git-hooks uninstall [hook-name] Remove hook(s)
  git-hooks list                  List available hooks
  git-hooks status                Show installation status

COMMANDS:
  install              Install all hooks
  install <name>       Install a specific hook
  uninstall            Remove all hooks
  uninstall <name>     Remove a specific hook
  list                 Show available hooks
  status               Show which hooks are installed

HOOKS:
  pre-commit           Run linting and tests before committing
  commit-msg           Validate commit message format (Conventional Commits)
  pre-push             Run checks before pushing to remote
  prepare-commit-msg   Add commit metadata to commit message
  post-merge           Run package install after merge
  post-checkout        Update dependencies when switching branches

EXAMPLES:
  git-hooks install          # Install all hooks
  git-hooks install pre-commit  # Install only pre-commit
  git-hooks uninstall        # Remove all hooks
  git-hooks list             # List available hooks
`);
}

const args = process.argv.slice(2);
const cmd = args[0];
const hookName = args[1];

switch (cmd) {
  case 'install':
    if (hookName) {
      if (!HOOKS[hookName]) {
        console.error(`Error: Unknown hook "${hookName}"`);
        console.error(`Available: ${Object.keys(HOOKS).join(', ')}`);
        process.exit(1);
      }
      installHook(hookName);
    } else {
      installAll();
    }
    break;

  case 'uninstall':
    if (hookName) {
      uninstallHook(hookName);
    } else {
      uninstallAll();
    }
    break;

  case 'list':
  case 'ls':
    printHooks();
    break;

  case 'status':
    printHooks();
    const installed = Object.keys(HOOKS).filter(isHookInstalled);
    console.log(`${installed.length}/${Object.keys(HOOKS).length} hooks installed.`);
    console.log('');
    break;

  case 'help':
  case '--help':
  case '-h':
    showUsage();
    break;

  default:
    if (!cmd) {
      showUsage();
    } else {
      console.error(`Unknown command: ${cmd}`);
      showUsage();
      process.exit(1);
    }
}
