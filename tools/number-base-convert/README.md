# number-base-convert

A CLI tool for converting numbers between different bases: binary (2), octal (8), decimal (10), hexadecimal (16), base36, and base62.

## Installation

```bash
git clone git@github.com:TataneSan/number-base-convert.git
cd number-base-convert
go build -o number-base-convert .
```

Or install directly:

```bash
go install github.com/TataneSan/number-base-convert@latest
```

## Usage

```
number-base-convert [options] <number>
```

### Options

| Flag | Description |
|------|-------------|
| `-f, --from <base>` | Source base (2, 8, 10, 16, 36, 62) |
| `-t, --to <base>` | Target base (2, 8, 10, 16, 36, 62) |
| `-a, --all` | Show conversion to all supported bases |
| `-F, --format <fmt>` | Output format: `text` (default), `json`, `table` |
| `-b, --batch` | Batch mode: read numbers from stdin, one per line |
| `-h, --help` | Show help message |

### Supported Bases

| Base | Name | Aliases |
|------|------|---------|
| 2 | Binary | `2`, `binary`, `bin` |
| 8 | Octal | `8`, `octal`, `oct` |
| 10 | Decimal | `10`, `decimal`, `dec` |
| 16 | Hexadecimal | `16`, `hex`, `hexadecimal` |
| 36 | Base36 | `36`, `base36` |
| 62 | Base62 | `62`, `base62` |

## Examples

### Basic conversion

```bash
# Decimal to binary
number-base-convert 255 --from 10 --to 2
# Output: 255 (base 10) = 11111111 (base 2)

# Binary to hex
number-base-convert 1010 --from 2 --to 16
# Output: 1010 (base 2) = a (base 16)

# Hex to decimal
number-base-convert FF --from 16 --to 10
# Output: FF (base 16) = 255 (base 10)
```

### Auto-detect input base with prefixes

```bash
# Binary prefix (0b)
number-base-convert 0b1010 --to 10
# Output: 1010 (base 2) = 10 (base 10)

# Octal prefix (0o)
number-base-convert 0o77 --to 10
# Output: 77 (base 8) = 63 (base 10)

# Hex prefix (0x)
number-base-convert 0xFF --to 10
# Output: FF (base 16) = 255 (base 10)
```

### Convert to all bases

```bash
number-base-convert 42 --from 10 --all
```

Output:
```
Conversions of 42 (base 10):
  base  2 (binary ): 101010
  base  8 (octal  ): 52
  base 10 (decimal): 42
  base 16 (hex    ): 2a
  base 36 (base36 ): 16
  base 62 (base62 ): 42
```

### Negative numbers

```bash
number-base-convert -42 --from 10 --to 16
# Output: -42 (base 10) = -2a (base 16)

number-base-convert -0xFF --to 10
# Output: -FF (base 16) = -255 (base 10)
```

### Batch mode

```bash
echo -e "255\n4096\n1024" | number-base-convert --batch --from 10 --to 2
```

Output:
```
255 (base 10) = 11111111 (base 2)
4096 (base 10) = 1000000000000 (base 2)
1024 (base 10) = 10000000000 (base 2)
```

### Output formats

#### JSON

```bash
number-base-convert 255 --from 10 --to 2 --format json
```

Output:
```json
[
  {
    "input_number": "255",
    "input_base": 10,
    "output_number": "11111111",
    "output_base": 2
  }
]
```

#### Table

```bash
number-base-convert 255 --from 10 --to 2 --format table
```

Output:
```
+---------+-----------+--------------+-----------+
| Input   | From Base | Output       | To Base   |
+---------+-----------+--------------+-----------+
| 255     | 10        | 11111111     | 2         |
+---------+-----------+--------------+-----------+
```

## Base62 Character Set

Base62 uses the following character set:

```
0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz
```

## License

MIT License. See [LICENSE](LICENSE) for details.
