"""Detect accidental repetitions in text: doubled words, repeated n-grams.

Two detectors:
  - doubled words: 'the the', 'and and' (case-insensitive)
  - repeated n-grams: sequences of N words occurring more than once

Exit codes:
    0 - no repetition found
    1 - I/O or CLI usage error
    2 - repetitions found (lint-style)
"""

import argparse
import json
import re
import sys
from collections import Counter

WORD_RE = re.compile(r"[\w'-]+", re.UNICODE)


def doubled(lines):
    findings = []
    for lineno, line in enumerate(lines, 1):
        words = list(WORD_RE.finditer(line))
        for a, b in zip(words, words[1:]):
            if a.group(0).lower() == b.group(0).lower():
                findings.append({
                    "line": lineno,
                    "col": a.start() + 1,
                    "type": "doubled-word",
                    "text": f"{a.group(0)} {b.group(0)}",
                })
        i = 0
    return findings


def repeated_ngrams(text, n, min_count=2):
    words = [m.group(0).lower() for m in WORD_RE.finditer(text)]
    grams = Counter(tuple(words[i:i + n]) for i in range(len(words) - n + 1))
    return [
        {"type": f"repeated-{n}gram", "text": " ".join(g), "count": c}
        for g, c in grams.most_common() if c >= min_count
    ]


def main(argv=None):
    p = argparse.ArgumentParser(
        prog="text-repetition-scan",
        description="Detect doubled words and repeated n-grams in text.",
    )
    p.add_argument("files", nargs="*", help="input files (default: stdin; '-' = stdin)")
    p.add_argument("--ngrams", type=int, metavar="N", default=0,
                   help="also report word N-grams occurring 2+ times")
    p.add_argument("--min-count", type=int, default=2,
                   help="minimum occurrences for n-grams (default: 2)")
    p.add_argument("--json", action="store_true", help="emit findings as JSON")
    args = p.parse_args(argv)

    files = args.files or ["-"]
    all_findings = []
    for path in files:
        if path == "-":
            text, label = sys.stdin.read(), "<stdin>"
        else:
            try:
                with open(path, encoding="utf-8") as fh:
                    text = fh.read()
                label = path
            except OSError as exc:
                print(f"text-repetition-scan: {path}: {exc}", file=sys.stderr)
                return 1
        findings = doubled(text.splitlines())
        if args.ngrams >= 2:
            findings += repeated_ngrams(text, args.ngrams, args.min_count)
        for f in findings:
            f["file"] = label
        all_findings.extend(findings)

    if args.json:
        print(json.dumps({"count": len(all_findings),
                          "findings": all_findings}, indent=2, ensure_ascii=False))
    else:
        for f in all_findings:
            loc = f"{f['file']}:{f['line']}:{f['col']}" if "line" in f else f["file"]
            extra = f" x{f['count']}" if "count" in f else ""
            print(f"{loc}: {f['type']}: {f['text']!r}{extra}")
    return 2 if all_findings else 0


if __name__ == "__main__":
    sys.exit(main())
