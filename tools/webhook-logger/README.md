# webhook-logger

Simple HTTP server that logs all incoming requests. Useful for testing webhooks locally.

## Features

- Logs all incoming HTTP requests (method, path, headers, body)
- Web UI to view logs in real-time
- REST API to retrieve logs programmatically
- Optional file logging for persistence
- Clear logs endpoint

## Installation

```bash
go install github.com/TataneSan/webhook-logger@latest
```

Or build from source:

```bash
git clone https://github.com/TataneSan/webhook-logger.git
cd webhook-logger
go build -o webhook-logger
```

## Usage

```bash
# Start server on default port (8080)
./webhook-logger

# Start on custom port
./webhook-logger -port 9090

# Log to file
./webhook-logger -log requests.log

# Both options
./webhook-logger -port 3000 -log webhook.log
```

## Endpoints

### Webhook Endpoints (logs requests)

- `POST /webhook/<any-path>` - Log any request
- `POST /hook/<any-path>` - Alternative webhook path

### API Endpoints

- `GET /api/logs` - List all logged requests
- `GET /api/log?id=1` - Get specific request by ID
- `POST /api/clear` - Clear all logs

### Web UI

- `GET /` - Web interface with real-time log viewing

## Examples

```bash
# Start the logger
./webhook-logger &

# Send a test webhook
curl -X POST http://localhost:8080/webhook/test \
  -H "Content-Type: application/json" \
  -d '{"event": "push", "repo": "my-project"}'

# View all logs via API
curl http://localhost:8080/api/logs | jq .

# View a specific request
curl http://localhost:8080/api/log?id=1 | jq .

# Clear all logs
curl -X POST http://localhost:8080/api/clear
```

## Response Format

All webhook endpoints return:

```json
{
  "status": "ok",
  "id": 1
}
```

## License

MIT
