import os
import io
import json
import hashlib
import tempfile
from flask import Flask, request, jsonify, send_file, Response
from PIL import Image, ImageDraw, ImageFont
import requests
import qrcode
import cairosvg

app = Flask(__name__)

PROXY_URL = os.environ.get('LLM_PROXY_URL', 'http://localhost:8088')
PORT = int(os.environ.get('IMAGE_API_PORT', 8091))
MAX_SIZE = 4096

PRICING = {
    'placeholder': 2,
    'text_card': 5,
    'banner': 8,
    'qr_code': 3,
    'chart': 10,
    'svg_icon': 5,
    'og_image': 8,
    'avatar': 5,
}

# ── Auth ──────────────────────────────────────────────────────────────
def require_api_key():
    auth = request.headers.get('Authorization', '')
    if not auth.startswith('Bearer '):
        return None, jsonify({'error': 'Missing API key'}), 401
    api_key = auth[7:]
    try:
        r = requests.get(f'{PROXY_URL}/v1/account',
                         headers={'Authorization': f'Bearer {api_key}'},
                         timeout=5)
        if r.status_code == 401:
            return None, jsonify({'error': 'Invalid API key'}), 401
        data = r.json()
        if data.get('credits_satoshis', 0) <= 0:
            return None, jsonify({'error': 'Insufficient credits'}), 402
        return api_key, None, None
    except Exception as e:
        return None, jsonify({'error': f'Auth failed: {str(e)}'}), 500


def get_font(size=40, bold=False):
    """Get a font, falling back gracefully."""
    font_paths = [
        '/usr/share/fonts/truetype/dejavu/DejaVuSans-Bold.ttf' if bold else '/usr/share/fonts/truetype/dejavu/DejaVuSans.ttf',
        '/usr/share/fonts/truetype/liberation/LiberationSans-Bold.ttf' if bold else '/usr/share/fonts/truetype/liberation/LiberationSans-Regular.ttf',
    ]
    for fp in font_paths:
        if os.path.exists(fp):
            return ImageFont.truetype(fp, size)
    return ImageFont.load_default()


def parse_color(c, default=(255, 255, 255)):
    """Parse hex color or name to RGB tuple."""
    if not c:
        return default
    c = c.strip().lstrip('#')
    color_map = {
        'white': (255, 255, 255), 'black': (0, 0, 0), 'red': (220, 53, 69),
        'green': (40, 167, 69), 'blue': (0, 123, 255), 'yellow': (255, 193, 7),
        'orange': (253, 126, 20), 'purple': (111, 66, 193), 'pink': (232, 62, 140),
        'gray': (108, 117, 125), 'grey': (108, 117, 125),
    }
    if c.lower() in color_map:
        return color_map[c.lower()]
    try:
        return tuple(int(c[i:i+2], 16) for i in (0, 2, 4))
    except:
        return default


# ── Landing Page ──────────────────────────────────────────────────────
@app.route('/')
def index():
    return '''<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>OpenClaw Image API - Generate Images, QR Codes, Banners</title>
    <style>
        * { margin: 0; padding: 0; box-sizing: border-box; }
        body { font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif; background: #0a0a0a; color: #e0e0e0; }
        .hero { text-align: center; padding: 80px 20px 40px; }
        .hero h1 { font-size: 2.5em; color: #fff; margin-bottom: 10px; }
        .hero h1 span { color: #8b5cf6; }
        .hero p { color: #888; font-size: 1.1em; max-width: 600px; margin: 0 auto; }
        .services { display: grid; grid-template-columns: repeat(auto-fit, minmax(280px, 1fr)); gap: 20px; max-width: 1000px; margin: 40px auto; padding: 0 20px; }
        .card { background: #1a1a1a; border: 1px solid #333; border-radius: 12px; padding: 30px; transition: border-color 0.3s; }
        .card:hover { border-color: #8b5cf6; }
        .card h3 { color: #8b5cf6; margin-bottom: 10px; }
        .card p { color: #999; font-size: 0.9em; line-height: 1.5; }
        .card .price { color: #10b981; font-weight: bold; margin-top: 15px; }
        .api-section { max-width: 900px; margin: 60px auto; padding: 0 20px; }
        .api-section h2 { color: #fff; margin-bottom: 20px; text-align: center; }
        .endpoint { background: #111; border: 1px solid #333; border-radius: 8px; padding: 20px; margin-bottom: 15px; }
        .method { background: #8b5cf6; color: #fff; padding: 3px 10px; border-radius: 4px; font-weight: bold; font-size: 0.85em; }
        .path { color: #8b5cf6; font-family: monospace; margin-left: 10px; }
        .desc { color: #888; margin-top: 10px; font-size: 0.9em; }
        .params { color: #999; font-size: 0.85em; margin-top: 8px; }
        .params code { color: #f59e0b; background: #1a1a1a; padding: 2px 6px; border-radius: 3px; }
        .code { background: #0d0d0d; border: 1px solid #222; border-radius: 6px; padding: 15px; margin: 10px 0; font-family: monospace; font-size: 0.8em; color: #10b981; overflow-x: auto; }
        .wallets { text-align: center; padding: 40px 20px; border-top: 1px solid #222; margin-top: 60px; }
        .wallets h3 { color: #8b5cf6; margin-bottom: 20px; }
        .wallet { background: #111; border: 1px solid #333; border-radius: 8px; padding: 15px; margin: 10px auto; max-width: 500px; font-family: monospace; font-size: 0.85em; word-break: break-all; }
    </style>
</head>
<body>
    <div class="hero">
        <h1>OpenClaw <span>Image API</span></h1>
        <p>Generate images, QR codes, banners, charts, and social media graphics programmatically. Pay-per-use with Bitcoin.</p>
    </div>

    <div class="services">
        <div class="card"><h3>Placeholder Images</h3><p>Custom size, color, and text placeholder images for development and design.</p><div class="price">2 sats per image</div></div>
        <div class="card"><h3>Text Cards</h3><p>Beautiful quote cards, announcements, and text-based images with custom styling.</p><div class="price">5 sats per card</div></div>
        <div class="card"><h3>Social Banners</h3><p>Twitter headers, LinkedIn banners, YouTube thumbnails with text overlays.</p><div class="price">8 sats per banner</div></div>
        <div class="card"><h3>QR Codes</h3><p>Generate QR codes for URLs, text, WiFi configs, vCards with custom colors.</p><div class="price">3 sats per code</div></div>
        <div class="card"><h3>Charts & Graphs</h3><p>Bar charts, pie charts, line graphs from JSON data. PNG/SVG output.</p><div class="price">10 sats per chart</div></div>
        <div class="card"><h3>SVG Icons</h3><p>Custom SVG icons and simple logos from text/shapes.</p><div class="price">5 sats per icon</div></div>
    </div>

    <div class="api-section">
        <h2>API Reference</h2>

        <div class="endpoint">
            <span class="method">GET</span><span class="path">/placeholder/{width}x{height}</span>
            <div class="desc">Generate a placeholder image with optional text and colors.</div>
            <div class="params">Params: <code>bg</code> (hex), <code>fg</code> (hex), <code>text</code> (string), <code>format</code> (png|jpg|webp)</div>
            <div class="code">curl "http://localhost:8091/placeholder/800x400?bg=1a1a2e&fg=e94560&text=Hello" -H "Authorization: Bearer KEY" -o banner.png</div>
        </div>

        <div class="endpoint">
            <span class="method">POST</span><span class="path">/generate/text-card</span>
            <div class="desc">Generate a text card with custom styling.</div>
            <div class="params">Body: <code>text</code>, <code>author</code>, <code>bg_color</code>, <code>text_color</code>, <code>width</code>, <code>height</code>, <code>font_size</code></div>
            <div class="code">curl -X POST http://localhost:8091/generate/text-card \\
  -H "Authorization: Bearer KEY" -H "Content-Type: application/json" \\
  -d '{"text":"Build in public","bg_color":"#0f0f23","text_color":"#e94560","width":1200,"height":630}' -o card.png</div>
        </div>

        <div class="endpoint">
            <span class="method">POST</span><span class="path">/generate/banner</span>
            <div class="desc">Generate social media banners (Twitter, LinkedIn, YouTube).</div>
            <div class="params">Body: <code>title</code>, <code>subtitle</code>, <code>template</code> (twitter|linkedin|youtube|custom), <code>bg_color</code>, <code>accent_color</code></div>
            <div class="code">curl -X POST http://localhost:8091/generate/banner \\
  -H "Authorization: Bearer KEY" -H "Content-Type: application/json" \\
  -d '{"title":"OpenClaw","subtitle":"AI Services","template":"twitter"}' -o banner.png</div>
        </div>

        <div class="endpoint">
            <span class="method">GET</span><span class="path">/qr</span>
            <div class="desc">Generate a QR code.</div>
            <div class="params">Params: <code>data</code> (required), <code>size</code>, <code>fg</code>, <code>bg</code>, <code>format</code> (png|svg)</div>
            <div class="code">curl "http://localhost:8091/qr?data=https://example.com&size=300&fg=8b5cf6" -H "Authorization: Bearer KEY" -o qr.png</div>
        </div>

        <div class="endpoint">
            <span class="method">POST</span><span class="path">/generate/chart</span>
            <div class="desc">Generate a chart from data.</div>
            <div class="params">Body: <code>type</code> (bar|pie|line), <code>data</code> (array), <code>labels</code> (array), <code>title</code>, <code>colors</code>, <code>width</code>, <code>height</code></div>
            <div class="code">curl -X POST http://localhost:8091/generate/chart \\
  -H "Authorization: Bearer KEY" -H "Content-Type: application/json" \\
  -d '{"type":"bar","data":[30,50,20,40],"labels":["Q1","Q2","Q3","Q4"],"title":"Revenue"}' -o chart.png</div>
        </div>
    </div>

    <div class="wallets">
        <h3>Add Credits (Pay with Crypto)</h3>
        <div class="wallet"><div style="color:#888;font-size:0.8em">BTC</div>1GXpyCzAhvVNcbjnkytRTe1QoVoAen7nfP</div>
        <div class="wallet"><div style="color:#888;font-size:0.8em">ETH</div>0xCF295d87E1534538FDac3b1f98746aF4A3E47352</div>
    </div>
</body>
</html>'''


@app.route('/health')
def health():
    return jsonify({'status': 'healthy', 'service': 'OpenClaw Image API', 'version': '1.0.0', 'pricing': PRICING})


# ── Placeholder Image ────────────────────────────────────────────────
@app.route('/placeholder/<int:width>x<int:height>')
def placeholder(width, height):
    api_key, err, code = require_api_key()
    if err: return err, code

    width = min(width, MAX_SIZE)
    height = min(height, MAX_SIZE)
    bg = parse_color(request.args.get('bg', '2d2d2d'), (45, 45, 45))
    fg = parse_color(request.args.get('fg', '888888'), (136, 136, 136))
    text = request.args.get('text', f'{width}x{height}')
    fmt = request.args.get('format', 'png').lower()

    img = Image.new('RGB', (width, height), bg)
    draw = ImageDraw.Draw(img)

    # Draw grid pattern
    grid_color = tuple(min(255, c + 15) for c in bg)
    for x in range(0, width, 50):
        draw.line([(x, 0), (x, height)], fill=grid_color, width=1)
    for y in range(0, height, 50):
        draw.line([(0, y), (width, y)], fill=grid_color, width=1)

    # Center text
    font_size = min(width, height) // 8
    font = get_font(font_size)
    bbox = draw.textbbox((0, 0), text, font=font)
    tw, th = bbox[2] - bbox[0], bbox[3] - bbox[1]
    draw.text(((width - tw) / 2, (height - th) / 2), text, fill=fg, font=font)

    return _serve_image(img, fmt, 'placeholder')


# ── Text Card ────────────────────────────────────────────────────────
@app.route('/generate/text-card', methods=['POST'])
def text_card():
    api_key, err, code = require_api_key()
    if err: return err, code

    data = request.get_json() or {}
    text = data.get('text', 'Hello World')
    author = data.get('author', '')
    bg_color = parse_color(data.get('bg_color', '#0f0f23'), (15, 15, 35))
    text_color = parse_color(data.get('text_color', '#ffffff'), (255, 255, 255))
    accent_color = parse_color(data.get('accent_color', '#8b5cf6'), (139, 92, 246))
    width = min(data.get('width', 1200), MAX_SIZE)
    height = min(data.get('height', 630), MAX_SIZE)
    font_size = data.get('font_size', 48)
    fmt = data.get('format', 'png')

    img = Image.new('RGB', (width, height), bg_color)
    draw = ImageDraw.Draw(img)

    # Accent line
    draw.rectangle([(60, 60), (64, height - 60)], fill=accent_color)

    # Word-wrap text
    font = get_font(font_size)
    margin = 100
    max_w = width - margin * 2
    lines = _wrap_text(draw, text, font, max_w)

    y = (height - len(lines) * (font_size + 10)) // 2
    for line in lines:
        bbox = draw.textbbox((0, 0), line, font=font)
        lw = bbox[2] - bbox[0]
        draw.text(((width - lw) / 2, y), line, fill=text_color, font=font)
        y += font_size + 10

    # Author
    if author:
        author_font = get_font(font_size // 2)
        draw.text((margin, height - 80), f'— {author}', fill=accent_color, font=author_font)

    return _serve_image(img, fmt, 'text-card')


# ── Banner ───────────────────────────────────────────────────────────
TEMPLATES = {
    'twitter': (1500, 500),
    'linkedin': (1584, 396),
    'youtube': (1280, 720),
    'og': (1200, 630),
    'instagram': (1080, 1080),
    'facebook': (1200, 630),
}

@app.route('/generate/banner', methods=['POST'])
def banner():
    api_key, err, code = require_api_key()
    if err: return err, code

    data = request.get_json() or {}
    title = data.get('title', 'OpenClaw')
    subtitle = data.get('subtitle', '')
    template = data.get('template', 'twitter')
    bg_color = parse_color(data.get('bg_color', '#0f0f23'), (15, 15, 35))
    accent_color = parse_color(data.get('accent_color', '#8b5cf6'), (139, 92, 246))
    text_color = parse_color(data.get('text_color', '#ffffff'), (255, 255, 255))
    fmt = data.get('format', 'png')

    w, h = TEMPLATES.get(template, (data.get('width', 1200), data.get('height', 630)))
    w, h = min(w, MAX_SIZE), min(h, MAX_SIZE)

    img = Image.new('RGB', (w, h), bg_color)
    draw = ImageDraw.Draw(img)

    # Gradient accent bar at bottom
    for i in range(8):
        alpha = 255 - i * 30
        color = tuple(max(0, c - i * 10) for c in accent_color)
        draw.rectangle([(0, h - 8 + i), (w, h - 7 + i)], fill=color)

    # Decorative circles
    for cx, cy, r in [(w - 100, 80, 60), (w - 180, 140, 30), (80, h - 100, 40)]:
        draw.ellipse([(cx - r, cy - r), (cx + r, cy + r)], outline=accent_color, width=2)

    # Title
    title_size = min(w, h) // 6
    title_font = get_font(title_size, bold=True)
    bbox = draw.textbbox((0, 0), title, font=title_font)
    tw = bbox[2] - bbox[0]
    ty = h // 2 - title_size - 10
    draw.text(((w - tw) / 2, ty), title, fill=text_color, font=title_font)

    # Subtitle
    if subtitle:
        sub_size = title_size // 2
        sub_font = get_font(sub_size)
        bbox = draw.textbbox((0, 0), subtitle, font=sub_font)
        sw = bbox[2] - bbox[0]
        draw.text(((w - sw) / 2, ty + title_size + 20), subtitle, fill=accent_color, font=sub_font)

    return _serve_image(img, fmt, 'banner')


# ── QR Code ──────────────────────────────────────────────────────────
@app.route('/qr')
def qr_code():
    api_key, err, code = require_api_key()
    if err: return err, code

    data = request.args.get('data', 'https://openclaw.ai')
    size = min(int(request.args.get('size', 300)), MAX_SIZE)
    fg = request.args.get('fg', '000000').lstrip('#')
    bg = request.args.get('bg', 'ffffff').lstrip('#')
    fmt = request.args.get('format', 'png').lower()

    qr = qrcode.QRCode(version=1, error_correction=qrcode.constants.ERROR_CORRECT_H, box_size=10, border=2)
    qr.add_data(data)
    qr.make(fit=True)

    if fmt == 'svg':
        matrix = qr.get_matrix()
        rows = len(matrix)
        box_size = size // (rows + 4)
        svg_parts = [f'<svg xmlns="http://www.w3.org/2000/svg" width="{size}" height="{size}" viewBox="0 0 {rows} {rows}">']
        svg_parts.append(f'<rect width="{rows}" height="{rows}" fill="#{bg}"/>')
        for r in range(rows):
            for c in range(rows):
                if matrix[r][c]:
                    svg_parts.append(f'<rect x="{c}" y="{r}" width="1" height="1" fill="#{fg}"/>')
        svg_parts.append('</svg>')
        return Response('\n'.join(svg_parts), mimetype='image/svg+xml')

    img = qr.make_image(fill_color=f'#{fg}', back_color=f'#{bg}').convert('RGB')
    img = img.resize((size, size), Image.LANCZOS)
    return _serve_image(img, 'png', 'qr')


# ── Chart ────────────────────────────────────────────────────────────
@app.route('/generate/chart', methods=['POST'])
def chart():
    api_key, err, code = require_api_key()
    if err: return err, code

    data = request.get_json() or {}
    chart_type = data.get('type', 'bar')
    values = data.get('data', [10, 20, 30, 40])
    labels = data.get('labels', [f'Item {i+1}' for i in range(len(values))])
    title = data.get('title', 'Chart')
    colors_hex = data.get('colors', ['#8b5cf6', '#10b981', '#f59e0b', '#ef4444', '#3b82f6', '#ec4899'])
    width = min(data.get('width', 800), MAX_SIZE)
    height = min(data.get('height', 500), MAX_SIZE)
    fmt = data.get('format', 'png')

    img = Image.new('RGB', (width, height), (15, 15, 35))
    draw = ImageDraw.Draw(img)

    colors = [parse_color(c, (139, 92, 246)) for c in colors_hex]

    # Title
    title_font = get_font(24, bold=True)
    draw.text((30, 20), title, fill=(255, 255, 255), font=title_font)

    chart_top = 70
    chart_bottom = height - 60
    chart_left = 80
    chart_right = width - 40
    chart_h = chart_bottom - chart_top
    chart_w = chart_right - chart_left

    if chart_type == 'bar':
        max_val = max(values) if values else 1
        bar_count = len(values)
        bar_w = max(20, (chart_w - 20 * bar_count) // bar_count)
        gap = 10

        for i, (val, label) in enumerate(zip(values, labels)):
            bar_h = int((val / max_val) * chart_h * 0.85)
            x = chart_left + i * (bar_w + gap) + gap
            y = chart_bottom - bar_h
            color = colors[i % len(colors)]
            draw.rectangle([(x, y), (x + bar_w, chart_bottom)], fill=color)

            # Value on top
            val_font = get_font(14)
            draw.text((x + bar_w // 2 - 10, y - 20), str(val), fill=(200, 200, 200), font=val_font)

            # Label below
            lbl_font = get_font(12)
            draw.text((x, chart_bottom + 5), label[:10], fill=(150, 150, 150), font=lbl_font)

    elif chart_type == 'pie':
        total = sum(values) or 1
        cx, cy = chart_left + chart_w // 2, chart_top + chart_h // 2
        radius = min(chart_w, chart_h) // 2 - 30
        start_angle = 0

        for i, (val, label) in enumerate(zip(values, labels)):
            sweep = (val / total) * 360
            color = colors[i % len(colors)]
            draw.pieslice([(cx - radius, cy - radius), (cx + radius, cy + radius)],
                         start=start_angle, end=start_angle + sweep, fill=color)
            start_angle += sweep

        # Legend
        lx = chart_right - 150
        ly = chart_top + 10
        legend_font = get_font(12)
        for i, (val, label) in enumerate(zip(values, labels)):
            draw.rectangle([(lx, ly), (lx + 15, ly + 15)], fill=colors[i % len(colors)])
            draw.text((lx + 20, ly), f'{label}: {val}', fill=(200, 200, 200), font=legend_font)
            ly += 22

    elif chart_type == 'line':
        max_val = max(values) if values else 1
        points = []
        step = chart_w // max(len(values) - 1, 1) if len(values) > 1 else chart_w

        for i, val in enumerate(values):
            x = chart_left + i * step
            y = chart_bottom - int((val / max_val) * chart_h * 0.85)
            points.append((x, y))

        if len(points) > 1:
            for i in range(len(points) - 1):
                color = colors[i % len(colors)]
                draw.line([points[i], points[i + 1]], fill=color, width=3)

        for i, (pt, val, label) in enumerate(zip(points, values, labels)):
            color = colors[i % len(colors)]
            draw.ellipse([(pt[0] - 5, pt[1] - 5), (pt[0] + 5, pt[1] + 5)], fill=color)
            val_font = get_font(12)
            draw.text((pt[0] - 10, pt[1] - 20), str(val), fill=(200, 200, 200), font=val_font)
            draw.text((pt[0] - 10, chart_bottom + 5), label[:8], fill=(150, 150, 150), font=val_font)

    # Axis lines
    draw.line([(chart_left, chart_top), (chart_left, chart_bottom)], fill=(80, 80, 80), width=1)
    draw.line([(chart_left, chart_bottom), (chart_right, chart_bottom)], fill=(80, 80, 80), width=1)

    return _serve_image(img, fmt, 'chart')


# ── SVG Icon ─────────────────────────────────────────────────────────
@app.route('/generate/svg-icon', methods=['POST'])
def svg_icon():
    api_key, err, code = require_api_key()
    if err: return err, code

    data = request.get_json() or {}
    text = data.get('text', 'OC')
    size = min(data.get('size', 200), MAX_SIZE)
    bg = data.get('bg_color', '#8b5cf6').lstrip('#')
    fg = data.get('text_color', '#ffffff').lstrip('#')
    shape = data.get('shape', 'circle')
    fmt = data.get('format', 'svg')

    r = size // 2
    svg = f'<svg xmlns="http://www.w3.org/2000/svg" width="{size}" height="{size}" viewBox="0 0 {size} {size}">'

    if shape == 'circle':
        svg += f'<circle cx="{r}" cy="{r}" r="{r}" fill="#{bg}"/>'
        svg += f'<text x="{r}" y="{r}" text-anchor="middle" dominant-baseline="central" font-family="Arial,sans-serif" font-size="{size//3}" font-weight="bold" fill="#{fg}">{text}</text>'
    elif shape == 'rounded':
        br = size // 8
        svg += f'<rect width="{size}" height="{size}" rx="{br}" fill="#{bg}"/>'
        svg += f'<text x="{r}" y="{r}" text-anchor="middle" dominant-baseline="central" font-family="Arial,sans-serif" font-size="{size//3}" font-weight="bold" fill="#{fg}">{text}</text>'
    else:
        svg += f'<rect width="{size}" height="{size}" fill="#{bg}"/>'
        svg += f'<text x="{r}" y="{r}" text-anchor="middle" dominant-baseline="central" font-family="Arial,sans-serif" font-size="{size//3}" font-weight="bold" fill="#{fg}">{text}</text>'

    svg += '</svg>'

    if fmt == 'png':
        png_data = cairosvg.svg2png(bytestring=svg.encode(), output_width=size, output_height=size)
        return send_file(io.BytesIO(png_data), mimetype='image/png', download_name='icon.png')

    return Response(svg, mimetype='image/svg+xml')


# ── Avatar ───────────────────────────────────────────────────────────
@app.route('/generate/avatar', methods=['POST'])
def avatar():
    api_key, err, code = require_api_key()
    if err: return err, code

    data = request.get_json() or {}
    name = data.get('name', 'User')
    size = min(data.get('size', 200), MAX_SIZE)
    bg = data.get('bg_color')
    fmt = data.get('format', 'png')

    # Deterministic color from name
    if not bg:
        h = int(hashlib.md5(name.encode()).hexdigest()[:6], 16)
        hue = h % 360
        import colorsys
        r, g, b = colorsys.hls_to_rgb(hue / 360, 0.45, 0.7)
        bg_color = (int(r * 255), int(g * 255), int(b * 255))
    else:
        bg_color = parse_color(bg, (139, 92, 246))

    img = Image.new('RGB', (size, size), bg_color)
    draw = ImageDraw.Draw(img)

    # Circle mask feel
    draw.ellipse([(0, 0), (size, size)], fill=bg_color)

    initials = ''.join(w[0].upper() for w in name.split()[:2])
    font = get_font(size // 2, bold=True)
    bbox = draw.textbbox((0, 0), initials, font=font)
    tw, th = bbox[2] - bbox[0], bbox[3] - bbox[1]
    draw.text(((size - tw) / 2, (size - th) / 2 - 5), initials, fill=(255, 255, 255), font=font)

    return _serve_image(img, fmt, 'avatar')


# ── Helpers ──────────────────────────────────────────────────────────
def _wrap_text(draw, text, font, max_width):
    words = text.split()
    lines = []
    current = ''
    for word in words:
        test = f'{current} {word}'.strip()
        bbox = draw.textbbox((0, 0), test, font=font)
        if bbox[2] - bbox[0] <= max_width:
            current = test
        else:
            if current:
                lines.append(current)
            current = word
    if current:
        lines.append(current)
    return lines or [text]


def _serve_image(img, fmt, name):
    buf = io.BytesIO()
    mime = 'image/png'
    ext = 'png'

    if fmt == 'jpg' or fmt == 'jpeg':
        img = img.convert('RGB')
        img.save(buf, 'JPEG', quality=90)
        mime = 'image/jpeg'
        ext = 'jpg'
    elif fmt == 'webp':
        img.save(buf, 'WEBP', quality=90)
        mime = 'image/webp'
        ext = 'webp'
    else:
        img.save(buf, 'PNG')
        mime = 'image/png'
        ext = 'png'

    buf.seek(0)
    return send_file(buf, mimetype=mime, download_name=f'{name}.{ext}')


if __name__ == '__main__':
    print(f'''
╔══════════════════════════════════════════════╗
║  OpenClaw Image API - RUNNING               ║
║  Port: {PORT}                                 ║
╚══════════════════════════════════════════════╝
    ''')
    app.run(host='0.0.0.0', port=PORT, debug=False)
