# csv-to-graph

Generate terminal bar charts from CSV data.

Converts CSV files into visual horizontal or vertical bar charts rendered directly in the terminal.

## Install

```bash
go install github.com/TataneSan/csv-to-graph@latest
```

Or build from source:

```bash
git clone https://github.com/TataneSan/csv-to-graph.git
cd csv-to-graph
go build -o csv-to-graph .
```

## Usage

```
csv-to-graph [options] < input.csv
csv-to-graph [options] -f file.csv
```

### Options

| Flag | Default | Description |
|------|---------|-------------|
| `-f` | `""` | Path to CSV file (default: stdin) |
| `-col` | `1` | Column index for values (0-based) |
| `-label` | `0` | Column index for labels (0-based) |
| `-width` | `50` | Maximum bar width in characters |
| `-vertical` | `false` | Render bars vertically |
| `-sort` | `none` | Sort order: `none`, `asc`, `desc` |
| `-top` | `0` | Show only top N bars (0 = all) |
| `-title` | `""` | Chart title |
| `-char` | `"█"` | Bar character |
| `-empty` | `"░"` | Empty bar character |

## Examples

### Basic horizontal chart

```bash
$ cat sales.csv
Product,Sales
Alpha,1200
Beta,800
Gamma,2100
Delta,450
```

```bash
$ csv-to-graph -f sales.csv
  Alpha  │█████████████████████████████░░░░░░░░░░░░░░░░░░░░░ 1200
  Beta   │███████████████████░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░ 800
  Gamma  │██████████████████████████████████████████████████ 2100
  Delta  │███████████░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░ 450
```

### Sorted descending, top 5

```bash
$ csv-to-graph -f data.csv -sort desc -top 5
```

### Vertical chart with title

```bash
$ csv-to-graph -f data.csv -vertical -title "Monthly Revenue"
```

### Custom width and columns

```bash
$ csv-to-graph -f data.csv -col 3 -label 1 -width 80
```

### Pipe from other commands

```bash
$ sort -t, -k2 -nr data.csv | csv-to-graph
```

## License

MIT
