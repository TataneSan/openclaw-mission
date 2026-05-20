const Database = require('better-sqlite3');
const path = require('path');
const crypto = require('crypto');

const DB_PATH = path.join(__dirname, '..', 'data', 'proxy.db');
const db = new Database(DB_PATH);

// Enable WAL mode for better concurrency
db.pragma('journal_mode = WAL');

// Initialize tables
db.exec(`
  CREATE TABLE IF NOT EXISTS api_keys (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    key_hash TEXT UNIQUE NOT NULL,
    key_prefix TEXT NOT NULL,
    client_name TEXT NOT NULL,
    email TEXT,
    credits_satoshis INTEGER DEFAULT 0,
    total_spent_satoshis INTEGER DEFAULT 0,
    total_requests INTEGER DEFAULT 0,
    total_tokens INTEGER DEFAULT 0,
    is_active INTEGER DEFAULT 1,
    rate_limit_per_min INTEGER DEFAULT 60,
    created_at TEXT DEFAULT (datetime('now')),
    last_used_at TEXT
  );

  CREATE TABLE IF NOT EXISTS usage_log (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    key_hash TEXT NOT NULL,
    model TEXT NOT NULL,
    tokens_prompt INTEGER DEFAULT 0,
    tokens_completion INTEGER DEFAULT 0,
    cost_satoshis INTEGER DEFAULT 0,
    latency_ms INTEGER DEFAULT 0,
    status_code INTEGER DEFAULT 200,
    created_at TEXT DEFAULT (datetime('now')),
    FOREIGN KEY (key_hash) REFERENCES api_keys(key_hash)
  );

  CREATE TABLE IF NOT EXISTS payments (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    key_hash TEXT NOT NULL,
    amount_satoshis INTEGER NOT NULL,
    tx_hash TEXT,
    crypto_type TEXT DEFAULT 'BTC',
    status TEXT DEFAULT 'pending',
    created_at TEXT DEFAULT (datetime('now')),
    confirmed_at TEXT,
    FOREIGN KEY (key_hash) REFERENCES api_keys(key_hash)
  );

  CREATE TABLE IF NOT EXISTS pricing (
    model TEXT PRIMARY KEY,
    cost_per_1k_prompt_tokens INTEGER NOT NULL,
    cost_per_1k_completion_tokens INTEGER NOT NULL,
    is_active INTEGER DEFAULT 1
  );

  CREATE INDEX IF NOT EXISTS idx_usage_key ON usage_log(key_hash);
  CREATE INDEX IF NOT EXISTS idx_usage_date ON usage_log(created_at);
  CREATE INDEX IF NOT EXISTS idx_payments_key ON payments(key_hash);
`);

// Default pricing (in satoshis per 1k tokens)
const defaultPricing = [
  { model: 'deepseek-v4-pro', prompt: 50, completion: 150 },
  { model: 'deepseek-v4-flash', prompt: 20, completion: 60 },
  { model: 'kimi-k2.6', prompt: 40, completion: 120 },
  { model: 'minimax-m2.7', prompt: 30, completion: 90 },
  { model: 'mimo-v2.5-pro', prompt: 45, completion: 135 },
  { model: 'skyclaw-v1', prompt: 35, completion: 105 },
];

const insertPricing = db.prepare(`
  INSERT OR IGNORE INTO pricing (model, cost_per_1k_prompt_tokens, cost_per_1k_completion_tokens)
  VALUES (?, ?, ?)
`);

for (const p of defaultPricing) {
  insertPricing.run(p.model, p.prompt, p.completion);
}

// Helper functions
function generateApiKey() {
  const prefix = 'oc_';
  const random = crypto.randomBytes(32).toString('hex');
  return prefix + random;
}

function hashKey(key) {
  return crypto.createHash('sha256').update(key).digest('hex');
}

function createApiKey(clientName, email, initialCreditsSatoshis = 0) {
  const key = generateApiKey();
  const keyHash = hashKey(key);
  const keyPrefix = key.substring(0, 8) + '...';

  db.prepare(`
    INSERT INTO api_keys (key_hash, key_prefix, client_name, email, credits_satoshis)
    VALUES (?, ?, ?, ?, ?)
  `).run(keyHash, keyPrefix, clientName, email || null, initialCreditsSatoshis);

  return { key, keyPrefix, clientName };
}

function validateKey(apiKey) {
  const keyHash = hashKey(apiKey);
  const row = db.prepare(`
    SELECT * FROM api_keys WHERE key_hash = ? AND is_active = 1
  `).get(keyHash);
  return row || null;
}

function deductCredits(keyHash, costSatoshis) {
  db.prepare(`
    UPDATE api_keys
    SET credits_satoshis = credits_satoshis - ?,
        total_spent_satoshis = total_spent_satoshis + ?,
        total_requests = total_requests + 1,
        last_used_at = datetime('now')
    WHERE key_hash = ?
  `).run(costSatoshis, costSatoshis, keyHash);
}

function addCredits(keyHash, amountSatoshis) {
  db.prepare(`
    UPDATE api_keys
    SET credits_satoshis = credits_satoshis + ?
    WHERE key_hash = ?
  `).run(amountSatoshis, keyHash);
}

function logUsage(keyHash, model, tokensPrompt, tokensCompletion, costSatoshis, latencyMs, statusCode) {
  db.prepare(`
    INSERT INTO usage_log (key_hash, model, tokens_prompt, tokens_completion, cost_satoshis, latency_ms, status_code)
    VALUES (?, ?, ?, ?, ?, ?, ?)
  `).run(keyHash, model, tokensPrompt, tokensCompletion, costSatoshis, latencyMs, statusCode);
}

function getPricing(model) {
  return db.prepare(`
    SELECT * FROM pricing WHERE model = ? AND is_active = 1
  `).get(model);
}

function getUsageStats(keyHash, days = 30) {
  return db.prepare(`
    SELECT
      model,
      COUNT(*) as requests,
      SUM(tokens_prompt) as total_prompt_tokens,
      SUM(tokens_completion) as total_completion_tokens,
      SUM(cost_satoshis) as total_cost_satoshis,
      AVG(latency_ms) as avg_latency_ms
    FROM usage_log
    WHERE key_hash = ? AND created_at >= datetime('now', '-' || ? || ' days')
    GROUP BY model
    ORDER BY total_cost_satoshis DESC
  `).all(keyHash, days);
}

function getGlobalStats() {
  return db.prepare(`
    SELECT
      COUNT(DISTINCT key_hash) as total_clients,
      SUM(total_requests) as total_requests,
      SUM(total_tokens) as total_tokens,
      SUM(total_spent_satoshis) as total_revenue_satoshis,
      SUM(credits_satoshis) as total_credits_remaining
    FROM api_keys
    WHERE is_active = 1
  `).get();
}

function listApiKeys() {
  return db.prepare(`
    SELECT id, key_prefix, client_name, email, credits_satoshis, total_spent_satoshis,
           total_requests, total_tokens, is_active, created_at, last_used_at
    FROM api_keys
    ORDER BY created_at DESC
  `).all();
}

function recordPayment(keyHash, amountSatoshis, cryptoType, txHash) {
  db.prepare(`
    INSERT INTO payments (key_hash, amount_satoshis, crypto_type, tx_hash, status)
    VALUES (?, ?, ?, ?, 'pending')
  `).run(keyHash, amountSatoshis, cryptoType, txHash || null);
}

module.exports = {
  db,
  createApiKey,
  validateKey,
  deductCredits,
  addCredits,
  logUsage,
  getPricing,
  getUsageStats,
  getGlobalStats,
  listApiKeys,
  recordPayment,
  hashKey
};
