"""Split a CSV into multiple chunk files.

Modes:
    --chunks N     even split into N files
    --size K       files of at most K data rows
    --by COLUMN    one file per distinct value of COLUMN (safe names)

The header is repeated in every output file.

Exit codes:
    0 - success
    1 - I/O or CLI usage error
    2 - unknown --by column
"""

import argparse
import csv
import json
import os
import re
import sys


def safe_name(value):
    s = re.sub(r"[^A-Za-z0-9._-]+", "_", value.strip()) or "empty"
    return s[:64]


def write_chunk(path, header, rows, delimiter):
    with open(path, "w", newline="", encoding="utf-8") as fh:
        w = csv.writer(fh, delimiter=delimiter, lineterminator="\n")
        w.writerow(header)
        w.writerows(rows)


def main(argv=None):
    p = argparse.ArgumentParser(
        prog="csv-split-rows",
        description="Split a CSV into chunk files (count, size or key column).",
    )
    p.add_argument("file", nargs="?", default="-", help="input CSV (default: stdin)")
    g = p.add_mutually_exclusive_group(required=True)
    g.add_argument("--chunks", type=int, help="number of output files")
    g.add_argument("--size", type=int, help="max data rows per file")
    g.add_argument("--by", metavar="COLUMN", help="split on distinct column values")
    p.add_argument("-o", "--out-dir", default=".", help="output directory (default: .)")
    p.add_argument("--prefix", default="part", help="output filename prefix")
    p.add_argument("-d", "--delimiter", default=",", help="field delimiter")
    p.add_argument("--json", action="store_true", help="emit a JSON manifest")
    args = p.parse_args(argv)

    if args.file == "-":
        fh = sys.stdin
    else:
        try:
            fh = open(args.file, newline="", encoding="utf-8")
        except OSError as exc:
            print(f"csv-split-rows: {args.file}: {exc}", file=sys.stderr)
            return 1
    try:
        rows = list(csv.reader(fh, delimiter=args.delimiter))
    finally:
        if fh is not sys.stdin:
            fh.close()
    if not rows:
        print("csv-split-rows: empty input", file=sys.stderr)
        return 1
    header, data = rows[0], rows[1:]

    chunks = []
    if args.by:
        if args.by not in header:
            print(f"csv-split-rows: unknown column: {args.by!r}", file=sys.stderr)
            return 2
        idx = header.index(args.by)
        groups = {}
        for row in data:
            key = row[idx] if idx < len(row) else ""
            groups.setdefault(key, []).append(row)
        for key in sorted(groups):
            chunks.append((f"{args.prefix}-{safe_name(key)}.csv", groups[key]))
    elif args.chunks:
        if args.chunks < 1:
            print("csv-split-rows: --chunks must be >= 1", file=sys.stderr)
            return 1
        n = min(args.chunks, max(len(data), 1))
        base, extra = divmod(len(data), n)
        pos = 0
        for i in range(n):
            k = base + (1 if i < extra else 0)
            chunks.append((f"{args.prefix}-{i+1:03d}.csv", data[pos:pos + k]))
            pos += k
    else:
        if args.size < 1:
            print("csv-split-rows: --size must be >= 1", file=sys.stderr)
            return 1
        for i in range(0, len(data), args.size):
            chunks.append((f"{args.prefix}-{i // args.size + 1:03d}.csv",
                           data[i:i + args.size]))

    try:
        os.makedirs(args.out_dir, exist_ok=True)
    except OSError as exc:
        print(f"csv-split-rows: {args.out_dir}: {exc}", file=sys.stderr)
        return 1

    manifest = []
    for name, chunk_rows in chunks:
        path = os.path.join(args.out_dir, name)
        try:
            write_chunk(path, header, chunk_rows, args.delimiter)
        except OSError as exc:
            print(f"csv-split-rows: {path}: {exc}", file=sys.stderr)
            return 1
        manifest.append({"file": path, "rows": len(chunk_rows)})

    if args.json:
        print(json.dumps({"chunks": manifest}, indent=2))
    else:
        for m in manifest:
            print(f"{m['file']}: {m['rows']} rows")
    return 0


if __name__ == "__main__":
    sys.exit(main())
