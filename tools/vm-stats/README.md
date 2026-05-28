# vm-stats

Collects and displays system statistics including CPU, memory, disk, and network usage.

## Features

- **CPU**: Usage breakdown (user, system, iowait, irq, nice), core count, load averages
- **Memory**: Total, used, free, available, buffers, cached, swap with usage percentages
- **Disk**: Filesystem usage per mount point with progress bars
- **Network**: Bytes transmitted/received per interface, error counts
- **Processes**: Total, running, sleeping, stopped, zombie counts
- **Multiple output formats**: Full report, compact one-line, JSON
- **Continuous monitoring**: Optional refresh interval for real-time stats

## Installation

```bash
go install github.com/TataneSan/vm-stats@latest
```

Or build from source:

```bash
git clone https://github.com/TataneSan/vm-stats.git
cd vm-stats
go build -o vm-stats .
```

## Usage

```
vm-stats [options]
```

### Options

| Flag | Description | Default |
|------|-------------|---------|
| `-format` | Output format: `full`, `compact`, `json` | `full` |
| `-interval` | Refresh interval in seconds | `1` |
| `-count` | Number of updates (`0` = infinite) | `1` |
| `-json`, `-j` | Output as JSON | `false` |
| `-compact`, `-c` | Compact one-line output | `false` |

### Examples

One-shot full report:
```bash
vm-stats
```

Monitor every 5 seconds:
```bash
vm-stats -interval 5 -count 0
```

Compact output (good for logging):
```bash
vm-stats -compact -interval 2
```

JSON output:
```bash
vm-stats -json
```

Run 10 iterations of compact output:
```bash
vm-stats -compact -interval 2 -count 10
```

## Output Example

```
14:32:05   CPU 12.3%  MEM 45.2%  LOAD 1.25 0.98 0.87  PROC 245  DISK 62.1%
14:32:07   CPU 8.1%   MEM 45.3%  LOAD 1.22 0.97 0.87  PROC 246  DISK 62.1%
```

## Requirements

- Linux (reads from `/proc` filesystem)
- Go 1.21+

## License

MIT
