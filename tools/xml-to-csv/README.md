# xml-to-csv

Convert XML files to CSV format.

## Install

```bash
pip install xml-to-csv
```

## Usage

```bash
xml-to-csv data.xml
xml-to-csv data.xml -o output.csv
xml-to-csv data.xml --tag record
```

## Examples

**Input** (`users.xml`):
```xml
<users>
  <user>
    <name>Alice</name>
    <age>30</age>
  </user>
  <user>
    <name>Bob</name>
    <age>25</age>
  </user>
</users>
```

**Output** (`users.csv`):
```
name,age
Alice,30
Bob,25
```

## License

MIT
