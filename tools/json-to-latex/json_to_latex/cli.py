"""CLI for json-to-latex: convert JSON arrays to LaTeX tables."""

import argparse
import json
import sys
from pathlib import Path


def escape_latex(text: str) -> str:
    return str(text).replace("&", r"\&").replace("%", r"\%").replace("$", r"\$")         .replace("#", r"\#").replace("_", r"\_").replace("{", r"\{").replace("}", r"\}")         .replace("~", r"	extasciitilde{}").replace("^", r"	extasciicircum{}")         .replace("-", r"	extendash{}")


def generate_latex(data: list[dict], *, booktabs: bool = False, center: bool = False) -> str:
    if not data:
        raise ValueError("No data to convert")

    headers = list(data[0].keys())
    n_cols = len(headers)
    alignment = "c" * n_cols if center else "l" * n_cols

    lines = []
    if booktabs:
        lines.append("\begin{table}[ht]")
        lines.append("\centering")
        lines.append(f"\begin{{tabular}}{{{alignment}}}")
        lines.append("\toprule")
    else:
        lines.append(f"\begin{{tabular}}{{|{'|'.join(alignment)}|}}")
        lines.append("\hline")

    lines.append(" & ".join(escape_latex(h) for h in headers) + " \\")
    if booktabs:
        lines.append("\midrule")
    else:
        lines.append("\hline")

    for row in data:
        cells = [escape_latex(str(row.get(h, ""))) for h in headers]
        lines.append(" & ".join(cells) + " \\")
        if not booktabs:
            lines.append("\hline")

    if booktabs:
        lines.append("\bottomrule")
    lines.append("\end{tabular}")
    if booktabs:
        lines.append("\end{table}")

    return "\n".join(lines)


def main() -> None:
    parser = argparse.ArgumentParser(prog="json-to-latex", description="Convert JSON arrays to LaTeX tables.")
    parser.add_argument("input", type=Path, help="Input JSON file (array of objects)")
    parser.add_argument("-o", "--output", type=Path, help="Output .tex file (default: input.tex)")
    parser.add_argument("--booktabs", action="store_true", help="Use booktabs style")
    parser.add_argument("--center", action="store_true", help="Center-align all columns")
    args = parser.parse_args()

    data = json.loads(args.input.read_text(encoding="utf-8"))
    if not isinstance(data, list):
        data = [data]

    output = args.output or args.input.with_suffix(".tex")
    try:
        latex = generate_latex(data, booktabs=args.booktabs, center=args.center)
        output.write_text(latex + "\n", encoding="utf-8")
        print(f"Converted {args.input} -> {output} ({len(data)} rows)")
    except (json.JSONDecodeError, ValueError) as exc:
        print(f"Error: {exc}", file=sys.stderr)
        sys.exit(1)


if __name__ == "__main__":
    main()
