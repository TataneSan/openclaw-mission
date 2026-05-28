# config-validator

Validate configuration files in multiple formats from the command line. Supports JSON, YAML, TOML, and INI with strict mode checks and JSON output for CI integration.

## Features

- **Multi-format support**: JSON, YAML, TOML, INI
- **Strict mode**: Additional checks for style issues (trailing whitespace, tabs in YAML, keys with spaces)
- **Directory scanning**: Recursively find and validate all config files in a directory
- **JSON output**: Machine-readable output for CI/CD pipelines
- **Exit codes**: Returns 1 if any file fails validation (CI-friendly)

## Installation

### From source

```bash
go install github.com/TataneSan/config-validator@latest
```

### Build manually

```bash
git clone https://github.com/TataneSan/config-validator.git
cd config-validator
go build -o config-validator .
```

## Usage

```
config-validator [options] <file|directory>
```

### Options

| Flag | Description |
|------|-------------|
| `-r` | Recursively search directories for config files |
| `-s` | Enable strict validation with additional style checks |
| `-o` | Output format: `text` (default) or `json` |
| `-e` | Comma-separated list of extensions to check |
| `-h` | Show help message |

### Examples

Validate a single file:

```bash
config-validator config.json
config-validator config.yaml
config-validator settings.toml
```

Strict validation with style checks:

```bash
config-validator -s config.yaml
```

Scan a directory recursively:

```bash
config-validator -r ./configs
```

Limit to specific file types:

```bash
config-validator -r -e .json,.yaml ./
```

JSON output for CI:

```bash
config-validator -o json config.json
```

## Output

### Text mode (default)

```
Config Validator - 3 files checked
Valid: 2 | Invalid: 1

[FAIL] config/broken.json (json)
  ERROR: invalid JSON: unexpected end of JSON input
[PASS] config/app.yaml (yaml)
[PASS] config/settings.toml (toml)
```

### JSON mode

```json
[
  {
    "file": "config/broken.json",
    "format": "json",
    "valid": false,
    "errors": ["invalid JSON: unexpected end of JSON input"]
  },
  {
    "file": "config/app.yaml",
    "format": "yaml",
    "valid": true
  }
]
```

## Strict Mode Checks

| Format | Checks |
|--------|--------|
| JSON | Keys with spaces, trailing commas |
| YAML | Tab characters (error), trailing whitespace |
| TOML | Duplicate sections, trailing whitespace |
| INI | Unclosed section headers |

## CI Integration

```yaml
# GitHub Actions example
- name: Validate configs
  run: |
    wget -O config-validator https://github.com/TataneSan/config-validator/releases/latest/download/config-validator-linux-amd64
    chmod +x config-validator
    ./config-validator -r -s ./config
```

## Supported Formats

| Extension | Format | Validation |
|-----------|--------|------------|
| `.json` | JSON | Full parse validation |
| `.yaml`, `.yml` | YAML | Full parse validation |
| `.toml` | TOML | Syntax validation |
| `.ini`, `.cfg`, `.conf` | INI | Full parse validation |

## License

MIT
