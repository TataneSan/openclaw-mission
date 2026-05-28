"""CLI for toml-to-sql-schema: generate SQL CREATE TABLE from TOML."""

import argparse
import sys
from pathlib import Path
from tomllib import loads

TYPE_MAP = {
    "string": "TEXT", "str": "TEXT", "text": "TEXT", "varchar": "VARCHAR(255)",
    "int": "INTEGER", "integer": "INTEGER", "long": "BIGINT", "bigint": "BIGINT",
    "float": "REAL", "double": "REAL", "real": "REAL", "decimal": "DECIMAL(19,4)",
    "bool": "BOOLEAN", "boolean": "BOOLEAN", "date": "DATE",
    "datetime": "TIMESTAMP", "timestamp": "TIMESTAMP", "time": "TIME",
    "binary": "BLOB", "blob": "BLOB", "json": "TEXT",
}


def infer_sql_type(value) -> str:
    if isinstance(value, bool):
        return "BOOLEAN"
    if isinstance(value, int):
        return "BIGINT" if abs(value) > 2**31 - 1 else "INTEGER"
    if isinstance(value, float):
        return "REAL"
    if isinstance(value, str):
        return "TEXT"
    if isinstance(value, list):
        return "TEXT"
    return "TEXT"


def generate_schema(data: dict, table_name: str | None = None) -> str:
    if not isinstance(data, dict):
        raise ValueError("TOML root must be a table")

    # Find the first non-nested table or use root
    first_val = next(iter(data.values()), None)
    if isinstance(first_val, dict) and not isinstance(first_val, dict):
        columns = data
        name = table_name or "table"
    elif isinstance(first_val, dict) and any(isinstance(v, dict) for v in first_val.values() if isinstance(v, dict)):
        # Nested tables - use first table
        for key, val in data.items():
            if isinstance(val, dict):
                columns = val
                name = table_name or key
                break
    else:
        columns = data
        name = table_name or "table"

    lines = [f"CREATE TABLE {name} ("]
    for col_name, col_def in columns.items():
        if isinstance(col_def, dict):
            col_type = str(col_def.get("type", "string")).lower()
            sql_type = TYPE_MAP.get(col_type, "TEXT")
            nullable = "NULL" if col_def.get("nullable", True) else "NOT NULL"
            pk = "PRIMARY KEY" if col_def.get("primary_key", False) or col_def.get("pk", False) else ""
            default = col_def.get("default")
            default_str = f"DEFAULT {default!r}" if default is not None else ""
            parts = [f"  {col_name} {sql_type}", nullable, pk, default_str]
            lines.append(", ".join(filter(None, parts)) + ",")
        else:
            sql_type = infer_sql_type(col_def)
            lines.append(f"  {col_name} {sql_type},")

    if lines[-1].endswith(","):
        lines[-1] = lines[-1][:-1]
    lines.append(");")
    return "\n".join(lines)


def main() -> None:
    parser = argparse.ArgumentParser(prog="toml-to-sql-schema", description="Generate SQL CREATE TABLE from TOML.")
    parser.add_argument("input", type=Path, help="Input TOML file")
    parser.add_argument("-o", "--output", type=Path, help="Output SQL file (default: stdout)")
    parser.add_argument("-t", "--table", type=str, help="Table name")
    args = parser.parse_args()

    toml_text = args.input.read_text(encoding="utf-8")
    try:
        data = loads(toml_text)
        sql = generate_schema(data, args.table)
    except (ValueError, Exception) as exc:
        print(f"Error: {exc}", file=sys.stderr)
        sys.exit(1)

    if args.output:
        args.output.write_text(sql + "\n", encoding="utf-8")
        print(f"Generated schema -> {args.output}", file=sys.stderr)
    else:
        print(sql)


if __name__ == "__main__":
    main()
