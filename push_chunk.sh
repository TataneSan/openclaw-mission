#!/bin/bash
# Usage: push_chunk.sh <chunk_number 0-4>
N="$1"
CHUNK="/tmp/chunk_0${N}"
TOOLS_DIR="/root/openclaw/tools"
LOG="/root/openclaw/push_chunk${N}.log"
: > "$LOG"

while read -r tool; do
    [ -n "$tool" ] || continue
    dir="$TOOLS_DIR/$tool"
    if [ ! -d "$dir/.git" ]; then echo "SKIP $tool (no git)" >> "$LOG"; continue; fi
    if [ -z "$(git -C "$dir" log -1 --oneline 2>/dev/null)" ]; then
        echo "SKIP $tool (no commits)" >> "$LOG"; continue
    fi
    if ! git -C "$dir" remote get-url origin >/dev/null 2>&1; then
        git -C "$dir" remote add origin "git@github.com:TataneSan/$tool.git" 2>/dev/null
    fi
    if git -C "$dir" push -u origin main >/dev/null 2>&1 || \
       git -C "$dir" push -u origin master:main >/dev/null 2>&1; then
        echo "OK   $tool" >> "$LOG"
        continue
    fi
    if gh repo create "TataneSan/$tool" --public >/dev/null 2>&1; then
        if git -C "$dir" push -u origin main >/dev/null 2>&1 || \
           git -C "$dir" push -u origin master:main >/dev/null 2>&1; then
            echo "OK   $tool created" >> "$LOG"
        else
            echo "FAIL $tool push-after-create" >> "$LOG"
        fi
    else
        echo "FAIL $tool" >> "$LOG"
    fi
done < "$CHUNK"
echo "ALLDONE" >> "$LOG"
