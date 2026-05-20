# OpenClaw - Quick Reference Card

## 🦞 Service Status
- **URL:** http://localhost:8088
- **Port:** 8088
- **Status:** 🟢 Operational

## 💰 Crypto Wallets

### BTC
```
1GXpyCzAhvVNcbjnkytRTe1QoVoAen7nfP
```

### ETH
```
0xCF295d87E1534538FDac3b1f98746aF4A3E47352
```

## 🔑 Admin Access
```bash
# Admin key
export ADMIN_KEY=admin_openclaw_2026

# Quick commands
curl -H "X-Admin-Key: $ADMIN_KEY" http://localhost:8088/admin/stats
```

## 📊 Management Commands

```bash
cd /root/openclaw/llm-proxy

# Create client
node manage.js create-client <name> <email> <credits>

# Add credits
node manage.js add-credits <prefix> <amount> [tx-hash] [BTC|ETH]

# List clients
node manage.js list-keys

# Statistics
node manage.js stats

# Health check
node manage.js health
```

## 🧪 Test Commands

```bash
# Health
curl http://localhost:8088/health

# Models
curl http://localhost:8088/v1/models

# Create test key
curl -X POST http://localhost:8088/admin/keys \
  -H "X-Admin-Key: admin_openclaw_2026" \
  -d '{"client_name":"test","initial_credits":10000}'

# Chat completion (replace KEY)
curl -X POST http://localhost:8088/v1/chat/completions \
  -H "Authorization: Bearer oc_KEY" \
  -d '{"model":"deepseek-v4-flash","messages":[{"role":"user","content":"Hi"}]}'
```

## 📁 Important Files

```
/root/openclaw/
├── README.md                    # Main documentation
├── revenue-log.md               # Detailed revenue tracking
├── QUICKREF.md                  # This file
├── wallets/
│   └── wallets.json             # Crypto wallets (chmod 600)
└── llm-proxy/
    ├── .env                     # Configuration
    ├── server.log               # Server logs
    ├── manage.js                # CLI tool
    ├── start.sh                 # Startup script
    ├── src/
    │   ├── server.js            # Main server
    │   ├── db.js                # Database layer
    │   └── proxy.js             # LLM proxy logic
    └── data/
        └── proxy.db             # SQLite database
```

## 💵 Pricing (satoshis per 1k tokens)

| Model | Prompt | Completion |
|-------|--------|------------|
| deepseek-v4-pro | 50 | 150 |
| deepseek-v4-flash | 20 | 60 |
| kimi-k2.6 | 40 | 120 |
| minimax-m2.7 | 30 | 90 |
| mimo-v2.5-pro | 45 | 135 |
| skyclaw-v1 | 35 | 105 |

## 🔄 Restart Service

```bash
# Kill existing
pkill -f "node src/server.js"

# Start fresh
cd /root/openclaw/llm-proxy
nohup node src/server.js > server.log 2>&1 &

# Or use startup script
./start.sh restart
```

## 📈 Revenue Tracking

```bash
# Total revenue (satoshis)
sqlite3 /root/openclaw/llm-proxy/data/proxy.db \
  "SELECT SUM(total_spent_satoshis) FROM api_keys;"

# Active clients
sqlite3 /root/openclaw/llm-proxy/data/proxy.db \
  "SELECT COUNT(*) FROM api_keys WHERE is_active=1;"

# Usage by model
sqlite3 /root/openclaw/llm-proxy/data/proxy.db \
  "SELECT model, SUM(tokens_prompt+tokens_completion) FROM usage_log GROUP BY model;"
```

---
**Generated:** 2026-05-20
**Agent:** OpenClaw
