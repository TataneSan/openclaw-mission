# text-acronym

[![Python](https://img.shields.io/badge/python-%E2%89%A53.9-blue)](https://www.python.org/)
[![License: MIT](https://img.shields.io/badge/license-MIT-green)](LICENSE)
[![zero deps](https://img.shields.io/badge/dependencies-zero-brightgreen)](https://pypi.org/)

Generate acronyms from phrases (`National Aeronautics and Space
Administration` -> `NASA`) and expand well-known acronyms back
(`NASA` -> `National Aeronautics and Space Administration`). English
and French stop words are skipped by default.

## Features

- Splits on spaces, hyphens, punctuation; first letters uppercased
- Stop-word filtering (of, the, de, et, ...) — disabled with `--all`
- Accent stripping for proper ASCII acronyms (`Réalité Virtuelle` -> `RV`)
- Built-in dictionary of ~60 common tech acronyms with `--reverse`
- `--list` dumps the whole dictionary
- Batch from arguments, file or stdin; JSON report; `-l` lowercase output
- Zero dependencies

## Install

```bash
pip install .
# or
pip install git+https://github.com/TataneSan/text-acronym.git
```

## Usage

```bash
# build an acronym
text-acronym National Aeronautics and Space Administration
# NASA

# French phrases work too, with french stop words
text-acronym Réalité Virtuelle
# RV
text-acronym Organisation des Nations Unies
# ONU

# include stop words as initials
text-acronym --all Zone Improvement Plan Office
# ZIPO

# expand known acronyms
text-acronym --reverse NASA YAML
# NASA  National Aeronautics and Space Administration
# YAML  YAML Ain't Markup Language

# batch from stdin
printf 'Command Line Interface\nSecure Shell\n' | text-acronym
# Command Line Interface  CLI
# Secure Shell  SSH

# dictionary
text-acronym --list

# JSON
text-acronym --json As Soon As Possible
```

## Exit codes

| Code | Meaning |
|------|---------|
| 0 | success |
| 1 | I/O or CLI error |
| 2 | one or more acronyms not found (`--reverse` mode) |

## License

MIT — see [LICENSE](LICENSE).
