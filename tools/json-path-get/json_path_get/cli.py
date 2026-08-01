"""Extract values from JSON using dot/bracket path expressions.

Path syntax:
    users.0.name        dot notation with numeric array index
    users[0].name       bracket index
    users[*].name       wildcard over array elements
    data.items[]        trailing [] expands an array

Exit codes:
    0  path found (or at least one match with wildcards)
    1  I/O or CLI error (invalid JSON, invalid path)
    2  path not found in the document
"""

import argparse
import json
import re
import sys


def _push(buf, tokens):
    if buf.isdigit() or (buf.startswith("-") and buf[1:].isdigit()):
        tokens.append(int(buf))
    else:
        tokens.append(buf)


def parse_path(path):
    """Parse 'a.b[0].c[*]' into a list of string keys, int indexes, and '*' wildcards."""
    tokens = []
    i = 0
    buf = ""
    while i < len(path):
        ch = path[i]
        if ch == ".":
            if buf:
                _push(buf, tokens)
                buf = ""
            i += 1
        elif ch == "[":
            if buf:
                _push(buf, tokens)
                buf = ""
            j = path.index("]", i)
            inner = path[i + 1:j].strip()
            if inner == "*":
                tokens.append("*")
            elif inner == "":
                tokens.append("*")
            elif inner.isdigit() or (inner.startswith("-") and inner[1:].isdigit()):
                tokens.append(int(inner))
            elif (inner.startswith("'") and inner.endswith("'")) or \
                 (inner.startswith('"') and inner.endswith('"')):
                tokens.append(inner[1:-1])
            else:
                tokens.append(inner)
            i = j + 1
        else:
            buf += ch
            i += 1
    if buf:
        _push(buf, tokens)
    return tokens


def get_values(doc, tokens):
    """Yield all values matching the token path."""
    if not tokens:
        yield doc
        return
    t, rest = tokens[0], tokens[1:]
    if t == "*":
        if isinstance(doc, dict):
            for k in sorted(doc.keys()):
                yield from get_values(doc[k], rest)
        elif isinstance(doc, list):
            for item in doc:
                yield from get_values(item, rest)
        return
    if isinstance(t, int):
        if isinstance(doc, list):
            idx = t if t >= 0 else len(doc) + t
            if 0 <= idx < len(doc):
                yield from get_values(doc[idx], rest)
        return
    # string key
    if isinstance(doc, dict) and t in doc:
        yield from get_values(doc[t], rest)


def parse_args(argv=None):
    p = argparse.ArgumentParser(
        prog="json-path-get",
        description="Extract values from JSON using path expressions.",
    )
    p.add_argument("file", nargs="?", default="-",
                   help="JSON input file (default: stdin, or '-').")
    p.add_argument("path", help="Path expression, e.g. users[0].name or items[*].id")
    p.add_argument("--raw", "-r", action="store_true",
                   help="Print strings without JSON quoting.")
    p.add_argument("--compact", "-c", action="store_true",
                   help="Compact JSON output (no indentation).")
    p.add_argument("--default", "-e", metavar="VALUE",
                   help="Fallback value printed when the path is missing (exit code stays 2).")
    p.add_argument("--first", action="store_true",
                   help="Only print the first match (with wildcards).")
    return p.parse_args(argv)


def main(argv=None):
    args = parse_args(argv)
    try:
        if args.file == "-":
            text = sys.stdin.read()
        else:
            with open(args.file, encoding="utf-8") as fh:
                text = fh.read()
    except OSError as e:
        print(f"error: {e}", file=sys.stderr)
        return 1

    try:
        doc = json.loads(text)
    except json.JSONDecodeError as e:
        print(f"error: invalid JSON: {e}", file=sys.stderr)
        return 1

    try:
        tokens = parse_path(args.path)
    except ValueError as e:
        print(f"error: invalid path: {e}", file=sys.stderr)
        return 1

    results = list(get_values(doc, tokens))
    if not results:
        if args.default is not None:
            print(args.default)
        else:
            print(f"error: path '{args.path}' not found", file=sys.stderr)
        return 2

    if args.first:
        results = results[:1]

    def emit(v):
        if args.raw and isinstance(v, str):
            print(v)
        else:
            print(json.dumps(v, indent=None if args.compact else 2, ensure_ascii=False))

    if len(results) == 1:
        emit(results[0])
    else:
        if args.compact or args.raw:
            for v in results:
                emit(v)
        else:
            emit(results)
    return 0


if __name__ == "__main__":
    sys.exit(main())
