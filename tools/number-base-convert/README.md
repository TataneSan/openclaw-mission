# number-base-convert

Convertit des entiers entre bases 2-64 (dec, hex, oct, bin, auto-détection de préfixes 0x/0o/0b).

## Installation

```bash
go build -o number-base-convert ./...
```

## Usage

```bash
# auto-détecte 255 comme décimal
number-base-convert 255

# convertit hex -> base 2
number-base-convert -from 16 -to 2 0xff

# batch
number-base-convert 100 0x64 0o144 0b1100100

# JSON
number-base-convert -json 255

# tronque en 8 bits (two's complement)
number-base-convert -w 8 -5
```

## Flags

| Flag | Description |
|------|-------------|
| `-from` | base d'entrée (0=auto, défaut auto) |
| `-to` | base de sortie (défaut 10) |
| `-w` | largeur en bits (mask two's complement) |
| `-json` | sortie JSON |

## Licence

MIT
