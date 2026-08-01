"""csv-nth-row: extract rows at given positions from a CSV file.

Positional specs are 1-based over data rows (header excluded when --header
is used). Supports single indices ('3'), inclusive ranges ('2:5'), negative
indices ('-1' = last row) and strided selection ('1::2' = every other row
starting at row 1).

Exit codes:
    0  success
    1  I/O or CLI error
"""

import argparse
import csv
import io
import json
import sys


def parse_spec(spec):
    """Parse one spec and return ('one', n) | ('range', a, b) | ('every', start, step)."""
    spec = spec.strip()
    if "::" in spec:
        start_s, step_s = spec.split("::", 1)
        try:
            start = int(start_s) if start_s else None
            step = int(step_s)
        except ValueError:
            raise ValueError("expected '<start>::<step>' with integers")
        if step <= 0:
            raise ValueError("step must be >= 1")
        if start is not None and start <= 0:
            raise ValueError("start must be >= 1 (rows are 1-based)")
        return ("every", start, step)
    if ":" in spec:
        a_s, b_s = spec.split(":", 1)
        try:
            return ("range", int(a_s), int(b_s))
        except ValueError:
            raise ValueError("expected '<a>:<b>' with integers")
    try:
        n = int(spec)
    except ValueError:
        raise ValueError("expected an integer, range or stride")
    if n == 0:
        raise ValueError("row indices are 1-based; 0 is invalid")
    return ("one", n)


def read_rows(path, delimiter):
    if path in (None, "-"):
        data = sys.stdin.read()
    else:
        try:
            with open(path, "r", newline="", encoding="utf-8",
                      errors="replace") as f:
                data = f.read()
        except OSError as exc:
            print(f"error: cannot read {path}: {exc}", file=sys.stderr)
            sys.exit(1)
    return list(csv.reader(io.StringIO(data), delimiter=delimiter))


def select(rows, specs):
    n = len(rows)
    picked = {}

    def resolve(idx):
        # 1-based -> 0-based, negatives from the end
        if idx < 0:
            return n + idx
        return idx - 1

    for spec in specs:
        try:
            parsed = parse_spec(spec)
        except ValueError as exc:
            print(f"error: invalid row spec {spec!r}: {exc}", file=sys.stderr)
            sys.exit(1)
        kind = parsed[0]
        if kind == "one":
            i = resolve(parsed[1])
            if 0 <= i < n:
                picked[i] = rows[i]
        elif kind == "range":
            a, b = resolve(parsed[1]), resolve(parsed[2])
            lo, hi = min(a, b), max(a, b)
            for i in range(max(0, lo), min(n - 1, hi) + 1):
                picked[i] = rows[i]
        else:  # every
            _, start, step = parsed
            start_i = 0 if start is None else start - 1
            for i in range(start_i, n, step):
                picked[i] = rows[i]
    indices = sorted(picked)
    return [picked[i] for i in indices], indices


def main():
    p = argparse.ArgumentParser(
        prog="csv-nth-row",
        description="Extract rows at given positions from a CSV file.",
        epilog="specs: '3' row 3 | '2:5' rows 2 to 5 | '-1' last row | "
               "'1::2' every 2nd row from row 1. Rows are 1-based over "
               "data rows (header excluded with --header).",
    )
    p.add_argument("specs", nargs="+",
                   help="row spec(s), e.g. 1 3 5:8 -1 1::10")
    p.add_argument("file", nargs="?", default="-",
                   help="CSV file (default: stdin)")
    p.add_argument("--header", action="store_true",
                   help="first line is a header: keep it in output, number "
                        "data rows from 1")
    p.add_argument("-d", "--delimiter", default=",",
                   help="field delimiter (default: ,)")
    p.add_argument("--with-index", action="store_true",
                   help="prefix each output row with its original position")
    p.add_argument("--json", action="store_true",
                   help="emit rows as a JSON document")
    args = p.parse_args()

    rows = read_rows(args.file, args.delimiter)
    header = None
    if args.header and rows:
        header, rows = rows[0], rows[1:]

    selected, indices = select(rows, args.specs)

    if args.json:
        out = {"header": header, "data_row_count": len(rows),
               "selected": [i + 1 for i in indices], "rows": selected}
        print(json.dumps(out, ensure_ascii=False, indent=2))
        return 0

    w = csv.writer(sys.stdout, lineterminator="\n")
    if header is not None:
        w.writerow((["row"] + header) if args.with_index else header)
    for i, row in zip(indices, selected):
        w.writerow(([i + 1] + row) if args.with_index else row)
    return 0


if __name__ == "__main__":
    sys.exit(main())
