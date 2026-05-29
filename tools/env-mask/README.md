# env-mask

CLI tool that masks sensitive values in `.env` files. Reads an environment file and replaces sensitive values (passwords, secrets, tokens, API keys) with `[MASKED]` while preserving the file structure, comments, and non-sensitive values.

## Features

- Automatic detection of sensitive keys (passwords, secrets, tokens, API keys, credentials)
- Custom key filtering with `--keys` flag
- Custom regex patterns with `--pattern` flag
- Configurable mask string (`--mask`)
- Strip mode that sets sensitive values to empty (`--strip`)
- Summary mode that lists masked keys without outputting the file (`--show`)
- JSON output format (`--format json`)
- Preserves comments, blank lines, and file structure

## Install

```bash
go install github.com/TataneSan/env-mask@latest
```

Or build from source:

```bash
git clone https://github.com/TataneSan/env-mask.git
cd env-mask
go build -o env-mask .
```

## Usage

```
env-mask [flags] <file.env>
```

### Flags

| Flag | Default | Description |
|------|---------|-------------|
| `--keys` | | Comma-separated list of exact key names to mask |
| `--pattern` | | Regex pattern to match sensitive key names (overrides defaults) |
| `--defaults` | `true` | Use built-in sensitive key patterns |
| `--mask` | `[MASKED]` | Replacement string for masked values |
| `--format` | `env` | Output format: `env`, `json` |
| `--strip` | `false` | Strip values entirely (set to empty string) |
| `--show` | `false` | Show which keys were masked (summary mode) |
| `--version` | `false` | Print version |

## Examples

### Mask sensitive values with defaults

```bash
env-mask .env
```

Output:
```
# Database config
DB_HOST=localhost
DB_PORT=5432
DB_PASSWORD=[MASKED]
# API settings
API_KEY=[MASKED]
# Auth
JWT_SECRET=[MASKED]
# Public values
APP_NAME=MyApp
APP_ENV=production
```

### Mask specific keys only

```bash
env-mask --defaults=false --keys PASSWORD,TOKEN .env
```

### Mask by regex pattern

```bash
env-mask --defaults=false --pattern '.*(PASS|SECRET|KEY|TOKEN).*' .env
```

### Custom mask string

```bash
env-mask --mask '***' .env
```

### Strip sensitive values (set to empty)

```bash
env-mask --strip .env
```

Output:
```
DB_PASSWORD=
API_KEY=
JWT_SECRET=
```

### Show summary of masked keys

```bash
env-mask --show .env
```

Output:
```
File: .env
Sensitive keys found: 4

Masked keys:
  - DB_PASSWORD
  - API_KEY
  - JWT_SECRET
  - SESSION_SECRET
```

### JSON output

```bash
env-mask --format json .env
```

### Write masked output to a new file

```bash
env-mask .env > .env.masked
```

## Default sensitive key patterns

The tool detects the following key patterns by default (case-insensitive):

- `password`, `passwd`
- `secret`, `secret_key`, `secret-key`
- `token`
- `api_key`, `apikey`, `api-key`
- `private_key`, `private-key`
- `access_key`, `access-key`
- `auth`
- `credential`, `credentials`
- `db_pass`, `db_password`, `database_password`
- `mysql_password`, `postgres_password`, `redis_password`, `smtp_password`
- `encryption_key`, `signing_key`
- `jwt_secret`, `session_secret`, `cookie_secret`
- `aws_secret`, `gcp_key`, `azure_key`

## License

MIT
