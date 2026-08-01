#!/bin/bash
# Push all local tools to GitHub TataneSan

TOOLS_DIR="/root/openclaw/tools"
LOG_FILE="/root/openclaw/push_log.txt"
> "$LOG_FILE"

# Get list of local tools not on GitHub
gh repo list TataneSan --limit 1000 | awk '{print $1}' | sed 's/TataneSan\///' | sort > /tmp/github_tools.txt
ls -la "$TOOLS_DIR" | grep -E "^d" | grep -v "2026" | awk '{print $9}' | grep -v "^[.]*$" | sort > /tmp/local_tools.txt

to_push=$(comm -23 /tmp/local_tools.txt /tmp/github_tools.txt)
total=$(echo "$to_push" | wc -l)
count=0
success=0
fail=0

for tool in $to_push; do
    count=$((count + 1))
    tool_dir="$TOOLS_DIR/$tool"

    # Skip if not a git repo
    if [ ! -d "$tool_dir/.git" ]; then
        echo "[$count/$total] SKIP $tool (no git)" >> "$LOG_FILE"
        fail=$((fail + 1))
        continue
    fi

    # Skip if no commits
    commits=$(git -C "$tool_dir" log --oneline 2>/dev/null | wc -l)
    if [ "$commits" -eq 0 ]; then
        echo "[$count/$total] SKIP $tool (no commits)" >> "$LOG_FILE"
        fail=$((fail + 1))
        continue
    fi

    # Check remote
    remote_url=$(git -C "$tool_dir" remote get-url origin 2>/dev/null)
    if [ -z "$remote_url" ]; then
        # Set remote
        git -C "$tool_dir" remote add origin "git@github.com:TataneSan/$tool.git"
    fi

    # Create repo on GitHub and push
    if gh repo create "TataneSan/$tool" --private --source "$tool_dir" --remote origin --push 2>>"$LOG_FILE"; then
        echo "[$count/$total] OK   $tool" >> "$LOG_FILE"
        success=$((success + 1))
    else
        # Try push directly if repo already exists
        if git -C "$tool_dir" push -u origin main 2>>"$LOG_FILE"; then
            echo "[$count/$total] OK   $tool (push only)" >> "$LOG_FILE"
            success=$((success + 1))
        else
            echo "[$count/$total] FAIL $tool" >> "$LOG_FILE"
            fail=$((fail + 1))
        fi
    fi

    # Progress every 50 tools
    if [ $((count % 50)) -eq 0 ]; then
        echo "Progress: $count/$total (OK: $success, FAIL: $fail)"
    fi
done

echo "Done: $count/$total (OK: $success, FAIL: $fail)" >> "$LOG_FILE"
echo "Done: $count/$total (OK: $success, FAIL: $fail)"
