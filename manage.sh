#!/bin/bash
# OpenClaw Services Manager
# Usage: ./manage.sh [start|stop|restart|status]

OPENCLAW_DIR="/root/openclaw"

start_proxy() {
    echo "Starting LLM Proxy..."
    cd "$OPENCLAW_DIR/llm-proxy" && bash start.sh start
}

start_scraper() {
    echo "Starting Web Scraper on port 8089..."
    cd "$OPENCLAW_DIR/web-scraper"
    if [ -f scraper.pid ] && kill -0 $(cat scraper.pid) 2>/dev/null; then
        echo "Web Scraper already running"
        return
    fi
    nohup node server.js > scraper.log 2>&1 &
    echo $! > scraper.pid
    sleep 1
    if kill -0 $(cat scraper.pid) 2>/dev/null; then
        echo "Web Scraper started (PID: $(cat scraper.pid))"
    else
        echo "Failed to start Web Scraper"
    fi
}

start_bot() {
    if [ ! -f "$OPENCLAW_DIR/telegram-bot/.env" ] || grep -q "YOUR_BOT_TOKEN_HERE" "$OPENCLAW_DIR/telegram-bot/.env"; then
        echo "Telegram Bot: No token configured. Edit telegram-bot/.env"
        return
    fi
    echo "Starting Telegram Bot..."
    cd "$OPENCLAW_DIR/telegram-bot"
    if [ -f bot.pid ] && kill -0 $(cat bot.pid) 2>/dev/null; then
        echo "Telegram Bot already running"
        return
    fi
    nohup node bot.js > bot.log 2>&1 &
    echo $! > bot.pid
    sleep 1
    if kill -0 $(cat bot.pid) 2>/dev/null; then
        echo "Telegram Bot started (PID: $(cat bot.pid))"
    else
        echo "Failed to start Telegram Bot"
    fi
}

stop_proxy() {
    echo "Stopping LLM Proxy..."
    cd "$OPENCLAW_DIR/llm-proxy" && bash start.sh stop
}

stop_scraper() {
    echo "Stopping Web Scraper..."
    if [ -f "$OPENCLAW_DIR/web-scraper/scraper.pid" ]; then
        kill $(cat "$OPENCLAW_DIR/web-scraper/scraper.pid") 2>/dev/null
        rm -f "$OPENCLAW_DIR/web-scraper/scraper.pid"
        echo "Web Scraper stopped"
    fi
}

stop_bot() {
    echo "Stopping Telegram Bot..."
    if [ -f "$OPENCLAW_DIR/telegram-bot/bot.pid" ]; then
        kill $(cat "$OPENCLAW_DIR/telegram-bot/bot.pid") 2>/dev/null
        rm -f "$OPENCLAW_DIR/telegram-bot/bot.pid"
        echo "Telegram Bot stopped"
    fi
}

start_converter() {
    echo "Starting PDF Converter on port 8090..."
    systemctl start openclaw-converter.service
    sleep 1
    if systemctl is-active --quiet openclaw-converter.service; then
        echo "PDF Converter started"
    else
        echo "Failed to start PDF Converter"
    fi
}

stop_converter() {
    echo "Stopping PDF Converter..."
    systemctl stop openclaw-converter.service
    echo "PDF Converter stopped"
}

start_image() {
    echo "Starting Image API on port 8091..."
    systemctl start openclaw-image.service
    sleep 1
    if systemctl is-active --quiet openclaw-image.service; then
        echo "Image API started"
    else
        echo "Failed to start Image API"
    fi
}

stop_image() {
    echo "Stopping Image API..."
    systemctl stop openclaw-image.service
    echo "Image API stopped"
}

status() {
    echo "=== OpenClaw Services Status ==="
    echo ""
    # LLM Proxy
    if curl -s http://localhost:8088/health > /dev/null 2>&1; then
        echo "[ONLINE]  LLM Proxy        - port 8088"
    else
        echo "[OFFLINE] LLM Proxy        - port 8088"
    fi
    # Web Scraper
    if curl -s http://localhost:8089/health > /dev/null 2>&1; then
        echo "[ONLINE]  Web Scraper      - port 8089"
    else
        echo "[OFFLINE] Web Scraper      - port 8089"
    fi
    # PDF Converter
    if curl -s http://localhost:8090/health > /dev/null 2>&1; then
        echo "[ONLINE]  PDF Converter    - port 8090"
    else
        echo "[OFFLINE] PDF Converter    - port 8090"
    fi
    # Image API
    if curl -s http://localhost:8091/health > /dev/null 2>&1; then
        echo "[ONLINE]  Image API        - port 8091"
    else
        echo "[OFFLINE] Image API        - port 8091"
    fi
    # Telegram Bot
    if [ -f "$OPENCLAW_DIR/telegram-bot/bot.pid" ] && kill -0 $(cat "$OPENCLAW_DIR/telegram-bot/bot.pid") 2>/dev/null; then
        echo "[ONLINE]  Telegram Bot     - PID $(cat $OPENCLAW_DIR/telegram-bot/bot.pid)"
    else
        echo "[OFFLINE] Telegram Bot     - no token configured"
    fi
    echo ""
}

case "$1" in
    start)
        start_proxy
        start_scraper
        start_converter
        start_image
        start_bot
        echo ""
        status
        ;;
    stop)
        stop_proxy
        stop_scraper
        stop_converter
        stop_image
        stop_bot
        ;;
    restart)
        stop_proxy
        stop_scraper
        stop_converter
        stop_image
        stop_bot
        sleep 2
        start_proxy
        start_scraper
        start_converter
        start_image
        start_bot
        echo ""
        status
        ;;
    status)
        status
        ;;
    *)
        echo "Usage: $0 {start|stop|restart|status}"
        exit 1
        ;;
esac
