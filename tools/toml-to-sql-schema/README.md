# toml-to-sql-schema

Generate SQL CREATE TABLE statements from TOML definitions.

## Install

```bash
pip install toml-to-sql-schema
```

## Usage

```bash
toml-to-sql-schema schema.toml
toml-to-sql-schema schema.toml -o schema.sql
toml-to-sql-schema schema.toml -t users
```

## Example

**Input** (`schema.toml`):
```toml
[users]

[users.id]
type = "integer"
primary_key = true

[users.name]
type = "varchar"
nullable = false

[users.email]
type = "string"
nullable = false

[users.active]
type = "boolean"
default = true
```

**Output**:
```sql
CREATE TABLE users (
  id INTEGER PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  email TEXT NOT NULL,
  active BOOLEAN DEFAULT True
);
```

## License

MIT
