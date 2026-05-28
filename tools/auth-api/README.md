# auth-api

REST API for JWT authentication with refresh tokens.

## Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | `/health` | Health check |
| POST | `/api/auth/login` | Login and get tokens |
| POST | `/api/auth/refresh` | Refresh access token |
| POST | `/api/auth/logout` | Revoke refresh token |

## Quick Start

### Docker

```bash
docker run -p 8080:8080 auth-api
```

### From source

```bash
go run ./cmd/main.go
```

Server starts on `:8080` by default. Override with `PORT` env var.

## Usage

### Login

```bash
curl -s -X POST http://localhost:8080/api/auth/login \
  -H 'Content-Type: application/json' \
  -d '{
    "user_id": "user123",
    "secret": "user123"
  }'
```

Response:
```json
{
  "access_token": "Bearer user123:1716900000.abc123",
  "refresh_token": "Bearer def456.xyz789"
}
```

### Refresh token

```bash
curl -s -X POST http://localhost:8080/api/auth/refresh \
  -H 'Content-Type: application/json' \
  -d '{
    "refresh_token": "Bearer def456.xyz789"
  }'
```

Response:
```json
{
  "access_token": "Bearer user123:1716900060.def456",
  "refresh_token": "Bearer ghi789.jkl012"
}
```

### Use access token

```bash
curl -s http://localhost:8080/api/protected \
  -H 'Authorization: Bearer user123:1716900060.def456'
```

### Logout

```bash
curl -s -X POST http://localhost:8080/api/auth/logout \
  -H 'Content-Type: application/json' \
  -d '{
    "refresh_token": "Bearer ghi789.jkl012"
  }'
```

Response:
```json
{"message": "token revoked"}
```

## Features

- JWT-style access tokens with automatic expiration
- Refresh token rotation (old token revoked on refresh)
- Refresh token revocation on logout
- In-memory token storage (suitable for single-instance deployments)
- Thread-safe token management

## Token Lifecycle

1. **Login**: Client sends credentials → receives access token + refresh token
2. **Access**: Client includes access token in `Authorization` header
3. **Refresh**: When access token expires, client sends refresh token → receives new pair
4. **Logout**: Client sends refresh token to revoke it

## Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| PORT | 8080 | Server port |

## License

MIT
