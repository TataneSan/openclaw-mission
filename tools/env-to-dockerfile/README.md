# env-to-dockerfile

Converts a .env file into Dockerfile ARG and ENV instructions.

Reads key=value pairs from a .env file and outputs Dockerfile-compatible ARG declarations for build-time and ENV declarations for runtime.

## Features

- Converts .env to Dockerfile ARG + ENV instructions
- Build-only (`-build`) or run-only (`-run`) modes
- Handles quoted values and comments
- Sanitizes keys to valid Docker ENV/ARG format
- Reads from file or stdin

## Install

```bash
go install github.com/TataneSan/env-to-dockerfile@latest
```

## Usage

```
env-to-dockerfile [flags] [file]
```

### Flags

| Flag | Default | Description |
|------|---------|-------------|
| `-build` | `false` | Output only ARG instructions (build-time) |
| `-run` | `false` | Output only ENV instructions (runtime) |

### Examples

Convert a .env file:

```bash
env-to-dockerfile
```

Use a custom file:

```bash
env-to-dockerfile config.env
```

Build-time only:

```bash
env-to-dockerfile -build
```

Runtime only:

```bash
env-to-dockerfile -run
```

### Input (.env)

```
DB_HOST=localhost
DB_PORT=5432
DB_NAME=myapp
API_KEY="secret-key-123"
```

### Output (default)

```dockerfile
# Build-time arguments
ARG DB_HOST
ARG DB_PORT
ARG DB_NAME
ARG API_KEY

# Runtime environment variables
ENV DB_HOST=${DB_HOST}
ENV DB_PORT=${DB_PORT}
ENV DB_NAME=${DB_NAME}
ENV API_KEY=${API_KEY}
```

## License

MIT
