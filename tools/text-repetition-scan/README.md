# text-repetition-scan

Proofreading helper: finds doubled words (`the the`) and repeated word
n-grams in prose. Lint-style exit code for CI.

## Features

- Doubled-word detection with file:line:col locations
- Optional n-gram repetition report: `--ngrams 3 --min-count 2`
- Case-insensitive matching, punctuation-aware tokens
- JSON output, multi-file, stdin
- Exit 2 when anything is found — wire it into docs CI
- Pure Python standard library, no dependencies

## Install

```bash
pip install .
pip install git+https://github.com/TataneSan/text-repetition-scan.git
```

## Usage

```bash
$ printf 'This is is a test.\nThe cat sat. The cat sat.\n' | text-repetition-scan - --ngrams 3
<stdin>:1:9: doubled-word: 'is is'
<stdin>: repeated-3gram: 'the cat sat' x2
$ echo $?
2

$ text-repetition-scan docs/*.md || echo 'fix repetitions'
```

## Exit codes

| Code | Meaning |
|------|---------|
| 0 | Clean |
| 1 | I/O or CLI error |
| 2 | Repetitions found |

## License

MIT
