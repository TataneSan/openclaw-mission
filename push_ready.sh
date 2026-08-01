#!/bin/bash
# Push every local tool (git repo with commits) that is missing on GitHub.
TOOLS_DIR="/root/openclaw/tools"
LOG="/root/openclaw/push_ready.log"
: > "$LOG"

for tool in "$@"; do
    dir="$TOOLS_DIR/$tool"
    [ -d "$dir/.git" ] || { echo "SKIP $tool (no git)" >> "$LOG"; continue; }
    commits=$(git -C "$dir" log --oneline 2>/dev/null | wc -l)
    [ "$commits" -gt 0 ] || { echo "SKIP $tool (no commits)" >> "$LOG"; continue; }

    git -C "$dir" remote get-url origin >/dev/null 2>&1 || git -C "$dir" remote add origin "git@github.com:TataneSan/$tool.git"

    if git -C "$dir" push -u origin main >/dev/null 2>&1 || \
       git -C "$dir" push -u origin master:main >/dev/null 2>&1; then
        echo "OK   $tool" >> "$LOG"
        continue
    fi

    if gh repo create "TataneSan/$tool" --public >/dev/null 2>&1; then
        if git -C "$dir" push -u origin main >/dev/null 2>&1 || \
           git -C "$dir" push -u origin master:main >/dev/null 2>&1; then
            echo "OK   $tool (created)" >> "$LOG"
        else
            echo "FAIL $tool (created but push failed)" >> "$LOG"
        fi
    else
        echo "FAIL $tool" >> "$LOG"
    fi
done
echo "DONE" >> "$LOG"
