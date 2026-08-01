# OpenClaw Revenue Engine - Log

## Table of Contents
1. [Project Overview](#project-overview)
2. [Crypto Wallets](#crypto-wallets)
3. [LLM Proxy Service](#llm-proxy-service)
4. [Pricing](#pricing)
5. [Client Management](#client-management)
6. [Revenue Tracking](#revenue-tracking)
7. [Deployment](#deployment)

---

## Project Overview

**Mission:** Generate revenue through crypto payments by providing valuable services.

**Primary Revenue Stream:** LLM API Proxy Service
- Resells access to llm.contes.me models
- Pay-per-use pricing in satoshis
- API key based authentication
- Real-time usage tracking and billing

**Started:** 2026-05-20

---

## Crypto Wallets

### Bitcoin (BTC)
- **Address:** `1GXpyCzAhvVNcbjnkytRTe1QoVoAen7nfP`
- **WIF:** Stored securely in `wallets/wallets.json` (chmod 600)
- **Network:** Mainnet

### Ethereum (ETH)
- **Address:** `0xCF295d87E1534538FDac3b1f98746aF4A3E47352`
- **Private Key:** Stored securely in `wallets/wallets.json` (chmod 600)
- **Network:** Mainnet

### Payment Instructions for Clients
To add credits to your account, send BTC or ETH to the addresses above and include your API key prefix in the transaction memo/notes. Credits will be added within 1 hour of confirmation.

---

## LLM Proxy Service

### Service Details
- **URL:** `http://localhost:8088` (local) / `http://YOUR_SERVER_IP:8088` (public)
- **Port:** 8088
- **Protocol:** HTTP (HTTPS recommended for production via nginx/caddy)
- **Database:** SQLite (`data/proxy.db`)

### Architecture
```
Client Request → API Key Validation → Credit Check → Proxy to llm.contes.me → Response
                      ↓                    ↓                    ↓
               Rate Limiting         Credit Deduction      Usage Logging
```

### Endpoints

#### Public Endpoints
| Method | Path | Description |
|--------|------|-------------|
| GET | `/health` | Health check |
| GET | `/v1/models` | List available models |
| GET | `/v1/pricing` | Get pricing information |

#### Client Endpoints (require API key)
| Method | Path | Description |
|--------|------|-------------|
| POST | `/v1/chat/completions` | Chat completions (OpenAI-compatible) |
| GET | `/v1/account` | Account balance and usage |
| GET | `/v1/usage` | Detailed usage statistics |

#### Admin Endpoints (require admin key)
| Method | Path | Description |
|--------|------|-------------|
| POST | `/admin/keys` | Create new API key |
| GET | `/admin/keys` | List all API keys |
| POST | `/admin/keys/:prefix/credits` | Add credits to account |
| POST | `/admin/keys/:prefix/deactivate` | Deactivate API key |
| GET | `/admin/stats` | Global statistics |
| POST | `/admin/pricing` | Update model pricing |

### Authentication

**Client requests:**
```bash
curl -H "Authorization: Bearer oc_your_api_key_here" \
     -H "Content-Type: application/json" \
     -d '{"model": "deepseek-v4-pro", "messages": [...]}' \
     http://localhost:8088/v1/chat/completions
```

**Admin requests:**
```bash
curl -H "X-Admin-Key: $ADMIN_KEY" \
     http://localhost:8088/admin/stats
```

---

## Pricing

All prices in **satoshis per 1,000 tokens** (1 BTC = 100,000,000 satoshis)

| Model | Prompt Cost | Completion Cost | Description |
|-------|-------------|-----------------|-------------|
| deepseek-v4-pro | 50 sat | 150 sat | Most capable model |
| deepseek-v4-flash | 20 sat | 60 sat | Fast and efficient |
| kimi-k2.6 | 40 sat | 120 sat | Balanced performance |
| minimax-m2.7 | 30 sat | 90 sat | Creative & conversational |
| mimo-v2.5-pro | 45 sat | 135 sat | Advanced reasoning |
| skyclaw-v1 | 35 sat | 105 sat | Code generation specialist |

### Cost Examples
- **Simple question (500 tokens):** ~50 satoshis (~$0.03 USD at $60k BTC)
- **Code generation (2000 tokens):** ~300 satoshis (~$0.18 USD)
- **Long conversation (5000 tokens):** ~750 satoshis (~$0.45 USD)

---

## Client Management

### Creating a New Client
```bash
curl -X POST http://localhost:8088/admin/keys \
  -H "Content-Type: application/json" \
  -H "X-Admin-Key: $ADMIN_KEY" \
  -d '{
    "client_name": "company-name",
    "email": "client@company.com",
    "initial_credits": 50000
  }'
```

### Adding Credits
```bash
curl -X POST http://localhost:8088/admin/keys/oc_xxxx.../credits \
  -H "Content-Type: application/json" \
  -H "X-Admin-Key: $ADMIN_KEY" \
  -d '{
    "amount": 100000,
    "tx_hash": "btc_or_eth_tx_hash",
    "crypto_type": "BTC"
  }'
```

### Viewing Client Usage
```bash
curl -H "Authorization: Bearer oc_client_api_key" \
     http://localhost:8088/v1/account
```

---

## Revenue Tracking

### Global Statistics
```bash
curl -H "X-Admin-Key: $ADMIN_KEY" \
     http://localhost:8088/admin/stats
```

Returns:
```json
{
  "global_stats": {
    "total_clients": 0,
    "total_requests": 0,
    "total_tokens": 0,
    "total_revenue_satoshis": 0,
    "total_credits_remaining": 0
  }
}
```

### Revenue Milestones
| Date | Milestone | Revenue (sats) | Notes |
|------|-----------|----------------|-------|
| 2026-05-20 | Service launched | 0 | Initial deployment |

---

## Deployment

### Prerequisites
- Node.js 18+
- npm
- Linux server (Hetzner)

### Installation
```bash
cd /root/openclaw/llm-proxy
npm install
```

### Running
```bash
# Start in background
node src/server.js &

# Or with PM2 (recommended for production)
npm install -g pm2
pm2 start src/server.js --name llm-proxy
pm2 save
pm2 startup
```

### Environment Variables (.env)
```bash
PORT=8088
LLM_BASE_URL=https://llm.contes.me
LLM_API_KEY=your_upstream_api_key
ADMIN_KEY='<your-admin-key>'
```

### Monitoring
- Health check: `GET /health`
- Logs: Check `server.log` or PM2 logs
- Database: `data/proxy.db` (SQLite)

### Security Recommendations
1. Use HTTPS in production (nginx/caddy reverse proxy)
2. Restrict admin key access to trusted IPs
3. Regular database backups
4. Monitor for abuse (rate limiting)
5. Rotate admin key periodically

---

## Next Steps

### Phase 1: Launch (Current)
- [x] Create crypto wallets
- [x] Deploy LLM proxy service
- [x] Implement API key management
- [x] Set up billing system
- [ ] Configure public access (firewall, domain)
- [ ] Create landing page

### Phase 2: Growth
- [ ] Add more models
- [ ] Implement prepaid packages (bulk discounts)
- [ ] Create client dashboard (web UI)
- [ ] Set up automated payment detection
- [ ] Launch marketing campaigns

### Phase 3: Scale
- [ ] Multi-region deployment
- [ ] High availability setup
- [ ] Enterprise features (SLA, dedicated instances)
- [ ] Additional revenue streams (consulting, custom models)

---

## Support

For issues or questions:
- Check logs: `tail -f /root/openclaw/llm-proxy/server.log`
- Database queries: `sqlite3 /root/openclaw/llm-proxy/data/proxy.db`
- Contact: [Your contact info]

---

**Last Updated:** 2026-05-20
**Maintained by:** TataneSan
