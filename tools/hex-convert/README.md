# hex-convert

Convert numbers between hexadecimal, decimal, binary, and octal formats.

## Features

- Convert between hex, decimal, binary, and octal
- Auto-detect input format (0x, 0b, 0o prefixes)
- Show all formats at once with `all` command
- Support for negative numbers
- Single binary, no dependencies

## Install

```bash
go install github.com/TataneSan/hex-convert@latest
```

Or build from source:

```bash
git clone https://github.com/TataneSan/hex-convert.git
cd hex-convert
go build -o hex-convert .
cp hex-convert /usr/local/bin/
```

## Usage

```
hex-convert <number> [--from fmt] [--to fmt]
hex-convert all <number> [--from fmt]
```

### Examples

Show all formats:
```bash
hex-convert all 255
```

Convert hex to decimal:
```bash
hex-convert FF --from hex --to dec
```

Convert decimal to hex:
```bash
hex-convert 42 --to hex
```

Convert binary to decimal:
```bash
hex-convert 101010 --from bin --to dec
```

Auto-detect with prefixes:
```bash
hex-convert 0xFF        # hex
hex-convert 0b1010      # binary
hex-convert 0o77        # octal
```

### Output

```
$ hex-convert all 255
  ╔══════════════════════════════════════╗
  ║          NUMBER CONVERSION           ║
  ╠══════════════════════════════════════╣
  ║  Input: 255 (dec)                    ║
  ╠══════════════════════════════════════╣
  ║  HEX:   0xff                         ║
  ║  DEC:   255                          ║
  ║  OCT:   0o377                        ║
  ║  BIN:   0b11111111                   ║
  ╚══════════════════════════════════════╝
```

## Formats

| Format | Flag  | Characters | Prefix |
|--------|-------|------------|--------|
| Hex    | hex   | 0-9, a-f   | 0x     |
| Dec    | dec   | 0-9        | —      |
| Bin    | bin   | 0-1        | 0b     |
| Oct    | oct   | 0-7        | 0o     |

## Requirements

- Go 1.21+

## License

MIT
