# csv-column-stats

Numeric statistics per CSV column, in a table or as machine-readable JSON.
Cells that don't parse as a number are skipped and reported, never counted
as zero. Usable as a CI gate via repeatable `--check` rules.

## Features

- Per-column statistics: `count` (numeric values), `empty`, `non_numeric`,
  `min`, `max`, `mean`, `median`, optional `stdev` (sample, n-1).
- Target columns by header name or 1-based index (`-c age,score` or `-c 2,3`).
- Non-numeric values are counted in `non_numeric` and excluded from min/max/....
- CI gate: `--check COLUMN:STAT[op]VALUE` (e.g. `age:mean>=18`), exit code 2 on failure.
- JSON or human table output.

## Install

```bash
pip install csv-column-stats
# or from source:
pip install git+https://github.com/TataneSan/csv-column-stats.git
```

## Usage

```bash
# full table for every column
csv-column-stats data.csv

# specific columns, include standard deviation
csv-column-stats data.csv -c age,score --stdev

# JSON report
csv-column-stats data.csv --json

# CI gate: fails (exit 2) if the mean age drops below 18
csv-column-stats data.csv --check 'age:mean>=18'

# several rules, by name or index (comma-separated or repeated flags)
csv-column-stats data.csv \
  --check 'age:count>=100' \
  --check 'score:min>=0,2:max<=100'

# read from stdin
cat data.csv | csv-column-stats --columns 2,3 --json
```

### Examples

Input `people.csv`:

```csv
name,age,score,notes
alice,30,88.5,
bob,25,92.1,x
carol,41,75.0,
dan,,60.5,broken
```

```text
$ csv-column-stats people.csv
name   count  empty  non_numeric  min   max   mean   median
-----  -----  -----  -----------  ---   ---   ----   ------
name       0      0            4  -     -     -      -
age        3      1            0  25    41    32     30
score      4      0            0  60.5  92.1  79.02  81.75
notes      0      2            2  -     -     -      -
```

JSON:

```json
$ csv-column-stats people.csv -c age --json
{
  "file": "people.csv",
  "rows": 4,
  "columns": [
    {
      "name": "age", "index": 2, "empty": 1, "non_numeric": 0,
      "count": 3, "min": 25.0, "max": 41.0,
      "mean": 32.0, "median": 30.0
    }
  ],
  "checks": []
}
```

CI failure:

```text
$ csv-column-stats people.csv --check 'age:mean>=40'
FAIL age:mean>=40.0 (actual: 32)   # exits with status 2
```

## Options

| Option | Description |
| ------ | ----------- |
| `-c, --columns LIST` | Comma-separated column selectors; default: all columns. |
| `-d, --delimiter C`  | Field delimiter (aliases: `tab`, `comma`, `semicolon`, `pipe`). |
| `-q, --quotechar C`  | Quote character (default `"`). |
| `--no-header`        | Treat the first row as data; address columns by index. |
| `--stdev`            | Include the sample standard deviation. |
| `--check RULE`       | CI assertion `COLUMN:STAT[op]VALUE`; repeatable, exit 2 on failure. |
| `--json`             | Emit a JSON report. |
| `--quiet`            | Suppress check diagnostics on stderr. |

Check operators: `>=`, `<=`, `>`, `<`, `==`, `!=`. Statistics: `count`,
`min`, `max`, `mean`, `median`, `stdev`.

## Exit codes

- `0` — success (all checks passed, if any)
- `1` — I/O or argument error
- `2` — check mode: at least one assertion failed

## Development

```bash
pip install -e .
python -m pytest tests -q
```

## License

MIT — see [LICENSE](LICENSE).
