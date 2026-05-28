# toml-to-sql

Convert TOML files to SQL INSERT statements.

## Install

```bash
go install github.com/TataneSan/toml-to-sql@latest
```

Or build from source:

```bash
git clone https://github.com/TataneSan/toml-to-sql.git
cd toml-to-sql
go build -o toml-to-sql .
```

## Usage

```
toml-to-sql [options] <file.toml>
```

### Options

| Flag       | Description                                      |
|------------|--------------------------------------------------|
| `-create`  | Generate CREATE TABLE statements before INSERTs  |

### Examples

Basic conversion:

```bash
toml-to-sql config.toml
```

With CREATE TABLE statements:

```bash
toml-to-sql -create config.toml
```

Pipe to a SQL file:

```bash
toml-to-sql -create config.toml > output.sql
```

## Input / Output

Given this `config.toml`:

```toml
[server]
host = "localhost"
port = 8080
debug = true

[[users]]
name = "alice"
age = 30

[[users]]
name = "bob"
age = 25
```

Running `toml-to-sql -create config.toml` outputs:

```sql
CREATE TABLE IF NOT EXISTS `server` (
  `debug` INTEGER,
  `host` TEXT,
  `port` INTEGER
);
INSERT INTO `server` (`debug`, `host`, `port`)
VALUES (1, 'localhost', 8080);
CREATE TABLE IF NOT EXISTS `users` (
  `age` INTEGER,
  `name` TEXT
);
INSERT INTO `users` (`age`, `name`)
VALUES (30, 'alice');
INSERT INTO `users` (`age`, `name`)
VALUES (25, 'bob');
```

## Features

- Converts TOML tables to SQL INSERT statements
- Optional CREATE TABLE generation with type inference
- Handles nested tables (flattened with dot notation)
- Handles arrays of tables (each element becomes a row)
- Arrays and inline tables are serialized as JSON
- Boolean values are converted to 1/0
- SQL string escaping (single quotes doubled)
- Sorted column output for deterministic results

## License

MIT
