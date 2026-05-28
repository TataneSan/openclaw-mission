# json-to-python

Generate Python dataclasses from JSON input.

## Features

- Infers Python types from JSON values (str, int, float, bool, List, Dict)
- Handles nested objects and arrays
- Generates `from_dict` class method for easy deserialization
- Python keyword and reserved name handling
- Reads from file or stdin

## Install

```bash
go install github.com/TataneSan/json-to-python@latest
```

## Usage

```bash
json-to-python [flags] <input.json>
```

### Flags

| Flag | Description | Default |
|------|-------------|---------|
| `-o` | Output file (default: stdout) | |
| `-n` | Dataclass name | `Model` |

### Examples

```bash
# Basic usage
json-to-python user.json

# Custom class name and output file
json-to-python user.json -n User -o models.py

# Pipe from stdin
cat user.json | json-to-python -
```

### Input

```json
{
    "name": "Alice",
    "age": 30,
    "active": true,
    "tags": ["dev", "go"],
    "address": {"city": "Paris"}
}
```

### Output

```python
from __future__ import annotations
from dataclasses import dataclass, field
from typing import Any, Dict, List


@dataclass
class User:
    active: bool = None
    address: Dict[str, Any] = field(default_factory=dict)
    age: int = None
    name: str = None
    tags: List[str] = field(default_factory=list)

    @classmethod
    def from_dict(cls, data: dict) -> User:
        return cls(**data)
```

## License

MIT
