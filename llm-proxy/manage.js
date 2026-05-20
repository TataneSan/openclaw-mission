#!/usr/bin/env node

const axios = require('axios');
require('dotenv').config({ path: require('path').join(__dirname, '.env') });

const BASE_URL = `http://localhost:${process.env.PORT || 8088}`;
const ADMIN_KEY = process.env.ADMIN_KEY;

const api = axios.create({
  baseURL: BASE_URL,
  headers: {
    'Content-Type': 'application/json',
    'X-Admin-Key': ADMIN_KEY
  }
});

const commands = {
  async createClient(args) {
    const [name, email, credits] = args;
    if (!name) {
      console.error('Usage: manage.js create-client <name> [email] [credits]');
      process.exit(1);
    }
    const res = await api.post('/admin/keys', {
      client_name: name,
      email: email || null,
      initial_credits: parseInt(credits) || 0
    });
    console.log('\n✅ API Key Created:');
    console.log('─'.repeat(50));
    console.log(`Client:     ${res.data.client_name}`);
    console.log(`API Key:    ${res.data.api_key}`);
    console.log(`Credits:    ${res.data.initial_credits_satoshis} satoshis`);
    console.log('─'.repeat(50));
    console.log('⚠️  Save this API key! It won\'t be shown again.\n');
  },

  async addCredits(args) {
    const [prefix, amount, txHash, crypto] = args;
    if (!prefix || !amount) {
      console.error('Usage: manage.js add-credits <key-prefix> <amount-sats> [tx-hash] [BTC|ETH]');
      process.exit(1);
    }
    const res = await api.post(`/admin/keys/${prefix}/credits`, {
      amount: parseInt(amount),
      tx_hash: txHash || null,
      crypto_type: crypto || 'BTC'
    });
    console.log('\n✅ Credits Added:');
    console.log(`Key:      ${res.data.key_prefix}`);
    console.log(`Added:    ${res.data.amount_added_satoshis} satoshis`);
    console.log(`Balance:  ${res.data.new_balance_satoshis} satoshis\n`);
  },

  async listKeys() {
    const res = await api.get('/admin/keys');
    console.log(`\n📋 API Keys (${res.data.total} total):\n`);
    console.log('Prefix'.padEnd(20) + 'Client'.padEnd(20) + 'Credits'.padEnd(15) + 'Requests'.padEnd(12) + 'Active');
    console.log('─'.repeat(80));
    for (const key of res.data.keys) {
      console.log(
        key.key_prefix.padEnd(20) +
        key.client_name.padEnd(20) +
        `${key.credits_satoshis} sats`.padEnd(15) +
        String(key.total_requests).padEnd(12) +
        (key.is_active ? '✅' : '❌')
      );
    }
    console.log('');
  },

  async deactivate(args) {
    const [prefix] = args;
    if (!prefix) {
      console.error('Usage: manage.js deactivate <key-prefix>');
      process.exit(1);
    }
    const res = await api.post(`/admin/keys/${prefix}/deactivate`);
    console.log(`\n✅ ${res.data.message}: ${res.data.key_prefix}\n`);
  },

  async stats() {
    const res = await api.get('/admin/stats');
    const s = res.data.global_stats;
    console.log('\n📊 Global Statistics:');
    console.log('─'.repeat(40));
    console.log(`Total Clients:     ${s.total_clients || 0}`);
    console.log(`Total Requests:    ${s.total_requests || 0}`);
    console.log(`Total Tokens:      ${s.total_tokens || 0}`);
    console.log(`Revenue:           ${s.total_revenue_satoshis || 0} satoshis`);
    console.log(`Credits Remaining: ${s.total_credits_remaining || 0} satoshis`);
    console.log(`Est. USD Value:    $${((s.total_revenue_satoshis || 0) * 0.0006).toFixed(2)}`);
    console.log('─'.repeat(40) + '\n');
  },

  async health() {
    const res = await axios.get(`${BASE_URL}/health`);
    console.log(`\n✅ Service Status: ${res.data.status}`);
    console.log(`   Version: ${res.data.version}`);
    console.log(`   Time: ${res.data.timestamp}\n`);
  },

  help() {
    console.log(`
🦞 OpenClaw LLM Proxy - Management CLI

Commands:
  create-client <name> [email] [credits]  Create new API key
  add-credits <prefix> <amount> [tx] [crypto]  Add credits
  list-keys                               List all API keys
  deactivate <prefix>                     Deactivate API key
  stats                                   Show global statistics
  health                                  Check service health
  help                                    Show this help

Examples:
  node manage.js create-client acme-corp john@acme.com 50000
  node manage.js add-credits oc_xxxx... 100000 txhash123 BTC
  node manage.js list-keys
  node manage.js stats
    `);
  }
};

const [,, command, ...args] = process.argv;

if (!command || !commands[command]) {
  commands.help();
} else {
  commands[command](args).catch(err => {
    console.error('Error:', err.response?.data || err.message);
    process.exit(1);
  });
}
