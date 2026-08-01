#!/usr/bin/env python3
"""url-encode-cli - percent-encode and decode URL components.

Values come from arguments or one per line on stdin. Default operation
is encode; -d decodes.

Component presets (safe character sets):
    query    - application/x-www-form-urlencoded compatible: only
               unreserved chars kept, spaces as %20
    path     - additionally keeps '/' and sub-delims common in paths
    form     - like query but encodes spaces as '+' (quote_plus)
    strict   - encodes everything except unreserved chars (RFC 3986)

Decoding validates percent triplets; --errors=replace swaps invalid
byte sequences for U+FFFD, --errors=strict (default) counts them.

Exit codes:
    0 - success
    1 - CLI or I/O error
    2 - --check mode: at least one value was not properly encoded
        (round-trip mismatch) or contained decoding errors
"""

import argparse
import json
import re
import sys
from urllib.parse import quote, quote_plus, unquote, unquote_plus

PRESETS = {
    "strict": ("", quote),
    "query": ("", quote),
    "path": ("/:@!$&'()*+,;=", quote),
    "form": ("", quote_plus),
}

def decode(text, form=False, errors="strict"):
    fn = unquote_plus if form else unquote
    return fn(text, errors=errors)


def encode(text, preset):
    safe, fn = PRESETS[preset]
    return fn(text, safe=safe)


def iter_inputs(args):
    if args:
        for a in args:
            yield a
    else:
        for line in sys.stdin:
            line = line.rstrip("\n")
            if line:
                yield line


def has_errors(encoded):
    """Detect malformed percent sequences (a '%' not followed by 2 hex)."""
    i, n = 0, len(encoded)
    while i < n:
        if encoded[i] == "%":
            if i + 2 >= n or not all(c in "0123456789abcdefABCDEF"
                                     for c in encoded[i + 1:i + 3]):
                return True
            i += 3
        else:
            i += 1
    return False


def main(argv=None):
    parser = argparse.ArgumentParser(
        prog="url-encode-cli",
        description="Percent-encode and decode URL components.",
    )
    parser.add_argument(
        "values", nargs="*", metavar="VALUE",
        help="values to encode/decode; stdin (one per line) when omitted",
    )
    parser.add_argument(
        "-d", "--decode", action="store_true",
        help="decode instead of encode",
    )
    parser.add_argument(
        "--preset", choices=list(PRESETS), default="query",
        help="safe character set for encoding (default: query)",
    )
    parser.add_argument(
        "--errors", choices=["strict", "replace"],
        default="strict",
        help="decoding error handling (default: strict)",
    )
    parser.add_argument(
        "--check", action="store_true",
        help="CI mode: exit 2 when decoding found malformed %% sequences, "
             "or an encoded value does not decode back exactly",
    )
    parser.add_argument(
        "--json", action="store_true",
        help="one JSON object per value: input/output",
    )
    args = parser.parse_args(argv)

    values = list(iter_inputs(args.values))
    if not values:
        print("url-encode-cli: no input values", file=sys.stderr)
        return 1

    results = []
    failures = 0
    for value in values:
        if args.decode:
            out = decode(value, form=args.preset == "form",
                         errors=args.errors)
            rt = encode(out, args.preset)
            bad = has_errors(value) or \
                (args.errors == "strict" and rt != value)
        else:
            out = encode(value, args.preset)
            back = decode(out, form=args.preset == "form",
                          errors="replace")
            bad = back != value
        failures += int(bad)
        results.append({"input": value, "output": out})

    exit_code = 2 if (args.check and failures) else 0
    for r in results:
        if args.json:
            print(json.dumps(r, ensure_ascii=False))
        else:
            print(r["output"])

    if args.check and failures:
        print(f"url-encode-cli: {failures} value(s) failed validation",
              file=sys.stderr)
    return exit_code


if __name__ == "__main__":
    sys.exit(main())
