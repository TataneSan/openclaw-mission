# csv-split-rows

Split a CSV into multiple chunk files: even split, max rows per file, or one
file per distinct value of a key column. The header row is repeated in every
output file.

## Features

- `--chunks N` even split, `--size K` fixed row batches, `--by COLUMN` grouping
- Safe output filenames from key values
- JSON manifest of produced files
- Custom delimiter, stdin support
- Pure Python standard library, no dependencies

## Install

```bash
pip install .
pip install git+https://github.com/TataneSan/csv-split-rows.git
```

## Usage

```bash
# 4 roughly even files
csv-split-rows big.csv --chunks 4 -o out/

# batches of 5000 rows
csv-split-rows big.csv --size 5000 --prefix batch -o out/

# one file per country
csv-split-rows users.csv --by country -o by-country/
by-country/part-france.csv: 1023 rows
by-country/part-spain.csv: 812 rows
```

## Exit codes

| Code | Meaning |
|------|---------|
| 0 | Success |
| 1 | I/O or CLI error |
| 2 | Unknown `--by` column |

## License

MIT
