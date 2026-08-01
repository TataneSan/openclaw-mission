#!/usr/bin/env bash
# Internal helper: init + push a tool repo. Usage: push_tool.sh <tool-name>
set -e
name="$1"
dir="/root/openclaw/tools/$name"
cd "$dir"
gh repo view "TataneSan/$name" >/dev/null 2>&1 || gh repo create "TataneSan/$name" --public --confirm >/dev/null 2>&1 || gh repo create "TataneSan/$name" --public >/dev/null
git init -q
git remote get-url origin >/dev/null 2>&1 || git remote add origin "git@github.com:TataneSan/$name.git"
git add -A
git -c user.name="TataneSan" -c user.email="tatsanesan@users.noreply.github.com" commit -q -m "feat: initial commit"
git branch -M main
git push -q -u origin main
echo "pushed: $name"
