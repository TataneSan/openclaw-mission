const express = require('express');
const cheerio = require('cheerio');
const https = require('https');
const http = require('http');
const crypto = require('crypto');
const { URL } = require('url');
const path = require('path');

const app = express();
const PORT = 8093;

// --- Helper: HTTP(S) fetch ---
function fetchUrl(url, maxRedirects = 5) {
  return new Promise((resolve, reject) => {
    if (maxRedirects <= 0) return reject(new Error('Too many redirects'));
    const mod = url.startsWith('https') ? https : http;
    const req = mod.get(url, {
      headers: {
        'User-Agent': 'Mozilla/5.0 (compatible; OpenClaw-SEO/1.0; +https://138.201.249.160)',
        'Accept': 'text/html,application/xhtml+xml,*/*',
      },
      timeout: 15000,
    }, (res) => {
      if (res.statusCode >= 300 && res.statusCode < 400 && res.headers.location) {
        const redirectUrl = new URL(res.headers.location, url).href;
        return fetchUrl(redirectUrl, maxRedirects - 1).then(resolve).catch(reject);
      }
      let data = '';
      res.on('data', chunk => data += chunk);
      res.on('end', () => resolve({ status: res.statusCode, headers: res.headers, body: data }));
    });
    req.on('error', reject);
    req.on('timeout', () => { req.destroy(); reject(new Error('Request timed out')); });
  });
}

// --- Helper: Call LLM Proxy ---
async function callLLM(prompt, maxTokens = 1024) {
  return new Promise((resolve, reject) => {
    const body = JSON.stringify({
      model: 'mistral-7b',
      messages: [{ role: 'user', content: prompt }],
      max_tokens: maxTokens,
      temperature: 0.3,
    });
    const req = http.request({
      hostname: 'localhost',
      port: 8088,
      path: '/v1/chat/completions',
      method: 'POST',
      headers: { 'Content-Type': 'application/json', 'x-api-key': 'seo-toolkit-internal' },
      timeout: 30000,
    }, (res) => {
      let data = '';
      res.on('data', chunk => data += chunk);
      res.on('end', () => {
        try {
          const j = JSON.parse(data);
          resolve(j.choices?.[0]?.message?.content || data);
        } catch { resolve(data); }
      });
    });
    req.on('error', (e) => resolve('LLM unavailable: ' + e.message));
    req.on('timeout', () => { req.destroy(); resolve('LLM timeout'); });
    req.write(body);
    req.end();
  });
}

// --- SEO Analysis Engine ---
function analyzePage(html, url) {
  const $ = cheerio.load(html);
  const parsedUrl = new URL(url);
  const issues = [];
  const scores = {};

  // 1. Title tag
  const title = $('title').text().trim();
  const titleLen = title.length;
  if (!title) issues.push({ severity: 'critical', category: 'title', message: 'Missing <title> tag' });
  else if (titleLen < 30) issues.push({ severity: 'warning', category: 'title', message: `Title too short (${titleLen} chars, recommend 50-60)` });
  else if (titleLen > 60) issues.push({ severity: 'warning', category: 'title', message: `Title too long (${titleLen} chars, recommend 50-60)` });
  scores.title = title ? (titleLen >= 30 && titleLen <= 60 ? 100 : titleLen > 0 ? 60 : 0) : 0;

  // 2. Meta description
  const metaDesc = $('meta[name="description"]').attr('content')?.trim() || '';
  const descLen = metaDesc.length;
  if (!metaDesc) issues.push({ severity: 'critical', category: 'meta', message: 'Missing meta description' });
  else if (descLen < 120) issues.push({ severity: 'warning', category: 'meta', message: `Meta description too short (${descLen} chars, recommend 150-160)` });
  else if (descLen > 160) issues.push({ severity: 'warning', category: 'meta', message: `Meta description too long (${descLen} chars, recommend 150-160)` });
  scores.meta = metaDesc ? (descLen >= 120 && descLen <= 160 ? 100 : 60) : 0;

  // 3. Headings structure
  const headings = {};
  for (let i = 1; i <= 6; i++) {
    headings[`h${i}`] = $(`h${i}`).map((_, el) => $(el).text().trim()).get();
  }
  if (!headings.h1.length) issues.push({ severity: 'critical', category: 'headings', message: 'No H1 heading found' });
  else if (headings.h1.length > 1) issues.push({ severity: 'warning', category: 'headings', message: `Multiple H1 tags (${headings.h1.length}), should have exactly 1` });
  scores.headings = headings.h1.length === 1 ? 100 : headings.h1.length > 0 ? 60 : 0;

  // 4. Images without alt text
  const images = $('img').map((_, el) => ({
    src: $(el).attr('src') || '',
    alt: $(el).attr('alt') || '',
    width: $(el).attr('width') || '',
    height: $(el).attr('height') || '',
  })).get();
  const imgNoAlt = images.filter(i => !i.alt);
  if (imgNoAlt.length > 0) issues.push({ severity: 'warning', category: 'images', message: `${imgNoAlt.length}/${images.length} images missing alt text` });
  scores.images = images.length === 0 ? 100 : Math.max(0, 100 - (imgNoAlt.length / images.length) * 100);

  // 5. Links analysis
  const internalLinks = [];
  const externalLinks = [];
  $('a[href]').each((_, el) => {
    const href = $(el).attr('href') || '';
    try {
      const linkUrl = new URL(href, url);
      if (linkUrl.hostname === parsedUrl.hostname) internalLinks.push(href);
      else externalLinks.push(href);
    } catch {}
  });
  scores.links = internalLinks.length > 0 ? 100 : 50;

  // 6. Canonical tag
  const canonical = $('link[rel="canonical"]').attr('href') || '';
  if (!canonical) issues.push({ severity: 'warning', category: 'technical', message: 'No canonical tag found' });
  scores.canonical = canonical ? 100 : 50;

  // 7. Open Graph tags
  const ogTags = {};
  $('meta[property^="og:"]').each((_, el) => {
    ogTags[$(el).attr('property')] = $(el).attr('content');
  });
  if (!ogTags['og:title']) issues.push({ severity: 'warning', category: 'social', message: 'Missing og:title' });
  if (!ogTags['og:description']) issues.push({ severity: 'warning', category: 'social', message: 'Missing og:description' });
  if (!ogTags['og:image']) issues.push({ severity: 'warning', category: 'social', message: 'Missing og:image' });
  scores.social = Object.keys(ogTags).length >= 3 ? 100 : Object.keys(ogTags).length * 30;

  // 8. Twitter Card tags
  const twitterTags = {};
  $('meta[name^="twitter:"]').each((_, el) => {
    twitterTags[$(el).attr('name')] = $(el).attr('content');
  });

  // 9. Structured data (JSON-LD)
  const jsonLd = [];
  $('script[type="application/ld+json"]').each((_, el) => {
    try { jsonLd.push(JSON.parse($(el).html())); } catch {}
  });
  scores.structured = jsonLd.length > 0 ? 100 : 30;

  // 10. Viewport meta
  const viewport = $('meta[name="viewport"]').attr('content') || '';
  if (!viewport) issues.push({ severity: 'critical', category: 'mobile', message: 'Missing viewport meta tag (not mobile-friendly)' });
  scores.mobile = viewport ? 100 : 0;

  // 11. Charset
  const charset = $('meta[charset]').attr('charset') || $('meta[http-equiv="Content-Type"]').attr('content') || '';

  // 12. Favicon
  const favicon = $('link[rel="icon"], link[rel="shortcut icon"]').attr('href') || '';

  // 13. Word count
  const bodyText = $('body').text().replace(/\s+/g, ' ').trim();
  const wordCount = bodyText.split(/\s+/).filter(w => w.length > 0).length;
  if (wordCount < 300) issues.push({ severity: 'warning', category: 'content', message: `Thin content: only ${wordCount} words (recommend 300+)` });
  scores.content = wordCount >= 300 ? 100 : Math.min(80, wordCount / 3);

  // 14. Keyword density (top 10 words)
  const stopWords = new Set(['the','a','an','is','are','was','were','be','been','being','have','has','had','do','does','did','will','would','could','should','may','might','shall','can','need','dare','ought','used','to','of','in','for','on','with','at','by','from','as','into','through','during','before','after','above','below','between','out','off','over','under','again','further','then','once','and','but','or','nor','not','so','yet','both','either','neither','each','every','all','any','few','more','most','other','some','such','no','only','own','same','than','too','very','just','that','this','these','those','it','its','i','me','my','we','our','you','your','he','him','his','she','her','they','them','their','what','which','who','whom','when','where','why','how','if','because','while','although','until','about','also','de','le','la','les','des','un','une','et','est','en','que','qui','dans','du','au','aux','pour','par','sur','pas','plus','ce','se','ou','il','ne','son','sa','ses','avec','tout','cette','ces','comme','mais','donc','alors']);
  const words = bodyText.toLowerCase().replace(/[^a-zàâéèêëïîôùûüÿçæœ\s]/g, '').split(/\s+/).filter(w => w.length > 2 && !stopWords.has(w));
  const freq = {};
  words.forEach(w => freq[w] = (freq[w] || 0) + 1);
  const topKeywords = Object.entries(freq).sort((a, b) => b[1] - a[1]).slice(0, 15).map(([word, count]) => ({
    word,
    count,
    density: +((count / words.length) * 100).toFixed(2),
  }));

  // Overall score
  const weights = { title: 15, meta: 15, headings: 15, images: 10, social: 10, mobile: 10, content: 10, canonical: 5, links: 5, structured: 5 };
  let totalScore = 0, totalWeight = 0;
  for (const [k, w] of Object.entries(weights)) {
    if (scores[k] !== undefined) { totalScore += scores[k] * w; totalWeight += w; }
  }
  const overallScore = Math.round(totalScore / totalWeight);

  return {
    url,
    overall_score: overallScore,
    scores,
    title: { text: title, length: titleLen },
    meta_description: { text: metaDesc, length: descLen },
    headings,
    images: { total: images.length, missing_alt: imgNoAlt.length, samples: images.slice(0, 10) },
    links: { internal: internalLinks.length, external: externalLinks.length },
    canonical,
    og_tags: ogTags,
    twitter_tags: twitterTags,
    structured_data: jsonLd,
    viewport,
    charset,
    favicon,
    word_count: wordCount,
    top_keywords: topKeywords,
    issues,
    issues_count: { critical: issues.filter(i => i.severity === 'critical').length, warning: issues.filter(i => i.severity === 'warning').length },
  };
}

// --- Auth middleware (reuse proxy.db with SHA256 hash matching) ---
let proxyDb;
try {
  const Database = require('better-sqlite3');
  proxyDb = new Database('/root/openclaw/llm-proxy/data/proxy.db', { readonly: true });
} catch { proxyDb = null; }

function hashKey(key) {
  return crypto.createHash('sha256').update(key).digest('hex');
}

function validateKey(apiKey) {
  if (!proxyDb) return null;
  const keyHash = hashKey(apiKey);
  return proxyDb.prepare('SELECT * FROM api_keys WHERE key_hash = ? AND is_active = 1').get(keyHash) || null;
}

function authMiddleware(req, res, next) {
  // Public endpoints
  if (req.path === '/' || req.path === '/health') return next();

  // Accept Authorization: Bearer oc_... or x-api-key header
  let apiKey = null;
  const authHeader = req.headers.authorization;
  if (authHeader && authHeader.startsWith('Bearer ')) {
    apiKey = authHeader.slice(7);
  }
  if (!apiKey) apiKey = req.headers['x-api-key'] || req.query.api_key;
  if (!apiKey) return res.status(401).json({ error: 'API key required. Get one at https://138.201.249.160:8088' });

  const keyData = validateKey(apiKey);
  if (!keyData) return res.status(403).json({ error: 'Invalid API key' });
  req.user = keyData;
  next();
}

app.use(express.json());
app.use(authMiddleware);

// ===================== ROUTES =====================

// Landing page
app.get('/', (req, res) => {
  res.send(`<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>OpenClaw SEO Toolkit — Automated SEO Analysis API</title>
<style>
  *{margin:0;padding:0;box-sizing:border-box}
  body{font-family:-apple-system,BlinkMacSystemFont,'Segoe UI',Roboto,sans-serif;background:#0a0a0f;color:#e0e0e0;min-height:100vh}
  .hero{text-align:center;padding:80px 20px 40px;background:linear-gradient(135deg,#0a0a0f 0%,#1a0a2e 100%)}
  .hero h1{font-size:3rem;color:#00d4aa;margin-bottom:10px}
  .hero .tag{color:#888;font-size:1.2rem}
  .features{max-width:900px;margin:40px auto;padding:0 20px}
  .features h2{color:#00d4aa;margin-bottom:20px;font-size:1.8rem}
  .features ul{list-style:none}
  .features li{padding:10px 0;border-bottom:1px solid #1a1a2e}
  .features li code{background:#1a1a2e;padding:2px 8px;border-radius:4px;color:#bb86fc}
  .demo{max-width:700px;margin:40px auto;padding:20px;background:#141420;border-radius:12px;border:1px solid #222}
  .demo h2{color:#00d4aa;margin-bottom:15px}
  .demo input{width:100%;padding:12px;border:1px solid #333;border-radius:8px;background:#0a0a0f;color:#e0e0e0;font-size:1rem;margin-bottom:10px}
  .demo button{background:#00d4aa;color:#0a0a0f;padding:12px 24px;border:none;border-radius:8px;font-weight:bold;cursor:pointer;font-size:1rem}
  .demo button:hover{background:#00b894}
  .result{margin-top:20px;padding:15px;background:#0a0a0f;border-radius:8px;display:none;max-height:400px;overflow-y:auto}
  .result pre{white-space:pre-wrap;font-size:0.85rem;color:#888}
  .pricing{max-width:600px;margin:40px auto;text-align:center}
  .pricing h2{color:#00d4aa;margin-bottom:15px}
  .pricing p{color:#888;margin-bottom:8px}
  .cta{text-align:center;padding:40px}
  .cta a{display:inline-block;background:#00d4aa;color:#0a0a0f;padding:14px 32px;border-radius:8px;text-decoration:none;font-weight:bold;font-size:1.1rem}
  .cta a:hover{background:#00b894}
  .wallet{text-align:center;padding:20px;color:#555;font-size:0.85rem}
  .wallet code{color:#00d4aa}
  .score-ring{display:inline-block;width:120px;height:120px;border-radius:50%;border:8px solid #222;position:relative;margin:20px auto}
  .score-ring .val{position:absolute;top:50%;left:50%;transform:translate(-50%,-50%);font-size:2rem;font-weight:bold;color:#00d4aa}
</style>
</head>
<body>
<div class="hero">
  <h1>🔍 OpenClaw SEO Toolkit</h1>
  <p class="tag">Automated SEO analysis · Meta tag generation · Content audit · Keyword research</p>
</div>

<div class="demo">
  <h2>Try it — Analyze any URL</h2>
  <input type="text" id="url" placeholder="https://example.com" />
  <button onclick="analyze()">Analyze SEO</button>
  <div class="result" id="result"><pre id="resultText"></pre></div>
</div>

<div class="features">
  <h2>API Endpoints</h2>
  <ul>
    <li><code>GET /api/v1/analyze?url=...</code> — Full SEO audit of a web page</li>
    <li><code>GET /api/v1/keywords?url=...</code> — Keyword density analysis</li>
    <li><code>POST /api/v1/meta/generate</code> — Generate optimized meta tags (title + description + OG)</li>
    <li><code>GET /api/v1/meta/check?url=...</code> — Check existing meta tags</li>
    <li><code>GET /api/v1/headings?url=...</code> — Heading structure analysis</li>
    <li><code>GET /api/v1/images?url=...</code> — Image alt text audit</li>
    <li><code>GET /api/v1/links?url=...</code> — Internal/external link analysis</li>
    <li><code>GET /api/v1/structured?url=...</code> — Schema.org structured data check</li>
    <li><code>POST /api/v1/content/audit</code> — AI-powered content quality audit</li>
    <li><code>POST /api/v1/content/rewrite</code> — AI SEO-optimized content rewrite</li>
    <li><code>GET /api/v1/robots?url=...</code> — robots.txt parser & checker</li>
    <li><code>GET /api/v1/sitemap?url=...</code> — Sitemap.xml parser</li>
  </ul>
</div>

<div class="pricing">
  <h2>Pricing</h2>
  <p>Basic analysis: 3 sats per request</p>
  <p>AI-powered audit/rewrite: 10 sats per request</p>
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
function analyze(){
  const url=document.getElementById('url').value;
  if(!url){alert('Enter a URL');return;}
  const r=document.getElementById('result');
  const rt=document.getElementById('resultText');
  r.style.display='block';
  rt.textContent='Analyzing...';
  fetch('/api/v1/analyze?url='+encodeURIComponent(url))
    .then(res=>res.json())
    .then(data=>{
      if(data.error){rt.textContent='Error: '+data.error;return;}
      let out='SEO SCORE: '+data.overall_score+'/100\\n\\n';
      out+='Title: '+data.title.text+' ('+data.title.length+' chars)\\n';
      out+='Meta: '+data.meta_description.text?.substring(0,100)+'... ('+data.meta_description.length+' chars)\\n';
      out+='H1: '+(data.headings.h1?.[0]||'MISSING')+'\\n';
      out+='Words: '+data.word_count+'\\n';
      out+='Images: '+data.images.total+' ('+data.images.missing_alt+' missing alt)\\n';
      out+='Links: '+data.links.internal+' internal, '+data.links.external+' external\\n\\n';
      out+='ISSUES ('+data.issues_count.critical+' critical, '+data.issues_count.warning+' warnings):\\n';
      data.issues.forEach(i=>{out+='  ['+i.severity.toUpperCase()+'] '+i.message+'\\n'});
      out+='\\nTOP KEYWORDS:\\n';
      data.top_keywords.slice(0,8).forEach(k=>{out+='  '+k.word+': '+k.count+' ('+k.density+'%)\\n'});
      rt.textContent=out;
    })
    .catch(e=>{rt.textContent='Error: '+e.message});
}
</script>
</body></html>`);
});

// Health
app.get('/health', (req, res) => {
  res.json({ status: 'healthy', service: 'OpenClaw SEO Toolkit', version: '1.0.0', timestamp: new Date().toISOString() });
});

// --- Full SEO Analysis ---
app.get('/api/v1/analyze', async (req, res) => {
  try {
    const url = req.query.url;
    if (!url) return res.status(400).json({ error: 'Missing ?url= parameter' });
    const { body } = await fetchUrl(url);
    const analysis = analyzePage(body, url);
    res.json(analysis);
  } catch (err) {
    res.status(500).json({ error: err.message });
  }
});

// --- Keyword Analysis ---
app.get('/api/v1/keywords', async (req, res) => {
  try {
    const url = req.query.url;
    if (!url) return res.status(400).json({ error: 'Missing ?url= parameter' });
    const { body } = await fetchUrl(url);
    const $ = cheerio.load(body);
    const bodyText = $('body').text().replace(/\s+/g, ' ').trim();
    const stopWords = new Set(['the','a','an','is','are','was','were','be','been','being','have','has','had','do','does','did','will','would','could','should','may','might','shall','can','need','dare','ought','used','to','of','in','for','on','with','at','by','from','as','into','through','during','before','after','above','below','between','out','off','over','under','again','further','then','once','and','but','or','nor','not','so','yet','both','either','neither','each','every','all','any','few','more','most','other','some','such','no','only','own','same','than','too','very','just','that','this','these','those','it','its','i','me','my','we','our','you','your','he','him','his','she','her','they','them','their','what','which','who','whom','when','where','why','how','if','because','while','although','until','about','also','de','le','la','les','des','un','une','et','est','en','que','qui','dans','du','au','aux','pour','par','sur','pas','plus','ce','se','ou','il','ne','son','sa','ses','avec','tout','cette','ces','comme','mais','donc','alors']);
    const words = bodyText.toLowerCase().replace(/[^a-zàâéèêëïîôùûüÿçæœ\s]/g, '').split(/\s+/).filter(w => w.length > 2 && !stopWords.has(w));
    const freq = {};
    words.forEach(w => freq[w] = (freq[w] || 0) + 1);
    const keywords = Object.entries(freq).sort((a, b) => b[1] - a[1]).map(([word, count]) => ({
      word, count, density: +((count / words.length) * 100).toFixed(2),
    }));
    // Bigrams
    const bigrams = {};
    for (let i = 0; i < words.length - 1; i++) {
      const bg = words[i] + ' ' + words[i + 1];
      bigrams[bg] = (bigrams[bg] || 0) + 1;
    }
    const topBigrams = Object.entries(bigrams).sort((a, b) => b[1] - a[1]).slice(0, 10).map(([phrase, count]) => ({ phrase, count }));
    res.json({ url, total_words: words.length, single_keywords: keywords.slice(0, 30), top_bigrams: topBigrams });
  } catch (err) {
    res.status(500).json({ error: err.message });
  }
});

// --- Meta Tag Generator (AI) ---
app.post('/api/v1/meta/generate', async (req, res) => {
  try {
    const { topic, url, content } = req.body;
    if (!topic && !url) return res.status(400).json({ error: 'Provide topic or url' });
    let context = topic || '';
    if (url) {
      const { body } = await fetchUrl(url);
      const $ = cheerio.load(body);
      context = $('body').text().replace(/\s+/g, ' ').trim().substring(0, 2000);
    }
    if (content) context = content;

    const prompt = `Generate optimized SEO meta tags for the following content/topic. Return ONLY valid JSON with these fields: title (50-60 chars), description (150-160 chars), og_title, og_description, og_image_alt, keywords (array of 5-8 keywords), schema_type (appropriate Schema.org type).

Content: ${context.substring(0, 1500)}`;

    const llmResponse = await callLLM(prompt);
    try {
      const json = JSON.parse(llmResponse.match(/\{[\s\S]*\}/)?.[0] || '{}');
      res.json({ ...json, generated_by: 'openclaw-seo-toolkit' });
    } catch {
      res.json({ raw_response: llmResponse, generated_by: 'openclaw-seo-toolkit' });
    }
  } catch (err) {
    res.status(500).json({ error: err.message });
  }
});

// --- Meta Tag Checker ---
app.get('/api/v1/meta/check', async (req, res) => {
  try {
    const url = req.query.url;
    if (!url) return res.status(400).json({ error: 'Missing ?url= parameter' });
    const { body, headers } = await fetchUrl(url);
    const $ = cheerio.load(body);

    const meta = {
      title: $('title').text().trim(),
      description: $('meta[name="description"]').attr('content') || '',
      robots: $('meta[name="robots"]').attr('content') || '',
      canonical: $('link[rel="canonical"]').attr('href') || '',
      viewport: $('meta[name="viewport"]').attr('content') || '',
      charset: $('meta[charset]').attr('charset') || '',
      og: {},
      twitter: {},
      http_equiv: {},
    };

    $('meta[property^="og:"]').each((_, el) => { meta.og[$(el).attr('property')] = $(el).attr('content'); });
    $('meta[name^="twitter:"]').each((_, el) => { meta.twitter[$(el).attr('name')] = $(el).attr('content'); });
    $('meta[http-equiv]').each((_, el) => { meta.http_equiv[$(el).attr('http-equiv')] = $(el).attr('content'); });

    meta.http_status = headers['status'] || 200;
    meta.content_type = headers['content-type'] || '';

    res.json({ url, meta });
  } catch (err) {
    res.status(500).json({ error: err.message });
  }
});

// --- Heading Structure ---
app.get('/api/v1/headings', async (req, res) => {
  try {
    const url = req.query.url;
    if (!url) return res.status(400).json({ error: 'Missing ?url= parameter' });
    const { body } = await fetchUrl(url);
    const $ = cheerio.load(body);
    const headings = [];
    for (let i = 1; i <= 6; i++) {
      $(`h${i}`).each((_, el) => {
        headings.push({ level: i, text: $(el).text().trim().substring(0, 200) });
      });
    }
    const issues = [];
    const h1s = headings.filter(h => h.level === 1);
    if (h1s.length === 0) issues.push('No H1 heading found');
    if (h1s.length > 1) issues.push(`Multiple H1 tags (${h1s.length})`);
    // Check hierarchy
    for (let i = 1; i < headings.length; i++) {
      if (headings[i].level > headings[i - 1].level + 1) {
        issues.push(`Heading hierarchy skip: H${headings[i - 1].level} → H${headings[i].level}`);
        break;
      }
    }
    res.json({ url, headings, count: headings.length, issues });
  } catch (err) {
    res.status(500).json({ error: err.message });
  }
});

// --- Image Audit ---
app.get('/api/v1/images', async (req, res) => {
  try {
    const url = req.query.url;
    if (!url) return res.status(400).json({ error: 'Missing ?url= parameter' });
    const { body } = await fetchUrl(url);
    const $ = cheerio.load(body);
    const images = [];
    $('img').each((_, el) => {
      images.push({
        src: $(el).attr('src') || '',
        alt: $(el).attr('alt') || null,
        title: $(el).attr('title') || null,
        width: $(el).attr('width') || null,
        height: $(el).attr('height') || null,
        loading: $(el).attr('loading') || null,
        has_alt: !!$(el).attr('alt'),
        alt_empty: $(el).attr('alt') === '',
      });
    });
    const missingAlt = images.filter(i => !i.has_alt);
    const emptyAlt = images.filter(i => i.alt_empty);
    const noDimensions = images.filter(i => !i.width || !i.height);
    res.json({
      url, total: images.length,
      missing_alt: missingAlt.length,
      empty_alt: emptyAlt.length,
      missing_dimensions: noDimensions.length,
      images,
      recommendations: [
        missingAlt.length > 0 ? `${missingAlt.length} images missing alt attribute` : null,
        emptyAlt.length > 0 ? `${emptyAlt.length} images with empty alt=""` : null,
        noDimensions.length > 0 ? `${noDimensions.length} images missing width/height (CLS issue)` : null,
      ].filter(Boolean),
    });
  } catch (err) {
    res.status(500).json({ error: err.message });
  }
});

// --- Link Analysis ---
app.get('/api/v1/links', async (req, res) => {
  try {
    const url = req.query.url;
    if (!url) return res.status(400).json({ error: 'Missing ?url= parameter' });
    const { body } = await fetchUrl(url);
    const $ = cheerio.load(body);
    const parsedUrl = new URL(url);
    const internal = [], external = [], broken = [];
    $('a[href]').each((_, el) => {
      const href = $(el).attr('href') || '';
      const text = $(el).text().trim().substring(0, 100);
      const rel = $(el).attr('rel') || '';
      const target = $(el).attr('target') || '';
      try {
        const linkUrl = new URL(href, url);
        const entry = { href: linkUrl.href, text, rel, target };
        if (linkUrl.hostname === parsedUrl.hostname) internal.push(entry);
        else external.push(entry);
      } catch {
        broken.push({ href, text });
      }
    });
    res.json({
      url,
      internal: { count: internal.length, links: internal.slice(0, 50) },
      external: { count: external.length, links: external.slice(0, 50) },
      broken: { count: broken.length, links: broken },
    });
  } catch (err) {
    res.status(500).json({ error: err.message });
  }
});

// --- Structured Data ---
app.get('/api/v1/structured', async (req, res) => {
  try {
    const url = req.query.url;
    if (!url) return res.status(400).json({ error: 'Missing ?url= parameter' });
    const { body } = await fetchUrl(url);
    const $ = cheerio.load(body);
    const jsonLd = [];
    $('script[type="application/ld+json"]').each((_, el) => {
      try { jsonLd.push(JSON.parse($(el).html())); } catch { jsonLd.push({ error: 'Invalid JSON-LD', raw: $(el).html()?.substring(0, 200) }); }
    });
    const microdata = $('[itemscope]').length;
    const rdfa = $('[typeof]').length;
    res.json({
      url,
      json_ld: { count: jsonLd.length, data: jsonLd },
      microdata_count: microdata,
      rdfa_count: rdfa,
      has_structured_data: jsonLd.length > 0 || microdata > 0 || rdfa > 0,
    });
  } catch (err) {
    res.status(500).json({ error: err.message });
  }
});

// --- AI Content Audit ---
app.post('/api/v1/content/audit', async (req, res) => {
  try {
    const { url, content, target_keyword } = req.body;
    let text = content || '';
    if (url) {
      const { body } = await fetchUrl(url);
      const $ = cheerio.load(body);
      text = $('body').text().replace(/\s+/g, ' ').trim().substring(0, 3000);
    }
    if (!text) return res.status(400).json({ error: 'Provide url or content' });

    const prompt = `You are an SEO expert. Audit the following web page content${target_keyword ? ` for the target keyword "${target_keyword}"` : ''}. 

Provide a JSON response with:
- score (0-100)
- strengths (array of strings)
- weaknesses (array of strings)
- recommendations (array of actionable improvements)
- keyword_usage (if target keyword provided: how many times used, density %, in headings yes/no)
- readability_score (0-100, estimated)
- content_type (blog/product/landing/news/other)

Content: ${text.substring(0, 2500)}`;

    const llmResponse = await callLLM(prompt, 1500);
    try {
      const json = JSON.parse(llmResponse.match(/\{[\s\S]*\}/)?.[0] || '{}');
      res.json({ url, audit: json, audited_by: 'openclaw-seo-toolkit' });
    } catch {
      res.json({ url, audit: { raw: llmResponse }, audited_by: 'openclaw-seo-toolkit' });
    }
  } catch (err) {
    res.status(500).json({ error: err.message });
  }
});

// --- AI Content Rewrite ---
app.post('/api/v1/content/rewrite', async (req, res) => {
  try {
    const { content, target_keyword, style } = req.body;
    if (!content) return res.status(400).json({ error: 'Provide content' });

    const prompt = `Rewrite the following content to be SEO-optimized${target_keyword ? ` for the keyword "${target_keyword}"` : ''}. Style: ${style || 'professional, engaging'}. Maintain the core message but improve readability, add relevant keywords naturally, and ensure proper heading structure. Return the rewritten content only.

Original: ${content.substring(0, 2000)}`;

    const rewritten = await callLLM(prompt, 2000);
    res.json({ original_length: content.length, rewritten, rewritten_by: 'openclaw-seo-toolkit' });
  } catch (err) {
    res.status(500).json({ error: err.message });
  }
});

// --- robots.txt parser ---
app.get('/api/v1/robots', async (req, res) => {
  try {
    let url = req.query.url;
    if (!url) return res.status(400).json({ error: 'Missing ?url= parameter' });
    const parsed = new URL(url);
    const robotsUrl = `${parsed.protocol}//${parsed.host}/robots.txt`;
    try {
      const { body } = await fetchUrl(robotsUrl);
      const lines = body.split('\n');
      const groups = [];
      let current = null;
      for (const line of lines) {
        const trimmed = line.trim();
        if (!trimmed || trimmed.startsWith('#')) continue;
        const [key, ...valParts] = trimmed.split(':');
        const value = valParts.join(':').trim();
        if (key.toLowerCase() === 'user-agent') {
          if (current) groups.push(current);
          current = { user_agent: value, allows: [], disallows: [], sitemaps: [] };
        } else if (current) {
          if (key.toLowerCase() === 'allow') current.allows.push(value);
          else if (key.toLowerCase() === 'disallow') current.disallows.push(value);
          else if (key.toLowerCase() === 'sitemap') current.sitemaps.push(value);
        }
      }
      if (current) groups.push(current);
      res.json({ url: robotsUrl, found: true, groups, raw: body.substring(0, 5000) });
    } catch {
      res.json({ url: robotsUrl, found: false, message: 'No robots.txt found' });
    }
  } catch (err) {
    res.status(500).json({ error: err.message });
  }
});

// --- Sitemap parser ---
app.get('/api/v1/sitemap', async (req, res) => {
  try {
    let url = req.query.url;
    if (!url) return res.status(400).json({ error: 'Missing ?url= parameter' });
    const parsed = new URL(url);
    // Try common sitemap locations
    const sitemapUrls = [
      `${parsed.protocol}//${parsed.host}/sitemap.xml`,
      `${parsed.protocol}//${parsed.host}/sitemap_index.xml`,
    ];
    // First check robots.txt for sitemap
    try {
      const { body: robotsBody } = await fetchUrl(`${parsed.protocol}//${parsed.host}/robots.txt`);
      const sitemapMatches = robotsBody.match(/Sitemap:\s*(.+)/gi);
      if (sitemapMatches) {
        sitemapMatches.forEach(m => {
          const sUrl = m.replace(/Sitemap:\s*/i, '').trim();
          if (!sitemapUrls.includes(sUrl)) sitemapUrls.unshift(sUrl);
        });
      }
    } catch {}

    for (const sUrl of sitemapUrls) {
      try {
        const { body } = await fetchUrl(sUrl);
        if (!body.includes('<?xml') && !body.includes('<urlset') && !body.includes('<sitemapindex')) continue;
        const $ = cheerio.load(body, { xmlMode: true });
        const urls = [];
        $('url').each((_, el) => {
          urls.push({
            loc: $(el).find('loc').text(),
            lastmod: $(el).find('lastmod').text() || null,
            changefreq: $(el).find('changefreq').text() || null,
            priority: $(el).find('priority').text() || null,
          });
        });
        const subSitemaps = [];
        $('sitemap').each((_, el) => {
          subSitemaps.push({ loc: $(el).find('loc').text(), lastmod: $(el).find('lastmod').text() || null });
        });
        return res.json({ sitemap_url: sUrl, found: true, url_count: urls.length, urls: urls.slice(0, 100), sub_sitemaps: subSitemaps });
      } catch {}
    }
    res.json({ tried: sitemapUrls, found: false, message: 'No sitemap.xml found' });
  } catch (err) {
    res.status(500).json({ error: err.message });
  }
});

// --- Start ---
app.listen(PORT, '0.0.0.0', () => {
  console.log(`OpenClaw SEO Toolkit running on port ${PORT}`);
});
