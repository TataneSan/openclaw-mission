const express = require('express');
const Database = require('better-sqlite3');
const cron = require('node-cron');
const https = require('https');
const http = require('http');
const path = require('path');

const app = express();
const PORT = 8092;

// --- DB Setup ---
const db = new Database(path.join(__dirname, 'data', 'crypto.db'));
db.pragma('journal_mode = WAL');

db.exec(`
  CREATE TABLE IF NOT EXISTS prices (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    coin_id TEXT NOT NULL,
    symbol TEXT NOT NULL,
    name TEXT,
    price_usd REAL,
    market_cap REAL,
    volume_24h REAL,
    change_24h REAL,
    change_7d REAL,
    ath REAL,
    ath_change REAL,
    fetched_at DATETIME DEFAULT (datetime('now'))
  );

  CREATE TABLE IF NOT EXISTS price_history (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    coin_id TEXT NOT NULL,
    price_usd REAL NOT NULL,
    volume REAL,
    recorded_at DATETIME DEFAULT (datetime('now'))
  );

  CREATE TABLE IF NOT EXISTS alerts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    api_key TEXT NOT NULL,
    coin_id TEXT NOT NULL,
    condition TEXT NOT NULL CHECK(condition IN ('above','below')),
    target_price REAL NOT NULL,
    triggered INTEGER DEFAULT 0,
    created_at DATETIME DEFAULT (datetime('now'))
  );

  CREATE TABLE IF NOT EXISTS signals (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    coin_id TEXT NOT NULL,
    signal_type TEXT NOT NULL,
    strength REAL,
    indicators TEXT,
    price_at_signal REAL,
    created_at DATETIME DEFAULT (datetime('now'))
  );

  CREATE INDEX IF NOT EXISTS idx_prices_coin ON prices(coin_id);
  CREATE INDEX IF NOT EXISTS idx_history_coin ON price_history(coin_id, recorded_at);
  CREATE INDEX IF NOT EXISTS idx_alerts_key ON alerts(api_key);
  CREATE INDEX IF NOT EXISTS idx_signals_coin ON signals(coin_id, created_at);
`);

// --- Helper: HTTPS fetch as Promise ---
function httpsGet(url) {
  return new Promise((resolve, reject) => {
    const mod = url.startsWith('https') ? https : http;
    mod.get(url, { headers: { 'User-Agent': 'OpenClaw-CryptoMonitor/1.0' } }, (res) => {
      let data = '';
      res.on('data', chunk => data += chunk);
      res.on('end', () => {
        try { resolve(JSON.parse(data)); }
        catch { resolve(data); }
      });
    }).on('error', reject);
  });
}

// --- Top coins to track ---
const TOP_COINS = [
  'bitcoin', 'ethereum', 'binancecoin', 'solana', 'ripple',
  'cardano', 'dogecoin', 'polkadot', 'avalanche-2', 'chainlink',
  'tron', 'litecoin', 'polygon-bridged', 'uniswap', 'stellar'
];

const COIN_SYMBOLS = {
  bitcoin: 'BTC', ethereum: 'ETH', binancecoin: 'BNB', solana: 'SOL',
  ripple: 'XRP', cardano: 'ADA', dogecoin: 'DOGE', polkadot: 'DOT',
  'avalanche-2': 'AVAX', chainlink: 'LINK', tron: 'TRX', litecoin: 'LTC',
  'polygon-bridged': 'MATIC', uniswap: 'UNI', stellar: 'XLM'
};

// --- Fetch & store prices from CoinGecko ---
async function fetchPrices() {
  try {
    const ids = TOP_COINS.join(',');
    const url = `https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&ids=${ids}&order=market_cap_desc&sparkline=false`;
    const data = await httpsGet(url);

    if (!Array.isArray(data)) {
      console.error('[fetchPrices] Unexpected response:', typeof data);
      return;
    }

    const insertPrice = db.prepare(`
      INSERT INTO prices (coin_id, symbol, name, price_usd, market_cap, volume_24h, change_24h, change_7d, ath, ath_change)
      VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    `);

    const insertHistory = db.prepare(`
      INSERT INTO price_history (coin_id, price_usd, volume)
      VALUES (?, ?, ?)
    `);

    const tx = db.transaction(() => {
      for (const coin of data) {
        insertPrice.run(
          coin.id, coin.symbol?.toUpperCase(), coin.name,
          coin.current_price, coin.market_cap, coin.total_volume,
          coin.price_change_percentage_24h, coin.price_change_percentage_7d_in_currency,
          coin.ath, coin.ath_change_percentage
        );
        insertHistory.run(coin.id, coin.current_price, coin.total_volume);
      }
    });
    tx();

    console.log(`[${new Date().toISOString()}] Prices updated: ${data.length} coins`);
    runSignalAnalysis();
    checkAlerts();
  } catch (err) {
    console.error('[fetchPrices] Error:', err.message);
  }
}

// --- Technical Analysis ---
function computeRSI(prices, period = 14) {
  if (prices.length < period + 1) return null;
  let gains = 0, losses = 0;
  for (let i = 1; i <= period; i++) {
    const diff = prices[i] - prices[i - 1];
    if (diff > 0) gains += diff; else losses -= diff;
  }
  let avgGain = gains / period;
  let avgLoss = losses / period;
  for (let i = period + 1; i < prices.length; i++) {
    const diff = prices[i] - prices[i - 1];
    avgGain = (avgGain * (period - 1) + (diff > 0 ? diff : 0)) / period;
    avgLoss = (avgLoss * (period - 1) + (diff < 0 ? -diff : 0)) / period;
  }
  if (avgLoss === 0) return 100;
  const rs = avgGain / avgLoss;
  return 100 - (100 / (1 + rs));
}

function computeSMA(prices, period) {
  if (prices.length < period) return null;
  const slice = prices.slice(-period);
  return slice.reduce((a, b) => a + b, 0) / period;
}

function computeEMA(prices, period) {
  if (prices.length < period) return null;
  const k = 2 / (period + 1);
  let ema = prices.slice(0, period).reduce((a, b) => a + b, 0) / period;
  for (let i = period; i < prices.length; i++) {
    ema = prices[i] * k + ema * (1 - k);
  }
  return ema;
}

function computeMACD(prices) {
  if (prices.length < 26) return null;
  const ema12 = computeEMA(prices, 12);
  const ema26 = computeEMA(prices, 26);
  if (!ema12 || !ema26) return null;
  const macdLine = ema12 - ema26;
  // Simplified signal
  return { macd: macdLine, signal: macdLine * 0.8, histogram: macdLine * 0.2 };
}

function computeBollinger(prices, period = 20, stdDev = 2) {
  if (prices.length < period) return null;
  const slice = prices.slice(-period);
  const sma = slice.reduce((a, b) => a + b, 0) / period;
  const variance = slice.reduce((sum, p) => sum + (p - sma) ** 2, 0) / period;
  const sd = Math.sqrt(variance);
  return { upper: sma + stdDev * sd, middle: sma, lower: sma - stdDev * sd };
}

function analyzeSignals(coinId, prices) {
  const currentPrice = prices[prices.length - 1];
  const signals = [];

  // RSI
  const rsi = computeRSI(prices);
  if (rsi !== null) {
    if (rsi < 30) signals.push({ type: 'BUY', indicator: 'RSI', value: rsi, reason: `RSI oversold (${rsi.toFixed(1)})` });
    else if (rsi > 70) signals.push({ type: 'SELL', indicator: 'RSI', value: rsi, reason: `RSI overbought (${rsi.toFixed(1)})` });
    else signals.push({ type: 'HOLD', indicator: 'RSI', value: rsi, reason: `RSI neutral (${rsi.toFixed(1)})` });
  }

  // MACD
  const macd = computeMACD(prices);
  if (macd) {
    if (macd.histogram > 0 && macd.macd > 0) signals.push({ type: 'BUY', indicator: 'MACD', value: macd.histogram, reason: 'MACD bullish crossover' });
    else if (macd.histogram < 0 && macd.macd < 0) signals.push({ type: 'SELL', indicator: 'MACD', value: macd.histogram, reason: 'MACD bearish crossover' });
    else signals.push({ type: 'HOLD', indicator: 'MACD', value: macd.histogram, reason: 'MACD neutral' });
  }

  // Bollinger
  const bb = computeBollinger(prices);
  if (bb) {
    if (currentPrice <= bb.lower) signals.push({ type: 'BUY', indicator: 'BB', value: currentPrice, reason: `Price at lower Bollinger Band (${bb.lower.toFixed(2)})` });
    else if (currentPrice >= bb.upper) signals.push({ type: 'SELL', indicator: 'BB', value: currentPrice, reason: `Price at upper Bollinger Band (${bb.upper.toFixed(2)})` });
    else signals.push({ type: 'HOLD', indicator: 'BB', value: currentPrice, reason: 'Price within Bollinger Bands' });
  }

  // Moving Averages
  const sma20 = computeSMA(prices, 20);
  const sma50 = computeSMA(prices, 50);
  if (sma20 && sma50) {
    if (sma20 > sma50) signals.push({ type: 'BUY', indicator: 'MA_CROSS', value: sma20 - sma50, reason: 'SMA20 above SMA50 (Golden Cross)' });
    else signals.push({ type: 'SELL', indicator: 'MA_CROSS', value: sma20 - sma50, reason: 'SMA20 below SMA50 (Death Cross)' });
  }

  // Overall
  const buyCount = signals.filter(s => s.type === 'BUY').length;
  const sellCount = signals.filter(s => s.type === 'SELL').length;
  let overall = 'HOLD';
  let strength = 0.5;
  if (buyCount > sellCount) { overall = 'BUY'; strength = 0.5 + (buyCount - sellCount) * 0.15; }
  else if (sellCount > buyCount) { overall = 'SELL'; strength = 0.5 + (sellCount - buyCount) * 0.15; }
  strength = Math.min(strength, 1.0);

  return { overall, strength: +strength.toFixed(2), signals, currentPrice, rsi, macd, bollinger: bb, sma20, sma50 };
}

function runSignalAnalysis() {
  try {
    const coinIds = db.prepare('SELECT DISTINCT coin_id FROM price_history').all().map(r => r.coin_id);
    const insertSignal = db.prepare(`
      INSERT INTO signals (coin_id, signal_type, strength, indicators, price_at_signal)
      VALUES (?, ?, ?, ?, ?)
    `);

    for (const coinId of coinIds) {
      const rows = db.prepare(
        'SELECT price_usd FROM price_history WHERE coin_id = ? ORDER BY recorded_at DESC LIMIT 100'
      ).all(coinId).reverse().map(r => r.price_usd);

      if (rows.length < 5) continue;

      const analysis = analyzeSignals(coinId, rows);
      insertSignal.run(
        coinId, analysis.overall, analysis.strength,
        JSON.stringify(analysis.signals),
        analysis.currentPrice
      );
    }
    console.log(`[${new Date().toISOString()}] Signal analysis complete for ${coinIds.length} coins`);
  } catch (err) {
    console.error('[runSignalAnalysis] Error:', err.message);
  }
}

// --- Alert checking ---
function checkAlerts() {
  try {
    const pending = db.prepare('SELECT * FROM alerts WHERE triggered = 0').all();
    const updateAlert = db.prepare('UPDATE alerts SET triggered = 1 WHERE id = ?');

    for (const alert of pending) {
      const latest = db.prepare(
        'SELECT price_usd FROM prices WHERE coin_id = ? ORDER BY fetched_at DESC LIMIT 1'
      ).get(alert.coin_id);

      if (!latest) continue;

      const triggered = (alert.condition === 'above' && latest.price_usd >= alert.target_price) ||
                         (alert.condition === 'below' && latest.price_usd <= alert.target_price);

      if (triggered) {
        updateAlert.run(alert.id);
        console.log(`[ALERT] ${alert.coin_id} ${alert.condition} $${alert.target_price} triggered at $${latest.price_usd}`);
      }
    }
  } catch (err) {
    console.error('[checkAlerts] Error:', err.message);
  }
}

// --- API Key auth (reuse proxy.db) ---
let proxyDb;
try {
  proxyDb = new Database('/root/openclaw/llm-proxy/data/proxy.db', { readonly: true });
} catch { proxyDb = null; }

function authMiddleware(req, res, next) {
  // Public endpoints
  if (req.path === '/' || req.path === '/health' || req.path === '/api/v1/market/overview') return next();

  const key = req.headers['x-api-key'] || req.query.api_key;
  if (!key) return res.status(401).json({ error: 'API key required. Get one at https://138.201.249.160:8088' });

  if (proxyDb) {
    const user = proxyDb.prepare('SELECT * FROM api_keys WHERE key = ?').get(key);
    if (!user) return res.status(403).json({ error: 'Invalid API key' });
    req.user = user;
  }
  next();
}

app.use(express.json());
app.use(authMiddleware);

// --- ROUTES ---

// Landing page
app.get('/', (req, res) => {
  res.send(`<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>OpenClaw Crypto Monitor — Trading Signals API</title>
<style>
  * { margin: 0; padding: 0; box-sizing: border-box; }
  body { font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif; background: #0a0a0f; color: #e0e0e0; min-height: 100vh; }
  .hero { text-align: center; padding: 80px 20px 40px; }
  .hero h1 { font-size: 3rem; color: #00d4aa; margin-bottom: 10px; }
  .hero .tag { color: #888; font-size: 1.2rem; }
  .prices { display: flex; flex-wrap: wrap; justify-content: center; gap: 16px; padding: 20px; max-width: 1200px; margin: 0 auto; }
  .price-card { background: #141420; border: 1px solid #222; border-radius: 12px; padding: 20px; min-width: 160px; text-align: center; }
  .price-card .symbol { color: #00d4aa; font-weight: bold; font-size: 1.1rem; }
  .price-card .price { font-size: 1.4rem; margin: 8px 0; }
  .price-card .change { font-size: 0.9rem; }
  .up { color: #00e676; } .down { color: #ff5252; }
  .features { max-width: 900px; margin: 40px auto; padding: 0 20px; }
  .features h2 { color: #00d4aa; margin-bottom: 20px; font-size: 1.8rem; }
  .features ul { list-style: none; }
  .features li { padding: 10px 0; border-bottom: 1px solid #1a1a2e; }
  .features li code { background: #1a1a2e; padding: 2px 8px; border-radius: 4px; color: #bb86fc; }
  .pricing { max-width: 600px; margin: 40px auto; text-align: center; }
  .pricing h2 { color: #00d4aa; margin-bottom: 15px; }
  .pricing p { color: #888; margin-bottom: 8px; }
  .cta { text-align: center; padding: 40px; }
  .cta a { display: inline-block; background: #00d4aa; color: #0a0a0f; padding: 14px 32px; border-radius: 8px; text-decoration: none; font-weight: bold; font-size: 1.1rem; }
  .cta a:hover { background: #00b894; }
  .wallet { text-align: center; padding: 20px; color: #555; font-size: 0.85rem; }
  .wallet code { color: #00d4aa; }
</style>
</head>
<body>
<div class="hero">
  <h1>📊 OpenClaw Crypto Monitor</h1>
  <p class="tag">Real-time prices · Technical analysis · Trading signals · Price alerts</p>
</div>

<div class="prices" id="prices"><p style="color:#555">Loading live prices...</p></div>

<div class="features">
  <h2>API Endpoints</h2>
  <ul>
    <li><code>GET /api/v1/market/overview</code> — Top coins with live prices (public)</li>
    <li><code>GET /api/v1/coins/:id/price</code> — Current price & 24h stats</li>
    <li><code>GET /api/v1/coins/:id/signals</code> — Trading signals (RSI, MACD, Bollinger, MA)</li>
    <li><code>GET /api/v1/coins/:id/history</code> — Price history</li>
    <li><code>POST /api/v1/alerts</code> — Create price alert</li>
    <li><code>GET /api/v1/alerts</code> — List your alerts</li>
    <li><code>GET /api/v1/signals/all</code> — All latest signals</li>
  </ul>
</div>

<div class="pricing">
  <h2>Pricing</h2>
  <p>Free tier: 100 requests/day (market overview always free)</p>
  <p>Standard: 5 sats per request</p>
  <p>Same API key as the LLM Proxy</p>
</div>

<div class="cta">
  <a href="https://138.201.249.160:8088">Get API Key →</a>
</div>

<div class="wallet">
  BTC: <code>1GXpyCzAhvVNcbjnkytRTe1QoVoAen7nfP</code> ·
  ETH: <code>0xCF295d87E1534538FDac3b1f98746aF4A3E47352</code>
</div>

<script>
fetch('/api/v1/market/overview').then(r=>r.json()).then(data=>{
  const el = document.getElementById('prices');
  if (!data.coins) { el.innerHTML = '<p style="color:#555">Prices loading...</p>'; return; }
  el.innerHTML = data.coins.slice(0,10).map(c => {
    const ch = c.change_24h || 0;
    const cls = ch >= 0 ? 'up' : 'down';
    return '<div class="price-card"><div class="symbol">' + (c.symbol||'').toUpperCase() + '</div>' +
      '<div class="price">$' + (c.price_usd ? c.price_usd.toLocaleString() : '-') + '</div>' +
      '<div class="change ' + cls + '">' + (ch>=0?'+':'') + ch.toFixed(2) + '%</div></div>';
  }).join('');
}).catch(()=>{});
</script>
</body></html>`);
});

// Health
app.get('/health', (req, res) => {
  res.json({ status: 'healthy', service: 'OpenClaw Crypto Monitor', version: '1.0.0', timestamp: new Date().toISOString() });
});

// Market overview (public)
app.get('/api/v1/market/overview', (req, res) => {
  try {
    const coins = db.prepare(`
      SELECT p.* FROM prices p
      INNER JOIN (SELECT coin_id, MAX(fetched_at) as max_time FROM prices GROUP BY coin_id) latest
      ON p.coin_id = latest.coin_id AND p.fetched_at = latest.max_time
      ORDER BY p.market_cap DESC
    `).all();
    res.json({ coins, count: coins.length, updated: coins[0]?.fetched_at });
  } catch (err) {
    res.status(500).json({ error: err.message });
  }
});

// Single coin price
app.get('/api/v1/coins/:id/price', (req, res) => {
  try {
    const price = db.prepare(
      'SELECT * FROM prices WHERE coin_id = ? ORDER BY fetched_at DESC LIMIT 1'
    ).get(req.params.id);
    if (!price) return res.status(404).json({ error: 'Coin not found' });
    res.json(price);
  } catch (err) {
    res.status(500).json({ error: err.message });
  }
});

// Trading signals for a coin
app.get('/api/v1/coins/:id/signals', (req, res) => {
  try {
    const rows = db.prepare(
      'SELECT price_usd FROM price_history WHERE coin_id = ? ORDER BY recorded_at DESC LIMIT 100'
    ).all(req.params.id).reverse().map(r => r.price_usd);

    if (rows.length < 5) return res.status(404).json({ error: 'Not enough data for analysis' });

    const analysis = analyzeSignals(req.params.id, rows);
    res.json({ coin_id: req.params.id, ...analysis, data_points: rows.length });
  } catch (err) {
    res.status(500).json({ error: err.message });
  }
});

// Price history
app.get('/api/v1/coins/:id/history', (req, res) => {
  try {
    const limit = Math.min(parseInt(req.query.limit) || 100, 1000);
    const rows = db.prepare(
      'SELECT price_usd, volume, recorded_at FROM price_history WHERE coin_id = ? ORDER BY recorded_at DESC LIMIT ?'
    ).all(req.params.id, limit);
    res.json({ coin_id: req.params.id, history: rows.reverse(), count: rows.length });
  } catch (err) {
    res.status(500).json({ error: err.message });
  }
});

// All latest signals
app.get('/api/v1/signals/all', (req, res) => {
  try {
    const signals = db.prepare(`
      SELECT s.* FROM signals s
      INNER JOIN (SELECT coin_id, MAX(created_at) as max_time FROM signals GROUP BY coin_id) latest
      ON s.coin_id = latest.coin_id AND s.created_at = latest.max_time
      ORDER BY s.strength DESC
    `).all();
    const parsed = signals.map(s => ({ ...s, indicators: JSON.parse(s.indicators || '[]') }));
    res.json({ signals: parsed, count: parsed.length });
  } catch (err) {
    res.status(500).json({ error: err.message });
  }
});

// Create alert
app.post('/api/v1/alerts', (req, res) => {
  try {
    const { coin_id, condition, target_price } = req.body;
    if (!coin_id || !condition || !target_price) return res.status(400).json({ error: 'Missing: coin_id, condition, target_price' });
    if (!['above', 'below'].includes(condition)) return res.status(400).json({ error: 'condition must be "above" or "below"' });

    const apiKey = req.headers['x-api-key'] || req.query.api_key;
    const result = db.prepare(
      'INSERT INTO alerts (api_key, coin_id, condition, target_price) VALUES (?, ?, ?, ?)'
    ).run(apiKey, coin_id, condition, target_price);

    res.json({ id: result.lastInsertRowid, coin_id, condition, target_price, status: 'active' });
  } catch (err) {
    res.status(500).json({ error: err.message });
  }
});

// List alerts
app.get('/api/v1/alerts', (req, res) => {
  try {
    const apiKey = req.headers['x-api-key'] || req.query.api_key;
    const alerts = db.prepare('SELECT * FROM alerts WHERE api_key = ? ORDER BY created_at DESC').all(apiKey);
    res.json({ alerts, count: alerts.length });
  } catch (err) {
    res.status(500).json({ error: err.message });
  }
});

// Top gainers/losers
app.get('/api/v1/market/movers', (req, res) => {
  try {
    const coins = db.prepare(`
      SELECT p.* FROM prices p
      INNER JOIN (SELECT coin_id, MAX(fetched_at) as max_time FROM prices GROUP BY coin_id) latest
      ON p.coin_id = latest.coin_id AND p.fetched_at = latest.max_time
      ORDER BY p.change_24h DESC
    `).all();

    res.json({
      top_gainers: coins.filter(c => c.change_24h > 0).slice(0, 5),
      top_losers: coins.filter(c => c.change_24h < 0).slice(-5).reverse()
    });
  } catch (err) {
    res.status(500).json({ error: err.message });
  }
});

// --- CRON: fetch prices every 5 minutes ---
cron.schedule('*/5 * * * *', () => {
  console.log('[CRON] Fetching prices...');
  fetchPrices();
});

// --- Start ---
fetchPrices(); // Initial fetch

app.listen(PORT, '0.0.0.0', () => {
  console.log(`OpenClaw Crypto Monitor running on port ${PORT}`);
});
