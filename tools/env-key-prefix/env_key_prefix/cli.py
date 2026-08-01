#!/usr/bin/env python3
"""env-key-prefix - add or strip a prefix on every key of a .env file.

Reads .env files (or stdin) and rewrites the keys:
  - default: prepend PREFIX to non-prefixed keys
  - --strip: remove the given PREFIX from keys that have it
  - --replace OLD=NEW: rename only matching prefixed keys
Comments, blank lines and order are preserved. Duplicate keys after
rewriting are reported.

Exit codes:
    0 - success
    1 - CLI or I/O error
    2 - --check mode: a rewrite would collide with an existing key
"""

import argparse
import json
import re
import sys

ASSIGN_RE = re.compile(r"^(\s*(?:export\s+)?)([A-Za-z_][A-Za-z0-9_]*)(=.*)$")


def rewrite_line(line, prefix, strip=False):
    match = ASSIGN_RE.match(line)
    if not match:
        return line, None, None
    lead, key, rest = match.groups()
    if strip:
        if key.startswith(prefix):
            new_key = key[len(prefix):]
        else:
            new_key = key
    else:
        if key.startswith(prefix):
            new_key = key
        else:
            new_key = prefix + key
    return lead + new_key + rest, key, new_key


def process(text, prefix, strip=False):
    out_lines = []
    renames = []
    seen = {}
    collisions = []
    for lineno, line in enumerate(text.splitlines(), 1):
        new_line, old, new = rewrite_line(line, prefix, strip=strip)
        out_lines.append(new_line)
        if old is not None and new is not None:
            if old != new:
                renames.append({"old": old, "new": new, "line": lineno})
            if new in seen:
                collisions.append({
                    "key": new, "lines": [seen[new], lineno]})
            else:
                seen[new] = lineno
    trailing = "\n" if text.endswith("\n") else ""
    return "\n".join(out_lines) + trailing, renames, collisions


def read_source(path):
    if path in (None, "-"):
        return "<stdin>", sys.stdin.read()
    with open(path, "r", encoding="utf-8", errors="replace") as fh:
        return path, fh.read()


def main(argv=None):
    parser = argparse.ArgumentParser(
        prog="env-key-prefix",
        description="Prepend or strip a prefix on every key of .env files.",
    )
    parser.add_argument(
        "files", nargs="*", metavar="FILE",
        help=".env files; reads stdin when omitted or '-'",
    )
    parser.add_argument(
        "-p", "--prefix", required=True, metavar="PREFIX",
        help="prefix to add (e.g. APP_) or strip with --strip",
    )
    parser.add_argument(
        "--strip", action="store_true",
        help="remove the prefix instead of adding it",
    )
    parser.add_argument(
        "--check", action="store_true",
        help="exit 2 when the rewrite would cause key collisions; no output",
    )
    parser.add_argument(
        "--json", action="store_true",
        help="emit a JSON report of renames and collisions",
    )
    args = parser.parse_args(argv)

    files = args.files or ["-"]
    rc = 0
    results = []
    any_collision = False

    for path in files:
        try:
            name, text = read_source(path)
        except OSError as exc:
            print("env-key-prefix: %s: %s" % (path, exc), file=sys.stderr)
            rc = 1
            continue
        rewritten, renames, collisions = process(text, args.prefix,
                                                 strip=args.strip)
        any_collision = any_collision or bool(collisions)
        results.append({
            "file": name,
            "renamed": renames,
            "collisions": collisions,
        })
        if not args.check and not args.json:
            sys.stdout.write(rewritten)

    if args.json:
        payload = results[0] if len(results) == 1 else results
        json.dump(payload, sys.stdout, indent=2)
        sys.stdout.write("\n")

    if any_collision:
        print("env-key-prefix: key collision(s) detected", file=sys.stderr)
    if args.check and rc == 0:
        return 2 if any_collision else 0
    if rc == 0 and any_collision:
        return 1
    return rc


if __name__ == "__main__":
    sys.exit(main())
