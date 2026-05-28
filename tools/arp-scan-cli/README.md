# arp-scan-cli

Linux ARP table viewer and network scanner. View, scan, watch, and manage ARP entries with vendor detection.

## Features

- **show** — Display the ARP table in a formatted table
- **scan** — Scan a network for active hosts via ARP
- **watch** — Monitor ARP table changes in real-time
- **get** — Lookup MAC address for a specific IP
- **export** — Export ARP table to JSON, CSV, or text
- **add** — Add a static ARP entry (requires root)
- **delete** — Remove an ARP entry (requires root)
- Vendor detection from OUI database (VMware, Intel, Cisco, etc.)

## Installation

```bash
go install github.com/tatanesan/arp-scan-cli@latest
```

Or build from source:

```bash
git clone https://github.com/TataneSan/arp-scan-cli.git
cd arp-scan-cli
go build -o arp-scan-cli .
sudo cp arp-scan-cli /usr/local/bin/
```

## Usage

```
arp-scan-cli <command> [arguments]
```

### Show ARP table

```bash
# Show all entries
arp-scan-cli show

# Filter by interface
arp-scan-cli show eth0
arp-scan-cli ls wlan0
```

### Scan a network

```bash
arp-scan-cli scan 192.168.1.0/24
```

Uses `arp-scan` if available, otherwise falls back to ping sweep.

### Watch for changes

```bash
# Watch all interfaces
arp-scan-cli watch

# Watch specific interface
arp-scan-cli watch eth0

# Custom interval (3 seconds)
arp-scan-cli watch -i 3 eth0
```

### Lookup by IP

```bash
arp-scan-cli get 192.168.1.1
```

### Export

```bash
arp-scan-cli export json
arp-scan-cli export csv
arp-scan-cli export text
```

### Add static entry (requires root)

```bash
sudo arp-scan-cli add 192.168.1.100 aa:bb:cc:dd:ee:ff
sudo arp-scan-cli add 192.168.1.100 aa:bb:cc:dd:ee:ff -i eth0
```

### Delete entry (requires root)

```bash
sudo arp-scan-cli del 192.168.1.100
```

## Output

### Show

```
IP ADDRESS       MAC ADDRESS        INTERFACE    VENDOR     STATUS
----------------------------------------------------------------------
192.168.1.1      00:1B:21:AA:BB:CC  eth0         Intel      Active
192.168.1.10     00:0C:29:DD:EE:FF  eth0         VMware     Active
192.168.1.20     incomplete         eth0         Unknown    Incomplete

3 entry(ies)
```

### Watch

```
Watching ARP table (interval: 2.0s)...

[+] 192.168.1.50 -> aa:bb:cc:dd:ee:ff (eth0)
[-] 192.168.1.30 removed

[14:23:45] 12 entries (no changes)
```

## Vendor Detection

Includes a built-in OUI database for common vendors:
- VMware, VirtualBox, QEMU/KVM
- Intel, Realtek, TP-Link
- Cisco, Dell, HP, Apple
- ASUS, Netgear, Synology
- And more

## Notes

- `scan`, `add`, and `delete` require root privileges
- `show`, `watch`, `get`, and `export` work without root
- Reads from `ip neigh` (preferred) or `/proc/net/arp` (fallback)
- Press `Ctrl+C` to stop watching

## Requirements

- Linux kernel with ARP support
- iproute2 (recommended) or /proc/net/arp
- Root privileges for modification commands
- Go 1.21+ to build from source

## License

MIT
