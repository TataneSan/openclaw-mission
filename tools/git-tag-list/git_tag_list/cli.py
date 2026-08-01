"""List and analyze git tags: versions, dates, gaps, and pre-release detection.

Exit codes:
    0  success
    1  I/O or CLI error (not a git repo, git missing)
    2  no tags found
"""

import argparse
import json
import re
import subprocess
import sys


SEMVER_RE = re.compile(r"^v?(\d+)\.(\d+)\.(\d+)(?:-([0-9A-Za-z.\-]+))?(?:\+[0-9A-Za-z.\-]+)?$")


def parse_args(argv=None):
    p = argparse.ArgumentParser(
        prog="git-tag-list",
        description="List and analyze git tags in a repository.",
    )
    p.add_argument("repo", nargs="?", default=".",
                   help="Path to the git repository (default: current directory).")
    p.add_argument("--semver", action="store_true",
                   help="Only include strict semver tags (v1.2.3 or 1.2.3).")
    p.add_argument("--sort", choices=("version", "date", "name"), default="version",
                   help="Sort order (default: version).")
    p.add_argument("--reverse", "-r", action="store_true",
                   help="Reverse the sort order.")
    p.add_argument("--limit", "-n", type=int, default=0,
                   help="Only show the N last tags.")
    p.add_argument("--grep", "-g", metavar="PATTERN",
                   help="Filter tag names with a regex.")
    p.add_argument("--gaps", action="store_true",
                   help="Show version gaps (jumps in minor/major numbers).")
    p.add_argument("--json", action="store_true",
                   help="Machine-readable JSON output.")
    return p.parse_args(argv)


def run_git(repo, *git_args):
    out = subprocess.run(
        ["git", "-C", repo] + list(git_args),
        capture_output=True, text=True,
    )
    if out.returncode != 0:
        raise SystemExit(f"error: git: {out.stderr.strip() or out.stdout.strip()}")
    return out.stdout


def collect(repo):
    fmt = "%(refname:short)%09%(creatordate:iso-strict)%09%(objecttype)"
    out = run_git(repo, "tag", "--list", f"--format={fmt}")
    tags = []
    for line in out.splitlines():
        if not line.strip():
            continue
        parts = line.split("\t")
        name = parts[0]
        date = parts[1] if len(parts) > 1 else ""
        objtype = parts[2] if len(parts) > 2 else ""
        m = SEMVER_RE.match(name)
        tags.append({
            "name": name,
            "date": date,
            "annotated": objtype == "tag",
            "semver": [int(m.group(1)), int(m.group(2)), int(m.group(3))] if m else None,
            "prerelease": m.group(4) if m else None,
        })
    return tags


def version_key(tag):
    if tag["semver"]:
        major, minor, patch = tag["semver"]
        pre = tag["prerelease"]
        # releases sort after their prereleases
        return (major, minor, patch, 0 if pre is None else -1, pre or "")
    return (-1, -1, -1, 0, tag["name"])


def detect_gaps(tags):
    semtags = sorted((t for t in tags if t["semver"] and t["prerelease"] is None),
                     key=version_key)
    gaps = []
    for prev, cur in zip(semtags, semtags[1:]):
        p, c = prev["semver"], cur["semver"]
        if c[0] > p[0] + 1:
            gaps.append(f"major jump: {prev['name']} -> {cur['name']}")
        elif c[0] == p[0] and (c[1] > p[1] + 1 or (c[1] > p[1] and p[2] != 0)):
            gaps.append(f"minor jump: {prev['name']} -> {cur['name']}")
    return gaps


def main(argv=None):
    args = parse_args(argv)
    try:
        tags = collect(args.repo)
    except SystemExit as e:
        print(e, file=sys.stderr)
        return 1
    except FileNotFoundError:
        print("error: git binary not found", file=sys.stderr)
        return 1

    if args.semver:
        tags = [t for t in tags if t["semver"]]
    if args.grep:
        rx = re.compile(args.grep)
        tags = [t for t in tags if rx.search(t["name"])]

    if not tags:
        print("error: no tags found", file=sys.stderr)
        return 2

    if args.sort == "version":
        tags.sort(key=version_key)
    elif args.sort == "date":
        tags.sort(key=lambda t: t["date"])
    else:
        tags.sort(key=lambda t: t["name"])
    if args.reverse:
        tags.reverse()
    if args.limit:
        tags = tags[-args.limit:]

    if args.json:
        payload = {"tags": tags, "count": len(tags)}
        if args.gaps:
            payload["gaps"] = detect_gaps(tags)
        print(json.dumps(payload, indent=2))
    else:
        for t in tags:
            extra = []
            if t["annotated"]:
                extra.append("annotated")
            if t["prerelease"]:
                extra.append(f"pre:{t['prerelease']}")
            suffix = f"  [{', '.join(extra)}]" if extra else ""
            date = t["date"][:10] if t["date"] else "----------"
            print(f"{t['name']:24} {date}{suffix}")
        if args.gaps:
            for g in detect_gaps(tags):
                print(f"gap: {g}")
    return 0


if __name__ == "__main__":
    sys.exit(main())
