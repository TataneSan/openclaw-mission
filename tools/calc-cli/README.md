# calc-cli

Scientific calculator for the command line.

**calc-cli** supports basic arithmetic, scientific functions, unit conversions, and interactive mode. Perfect for quick calculations in the terminal without leaving your workflow.

## Features

- Basic arithmetic: `+`, `-`, `*`, `/`
- Exponentiation: `**` or `^`
- Scientific functions: `sqrt`, `sin`, `cos`, `tan`, `log`, `exp`, and more
- Constants: `pi`, `e`
- Unit conversions: length, weight, temperature, data
- Interactive mode with prompt
- Expression parser with proper operator precedence

## Install

```bash
# From source
go install github.com/TataneSan/calc-cli@latest

# Or build manually
git clone https://github.com/TataneSan/calc-cli.git
cd calc-cli
go build -o calc-cli .
```

## Usage

### One-shot evaluation

```bash
calc-cli "2 + 3"              # 5
calc-cli "10 * 5 - 3"         # 47
calc-cli "2 ** 10"            # 1024
calc-cli "sqrt(16)"           # 4
calc-cli "sin(pi/2)"          # 1
calc-cli "log10(100)"         # 2
calc-cli "pi * 2"             # 6.283185307
```

### Unit conversion

```bash
calc-cli convert 100 c f      # 212 f
calc-cli convert 5 km mi      # 3.10686 mi
calc-cli convert 1 gb mb      # 1024 mb
calc-cli convert 72 f c       # 22.2222 c
calc-cli convert 10 kg lb     # 22.0462 lb
```

### Interactive mode

```bash
calc-cli
> 2 + 3
= 5
> sqrt(16)
= 4
> sin(pi/2)
= 1
> quit
```

## Supported Operations

| Operator | Description |
|----------|-------------|
| `+` | Addition |
| `-` | Subtraction |
| `*` | Multiplication |
| `/` | Division |
| `**` or `^` | Exponentiation |
| `()` | Grouping |

## Functions

| Function | Description |
|----------|-------------|
| `sqrt(x)` | Square root |
| `sin(x)` | Sine (radians) |
| `cos(x)` | Cosine (radians) |
| `tan(x)` | Tangent (radians) |
| `asin(x)` | Arc sine |
| `acos(x)` | Arc cosine |
| `atan(x)` | Arc tangent |
| `log(x)` | Natural logarithm |
| `log2(x)` | Base-2 logarithm |
| `log10(x)` | Base-10 logarithm |
| `exp(x)` | Exponential (e^x) |
| `abs(x)` | Absolute value |
| `ceil(x)` | Ceiling |
| `floor(x)` | Floor |

## Constants

| Constant | Value |
|----------|-------|
| `pi` | 3.14159... |
| `e` | 2.71828... |

## Unit Conversions

### Length
`mm`, `cm`, `m`, `km`, `in`, `ft`, `yd`, `mi`

### Weight
`mg`, `g`, `kg`, `t`, `oz`, `lb`

### Temperature
`c` (Celsius), `f` (Fahrenheit), `k` (Kelvin)

### Data
`b`, `kb`, `mb`, `gb`, `tb`

## Examples

```bash
# Compound expressions
calc-cli "(2 + 3) * 4"        # 20
calc-cli "2 ** 10 / 1024"     # 1

# Trigonometry with constants
calc-cli "cos(pi)"             # -1
calc-cli "tan(pi/4)"           # 1

# Nested functions
calc-cli "log2(2 ** 10)"      # 10

# Practical calculations
calc-cli "100 * 0.08 + 100"   # 108 (8% tip on $100)
calc-cli "500 / 12"           # 41.6667 (monthly payment)
```

## License

MIT
