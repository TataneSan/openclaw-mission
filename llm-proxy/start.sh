#!/bin/bash

# OpenClaw LLM Proxy - Startup Script
# Place in /etc/init.d/ or use with systemd

PROXY_DIR="/root/openclaw/llm-proxy"
LOG_FILE="$PROXY_DIR/server.log"
PID_FILE="$PROXY_DIR/server.pid"

start() {
    echo "🦞 Starting OpenClaw LLM Proxy..."
    cd "$PROXY_DIR"

    if [ -f "$PID_FILE" ] && kill -0 $(cat "$PID_FILE") 2>/dev/null; then
        echo "⚠️  Server already running (PID: $(cat $PID_FILE))"
        return 1
    fi

    nohup node src/server.js > "$LOG_FILE" 2>&1 &
    echo $! > "$PID_FILE"
    sleep 2

    if kill -0 $(cat "$PID_FILE") 2>/dev/null; then
        echo "✅ Server started (PID: $(cat $PID_FILE))"
        echo "   Port: $(grep PORT .env | cut -d= -f2)"
        echo "   Log:  $LOG_FILE"
    else
        echo "❌ Failed to start server"
        cat "$LOG_FILE"
        return 1
    fi
}

stop() {
    echo "🛑 Stopping OpenClaw LLM Proxy..."

    if [ ! -f "$PID_FILE" ]; then
        echo "⚠️  No PID file found"
        return 1
    fi

    PID=$(cat "$PID_FILE")
    if kill -0 "$PID" 2>/dev/null; then
        kill "$PID"
        rm -f "$PID_FILE"
        echo "✅ Server stopped"
    else
        echo "⚠️  Process not running"
        rm -f "$PID_FILE"
    fi
}

restart() {
    stop
    sleep 1
    start
}

status() {
    if [ -f "$PID_FILE" ] && kill -0 $(cat "$PID_FILE") 2>/dev/null; then
        echo "✅ Server running (PID: $(cat $PID_FILE))"
        curl -s http://localhost:$(grep PORT .env | cut -d= -f2)/health | jq .
    else
        echo "❌ Server not running"
    fi
}

case "$1" in
    start)   start ;;
    stop)    stop ;;
    restart) restart ;;
    status)  status ;;
    *)
        echo "Usage: $0 {start|stop|restart|status}"
        exit 1
        ;;
esac
