# json-to-latex

Convert JSON arrays to LaTeX tabular tables.

## Install

```bash
pip install json-to-latex
```

## Usage

```bash
json-to-latex data.json
json-to-latex data.json -o table.tex
json-to-latex data.json --booktabs
```

## Examples

**Input**:
```json
[{"name": "Alice", "score": 95}, {"name": "Bob", "score": 87}]
```

**Output**:
```latex
\begin{tabular}{|l|l|}
\hline
name & score \\
\hline
Alice & 95 \\
\hline
Bob & 87 \\
\hline
\end{tabular}
```

## License

MIT
