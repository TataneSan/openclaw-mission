# json-path-get

Extract values from JSON documents using simple dot/bracket path expressions.

Think of it as a zero-dependency, mini `jq` for the 90% case: pull a value out of
a JSON blob without writing a script.

## Path syntax

| Expression | Meaning |
|------------|---------|
| `users.0.name` | Dot notation with numeric array index |
| `users[0].name` | Bracket index |
| `users[*].name` | Wildcard: every array element |
| `data.items[]` | Trailing `[]` expands an array |
| `a['key with spaces']` | Quoted dict keys |

## Install

```bash
pip install .
# or
pip install git+https://github.com/TataneSan/json-path-get.git
```

## Usage

```bash
json-path-get data.json users[0].name
# "alice"

cat api-response.json | json-path-get - data.items[*].id
# [ 101, 102, 103 ]

json-path-get config.json databases[0].host --raw
# db.internal

json-path-get data.json missing.path --default "N/A"
# N/A        (exit code 2)
```

### Examples

```bash
# first match only
json-path-get data.json 'tags[*]' --first --raw

# compact output for scripting
json-path-get data.json 'servers[*].ip' --compact

# negative index (last element)
json-path-get data.json 'history[-1]'
```

## Exit codes

| Code | Meaning |
|------|---------|
| 0 | Path found |
| 1 | I/O or CLI error (invalid JSON / path) |
| 2 | Path not found in the document |

## License

MIT
