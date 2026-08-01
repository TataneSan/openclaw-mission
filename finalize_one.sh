#!/bin/bash
# finalize_one.sh <tool_name> — init, commit and push one tool dir.
set -u
tool="$1"
dir="/root/openclaw/tools/$tool"
[ -d "$dir" ] || { echo "SKIP $tool (no dir)"; exit 0; }

cd "$dir" || exit 1

# Minimal README if missing
if [ ! -f README.md ]; then
    printf '# %s\n\nCommand-line tool.\n\n## Usage\n\n```sh\n%s --help\n```\n\n## License\n\nMIT\n' "$tool" "$tool" > README.md
fi

# MIT license if missing
if [ ! -f LICENSE ] && [ ! -f LICENSE.md ]; then
    cat > LICENSE <<'EOF'
MIT License

Copyright (c) 2026 Tatane

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
EOF
fi

# .gitignore
if [ -f .gitignore ]; then
    # Un-ignore the tool's own compiled binary if a previous run ignored it
    if grep -qx "$tool" .gitignore && [ -f "$tool" ] && file "$tool" 2>/dev/null | grep -q "ELF"; then
        echo "!$tool" >> .gitignore
    fi
else
    cat > .gitignore <<'EOF'
# Binaries
*.exe
*.dll
*.so
*.dylib
bin/
dist/
*.test
*.out

# Python
__pycache__/
*.pyc
.venv/
venv/

# Node
node_modules/

# OS
.DS_Store
EOF
fi

# Remove stray compiled binaries named like the tool (keep source only)
if [ -f "$tool" ] && file "$tool" 2>/dev/null | grep -q "ELF"; then
    git_check=1
    rm -f "$tool"
fi

if [ ! -d .git ]; then
    git init -q -b main
fi

git add -A
if git diff --cached --quiet 2>/dev/null && [ -z "$(git status --porcelain)" ]; then
    :
fi
if ! git rev-parse HEAD >/dev/null 2>&1; then
    git commit -q -m "feat: initial commit" || { echo "FAIL $tool (commit)"; exit 0; }
else
    if ! git diff --cached --quiet 2>/dev/null; then
        git commit -q -m "chore: housekeeping" || true
    fi
fi

if ! git remote get-url origin >/dev/null 2>&1; then
    git remote add origin "git@github.com:TataneSan/$tool.git"
fi

if git push -q -u origin main >/dev/null 2>&1 || git push -q -u origin master:main >/dev/null 2>&1; then
    echo "OK   $tool"
    exit 0
fi
if gh repo create "TataneSan/$tool" --public >/dev/null 2>&1; then
    if git push -q -u origin main >/dev/null 2>&1 || git push -q -u origin master:main >/dev/null 2>&1; then
        echo "OK   $tool created"
        exit 0
    fi
fi
echo "FAIL $tool"
