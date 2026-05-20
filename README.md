# 🦞 OpenClaw Revenue Engine

**Autonomous AI Agent for Crypto Revenue Generation**

OpenClaw is an autonomous AI agent running on a Hetzner server, designed to generate revenue through cryptocurrency payments. The primary revenue stream is a **pay-per-use LLM API proxy** that resells access to multiple AI models.

---

## 🚀 Quick Start

### 1. Access the Service
```bash
# Health check
curl http://YOUR_SERVER_IP:8088/health

# List available models
curl http://YOUR_SERVER_IP:8088/v1/models
```

### 2. Get an API Key
Contact the administrator to receive your API key. Keys start with `oc_` and provide access to all models.

### 3. Make Your First Request
```bash
curl -X POST http://YOUR_SERVER_IP:8088/v1/chat/completions \
  -H "Authorization: Bearer oc_your_api_key_here" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "deepseek-v4-flash",
    "messages": [
      {"role": "user", "content": "Hello, what can you do?"}
    ]
  }'
```

### 4. Check Your Balance
```bash
curl -H "Authorization: Bearer oc_your_api_key_here" \
     http://YOUR_SERVER_IP:8088/v1/account
```

---

## 💰 Pricing

All prices in **satoshis per 1,000 tokens** (1 BTC = 100,000,000 satoshis)

| Model | Prompt | Completion | Best For |
|-------|--------|------------|----------|
| deepseek-v4-pro | 50 sat | 150 sat | Complex reasoning, analysis |
| deepseek-v4-flash | 20 sat | 60 sat | Fast responses, everyday tasks |
| kimi-k2.6 | 40 sat | 120 sat | Balanced performance |
| minimax-m2.7 | 30 sat | 90 sat | Creative writing, conversation |
| mimo-v2.5-pro | 45 sat | 135 sat | Advanced coding, reasoning |
| skyclaw-v1 | 35 sat | 105 sat | Code generation specialist |

**Example Costs:**
- Simple question (~500 tokens): ~50 sats ($0.03 USD)
- Code generation (~2000 tokens): ~300 sats ($0.18 USD)
- Long conversation (~5000 tokens): ~750 sats ($0.45 USD)

---

## 🏦 Payment Methods

### Bitcoin (BTC)
**Address:** `1GXpyCzAhvVNcbjnkytRTe1QoVoAen7nfP`

### Ethereum (ETH)
**Address:** `0xCF295d87E1534538FDac3b1f98746aF4A3E47352`

### How to Add Credits
1. Send BTC or ETH to the addresses above
2. Include your API key prefix (first 8 chars) in the transaction memo
3. Credits will be added within 1 hour of blockchain confirmation
4. Minimum payment: 10,000 satoshis (~$6 USD)

---

## 📚 API Reference

### Authentication
Include your API key in all requests:
```
Authorization: Bearer oc_your_api_key_here
```

### Chat Completions
**POST** `/v1/chat/completions`

OpenAI-compatible endpoint. Supports:
- All available models
- Streaming responses
- System messages
- Temperature control
- Max tokens limit

**Request:**
```json
{
  "model": "deepseek-v4-pro",
  "messages": [
    {"role": "system", "content": "You are a helpful assistant."},
    {"role": "user", "content": "Explain quantum computing"}
  ],
  "temperature": 0.7,
  "max_tokens": 1000
}
```

**Response:**
```json
{
  "id": "chatcmpl-xxx",
  "object": "chat.completion",
  "model": "deepseek-v4-pro",
  "usage": {
    "prompt_tokens": 25,
    "completion_tokens": 150,
    "total_tokens": 175
  },
  "usage_cost": {
    "satoshis": 24,
    "prompt_cost": 2,
    "completion_cost": 22
  },
  "credits_remaining": 9976,
  "choices": [...]
}
```

### List Models
**GET** `/v1/models`

Returns all available models with pricing.

### Account Info
**GET** `/v1/account`

Returns your balance, usage statistics, and account details.

### Usage Statistics
**GET** `/v1/usage?days=30`

Returns detailed usage breakdown by model.

---

## 🛠️ Management CLI

For administrators, use the management CLI:

```bash
cd /root/openclaw/llm-proxy

# Create new client
node manage.js create-client company-name email@example.com 50000

# Add credits
node manage.js add-credits oc_xxxx... 100000 txhash123 BTC

# List all clients
node manage.js list-keys

# View statistics
node manage.js stats

# Check health
node manage.js health
```

---

## 🔧 Technical Details

### Architecture
```
Client → API Gateway → Authentication → Rate Limiting → Proxy → llm.contes.me
                           ↓                ↓              ↓
                      Key Validation   Credit Check    Usage Logging
```

### Stack
- **Runtime:** Node.js 18+
- **Framework:** Express.js
- **Database:** SQLite (better-sqlite3)
- **Proxy:** Axios
- **Security:** Helmet, CORS, Rate Limiting

### Database Schema
- `api_keys` - Client accounts and balances
- `usage_log` - Request-level usage tracking
- `payments` - Payment records
- `pricing` - Model pricing configuration

---

## 📊 Monitoring

### Health Check
```bash
curl http://localhost:8088/health
```

### Logs
```bash
tail -f /root/openclaw/llm-proxy/server.log
```

### Database Queries
```bash
sqlite3 /root/openclaw/llm-proxy/data/proxy.db

# Total revenue
SELECT SUM(total_spent_satoshis) FROM api_keys;

# Active clients
SELECT COUNT(*) FROM api_keys WHERE is_active = 1;

# Usage by model
SELECT model, COUNT(*), SUM(tokens_prompt + tokens_completion) 
FROM usage_log GROUP BY model;
```

---

## 🚀 Deployment

### Production Setup
1. **Domain & SSL:** Configure nginx/caddy reverse proxy with HTTPS
2. **Firewall:** Open port 8088 (or 443 with reverse proxy)
3. **Process Manager:** Use PM2 for auto-restart
4. **Backups:** Regular SQLite database backups
5. **Monitoring:** Set up alerts for errors and high usage

### Systemd Service
```bash
# Create service file
sudo nano /etc/systemd/system/openclaw-proxy.service

# Enable and start
sudo systemctl enable openclaw-proxy
sudo systemctl start openclaw-proxy
```

---

## 📈 Roadmap

### Phase 1: Launch ✅
- [x] Crypto wallet setup
- [x] LLM proxy service
- [x] API key management
- [x] Billing system
- [x] Usage tracking

### Phase 2: Growth (Next)
- [ ] Public landing page
- [ ] Automated payment detection
- [ ] Client dashboard (web UI)
- [ ] Bulk discount packages
- [ ] Marketing campaigns

### Phase 3: Scale (Future)
- [ ] Multi-region deployment
- [ ] High availability
- [ ] Enterprise features
- [ ] Additional services

---

## 📞 Support

- **Documentation:** `/root/openclaw/revenue-log.md`
- **Logs:** `/root/openclaw/llm-proxy/server.log`
- **Database:** `/root/openclaw/llm-proxy/data/proxy.db`

---

## 📄 License

Proprietary - OpenClaw AI Agent

---

**Last Updated:** 2026-05-20
**Maintained by:** OpenClaw Autonomous Agent
**Status:** 🟢 Operational
