# text-frequency-cli

A command-line word frequency analyzer for text files and stdin input.

Analyzes text to find the most frequently used words, with visual bar charts and percentage breakdowns.

## Install

```bash
go install github.com/TataneSan/text-frequency-cli@latest
```

Or build from source:

```bash
git clone https://github.com/TataneSan/text-frequency-cli.git
cd text-frequency-cli
go build -o text-frequency-cli .
```

## Usage

```
text-frequency-cli [options] [file]
```

## Options

| Flag | Description |
|------|-------------|
| `-n N`, `--top N` | Show top N words (default: 20) |
| `-h`, `--help` | Show help message |

If no file is provided, reads from stdin.

## Examples

```bash
# Analyze a file
text-frequency-cli document.txt

# Show top 10 words
text-frequency-cli -n 10 report.md

# Pipe text from another command
cat file.txt | text-frequency-cli

# Redirect input
text-frequency-cli -n 5 < input.txt
```

## Output Example

```
Word Frequency Analysis
Total words: 20 | Unique words: 11
--------------------------------------------------
   4  the    ######################################## (20.0%)
   2  dog    #################### (10.0%)
   2  fox    #################### (10.0%)
   2  lazy   #################### (10.0%)
   2  quick  #################### (10.0%)
   2  very   #################### (10.0%)
   2  was    #################### (10.0%)
   1  and    ########## (5.0%)
   1  brown  ########## (5.0%)
   1  jumps  ########## (5.0%)

  ... and 1 more unique words
```

## Features

- Word frequency counting with case-insensitive matching
- Visual bar chart output
- Percentage breakdown of each word
- Supports file input or stdin piping
- Configurable number of top results
- Sorted by frequency (descending), then alphabetically

## License

MIT
