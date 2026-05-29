# file-signature

Cryptographic file signature tool supporting Ed25519 and RSA keys. Sign files, verify integrity, and detect tampering.

## Features

- **Ed25519 keys** — fast, modern, compact signatures
- **RSA keys** — 2048/4096-bit support for compatibility
- **SHA-256 hashing** — secure digest before signing
- **PEM key format** — PKCS#8 private keys, PKIX public keys
- **Single binary** — no dependencies, cross-platform

## Installation

```bash
go install github.com/TataneSan/file-signature@latest
```

Or build from source:

```bash
git clone https://github.com/TataneSan/file-signature.git
cd file-signature
go build -o file-signature .
```

## Usage

```
file-signature <command> [options]
```

### Commands

| Command | Description |
|---------|-------------|
| `generate <type> <dir>` | Generate a new key pair (ed25519 or rsa) |
| `sign <file> <key> [out]` | Sign a file with a private key |
| `verify <file> <sig> <key>` | Verify a file signature |

### Examples

Generate an Ed25519 key pair:

```bash
file-signature generate ed25519 ./keys
```

Generate an RSA 4096-bit key pair:

```bash
file-signature generate rsa ./keys 4096
```

Sign a file:

```bash
file-signature sign archive.tar keys/private.pem
# Creates archive.tar.sig
```

Sign with custom output path:

```bash
file-signature sign archive.tar keys/private.pem archive.sig
```

Verify a signature:

```bash
file-signature verify archive.tar archive.tar.sig keys/public.pem
```

Exit codes: `0` = valid, `2` = invalid signature, `1` = usage error.

## Workflow

1. Generate a key pair on a secure machine
2. Share the public key with verifiers
3. Sign releases with the private key
4. Verifiers check signatures before use

```bash
# Creator
file-signature generate ed25519 ./keys
file-signature sign release-v1.0.tar.gz ./keys/private.pem

# Verifier (only needs public.pem)
file-signature verify release-v1.0.tar.gz release-v1.0.tar.gz.sig ./keys/public.pem
```

## Key Format

Keys are stored in standard PEM format:

- Private keys: PKCS#8 (`-----BEGIN PRIVATE KEY-----`)
- Public keys: PKIX (`-----BEGIN PUBLIC KEY-----`)

Private keys are written with `0600` permissions.

## License

MIT
