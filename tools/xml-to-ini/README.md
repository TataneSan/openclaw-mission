# xml-to-ini

Convert XML files to INI format. XML elements become INI sections, attributes and child elements become key-value pairs.

## Install

```bash
go install github.com/TataneSan/xml-to-ini@latest
```

Or build from source:

```bash
git clone https://github.com/TataneSan/xml-to-ini.git
cd xml-to-ini
go build -o xml-to-ini .
```

## Usage

```
xml-to-ini [options] <xml-file>
```

### Options

| Flag | Description |
|------|-------------|
| `-o, --output <file>` | Output file (default: same name with `.ini` extension) |
| `-h, --help` | Show help message |

### Examples

Convert a file (output gets `.ini` extension automatically):

```bash
xml-to-ini config.xml
# Creates config.ini
```

Specify output file:

```bash
xml-to-ini -o output.ini config.xml
```

Read from stdin:

```bash
cat data.xml | xml-to-ini -
```

## Conversion Rules

- The root XML element becomes the first INI section (its attributes become key-value pairs)
- Child elements **without** sub-children become key-value pairs in the current section
- Child elements **with** sub-children become new INI sections
- XML attributes are preserved as key-value pairs
- Sections and keys are sorted alphabetically in the output
- Hyphens, spaces, and dots in element names are converted to underscores

### Example

**Input (config.xml):**

```xml
<config version="1.0">
  <database>
    <host>localhost</host>
    <port>5432</port>
    <name>mydb</name>
  </database>
  <server port="8080">
    <host>0.0.0.0</host>
    <debug>true</debug>
  </server>
</config>
```

**Output (config.ini):**

```ini
[config]
version = 1.0

[database]
host = localhost
name = mydb
port = 5432

[server]
debug = true
host = 0.0.0.0
port = 8080
```

## License

MIT
