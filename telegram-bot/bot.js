require('dotenv').config();
const { Telegraf, Markup } = require('telegraf');
const axios = require('axios');

const BOT_TOKEN = process.env.TELEGRAM_BOT_TOKEN;
const PROXY_URL = process.env.LLM_PROXY_URL || 'http://localhost:8088';
const ADMIN_KEY = process.env.LLM_PROXY_ADMIN_KEY;
const DEFAULT_MODEL = process.env.DEFAULT_MODEL || 'deepseek-v4-flash';
const FREE_TRIAL = parseInt(process.env.FREE_TRIAL_CREDITS) || 500;
const MAX_MSG = parseInt(process.env.MAX_MESSAGE_LENGTH) || 4000;
const HIST_LEN = parseInt(process.env.CONVERSATION_HISTORY_LENGTH) || 10;

if (!BOT_TOKEN || BOT_TOKEN === 'YOUR_BOT_TOKEN_HERE') {
  console.error('ERROR: Set TELEGRAM_BOT_TOKEN in .env');
  process.exit(1);
}

const bot = new Telegraf(BOT_TOKEN);

// In-memory user sessions: { userId: { apiKey, model, history: [{role, content}] } }
const sessions = new Map();

// ==========================================
// PROXY API HELPERS
// ==========================================

async function createProxyKey(clientName, email) {
  try {
    const res = await axios.post(`${PROXY_URL}/admin/keys`, {
      client_name: clientName,
      email: email,
      initial_credits: FREE_TRIAL
    }, {
      headers: { 'X-Admin-Key': ADMIN_KEY }
    });
    return res.data;
  } catch (err) {
    console.error('Failed to create proxy key:', err.message);
    return null;
  }
}

async function getAccountInfo(apiKey) {
  try {
    const res = await axios.get(`${PROXY_URL}/v1/account`, {
      headers: { 'Authorization': `Bearer ${apiKey}` }
    });
    return res.data;
  } catch (err) {
    if (err.response && err.response.status === 401) return null;
    console.error('Account info error:', err.message);
    return null;
  }
}

async function chatCompletion(apiKey, model, messages) {
  try {
    const res = await axios.post(`${PROXY_URL}/v1/chat/completions`, {
      model: model,
      messages: messages,
      max_tokens: 2048,
      temperature: 0.7
    }, {
      headers: { 'Authorization': `Bearer ${apiKey}` },
      timeout: 120000
    });
    return res.data;
  } catch (err) {
    if (err.response) {
      return { error: err.response.data };
    }
    return { error: { message: err.message } };
  }
}

// ==========================================
// USER SESSION MANAGEMENT
// ==========================================

function getSession(userId) {
  if (!sessions.has(userId)) {
    sessions.set(userId, {
      apiKey: null,
      model: DEFAULT_MODEL,
      history: []
    });
  }
  return sessions.get(userId);
}

function addToHistory(userId, role, content) {
  const session = getSession(userId);
  session.history.push({ role, content });
  if (session.history.length > HIST_LEN * 2) {
    session.history = session.history.slice(-HIST_LEN * 2);
  }
}

// ==========================================
// BOT COMMANDS
// ==========================================

// /start - Register and get free credits
bot.start(async (ctx) => {
  const userId = ctx.from.id.toString();
  const username = ctx.from.username || ctx.from.first_name || 'user';
  const session = getSession(userId);

  if (session.apiKey) {
    return ctx.reply(
      `Welcome back, ${username}!\n\n` +
      `Your API key: \`${session.apiKey.substring(0, 12)}...\`\n` +
      `Current model: ${session.model}\n\n` +
      `Just send me a message to chat!\n` +
      `Use /help for all commands.`,
      { parse_mode: 'Markdown' }
    );
  }

  // Create new API key
  const result = await createProxyKey(
    `telegram_${username}_${userId}`,
    null
  );

  if (!result) {
    return ctx.reply('Service temporarily unavailable. Please try again later.');
  }

  session.apiKey = result.api_key;

  ctx.reply(
    `Welcome to OpenClaw AI, ${username}!\n\n` +
    `You've received ${FREE_TRIAL} free satoshis in credits.\n\n` +
    `Your API key: \`${result.api_key.substring(0, 16)}...\`\n\n` +
    `Just send me any message to start chatting!\n\n` +
    `Commands:\n` +
    `/balance - Check credits\n` +
    `/model - Switch AI model\n` +
    `/clear - Clear conversation\n` +
    `/buy - Add more credits\n` +
    `/help - Show all commands`,
    { parse_mode: 'Markdown' }
  );
});

// /help - Show commands
bot.help((ctx) => {
  ctx.reply(
    `OpenClaw AI Assistant - Commands\n\n` +
    `/start - Register & get free credits\n` +
    `/balance - Check your credit balance\n` +
    `/model - Switch between AI models\n` +
    `/clear - Clear conversation history\n` +
    `/buy - Purchase more credits (BTC/ETH)\n` +
    `/key - Show your API key\n\n` +
    `Models available:\n` +
    `  deepseek-v4-pro - Best quality\n` +
    `  deepseek-v4-flash - Fast & cheap\n` +
    `  kimi-k2.6 - Balanced\n` +
    `  minimax-m2.7 - Creative\n` +
    `  mimo-v2.5-pro - Advanced reasoning\n` +
    `  skyclaw-v1 - Code specialist\n\n` +
    `Just send a message to chat!`
  );
});

// /balance - Check credits
bot.command('balance', async (ctx) => {
  const session = getSession(ctx.from.id.toString());
  if (!session.apiKey) {
    return ctx.reply('Please /start first to register.');
  }

  const info = await getAccountInfo(session.apiKey);
  if (!info) {
    return ctx.reply('Could not fetch account info. Try /start to re-register.');
  }

  const balanceBTC = (info.credits_satoshis / 100000000).toFixed(8);
  ctx.reply(
    `Your Account\n\n` +
    `Credits: ${info.credits_satoshis} satoshis (~${balanceBTC} BTC)\n` +
    `Total spent: ${info.total_spent_satoshis} satoshis\n` +
    `Total requests: ${info.total_requests}\n` +
    `Model: ${session.model}\n\n` +
    `Need more? Use /buy`
  );
});

// /model - Switch model
bot.command('model', (ctx) => {
  const session = getSession(ctx.from.id.toString());
  if (!session.apiKey) {
    return ctx.reply('Please /start first to register.');
  }

  ctx.reply(
    `Current model: ${session.model}\n\nSelect a new model:`,
    Markup.inlineKeyboard([
      [
        Markup.button.callback('DeepSeek V4 Pro', 'model_deepseek-v4-pro'),
        Markup.button.callback('DeepSeek V4 Flash', 'model_deepseek-v4-flash')
      ],
      [
        Markup.button.callback('Kimi K2.6', 'model_kimi-k2.6'),
        Markup.button.callback('MiniMax M2.7', 'model_minimax-m2.7')
      ],
      [
        Markup.button.callback('MiMo V2.5 Pro', 'model_mimo-v2.5-pro'),
        Markup.button.callback('SkyClaw V1', 'model_skyclaw-v1')
      ]
    ])
  );
});

// Model selection callbacks
bot.action(/model_(.+)/, async (ctx) => {
  const session = getSession(ctx.from.id.toString());
  const model = ctx.match[1];
  session.model = model;
  session.history = []; // Clear history on model switch
  await ctx.answerCbQuery(`Switched to ${model}`);
  await ctx.editMessageText(`Model switched to: ${model}\nConversation cleared. Send a message to start!`);
});

// /clear - Clear history
bot.command('clear', (ctx) => {
  const session = getSession(ctx.from.id.toString());
  session.history = [];
  ctx.reply('Conversation cleared. Send a new message to start fresh!');
});

// /key - Show API key
bot.command('key', (ctx) => {
  const session = getSession(ctx.from.id.toString());
  if (!session.apiKey) {
    return ctx.reply('Please /start first to register.');
  }
  ctx.reply(
    `Your API key:\n\`${session.apiKey}\`\n\n` +
    `You can use this key with the OpenClaw API directly:\n` +
    `POST ${PROXY_URL}/v1/chat/completions\n` +
    `Header: Authorization: Bearer ${session.apiKey}`,
    { parse_mode: 'Markdown' }
  );
});

// /buy - Purchase credits
bot.command('buy', (ctx) => {
  ctx.reply(
    `Purchase Credits\n\n` +
    `Send BTC or ETH to the addresses below:\n\n` +
    `Bitcoin (BTC):\n\`1GXpyCzAhvVNcbjnkytRTe1QoVoAen7nfP\`\n\n` +
    `Ethereum (ETH):\n\`0xCF295d87E1534538FDac3b1f98746aF4A3E47352\`\n\n` +
    `After sending, use:\n/confirm <tx_hash> <amount_sats>\n\n` +
    `Credits are added manually within 1 hour.\n` +
    `Minimum: 1000 satoshis (~$0.60)`,
    { parse_mode: 'Markdown' }
  );
});

// /confirm - Confirm payment
bot.command('confirm', async (ctx) => {
  const session = getSession(ctx.from.id.toString());
  if (!session.apiKey) {
    return ctx.reply('Please /start first to register.');
  }

  const parts = ctx.message.text.split(' ').slice(1);
  if (parts.length < 2) {
    return ctx.reply('Usage: /confirm <tx_hash> <amount_sats>\nExample: /confirm abc123... 5000');
  }

  const [txHash, amountStr] = parts;
  const amount = parseInt(amountStr);

  if (!amount || amount < 100) {
    return ctx.reply('Invalid amount. Minimum 100 satoshis.');
  }

  ctx.reply(
    `Payment submitted!\n\n` +
    `TX: ${txHash}\n` +
    `Amount: ${amount} satoshis\n\n` +
    `Credits will be added within 1 hour after confirmation.\n` +
    `Thank you!`
  );
});

// ==========================================
// MAIN CHAT HANDLER
// ==========================================

bot.on('text', async (ctx) => {
  const userId = ctx.from.id.toString();
  const session = getSession(userId);

  if (!session.apiKey) {
    return ctx.reply('Welcome! Please /start to register and get free credits.');
  }

  const userMessage = ctx.message.text;

  // Add user message to history
  addToHistory(userId, 'user', userMessage);

  // Show typing indicator
  await ctx.replyWithChatAction('typing');

  // Build messages array
  const messages = [
    {
      role: 'system',
      content: 'You are OpenClaw AI, a helpful and knowledgeable assistant. Be concise, accurate, and helpful. Respond in the same language the user writes in.'
    },
    ...session.history
  ];

  // Call LLM proxy
  const result = await chatCompletion(session.apiKey, session.model, messages);

  if (result.error) {
    const errMsg = result.error.message || result.error.error || 'Unknown error';
    if (errMsg.includes('Insufficient credits') || errMsg.includes('402')) {
      return ctx.reply(
        `You've run out of credits!\n\n` +
        `Use /buy to add more credits, or use /balance to check.`
      );
    }
    return ctx.reply(`Error: ${errMsg}\n\nTry again or use /clear to reset.`);
  }

  const assistantMessage = result.choices?.[0]?.message?.content || 'No response generated.';

  // Add to history
  addToHistory(userId, 'assistant', assistantMessage);

  // Split long messages
  const chunks = splitMessage(assistantMessage, MAX_MSG);
  for (const chunk of chunks) {
    await ctx.reply(chunk, { parse_mode: 'Markdown' }).catch(() => ctx.reply(chunk));
  }
});

// ==========================================
// HELPERS
// ==========================================

function splitMessage(text, maxLen) {
  if (text.length <= maxLen) return [text];
  const chunks = [];
  let remaining = text;
  while (remaining.length > 0) {
    if (remaining.length <= maxLen) {
      chunks.push(remaining);
      break;
    }
    // Try to split at newline
    let splitIdx = remaining.lastIndexOf('\n', maxLen);
    if (splitIdx < maxLen * 0.5) splitIdx = maxLen;
    chunks.push(remaining.substring(0, splitIdx));
    remaining = remaining.substring(splitIdx);
  }
  return chunks;
}

// ==========================================
// LAUNCH
// ==========================================

bot.launch().then(() => {
  console.log(`
╔══════════════════════════════════════════╗
║  OpenClaw Telegram Bot - RUNNING        ║
║  Model: ${DEFAULT_MODEL.padEnd(30)}║
║  Free trial: ${FREE_TRIAL} sats${' '.repeat(18)}║
╚══════════════════════════════════════════╝
  `);
});

// Graceful stop
process.once('SIGINT', () => bot.stop('SIGINT'));
process.once('SIGTERM', () => bot.stop('SIGTERM'));
