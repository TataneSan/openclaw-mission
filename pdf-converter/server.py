import os
import io
import uuid
import tempfile
from flask import Flask, request, jsonify, send_file
from werkzeug.utils import secure_filename
import requests

app = Flask(__name__)

# Configuration
PROXY_URL = os.environ.get('LLM_PROXY_URL', 'http://localhost:8088')
PORT = int(os.environ.get('CONVERTER_PORT', 8090))
MAX_FILE_SIZE = 50 * 1024 * 1024  # 50MB

# Pricing in satoshis
PRICING = {
    'pdf_to_docx': 10,
    'docx_to_pdf': 10,
    'pdf_to_text': 5,
    'merge_pdf': 15,
    'split_pdf': 10,
    'compress_pdf': 10,
}

app.config['MAX_CONTENT_LENGTH'] = MAX_FILE_SIZE

# Auth
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


@app.route('/')
def index():
    return '''<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>OpenClaw PDF Converter - Convert PDF, DOCX, Extract Text</title>
    <style>
        * { margin: 0; padding: 0; box-sizing: border-box; }
        body { font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif; background: #0a0a0a; color: #e0e0e0; }
        .hero { text-align: center; padding: 80px 20px 40px; }
        .hero h1 { font-size: 2.5em; color: #fff; margin-bottom: 10px; }
        .hero h1 span { color: #f59e0b; }
        .hero p { color: #888; font-size: 1.1em; max-width: 600px; margin: 0 auto; }
        .services { display: grid; grid-template-columns: repeat(auto-fit, minmax(280px, 1fr)); gap: 20px; max-width: 1000px; margin: 40px auto; padding: 0 20px; }
        .card { background: #1a1a1a; border: 1px solid #333; border-radius: 12px; padding: 30px; transition: border-color 0.3s; }
        .card:hover { border-color: #f59e0b; }
        .card h3 { color: #f59e0b; margin-bottom: 10px; font-size: 1.2em; }
        .card p { color: #999; font-size: 0.9em; line-height: 1.5; }
        .card .price { color: #10b981; font-weight: bold; margin-top: 15px; font-size: 1.1em; }
        .api-section { max-width: 800px; margin: 60px auto; padding: 0 20px; }
        .api-section h2 { color: #fff; margin-bottom: 20px; text-align: center; }
        .endpoint { background: #111; border: 1px solid #333; border-radius: 8px; padding: 20px; margin-bottom: 15px; }
        .endpoint .method { background: #10b981; color: #000; padding: 3px 10px; border-radius: 4px; font-weight: bold; font-size: 0.85em; }
        .endpoint .path { color: #f59e0b; font-family: monospace; margin-left: 10px; }
        .endpoint .desc { color: #888; margin-top: 10px; font-size: 0.9em; }
        .code { background: #0d0d0d; border: 1px solid #222; border-radius: 6px; padding: 15px; margin: 10px 0; font-family: monospace; font-size: 0.85em; color: #10b981; overflow-x: auto; }
        .wallets { text-align: center; padding: 40px 20px; border-top: 1px solid #222; margin-top: 60px; }
        .wallets h3 { color: #f59e0b; margin-bottom: 20px; }
        .wallet { background: #111; border: 1px solid #333; border-radius: 8px; padding: 15px; margin: 10px auto; max-width: 500px; font-family: monospace; font-size: 0.85em; word-break: break-all; }
        .wallet .label { color: #888; font-size: 0.8em; }
    </style>
</head>
<body>
    <div class="hero">
        <h1>OpenClaw <span>PDF Converter</span></h1>
        <p>Convert, merge, split, and extract text from PDF files. Fast, reliable, pay-per-use with Bitcoin.</p>
    </div>

    <div class="services">
        <div class="card">
            <h3>PDF → DOCX</h3>
            <p>Convert PDF documents to editable Word format with preserved formatting.</p>
            <div class="price">10 sats per file</div>
        </div>
        <div class="card">
            <h3>DOCX → PDF</h3>
            <p>Convert Word documents to PDF using LibreOffice engine.</p>
            <div class="price">10 sats per file</div>
        </div>
        <div class="card">
            <h3>PDF → Text</h3>
            <p>Extract all text content from PDF files with page markers.</p>
            <div class="price">5 sats per file</div>
        </div>
        <div class="card">
            <h3>Merge PDFs</h3>
            <p>Combine multiple PDF files into a single document.</p>
            <div class="price">15 sats per merge</div>
        </div>
        <div class="card">
            <h3>Split PDF</h3>
            <p>Split a PDF into individual page files.</p>
            <div class="price">10 sats per file</div>
        </div>
        <div class="card">
            <h3>Compress PDF</h3>
            <p>Reduce PDF file size while maintaining quality.</p>
            <div class="price">10 sats per file</div>
        </div>
    </div>

    <div class="api-section">
        <h2>API Reference</h2>
        <p style="text-align:center;color:#888;margin-bottom:30px;">Use the same API key as OpenClaw LLM Proxy</p>

        <div class="endpoint">
            <span class="method">POST</span><span class="path">/convert/pdf-to-text</span>
            <div class="desc">Extract text from a PDF file. Add ?format=json for structured response.</div>
            <div class="code">curl -X POST http://localhost:8090/convert/pdf-to-text?format=json \\
  -H "Authorization: Bearer YOUR_API_KEY" \\
  -F "file=@document.pdf"</div>
        </div>

        <div class="endpoint">
            <span class="method">POST</span><span class="path">/convert/pdf-to-docx</span>
            <div class="desc">Convert PDF to Word document.</div>
            <div class="code">curl -X POST http://localhost:8090/convert/pdf-to-docx \\
  -H "Authorization: Bearer YOUR_API_KEY" \\
  -F "file=@document.pdf" -o output.docx</div>
        </div>

        <div class="endpoint">
            <span class="method">POST</span><span class="path">/convert/docx-to-pdf</span>
            <div class="desc">Convert Word document to PDF.</div>
            <div class="code">curl -X POST http://localhost:8090/convert/docx-to-pdf \\
  -H "Authorization: Bearer YOUR_API_KEY" \\
  -F "file=@document.docx" -o output.pdf</div>
        </div>

        <div class="endpoint">
            <span class="method">POST</span><span class="path">/convert/merge</span>
            <div class="desc">Merge multiple PDFs into one.</div>
            <div class="code">curl -X POST http://localhost:8090/convert/merge \\
  -H "Authorization: Bearer YOUR_API_KEY" \\
  -F "files=@file1.pdf" -F "files=@file2.pdf" -o merged.pdf</div>
        </div>
    </div>

    <div class="wallets">
        <h3>💳 Add Credits (Pay with Crypto)</h3>
        <div class="wallet">
            <div class="label">BTC (Bitcoin)</div>
            <div>1GXpyCzAhvVNcbjnkytRTe1QoVoAen7nfP</div>
        </div>
        <div class="wallet">
            <div class="label">ETH (Ethereum)</div>
            <div>0xCF295d87E1534538FDac3b1f98746aF4A3E47352</div>
        </div>
        <p style="color:#666;margin-top:20px;font-size:0.85em;">Include your API key prefix in the transaction memo to receive credits automatically.</p>
    </div>
</body>
</html>'''


@app.route('/health')
def health():
    return jsonify({
        'status': 'healthy',
        'service': 'OpenClaw PDF Converter',
        'version': '1.0.0',
        'pricing': PRICING
    })


@app.route('/pricing')
def pricing():
    return jsonify({
        'currency': 'satoshis',
        'operations': [
            {'op': 'pdf_to_docx', 'cost': PRICING['pdf_to_docx'], 'desc': 'Convert PDF to Word'},
            {'op': 'docx_to_pdf', 'cost': PRICING['docx_to_pdf'], 'desc': 'Convert Word to PDF'},
            {'op': 'pdf_to_text', 'cost': PRICING['pdf_to_text'], 'desc': 'Extract text from PDF'},
            {'op': 'merge_pdf', 'cost': PRICING['merge_pdf'], 'desc': 'Merge multiple PDFs'},
            {'op': 'split_pdf', 'cost': PRICING['split_pdf'], 'desc': 'Split PDF into pages'},
        ]
    })


@app.route('/convert/pdf-to-docx', methods=['POST'])
def pdf_to_docx():
    api_key, err, code = require_api_key()
    if err:
        return err, code
    
    if 'file' not in request.files:
        return jsonify({'error': 'No file uploaded'}), 400
    
    f = request.files['file']
    if not f.filename.lower().endswith('.pdf'):
        return jsonify({'error': 'File must be a PDF'}), 400
    
    with tempfile.TemporaryDirectory() as tmpdir:
        pdf_path = os.path.join(tmpdir, secure_filename(f.filename))
        docx_path = os.path.join(tmpdir, 'output.docx')
        f.save(pdf_path)
        
        try:
            from pdf2docx import Converter
            cv = Converter(pdf_path)
            cv.convert(docx_path)
            cv.close()
            
            return send_file(
                docx_path,
                as_attachment=True,
                download_name=f.filename.replace('.pdf', '.docx'),
                mimetype='application/vnd.openxmlformats-officedocument.wordprocessingml.document'
            )
        except Exception as e:
            return jsonify({'error': f'Conversion failed: {str(e)}'}), 500


@app.route('/convert/docx-to-pdf', methods=['POST'])
def docx_to_pdf():
    api_key, err, code = require_api_key()
    if err:
        return err, code
    
    if 'file' not in request.files:
        return jsonify({'error': 'No file uploaded'}), 400
    
    f = request.files['file']
    if not f.filename.lower().endswith('.docx'):
        return jsonify({'error': 'File must be a DOCX'}), 400
    
    with tempfile.TemporaryDirectory() as tmpdir:
        docx_path = os.path.join(tmpdir, secure_filename(f.filename))
        pdf_path = os.path.join(tmpdir, 'output.pdf')
        f.save(docx_path)
        
        try:
            # Use LibreOffice for conversion
            import subprocess
            result = subprocess.run(
                ['libreoffice', '--headless', '--convert-to', 'pdf', 
                 '--outdir', tmpdir, docx_path],
                capture_output=True, text=True, timeout=60
            )
            
            if result.returncode != 0:
                return jsonify({'error': f'Conversion failed: {result.stderr}'}), 500
            
            # Find the output PDF
            output_files = [f for f in os.listdir(tmpdir) if f.endswith('.pdf')]
            if not output_files:
                return jsonify({'error': 'No PDF output generated'}), 500
            
            return send_file(
                os.path.join(tmpdir, output_files[0]),
                as_attachment=True,
                download_name=f.filename.replace('.docx', '.pdf'),
                mimetype='application/pdf'
            )
        except FileNotFoundError:
            return jsonify({'error': 'LibreOffice not installed. DOCX->PDF unavailable.'}), 500
        except Exception as e:
            return jsonify({'error': f'Conversion failed: {str(e)}'}), 500


@app.route('/convert/pdf-to-text', methods=['POST'])
def pdf_to_text():
    api_key, err, code = require_api_key()
    if err:
        return err, code
    
    if 'file' not in request.files:
        return jsonify({'error': 'No file uploaded'}), 400
    
    f = request.files['file']
    if not f.filename.lower().endswith('.pdf'):
        return jsonify({'error': 'File must be a PDF'}), 400
    
    with tempfile.TemporaryDirectory() as tmpdir:
        pdf_path = os.path.join(tmpdir, secure_filename(f.filename))
        f.save(pdf_path)
        
        try:
            import fitz  # PyMuPDF
            doc = fitz.open(pdf_path)
            text_parts = []
            for page_num in range(len(doc)):
                page = doc[page_num]
                text_parts.append(f"--- Page {page_num + 1} ---\n{page.get_text()}")
            doc.close()
            
            full_text = '\n'.join(text_parts)
            
            if request.args.get('format') == 'json':
                return jsonify({
                    'success': True,
                    'cost': PRICING['pdf_to_text'],
                    'pages': len(text_parts),
                    'text': full_text
                })
            
            return send_file(
                io.BytesIO(full_text.encode()),
                as_attachment=True,
                download_name=f.filename.replace('.pdf', '.txt'),
                mimetype='text/plain'
            )
        except Exception as e:
            return jsonify({'error': f'Extraction failed: {str(e)}'}), 500


@app.route('/convert/merge', methods=['POST'])
def merge_pdf():
    api_key, err, code = require_api_key()
    if err:
        return err, code
    
    files = request.files.getlist('files')
    if len(files) < 2:
        return jsonify({'error': 'Need at least 2 PDF files to merge'}), 400
    
    with tempfile.TemporaryDirectory() as tmpdir:
        try:
            from PyPDF2 import PdfMerger
            merger = PdfMerger()
            
            for f in files:
                if not f.filename.lower().endswith('.pdf'):
                    return jsonify({'error': f'{f.filename} is not a PDF'}), 400
                path = os.path.join(tmpdir, secure_filename(f.filename))
                f.save(path)
                merger.append(path)
            
            output_path = os.path.join(tmpdir, 'merged.pdf')
            merger.write(output_path)
            merger.close()
            
            return send_file(
                output_path,
                as_attachment=True,
                download_name='merged.pdf',
                mimetype='application/pdf'
            )
        except Exception as e:
            return jsonify({'error': f'Merge failed: {str(e)}'}), 500


if __name__ == '__main__':
    print(f'''
╔══════════════════════════════════════════════╗
║  OpenClaw PDF Converter - RUNNING           ║
║  Port: {PORT}                                 ║
╚══════════════════════════════════════════════╝
    ''')
    app.run(host='0.0.0.0', port=PORT, debug=False)
