"""text-acronym: generate acronyms from phrases, expand known ones.

Splits phrases on whitespace/hyphens/underscores, skips stop words and
takes the first letter of each remaining word. Comes with a small
built-in dictionary to expand well-known tech acronyms, plus --reverse
to look up an acronym.

Exit codes:
    0 - success
    1 - I/O or CLI error
    2 - acronym not found in dictionary (--reverse mode)
"""

import argparse
import json
import re
import sys
import unicodedata
from typing import List, Optional

from . import __version__

STOP_WORDS = {
    # english
    "a", "an", "and", "are", "as", "at", "be", "by", "for", "from", "in",
    "is", "it", "of", "on", "or", "the", "to", "with",
    # french
    "le", "la", "les", "de", "des", "du", "un", "une", "et", "ou", "en",
    "au", "aux", "dans", "sur", "par", "pour", "avec", "se", "sa", "son",
}

DICTIONARY = {
    "API": "Application Programming Interface",
    "ASCII": "American Standard Code for Information Interchange",
    "CLI": "Command Line Interface",
    "CPU": "Central Processing Unit",
    "CSS": "Cascading Style Sheets",
    "CSV": "Comma-Separated Values",
    "DNS": "Domain Name System",
    "DOM": "Document Object Model",
    "FAQ": "Frequently Asked Questions",
    "FTP": "File Transfer Protocol",
    "GIF": "Graphics Interchange Format",
    "GNU": "GNU's Not Unix",
    "GPU": "Graphics Processing Unit",
    "GUI": "Graphical User Interface",
    "HTML": "HyperText Markup Language",
    "HTTP": "HyperText Transfer Protocol",
    "HTTPS": "HyperText Transfer Protocol Secure",
    "IDE": "Integrated Development Environment",
    "IP": "Internet Protocol",
    "ISBN": "International Standard Book Number",
    "ISO": "International Organization for Standardization",
    "JPEG": "Joint Photographic Experts Group",
    "JSON": "JavaScript Object Notation",
    "LAN": "Local Area Network",
    "LASER": "Light Amplification by Stimulated Emission of Radiation",
    "LCD": "Liquid Crystal Display",
    "LED": "Light Emitting Diode",
    "MAC": "Media Access Control",
    "MIME": "Multipurpose Internet Mail Extensions",
    "MVP": "Minimum Viable Product",
    "NASA": "National Aeronautics and Space Administration",
    "NATO": "North Atlantic Treaty Organization",
    "OCR": "Optical Character Recognition",
    "PDF": "Portable Document Format",
    "PIN": "Personal Identification Number",
    "PNG": "Portable Network Graphics",
    "RADAR": "Radio Detection and Ranging",
    "RAM": "Random Access Memory",
    "REST": "Representational State Transfer",
    "RGB": "Red Green Blue",
    "ROM": "Read-Only Memory",
    "RSS": "Really Simple Syndication",
    "SCUBA": "Self-Contained Underwater Breathing Apparatus",
    "SDK": "Software Development Kit",
    "SMTP": "Simple Mail Transfer Protocol",
    "SQL": "Structured Query Language",
    "SSH": "Secure Shell",
    "SSL": "Secure Sockets Layer",
    "SVG": "Scalable Vector Graphics",
    "TCP": "Transmission Control Protocol",
    "TLS": "Transport Layer Security",
    "TTFB": "Time To First Byte",
    "UDP": "User Datagram Protocol",
    "UI": "User Interface",
    "URI": "Uniform Resource Identifier",
    "URL": "Uniform Resource Locator",
    "USB": "Universal Serial Bus",
    "UTC": "Coordinated Universal Time",
    "UUID": "Universally Unique Identifier",
    "UX": "User Experience",
    "VPN": "Virtual Private Network",
    "WAN": "Wide Area Network",
    "W3C": "World Wide Web Consortium",
    "XML": "eXtensible Markup Language",
    "YAML": "YAML Ain't Markup Language",
    "ZIP": "Zone Improvement Plan",
}

_WORD_RE = re.compile(r"[^\W_]+", re.UNICODE)


def _strip_accents(word: str) -> str:
    return "".join(
        c for c in unicodedata.normalize("NFD", word)
        if unicodedata.category(c) != "Mn"
    )


def words_of(phrase: str) -> List[str]:
    return _WORD_RE.findall(phrase)


def build_acronym(phrase: str, include_stop_words: bool = False) -> dict:
    words = words_of(phrase)
    kept = []
    skipped = []
    for w in words:
        if not include_stop_words and w.lower() in STOP_WORDS:
            skipped.append(w)
            continue
        kept.append(w)
    acronym = "".join(_strip_accents(w)[0] for w in kept).upper() if kept else ""
    return {
        "phrase": phrase.strip(),
        "words": kept,
        "skipped": skipped,
        "acronym": acronym,
        "length": len(acronym),
    }


def expand(acronym: str) -> Optional[str]:
    return DICTIONARY.get(acronym.strip().upper())


def _read_phrases(args_texts: List[str], file: Optional[str]) -> List[str]:
    if file:
        if file == "-":
            return [l.strip() for l in sys.stdin if l.strip()]
        try:
            with open(file, "r", encoding="utf-8") as fh:
                return [l.strip() for l in fh if l.strip()]
        except OSError as exc:
            raise SystemExit("error: cannot read %s: %s" % (file, exc))
    if args_texts:
        return [" ".join(args_texts)]
    if not sys.stdin.isatty():
        return [l.strip() for l in sys.stdin if l.strip()]
    return []


def main(argv: Optional[List[str]] = None) -> int:
    parser = argparse.ArgumentParser(
        prog="text-acronym",
        description="Generate acronyms from phrases, or expand well-known acronyms.",
        epilog="Exit codes: 0 ok, 1 error, 2 not found in --reverse mode.",
    )
    parser.add_argument(
        "texts",
        nargs="*",
        metavar="PHRASE",
        help="phrase(s) to condense (joined together if several). Reads stdin when omitted.",
    )
    parser.add_argument(
        "-f", "--file",
        metavar="FILE",
        help="read phrases from FILE, one per line (- for stdin)",
    )
    parser.add_argument(
        "-a", "--all",
        action="store_true",
        help="include stop words (of, the, de, ...) as initials",
    )
    parser.add_argument(
        "-r", "--reverse",
        action="store_true",
        help="expand known acronyms instead of building them",
    )
    parser.add_argument("--json", action="store_true", help="emit a JSON report")
    parser.add_argument(
        "-l", "--lower",
        action="store_true",
        help="output the acronym in lowercase",
    )
    parser.add_argument(
        "--list",
        action="store_true",
        dest="list_dict",
        help="print the built-in acronym dictionary and exit",
    )
    parser.add_argument("--version", action="version", version="%(prog)s " + __version__)
    args = parser.parse_args(argv)

    if args.list_dict:
        for key in sorted(DICTIONARY):
            print("%-8s %s" % (key, DICTIONARY[key]))
        return 0

    try:
        phrases = _read_phrases(args.texts, args.file)
    except SystemExit as exc:
        print(exc, file=sys.stderr)
        return 1

    if not phrases:
        parser.error("no input (pass a phrase, --file, or pipe stdin)")
        return 1

    if args.reverse:
        misses = 0
        reports = []
        for item in phrases:
            hit = expand(item)
            reports.append({"acronym": item.strip().upper(), "expansion": hit, "found": hit is not None})
            if hit is None:
                misses += 1
        if args.json:
            payload = reports if len(reports) > 1 else reports[0]
            print(json.dumps(payload, indent=2, ensure_ascii=False))
        else:
            for r in reports:
                label = r["expansion"] or "unknown"
                print("%s  %s" % (r["acronym"], label))
        return 2 if misses else 0

    reports = []
    for phrase in phrases:
        if not phrase:
            continue
        r = build_acronym(phrase, include_stop_words=args.all)
        if args.lower:
            r["acronym"] = r["acronym"].lower()
        reports.append(r)

    if args.json:
        payload = reports if len(reports) > 1 else reports[0]
        print(json.dumps(payload, indent=2, ensure_ascii=False))
    else:
        single = len(reports) == 1
        for r in reports:
            if single:
                print(r["acronym"])
            else:
                print("%s  %s" % (r["phrase"], r["acronym"]))
    return 0


if __name__ == "__main__":
    sys.exit(main())
