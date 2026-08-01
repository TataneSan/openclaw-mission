"""csv-column-stats: numeric statistics per CSV column.

Reads a CSV document (file or stdin) and computes numeric statistics
(count, empty values, non-numeric cells, min, max, mean, median, stdev)
for selected columns. Non-numeric cells are skipped and reported, never
counted as zero. Also usable as a CI gate with --check.

Exit codes:
  0  success
  1  I/O or argument error
  2  check mode: at least one assertion failed
"""

import argparse
import csv
import io
import json
import re
import statistics
import sys

from . import __version__

STATS = ("count", "min", "max", "mean", "median", "stdev")

NUMBER_RE = re.compile(r"^[+-]?(?:\d+(?:\.\d*)?|\.\d+)(?:[eE][+-]?\d+)?$")

CHECK_RE = re.compile(
    r"^([^:=!<>]+?)\s*:\s*(count|min|max|mean|median|stdev)\s*(>=|<=|==|!=|>|<)\s*(.+?)\s*$"
)


def parse_delimiter(value):
    aliases = {"\\t": "\t", "tab": "\t", "comma": ",", "semicolon": ";", "pipe": "|"}
    if value in aliases:
        return aliases[value]
    if len(value) != 1:
        raise ValueError(
            "delimiter must be a single character (aliases: tab, comma, semicolon, pipe)"
        )
    return value


def to_number(text):
    text = text.strip()
    if not NUMBER_RE.match(text):
        return None
    try:
        value = float(text)
    except ValueError:
        return None
    if value != value or value in (float("inf"), float("-inf")):
        return None
    return value


def resolve_columns(header, selectors):
    """Selectors are header names or 1-based indices. Returns index list."""
    picked = []
    n = len(header)
    for sel in selectors:
        if sel.lstrip("-").isdigit():
            idx = int(sel)
            if idx < 1 or idx > n:
                raise ValueError(
                    f"column index out of range: {sel} (header has {n} columns)"
                )
            pos = idx - 1
        else:
            matches = [i for i, name in enumerate(header) if name == sel]
            if not matches:
                raise ValueError(f"unknown column name: {sel!r}")
            if len(matches) > 1:
                raise ValueError(
                    f"ambiguous column name {sel!r}: matches {len(matches)} header cells"
                )
            pos = matches[0]
        if pos in picked:
            raise ValueError(f"duplicate selector: {sel!r}")
        picked.append(pos)
    return picked


def parse_check_rule(rule):
    m = CHECK_RE.match(rule.strip())
    if not m:
        raise ValueError(
            f"invalid --check rule {rule!r}; expected COLUMN:STATOP VALUE "
            "(e.g. age:mean>=18)"
        )
    col, stat, op, raw = m.groups()
    try:
        expected = float(raw)
    except ValueError:
        raise ValueError(f"invalid numeric threshold in rule {rule!r}: {raw!r}")
    return col.strip(), stat, op, expected


def compare(actual, op, expected):
    if op == ">=":
        return actual >= expected
    if op == "<=":
        return actual <= expected
    if op == ">":
        return actual > expected
    if op == "<":
        return actual < expected
    if op == "==":
        return actual == expected
    return actual != expected


def read_rows(fp, Delimiter, quotechar):
    reader = csv.reader(fp, delimiter=Delimiter, quotechar=quotechar)
    rows = [row for row in reader]
    return rows


def open_input(path):
    if path not in ("-", None):
        return open(path, newline="", encoding="utf-8-sig")
    return sys.stdin


def compute_column_stats(name, index, column_values, want_stdev):
    values = []
    empty = 0
    for raw in column_values:
        cell = (raw or "").strip()
        if cell == "":
            empty += 1
            continue
        num = to_number(cell)
        if num is not None:
            values.append(num)
    entry = {
        "name": name,
        "index": index,
        "empty": empty,
        "non_numeric": len(column_values) - len(values) - empty,
        "count": len(values),
    }
    if values:
        entry["min"] = min(values)
        entry["max"] = max(values)
        entry["mean"] = sum(values) / len(values)
        entry["median"] = statistics.median(values)
        if want_stdev:
            entry["stdev"] = statistics.stdev(values) if len(values) > 1 else 0.0
    return entry


def build_parser():
    p = argparse.ArgumentParser(
        prog="csv-column-stats",
        description=(
            "Compute numeric statistics (min/max/mean/median/stdev) per CSV column."
        ),
    )
    p.add_argument(
        "file",
        nargs="?",
        default="-",
        help="CSV file to read (default: stdin; '-' also means stdin)",
    )
    p.add_argument(
        "-c",
        "--columns",
        default=None,
        help="comma-separated column selectors (header names or 1-based indices); default: all",
    )
    p.add_argument(
        "-d",
        "--delimiter",
        default=",",
        help="field delimiter (default: ','; aliases: tab, comma, semicolon, pipe)",
    )
    p.add_argument(
        "-q",
        "--quotechar",
        default='"',
        help="quote character (default: double quote)",
    )
    p.add_argument(
        "--no-header",
        action="store_true",
        help="treat the first row as data; columns are addressed by index",
    )
    p.add_argument(
        "--stdev",
        action="store_true",
        help="include the sample standard deviation (n-1)",
    )
    p.add_argument(
        "--check",
        action="append",
        default=[],
        metavar="RULE",
        help=(
            "CI rule of the form COLUMN:STAT[op]VALUE (e.g. age:mean>=18). "
            "Repeatable; exit 2 when any assertion fails."
        ),
    )
    p.add_argument(
        "--json", action="store_true", help="print a JSON report instead of a table"
    )
    p.add_argument(
        "--quiet",
        action="store_true",
        help="with --check, suppress the human-readable diagnosis",
    )
    p.add_argument("--version", action="version", version=f"%(prog)s {__version__}")
    return p


def rules_need_stdev(rule_args):
    for value in rule_args or []:
        for piece in value.split(","):
            piece = piece.strip()
            if piece and f":stdev" in piece:
                return True
    return False


def expand_rule_args(rule_args):
    out = []
    for value in rule_args or []:
        for piece in value.split(","):
            piece = piece.strip()
            if piece:
                out.append(piece)
    return out


def fmt(value):
    if value is None:
        return "-"
    if isinstance(value, float):
        if float(value).is_integer() and abs(value) < 1e15:
            return str(int(value))
        return f"{value:.6g}"
    return str(value)


def print_table(results, include_stdev):
    cols = ["name", "count", "empty", "non_numeric", "min", "max", "mean", "median"]
    if include_stdev:
        cols.append("stdev")
    widths = {}
    for c in cols:
        widths[c] = max([len(c)] + [len(fmt(r.get(c))) for r in results])
    print("  ".join(c.ljust(widths[c]) for c in cols))
    print("  ".join("-" * widths[c] for c in cols))
    for r in results:
        line = []
        for c in cols:
            text = fmt(r.get(c))
            if c == "name":
                line.append(text.ljust(widths[c]))
            else:
                line.append(text.rjust(widths[c]))
        print("  ".join(line))


def main(argv=None):
    args = build_parser().parse_args(argv)

    try:
        delimiter = parse_delimiter(args.delimiter)
    except ValueError as exc:
        print(f"csv-column-stats: error: {exc}", file=sys.stderr)
        return 1
    if len(args.quotechar) != 1:
        print(
            "csv-column-stats: error: quotechar must be a single character",
            file=sys.stderr,
        )
        return 1

    try:
        fp = open_input(args.file)
    except OSError as exc:
        print(f"csv-column-stats: error: {exc}", file=sys.stderr)
        return 1

    try:
        with fp:
            data = fp.read()
    except OSError as exc:
        print(f"csv-column-stats: error: {exc}", file=sys.stderr)
        return 1

    rows = read_rows(io.StringIO(data), delimiter, args.quotechar)
    if not rows:
        print("csv-column-stats: error: empty CSV document", file=sys.stderr)
        return 1

    if args.no_header:
        header = [str(i + 1) for i in range(len(rows[0]))]
    else:
        header, rows = rows[0], rows[1:]

    selectors = [
        s.strip() for s in (args.columns or "").split(",") if s.strip() != ""
    ]
    try:
        picked = resolve_columns(header, selectors)
    except ValueError as exc:
        print(f"csv-column-stats: error: {exc}", file=sys.stderr)
        return 1
    indices = picked if picked else list(range(len(header)))

    want_stdev = args.stdev or rules_need_stdev(args.check)

    results = []
    stat_by_name = {}
    for i in indices:
        column_values = [(row[i] if i < len(row) else "") for row in rows]
        entry = compute_column_stats(header[i], i + 1, column_values, want_stdev)
        results.append(entry)
        stat_by_name[header[i]] = entry

    rules = []
    for rule_text in expand_rule_args(args.check):
        try:
            rules.append(parse_check_rule(rule_text))
        except ValueError as exc:
            print(f"csv-column-stats: error: {exc}", file=sys.stderr)
            return 1

    check_results = []
    failed = 0
    for col, stat, op, expected in rules:
        entry = stat_by_name.get(col)
        if entry is None:
            try:
                pos = resolve_columns(header, [col])[0]
            except ValueError as exc:
                print(f"csv-column-stats: error: {exc}", file=sys.stderr)
                return 1
            column_values = [(row[pos] if pos < len(row) else "") for row in rows]
            entry = compute_column_stats(header[pos], pos + 1, column_values, want_stdev)
            results.append(entry)
            stat_by_name[header[pos]] = entry
        actual = entry.get(stat)
        if stat == "count":
            actual = entry["count"]
        ok = actual is not None and compare(actual, op, expected)
        failed += 0 if ok else 1
        check_results.append(
            {
                "column": col,
                "stat": stat,
                "op": op,
                "expected": expected,
                "actual": actual,
                "ok": ok,
            }
        )

    include_stdev = args.stdev or any(c["stat"] == "stdev" for c in check_results)

    report = {
        "file": args.file,
        "rows": len(rows),
        "columns": results,
        "checks": check_results,
    }

    if args.json:
        print(json.dumps(report, indent=2, ensure_ascii=False))
    else:
        if rules:
            if not args.quiet:
                for c in check_results:
                    status = "OK  " if c["ok"] else "FAIL"
                    print(
                        f"{status} {c['column']}:{c['stat']}{c['op']}{c['expected']} "
                        f"(actual: {fmt(c['actual'])})",
                        file=sys.stderr,
                    )
        else:
            print_table(results, include_stdev)

    if rules:
        return 2 if failed else 0
    return 0


if __name__ == "__main__":
    sys.exit(main())
