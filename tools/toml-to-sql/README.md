# toml-to-sql

Converts TOML files to SQL `CREATE TABLE` and `INSERT` statements.

Supports PostgreSQL and MySQL dialects, automatic type inference, and reads from files or stdin.

## Installation

```bash
# From source
go install github.com/TataneSan/toml-to-sql@latest

# Or download the binary from Releases
```

## Usage

```
toml-to-sql [OPTIONS] [FILE]
```

### Options

| Flag | Description | Default |
|------|-------------|---------|
| `-f, --format FORMAT` | Output format: `text`, `json` | `text` |
| `-t, --table NAME` | Override table name | TOML table name |
| `-d, --dialect DIALECT` | SQL dialect: `postgres`, `mysql` | `postgres` |
| `--file FILE` | Input file path | stdin |
| `-h, --help` | Show help message | — |

## Examples

### Basic usage from stdin

```bash
cat data.toml | toml-to-sql
```

Given `data.toml`:
```toml
[[products]]
name = "Widget"
price = 9.99
in_stock = true

[[products]]
name = "Gadget"
price = 24.99
in_stock = false
```

Output:
```sql
-- Table: products
CREATE TABLE "products" (
    "in_stock" BOOLEAN,
    "name" TEXT,
    "price" REAL
);

INSERT INTO "products" ("in_stock", "name", "price") VALUES (TRUE, 'Widget', 9.99);
INSERT INTO "products" ("in_stock", "name", "price") VALUES (FALSE, 'Gadget', 24.99);
```

### From a file

```bash
toml-to-sql data.toml
```

### MySQL dialect

```bash
toml-to-sql --dialect mysql data.toml
```

Output:
```sql
-- Table: products
CREATE TABLE `products` (
    `in_stock` TINYINT(1),
    `name` TEXT,
    `price` REAL
);

INSERT INTO `products` (`in_stock`, `name`, `price`) VALUES (1, 'Widget', 9.99);
INSERT INTO `products` (`in_stock`, `name`, `price`) VALUES (0, 'Gadget', 24.99);
```

### Override table name

```bash
toml-to-sql --table inventory data.toml
```

## Features

- **Auto type inference**: Detects `BOOLEAN`, `INTEGER`, `REAL`, and `TEXT` types
- **Multiple dialects**: PostgreSQL (default) and MySQL
- **Array of tables**: Supports `[[array]]` TOML syntax
- **Single tables**: Supports `[table]` TOML syntax
- **stdin support**: Pipe TOML data directly
- **Sorted columns**: Output columns in alphabetical order for consistency

## Type Inference Rules

| TOML Type | PostgreSQL | MySQL |
|-----------|-----------|-------|
| boolean | BOOLEAN | TINYINT(1) |
| integer | INTEGER | INTEGER |
| float | REAL | REAL |
| string | TEXT | TEXT |

If any value in a column is a string, the entire column is typed as `TEXT`.

## License

MIT — see [LICENSE](LICENSE) for details.
