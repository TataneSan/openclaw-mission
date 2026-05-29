# env-mask

Mask sensitive values in `.env` files before sharing or logging. Detects passwords, secrets, tokens, and keys by pattern matching and replaces their values with a placeholder.

## Features

- **Auto-detection** of 25+ sensitive key patterns (password, secret, token, api_key, etc.)
- **Case-insensitive** matching
- **Configurable mask style**: `****`, `••••`, or `<masked>`
- **Whitelist** keys to skip masking
- **Custom patterns** for project-specific sensitive keys
- **Stdin support** — pipe any `.env` content through the tool
- **Preserves** comments, empty lines, and non-sensitive keys
- **Supports** both `KEY=VALUE` and `export KEY=VALUE` formats
- Single binary, no dependencies

## Install

```bash
go install github.com/TataneSan/env-mask@latest
```

Or build from source:

```bash
git clone https://github.com/TataneSan/env-mask.git
cd env-mask
go build -o env-mask .
cp env-mask /usr/local/bin/
```

## Usage

```
env-mask [OPTIONS] [FILE]
cat .env | env-mask [OPTIONS]
```

### Options

| Flag | Description |
|------|-------------|
| `-p, --pattern PATTERNS` | Comma-separated sensitive key patterns |
| `-w, --whitelist KEYS` | Comma-separated keys to skip masking |
| `-s, --style STYLE` | Mask style: `asterisk`, `dots`, `placeholder` |
| `-h, --help` | Show help |

### Examples

```bash
# Mask sensitive values in .env
env-mask .env

# Use placeholder style
env-mask -s placeholder .env

# Pipe from stdin
cat .env | env-mask

# Whitelist specific keys
env-mask -w "DB_PASSWORD,API_KEY" .env

# Custom patterns
env-mask -p "secret,key,token" .env
```

### Example

Given this `.env` file:

```
DB_HOST=localhost
DB_PORT=5432
DB_PASSWORD=supers3cret
API_KEY=abc123xyz
APP_NAME=myapp
SECRET_TOKEN=tok_98765
export AWS_SECRET_KEY=wJalrXUtn
```

Running `env-mask .env` outputs:

```
DB_HOST=localhost
DB_PORT=5432
DB_PASSWORD=****
API_KEY=****
APP_NAME=myapp
SECRET_TOKEN=****
export AWS_SECRET_KEY=****
```

## Default Sensitive Patterns

The tool matches keys containing any of these patterns (case-insensitive):

password, passwd, pass, secret, token, api_key, apikey, api-key, private_key, privatekey, private-key, access_key, accesskey, access-key, secret_key, secretkey, secret-key, auth, credential, credentials, encryption_key, signing_key, jwt_secret, db_password, database_password, smtp_password, email_password, aws_secret, gcp_secret, azure_secret, ssh_key, ssh-key, cookie_secret, session_secret, master_key, masterkey, client_secret, clientsecret

## License

MIT
