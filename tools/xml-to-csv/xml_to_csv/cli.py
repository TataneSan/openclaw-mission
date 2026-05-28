"""CLI for xml-to-csv: convert XML files to CSV."""

import argparse
import csv
import sys
import xml.etree.ElementTree as ET
from pathlib import Path


def extract_records(root: ET.Element, record_tag: str | None) -> tuple[list[str], list[list[str]]]:
    """Extract header and rows from XML."""
    if record_tag:
        elements = root.findall(f".//{record_tag}")
    else:
        # Auto-detect: use direct children that have sub-elements or text siblings
        candidates = list(root)
        if candidates and list(candidates[0]):
            elements = candidates
        else:
            # Look for repeated tags at any depth
            tag_counts: dict[str, int] = {}
            for elem in root.iter():
                if elem.tag not in tag_counts:
                    tag_counts[elem.tag] = 0
                tag_counts[elem.tag] += 1
            # Find tag with count > 1 that has children or text
            record_tag = None
            for tag, count in tag_counts.items():
                if count > 1:
                    record_tag = tag
                    break
            if record_tag:
                elements = root.findall(f".//{record_tag}")
            else:
                elements = [root]

    if not elements:
        raise ValueError(f"No records found with tag '{record_tag or 'auto'}'")

    # Flatten each element into a dict
    all_keys: list[str] = []
    rows: list[list[str]] = []

    for elem in elements:
        record: dict[str, str] = {}
        for child in elem:
            key = child.tag
            if key not in all_keys:
                all_keys.append(key)
            record[key] = child.text or ""
        # Also capture attributes
        for attr_name, attr_val in elem.attrib.items():
            key = f"@{attr_name}"
            if key not in all_keys:
                all_keys.append(key)
            record[key] = attr_val
        rows.append([record.get(k, "") for k in all_keys])

    return all_keys, rows


def convert(input_path: Path, output_path: Path, *, record_tag: str | None = None,
            delimiter: str = ",") -> None:
    tree = ET.parse(str(input_path))
    root = tree.getroot()
    headers, rows = extract_records(root, record_tag)

    with output_path.open("w", newline="", encoding="utf-8") as f:
        writer = csv.writer(f, delimiter=delimiter)
        writer.writerow(headers)
        writer.writerows(rows)

    print(f"Converted {input_path} -> {output_path} ({len(rows)} rows, {len(headers)} columns)")


def main() -> None:
    parser = argparse.ArgumentParser(prog="xml-to-csv", description="Convert XML files to CSV.")
    parser.add_argument("input", type=Path, help="Input XML file")
    parser.add_argument("-o", "--output", type=Path, help="Output CSV file (default: input.csv)")
    parser.add_argument("-t", "--tag", type=str, help="XML tag for record rows (auto-detect if omitted)")
    parser.add_argument("-d", "--delimiter", default=",", help="CSV delimiter (default: comma)")

    args = parser.parse_args()
    output = args.output or args.input.with_suffix(".csv")
    try:
        convert(args.input, output, record_tag=args.tag, delimiter=args.delimiter)
    except (FileNotFoundError, ET.ParseError, ValueError) as exc:
        print(f"Error: {exc}", file=sys.stderr)
        sys.exit(1)


if __name__ == "__main__":
    main()
