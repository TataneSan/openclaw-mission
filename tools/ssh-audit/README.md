# ssh-audit

Audit SSH server security configuration from the command line. Connects to a remote SSH server, performs a KEX_INIT exchange, and reports on the security of supported algorithms with color-coded ratings.

## Features

- Enumerates key exchange algorithms, ciphers, MACs, and host key types
- Color-coded security ratings: excellent, good, fair, weak, broken
- Detects AEAD ciphers, encrypt-then-MAC, and post-quantum KEX support
- Overall PASS/FAIL/WARNING verdict
- Recommendations for hardening sshd_config
- JSON output mode for scripting and CI pipelines
- No external dependencies (pure Go, stdlib only)

## Install

```bash
go install github.com/TataneSan/ssh-audit@latest
```

Or build from source:

```bash
git clone https://github.com/TataneSan/ssh-audit.git
cd ssh-audit
go build -o ssh-audit .
```

## Usage

```
ssh-audit [options] <host>[:port]
```

### Options

| Flag | Description |
|------|-------------|
| `-json` | Output results as JSON |
| `-p PORT`, `--port PORT` | Specify port (default: 22) |
| `-h`, `--help` | Show help message |

### Examples

Audit localhost:

```bash
ssh-audit localhost
```

Audit a specific host and port:

```bash
ssh-audit example.com:2222
```

Or with the port flag:

```bash
ssh-audit -p 2222 myserver.com
```

JSON output for CI/automation:

```bash
ssh-audit -json localhost
```

## Sample Output

```
╔══════════════════════════════════════════════════════════════╗
║                    SSH SERVER AUDIT                         ║
╚══════════════════════════════════════════════════════════════╝

Target:   localhost:22
Server:   OpenSSH_9.6
Response: 12ms

Overall:  WARNING (2 weak)

  AEAD Ciphers:     YES
  Encrypt-then-MAC: YES
  Post-Quantum KEX: NO

── Key Exchange Algorithms ───────────────────────────────────
  curve25519-sha256@libssh.org                 EXCELLENT  RFC 8731
  sntrup761x25519-sha512@openssh.com           EXCELLENT  post-quantum hybrid
  diffie-hellman-group16-sha512                GOOD       4096-bit MODP
  diffie-hellman-group-exchange-sha256         GOOD       2048-bit minimum
  diffie-hellman-group14-sha256                GOOD       2048-bit MODP
  diffie-hellman-group14-sha1                  WEAK       SHA-1 deprecated

── Ciphers ───────────────────────────────────────────────────
  chacha20-poly1305@openssh.com                EXCELLENT  AEAD
  aes256-gcm@openssh.com                       EXCELLENT  AEAD GCM
  aes128-gcm@openssh.com                       EXCELLENT  AEAD GCM
  aes256-ctr                                   GOOD       AES-256 CTR
  aes128-ctr                                   GOOD       AES-128 CTR
```

## Security Ratings

| Rating | Meaning |
|--------|---------|
| **Excellent** | Modern, strong algorithms (AEAD, post-quantum, Ed25519) |
| **Good** | Secure algorithms widely considered safe |
| **Fair** | Acceptable but with known caveats (NIST curves, CBC mode) |
| **Weak** | Deprecated algorithms (SHA-1, truncated MACs) |
| **Broken** | Must be disabled (RC4, 3DES, DSA, MD5) |

## Algorithm Databases

### Key Exchange
- Excellent: curve25519-sha256, sntrup761x25519 (post-quantum)
- Good: DH group14/16/18 with SHA-256/512, group-exchange-sha256
- Fair: ECDH NIST curves (P-256, P-384, P-521)
- Weak: DH group14-sha1, group-exchange-sha1
- Broken: DH group1-sha1 (1024-bit, Logjam)

### Ciphers
- Excellent: chacha20-poly1305, aes-*-gcm (AEAD)
- Good: aes-*-ctr
- Fair: aes-*-cbc (no AEAD)
- Weak: blowfish-cbc (Sweet32)
- Broken: 3des-cbc, arcfour, cast128-cbc

### MACs
- Excellent: hmac-sha2-*-etm (encrypt-then-MAC)
- Good: hmac-sha2-256/512, umac-128
- Fair: hmac-sha1-etm
- Weak: hmac-sha1, umac-64
- Broken: hmac-md5, hmac-md5-96, hmac-sha1-96

### Host Keys
- Excellent: ssh-ed25519
- Good: rsa-sha2-256/512, ecdsa-sha2-nistp*
- Weak: ssh-rsa (SHA-1)
- Broken: ssh-dss (DSA 1024-bit)

## License

MIT
