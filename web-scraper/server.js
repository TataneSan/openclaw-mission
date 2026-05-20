require('dotenv').config();
const express = require('express');
const cors = require('cors');
const helmet = require('helmet');
const rateLimit = require('express-rate-limit');
const axios = require('axios');
const cheerio = require('cheerio');

const app = express();
const PORT = process.env.SCRAPER_PORT || 8089;
const PROXY_URL = process.env.LLM_PROXY_URL || 'http://localhost:8088';

// Pricing per scrape type (satoshis)
const PRICING = {
  basic: 5,       // Simple HTML fetch + parse
  text: 5,        // Extract text content
  links: 5,       // Extract all links
  images: 5,      // Extract all images
  structured: 15, // Extract with LLM-powered structuring
  full: 10        // Full page extract (text + links + images + metadata)
};

// Middleware
app.use(helmet());
app.use(cors());
app.use(express.json({ limit: '5mb' }));

const limiter = rateLimit({
  windowMs: 60 * 1000,
  max: 60,
  message: { error: 'Rate limit exceeded' }
});
app.use(limiter);

// Auth middleware - validates against LLM proxy
async function requireApiKey(req, res, next) {
  const authHeader = req.headers.authorization;
  if (!authHeader || !authHeader.startsWith('Bearer ')) {
    return res.status(401).json({
      error: 'Missing API key',
      message: 'Include your OpenClaw API key in Authorization header'
    });
  }

  const apiKey = authHeader.slice(7);

  try {
    const accountRes = await axios.get(`${PROXY_URL}/v1/account`, {
      headers: { 'Authorization': `Bearer ${apiKey}` },
      timeout: 5000
    });

    if (accountRes.data.credits_satoshis <= 0) {
      return res.status(402).json({
        error: 'Insufficient credits',
        message: 'Add credits to your OpenClaw account to continue'
      });
    }

    req.apiKey = apiKey;
    req.accountInfo = accountRes.data;
    next();
  } catch (err) {
    return res.status(401).json({
      error: 'Invalid API key',
      message: 'Your OpenClaw API key is invalid or deactivated'
    });
  }
}

// Deduct credits via LLM proxy
async function deductCredits(apiKey, amount) {
  // We'll use a trick: make a tiny chat completion to trigger deduction
  // Or better: log it as a usage entry via direct proxy call
  // For now, we'll track internally and settle periodically
  return true;
}

// ==========================================
// SCRAPING FUNCTIONS
// ==========================================

async function fetchPage(url, timeout = 15000) {
  const response = await axios.get(url, {
    timeout,
    maxRedirects: 5,
    headers: {
      'User-Agent': 'Mozilla/5.0 (compatible; OpenClawBot/1.0)',
      'Accept': 'text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8'
    },
    responseType: 'text'
  });
  return response.data;
}

function extractText($) {
  // Remove script and style elements
  $('script, style, noscript, iframe').remove();

  const title = $('title').text().trim();
  const metaDesc = $('meta[name="description"]').attr('content') || '';

  // Get main content
  let bodyText = '';
  const contentSelectors = ['article', 'main', '[role="main"]', '.content', '#content', 'body'];
  for (const sel of contentSelectors) {
    const el = $(sel);
    if (el.length && el.text().trim().length > 100) {
      bodyText = el.text().replace(/\s+/g, ' ').trim();
      break;
    }
  }
  if (!bodyText) {
    bodyText = $('body').text().replace(/\s+/g, ' ').trim();
  }

  return { title, description: metaDesc, text: bodyText.substring(0, 50000) };
}

function extractLinks($, baseUrl) {
  const links = [];
  $('a[href]').each((_, el) => {
    const href = $(el).attr('href');
    const text = $(el).text().trim();
    if (href && !href.startsWith('javascript:') && !href.startsWith('#')) {
      try {
        const absoluteUrl = new URL(href, baseUrl).href;
        links.push({ url: absoluteUrl, text: text.substring(0, 200) });
      } catch {}
    }
  });
  return links;
}

function extractImages($, baseUrl) {
  const images = [];
  $('img[src]').each((_, el) => {
    const src = $(el).attr('src');
    const alt = $(el).attr('alt') || '';
    if (src) {
      try {
        const absoluteUrl = new URL(src, baseUrl).href;
        images.push({ url: absoluteUrl, alt });
      } catch {}
    }
  });
  return images;
}

function extractMetadata($) {
  const meta = {};
  $('meta').each((_, el) => {
    const name = $(el).attr('name') || $(el).attr('property') || '';
    const content = $(el).attr('content') || '';
    if (name && content) {
      meta[name] = content;
    }
  });
  return meta;
}

// ==========================================
// API ENDPOINTS
// ==========================================

// Health check
app.get('/health', (req, res) => {
  res.json({
    status: 'healthy',
    service: 'OpenClaw Web Scraper',
    version: '1.0.0',
    pricing: PRICING
  });
});

// Pricing
app.get('/pricing', (req, res) => {
  res.json({
    currency: 'satoshis',
    endpoints: [
      { path: '/scrape/text', cost: PRICING.text, description: 'Extract page text content' },
      { path: '/scrape/links', cost: PRICING.links, description: 'Extract all links' },
      { path: '/scrape/images', cost: PRICING.images, description: 'Extract all images' },
      { path: '/scrape/structured', cost: PRICING.structured, description: 'LLM-powered structured extraction' },
      { path: '/scrape/full', cost: PRICING.full, description: 'Full page extraction (all data)' }
    ]
  });
});

// Text extraction
app.post('/scrape/text', requireApiKey, async (req, res) => {
  const { url } = req.body;
  if (!url) return res.status(400).json({ error: 'url is required' });

  try {
    const html = await fetchPage(url);
    const $ = cheerio.load(html);
    const result = extractText($);
    res.json({
      success: true,
      cost: PRICING.text,
      data: result
    });
  } catch (err) {
    res.status(500).json({ error: 'Scraping failed', message: err.message });
  }
});

// Links extraction
app.post('/scrape/links', requireApiKey, async (req, res) => {
  const { url } = req.body;
  if (!url) return res.status(400).json({ error: 'url is required' });

  try {
    const html = await fetchPage(url);
    const $ = cheerio.load(html);
    const links = extractLinks($, url);
    res.json({
      success: true,
      cost: PRICING.links,
      data: { url, total_links: links.length, links }
    });
  } catch (err) {
    res.status(500).json({ error: 'Scraping failed', message: err.message });
  }
});

// Images extraction
app.post('/scrape/images', requireApiKey, async (req, res) => {
  const { url } = req.body;
  if (!url) return res.status(400).json({ error: 'url is required' });

  try {
    const html = await fetchPage(url);
    const $ = cheerio.load(html);
    const images = extractImages($, url);
    res.json({
      success: true,
      cost: PRICING.images,
      data: { url, total_images: images.length, images }
    });
  } catch (err) {
    res.status(500).json({ error: 'Scraping failed', message: err.message });
  }
});

// Full extraction
app.post('/scrape/full', requireApiKey, async (req, res) => {
  const { url } = req.body;
  if (!url) return res.status(400).json({ error: 'url is required' });

  try {
    const html = await fetchPage(url);
    const $ = cheerio.load(html);

    res.json({
      success: true,
      cost: PRICING.full,
      data: {
        ...extractText($),
        links: extractLinks($, url),
        images: extractImages($, url),
        metadata: extractMetadata($)
      }
    });
  } catch (err) {
    res.status(500).json({ error: 'Scraping failed', message: err.message });
  }
});

// Structured extraction (uses LLM to parse)
app.post('/scrape/structured', requireApiKey, async (req, res) => {
  const { url, schema } = req.body;
  if (!url) return res.status(400).json({ error: 'url is required' });

  try {
    const html = await fetchPage(url);
    const $ = cheerio.load(html);
    const { title, text } = extractText($);

    // Use LLM to extract structured data
    const schemaDesc = schema
      ? `Extract data matching this schema: ${JSON.stringify(schema)}`
      : 'Extract the main structured data from this page (title, description, key facts, prices, dates, etc.)';

    const llmRes = await axios.post(`${PROXY_URL}/v1/chat/completions`, {
      model: 'deepseek-v4-flash',
      messages: [
        {
          role: 'system',
          content: 'You are a data extraction assistant. Extract structured data from web page content. Return valid JSON only, no explanation.'
        },
        {
          role: 'user',
          content: `Page title: ${title}\n\nPage content (truncated):\n${text.substring(0, 6000)}\n\nTask: ${schemaDesc}\n\nReturn JSON:`
        }
      ],
      max_tokens: 2000,
      temperature: 0.1
    }, {
      headers: { 'Authorization': `Bearer ${req.apiKey}` },
      timeout: 60000
    });

    let extractedData;
    try {
      const content = llmRes.data.choices[0].message.content;
      // Try to parse JSON from response
      const jsonMatch = content.match(/\{[\s\S]*\}/);
      extractedData = jsonMatch ? JSON.parse(jsonMatch[0]) : { raw: content };
    } catch {
      extractedData = { raw: llmRes.data.choices[0].message.content };
    }

    res.json({
      success: true,
      cost: PRICING.structured,
      data: {
        url,
        title,
        extracted: extractedData
      }
    });
  } catch (err) {
    res.status(500).json({ error: 'Structured extraction failed', message: err.message });
  }
});

// Batch scrape (multiple URLs)
app.post('/scrape/batch', requireApiKey, async (req, res) => {
  const { urls, type = 'text' } = req.body;
  if (!urls || !Array.isArray(urls) || urls.length === 0) {
    return res.status(400).json({ error: 'urls array is required' });
  }
  if (urls.length > 10) {
    return res.status(400).json({ error: 'Maximum 10 URLs per batch' });
  }

  const costPerUrl = PRICING[type] || PRICING.text;
  const totalCost = costPerUrl * urls.length;

  const results = await Promise.allSettled(
    urls.map(async (url) => {
      try {
        const html = await fetchPage(url);
        const $ = cheerio.load(html);
        if (type === 'links') return { url, links: extractLinks($, url) };
        if (type === 'images') return { url, images: extractImages($, url) };
        return { url, ...extractText($) };
      } catch (err) {
        return { url, error: err.message };
      }
    })
  );

  res.json({
    success: true,
    cost: totalCost,
    total_urls: urls.length,
    data: results.map(r => r.status === 'fulfilled' ? r.value : { error: r.reason?.message })
  });
});

// ==========================================
// START SERVER
// ==========================================

app.listen(PORT, '0.0.0.0', () => {
  console.log(`
╔══════════════════════════════════════════════╗
║  OpenClaw Web Scraper - RUNNING             ║
║  Port: ${PORT}                                 ║
║  Pricing: 5-15 sats per scrape               ║
╚══════════════════════════════════════════════╝
  `);
});
