require('dotenv').config();
const express = require('express');
const cors = require('cors');
const helmet = require('helmet');
const rateLimit = require('express-rate-limit');
const morgan = require('morgan');
const { createApiKey, validateKey, deductCredits, addCredits, logUsage, getPricing, getUsageStats, getGlobalStats, listApiKeys, recordPayment, hashKey } = require('./db');
const { proxyChatCompletion, proxyStreamCompletion, listAvailableModels } = require('./proxy');

const path = require('path');
const app = express();
const PORT = process.env.PORT || 8080;
const ADMIN_KEY = process.env.ADMIN_KEY || 'admin_' + require('crypto').randomBytes(16).toString('hex');

// Middleware
app.use(helmet());
app.use(cors());
app.use(express.json({ limit: '10mb' }));
app.use(morgan('combined'));

// Rate limiting
const limiter = rateLimit({
  windowMs: 60 * 1000, // 1 minute
  max: 100, // 100 requests per minute per IP
  message: { error: 'Rate limit exceeded. Please slow down.' }
});
app.use(limiter);

// Admin authentication middleware
function requireAdmin(req, res, next) {
  const adminKey = req.headers['x-admin-key'] || req.query.admin_key;
  if (!adminKey || adminKey !== ADMIN_KEY) {
    return res.status(401).json({ error: 'Invalid admin key' });
  }
  next();
}

// Client API key authentication middleware
async function requireApiKey(req, res, next) {
  const authHeader = req.headers.authorization;
  if (!authHeader || !authHeader.startsWith('Bearer ')) {
    return res.status(401).json({
      error: 'Missing API key',
      message: 'Include your API key in the Authorization header: Bearer oc_...'
    });
  }

  const apiKey = authHeader.slice(7);
  const keyData = validateKey(apiKey);

  if (!keyData) {
    return res.status(401).json({
      error: 'Invalid API key',
      message: 'Your API key is invalid or has been deactivated'
    });
  }

  if (keyData.credits_satoshis <= 0) {
    return res.status(402).json({
      error: 'Insufficient credits',
      message: 'Your account has no remaining credits. Please add more credits to continue.',
      credits_remaining: 0
    });
  }

  req.apiKey = apiKey;
  req.keyData = keyData;
  req.keyHash = keyData.key_hash;
  next();
}

// ============================================
// PUBLIC ENDPOINTS
// ============================================

// Health check
app.get('/health', (req, res) => {
  res.json({
    status: 'healthy',
    service: 'OpenClaw LLM Proxy',
    version: '1.0.0',
    timestamp: new Date().toISOString()
  });
});

// List available models (public)
app.get('/v1/models', (req, res) => {
  res.json({
    object: 'list',
    data: listAvailableModels()
  });
});

// Get pricing (public)
app.get('/v1/pricing', (req, res) => {
  const models = listAvailableModels();
  res.json({
    currency: 'satoshis',
    unit: 'per 1k tokens',
    models: models.map(m => ({
      id: m.id,
      name: m.name,
      prompt_cost: m.pricing.prompt,
      completion_cost: m.pricing.completion
    }))
  });
});

// ============================================
// CLIENT ENDPOINTS (require API key)
// ============================================

// Chat completions (OpenAI-compatible)
app.post('/v1/chat/completions', requireApiKey, async (req, res) => {
  const { keyData, keyHash } = req;
  const { model, stream } = req.body;

  // Get pricing for the requested model
  const pricing = getPricing(model);
  if (!pricing) {
    return res.status(400).json({
      error: 'Invalid model',
      message: `Model '${model}' is not available. Use GET /v1/models to see available models.`
    });
  }

  // Check if streaming is requested
  if (stream) {
    // For streaming, we need to estimate cost upfront and deduct
    // We'll estimate based on typical request size
    const estimatedPromptTokens = Math.ceil(JSON.stringify(req.body.messages).length / 4);
    const estimatedCost = Math.ceil((estimatedPromptTokens / 1000) * pricing.cost_per_1k_prompt_tokens) + 100; // Buffer

    if (keyData.credits_satoshis < estimatedCost) {
      return res.status(402).json({
        error: 'Insufficient credits for streaming',
        message: `Estimated cost: ${estimatedCost} satoshis. Your balance: ${keyData.credits_satoshis} satoshis.`,
        credits_remaining: keyData.credits_satoshis
      });
    }

    // Deduct estimated cost
    deductCredits(keyHash, estimatedCost);

    // Start streaming
    const streamResult = await proxyStreamCompletion(req.body, req.apiKey, res);

    // Log usage after stream completes (best effort)
    if (streamResult && streamResult.success) {
      const actualCost = Math.ceil(
        (streamResult.usage.prompt_tokens / 1000) * pricing.cost_per_1k_prompt_tokens +
        (streamResult.usage.completion_tokens / 1000) * pricing.cost_per_1k_completion_tokens
      );

      // Adjust credits (refund difference if overcharged)
      const adjustment = estimatedCost - actualCost;
      if (adjustment > 0) {
        addCredits(keyHash, adjustment);
      } else if (adjustment < 0) {
        deductCredits(keyHash, Math.abs(adjustment));
      }

      logUsage(keyHash, streamResult.resolvedModel, streamResult.usage.prompt_tokens,
               streamResult.usage.completion_tokens, actualCost, streamResult.latencyMs, 200);
    }
    return;
  }

  // Non-streaming request
  const result = await proxyChatCompletion(req.body, req.apiKey);

  if (result.success) {
    // Calculate actual cost
    const actualCost = Math.ceil(
      (result.usage.prompt_tokens / 1000) * pricing.cost_per_1k_prompt_tokens +
      (result.usage.completion_tokens / 1000) * pricing.cost_per_1k_completion_tokens
    );

    // Check if we have enough credits (we already deducted nothing yet for non-streaming)
    if (keyData.credits_satoshis < actualCost) {
      return res.status(402).json({
        error: 'Insufficient credits',
        message: `Request cost: ${actualCost} satoshis. Your balance: ${keyData.credits_satoshis} satoshis.`,
        credits_remaining: keyData.credits_satoshis
      });
    }

    // Deduct credits
    deductCredits(keyHash, actualCost);

    // Log usage
    logUsage(keyHash, result.resolvedModel, result.usage.prompt_tokens,
             result.usage.completion_tokens, actualCost, result.latencyMs, 200);

    // Add usage info to response
    result.data.usage_cost = {
      satoshis: actualCost,
      prompt_cost: Math.ceil((result.usage.prompt_tokens / 1000) * pricing.cost_per_1k_prompt_tokens),
      completion_cost: Math.ceil((result.usage.completion_tokens / 1000) * pricing.cost_per_1k_completion_tokens)
    };
    result.data.credits_remaining = keyData.credits_satoshis - actualCost;

    res.json(result.data);
  } else {
    // Log failed request
    logUsage(keyHash, result.resolvedModel, 0, 0, 0, result.latencyMs, result.statusCode);
    res.status(result.statusCode).json(result.error);
  }
});

// Get account balance and usage
app.get('/v1/account', requireApiKey, (req, res) => {
  const { keyData, keyHash } = req;
  const usage = getUsageStats(keyHash);

  res.json({
    client_name: keyData.client_name,
    email: keyData.email,
    credits_satoshis: keyData.credits_satoshis,
    total_spent_satoshis: keyData.total_spent_satoshis,
    total_requests: keyData.total_requests,
    total_tokens: keyData.total_tokens,
    created_at: keyData.created_at,
    last_used_at: keyData.last_used_at,
    usage_by_model: usage
  });
});

// Get detailed usage statistics
app.get('/v1/usage', requireApiKey, (req, res) => {
  const { keyHash } = req;
  const days = parseInt(req.query.days) || 30;
  const usage = getUsageStats(keyHash, days);

  res.json({
    period_days: days,
    usage_by_model: usage
  });
});

// ============================================
// ADMIN ENDPOINTS
// ============================================

// Create new API key
app.post('/admin/keys', requireAdmin, (req, res) => {
  const { client_name, email, initial_credits } = req.body;

  if (!client_name) {
    return res.status(400).json({ error: 'client_name is required' });
  }

  const result = createApiKey(client_name, email, initial_credits || 0);

  res.json({
    message: 'API key created successfully',
    api_key: result.key,
    key_prefix: result.keyPrefix,
    client_name: result.clientName,
    initial_credits_satoshis: initial_credits || 0,
    warning: 'Save this API key securely. It will not be shown again.'
  });
});

// List all API keys
app.get('/admin/keys', requireAdmin, (req, res) => {
  const keys = listApiKeys();
  res.json({
    total: keys.length,
    keys
  });
});

// Add credits to an API key
app.post('/admin/keys/:keyPrefix/credits', requireAdmin, (req, res) => {
  const { keyPrefix } = req.params;
  const { amount, tx_hash, crypto_type } = req.body;

  if (!amount || amount <= 0) {
    return res.status(400).json({ error: 'amount must be a positive integer (in satoshis)' });
  }

  // Find key by prefix
  const db = require('./db').db;
  const keyData = db.prepare('SELECT * FROM api_keys WHERE key_prefix = ?').get(keyPrefix);

  if (!keyData) {
    return res.status(404).json({ error: 'API key not found' });
  }

  // Add credits
  addCredits(keyData.key_hash, amount);

  // Record payment
  recordPayment(keyData.key_hash, amount, crypto_type || 'BTC', tx_hash);

  res.json({
    message: 'Credits added successfully',
    key_prefix: keyPrefix,
    amount_added_satoshis: amount,
    new_balance_satoshis: keyData.credits_satoshis + amount
  });
});

// Deactivate an API key
app.post('/admin/keys/:keyPrefix/deactivate', requireAdmin, (req, res) => {
  const { keyPrefix } = req.params;
  const db = require('./db').db;

  const result = db.prepare('UPDATE api_keys SET is_active = 0 WHERE key_prefix = ?').run(keyPrefix);

  if (result.changes === 0) {
    return res.status(404).json({ error: 'API key not found' });
  }

  res.json({
    message: 'API key deactivated',
    key_prefix: keyPrefix
  });
});

// Get global statistics
app.get('/admin/stats', requireAdmin, (req, res) => {
  const stats = getGlobalStats();
  res.json({
    global_stats: stats,
    timestamp: new Date().toISOString()
  });
});

// Update pricing for a model
app.post('/admin/pricing', requireAdmin, (req, res) => {
  const { model, prompt_cost, completion_cost } = req.body;

  if (!model || prompt_cost === undefined || completion_cost === undefined) {
    return res.status(400).json({ error: 'model, prompt_cost, and completion_cost are required' });
  }

  const db = require('./db').db;
  db.prepare(`
    INSERT OR REPLACE INTO pricing (model, cost_per_1k_prompt_tokens, cost_per_1k_completion_tokens, is_active)
    VALUES (?, ?, ?, 1)
  `).run(model, prompt_cost, completion_cost);

  res.json({
    message: 'Pricing updated',
    model,
    prompt_cost_satoshis_per_1k: prompt_cost,
    completion_cost_satoshis_per_1k: completion_cost
  });
});

// ============================================
// LANDING PAGE
// ============================================
app.use('/static', express.static(path.join(__dirname, '..', '..', 'landing')));
app.get('/', (req, res) => {
  res.sendFile(path.join(__dirname, '..', '..', 'landing', 'index.html'));
});

// ============================================
// START SERVER
// ============================================

app.listen(PORT, '0.0.0.0', () => {
  console.log(`
╔════════════════════════════════════════════════════════════╗
║                                                            ║
║   🦞 OpenClaw LLM Proxy - Revenue Engine                  ║
║                                                            ║
║   Server running on port ${PORT}                            ║
║                                                            ║
║   Endpoints:                                               ║
║   POST /v1/chat/completions  - Chat completions            ║
║   GET  /v1/models            - List available models       ║
║   GET  /v1/pricing           - Get pricing info            ║
║   GET  /v1/account           - Account balance & usage     ║
║   GET  /v1/usage             - Detailed usage stats        ║
║                                                            ║
║   Admin:                                                   ║
║   POST /admin/keys           - Create API key              ║
║   GET  /admin/keys           - List all keys               ║
║   POST /admin/keys/:id/credits - Add credits               ║
║   GET  /admin/stats          - Global statistics           ║
║                                                            ║
║   Admin Key: ${ADMIN_KEY.substring(0, 20)}...               ║
║                                                            ║
╚════════════════════════════════════════════════════════════╝
  `);
});

module.exports = app;
