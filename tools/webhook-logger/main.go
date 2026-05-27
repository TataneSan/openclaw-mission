package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

type LogEntry struct {
	ID        int       `json:"id"`
	Timestamp string    `json:"timestamp"`
	Method    string    `json:"method"`
	Path      string    `json:"path"`
	Headers   Headers   `json:"headers"`
	Body      string    `json:"body"`
	Remote    string    `json:"remote"`
}

type Headers map[string]string

type WebhookLogger struct {
	mu     sync.Mutex
	entries []LogEntry
	count  int
	logFile *os.File
}

func NewWebhookLogger(filename string) (*WebhookLogger, error) {
	var logFile *os.File
	var err error

	if filename != "" {
		logFile, err = os.Create(filename)
		if err != nil {
			return nil, fmt.Errorf("failed to create log file: %w", err)
		}
	}

	return &WebhookLogger{
		entries: make([]LogEntry, 0),
		logFile: logFile,
	}, nil
}

func (w *WebhookLogger) Close() {
	if w.logFile != nil {
		w.logFile.Close()
	}
}

func (w *WebhookLogger) AddEntry(entry LogEntry) {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.count++
	entry.ID = w.count
	w.entries = append(w.entries, entry)

	// Write to log file if open
	if w.logFile != nil {
		data, _ := json.MarshalIndent(entry, "", "  ")
		w.logFile.WriteString(fmt.Sprintf("--- Request #%d ---\n%s\n\n", entry.ID, string(data)))
		w.logFile.Sync()
	}
}

func (w *WebhookLogger) GetEntries() []LogEntry {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.entries
}

func (w *WebhookLogger) GetEntry(id int) (*LogEntry, bool) {
	w.mu.Lock()
	defer w.mu.Unlock()

	for i := range w.entries {
		if w.entries[i].ID == id {
			return &w.entries[i], true
		}
	}
	return nil, false
}

func (w *WebhookLogger) Clear() {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.entries = make([]LogEntry, 0)
	w.count = 0
}

func mainHandler(logger *WebhookLogger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Read body
		bodyBytes, _ := io.ReadAll(r.Body)
		body := string(bodyBytes)

		// Create log entry
		entry := LogEntry{
			Timestamp: time.Now().Format(time.RFC3339),
			Method:    r.Method,
			Path:      r.URL.Path,
			Headers:   make(Headers),
			Body:      body,
			Remote:    r.RemoteAddr,
		}

		// Copy headers
		for key, values := range r.Header {
			if len(values) > 0 {
				entry.Headers[key] = values[0]
			}
		}

		// Add to logger
		logger.AddEntry(entry)

		// Log to console
		fmt.Printf("[%d] %s %s from %s\n", entry.ID, entry.Method, entry.Path, entry.Remote)
		if entry.Body != "" {
			fmt.Printf("    Body: %s\n", truncate(entry.Body, 100))
		}

		// Return success response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "ok",
			"id":     entry.ID,
		})
	}
}

func listHandler(logger *WebhookLogger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(logger.GetEntries())
	}
}

func entryHandler(logger *WebhookLogger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			http.Error(w, "missing id parameter", http.StatusBadRequest)
			return
		}

		var id int
		if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}

		entry, ok := logger.GetEntry(id)
		if !ok {
			http.Error(w, "entry not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(entry)
	}
}

func clearHandler(logger *WebhookLogger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Clear()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "cleared"})
	}
}

func indexHandler(logger *WebhookLogger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		html := "<!DOCTYPE html>\n" +
			"<html>\n<head>\n" +
			"<title>Webhook Logger</title>\n" +
			"<style>\n" +
			"body { font-family: monospace; margin: 20px; background: #1a1a1a; color: #eee; }\n" +
			"h1 { color: #4CAF50; }\n" +
			".request { background: #2a2a2a; margin: 10px 0; padding: 15px; border-radius: 5px; }\n" +
			".method { font-weight: bold; padding: 2px 6px; border-radius: 3px; }\n" +
			"GET { background: #2196F3; } POST { background: #4CAF50; }\n" +
			"PUT { background: #FF9800; } DELETE { background: #f44336; }\n" +
			".path { color: #4CAF50; }\n" +
			".body { background: #000; padding: 10px; margin-top: 10px; white-space: pre-wrap; }\n" +
			".clear { background: #f44336; color: white; border: none; padding: 10px 20px; cursor: pointer; border-radius: 5px; }\n" +
			".refresh { background: #2196F3; color: white; border: none; padding: 10px 20px; cursor: pointer; border-radius: 5px; margin-left: 10px; }\n" +
			"</style>\n</head>\n<body>\n" +
			"<h1>Webhook Logger</h1>\n" +
			"<p>Send requests to any path and they will be logged here.</p>\n" +
			"<button class=\"clear\" onclick=\"clearLogs()\">Clear Logs</button>\n" +
			"<button class=\"refresh\" onclick=\"refresh()\">Refresh</button>\n" +
			"<div id=\"logs\"></div>\n" +
			"<script>\n" +
			"async function refresh() {\n" +
			"  const res = await fetch('/api/logs');\n" +
			"  const logs = await res.json();\n" +
			"  const container = document.getElementById('logs');\n" +
			"  container.innerHTML = logs.map(log => {\n" +
			"    var bodyHtml = log.body ? '<div class=\"body\">' + escapeHtml(log.body) + '</div>' : '';\n" +
			"    return '<div class=\"request\">' +\n" +
			"      '<span class=\"method\">' + log.method + '</span> ' +\n" +
			"      '<span class=\"path\">' + log.path + '</span> ' +\n" +
			"      '<span>' + log.timestamp + '</span> ' +\n" +
			"      '<span>from ' + log.remote + '</span>' +\n" +
			"      bodyHtml +\n" +
			"      '</div>';\n" +
			"  }).join('');\n" +
			"}\n" +
			"function escapeHtml(text) {\n" +
			"  var div = document.createElement('div');\n" +
			"  div.textContent = text;\n" +
			"  return div.innerHTML;\n" +
			"}\n" +
			"async function clearLogs() {\n" +
			"  await fetch('/api/clear', {method: 'POST'});\n" +
			"  refresh();\n" +
			"}\n" +
			"refresh();\n" +
			"setInterval(refresh, 2000);\n" +
			"</script>\n</body>\n</html>"

		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(html))
	}
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}

func main() {
	port := ":8080"
	logFile := ""

	// Parse args
	for i := 1; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "-port":
			if i+1 < len(os.Args) {
				port = ":" + os.Args[i+1]
				i++
			}
		case "-log":
			if i+1 < len(os.Args) {
				logFile = os.Args[i+1]
				i++
			}
		case "-h", "--help":
			fmt.Println("webhook-logger - Log incoming HTTP requests")
			fmt.Println()
			fmt.Println("Usage:")
			fmt.Println("  webhook-logger [options]")
			fmt.Println()
			fmt.Println("Options:")
			fmt.Println("  -port PORT   Port to listen on (default: 8080)")
			fmt.Println("  -log FILE    Write logs to file")
			fmt.Println("  -h, --help   Show this help")
			return
		}
	}

	logger, err := NewWebhookLogger(logFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	defer logger.Close()

	// Setup routes
	http.HandleFunc("/", indexHandler(logger))
	http.HandleFunc("/api/logs", listHandler(logger))
	http.HandleFunc("/api/log", entryHandler(logger))
	http.HandleFunc("/api/clear", clearHandler(logger))

	// Catch-all for webhook endpoints
	http.HandleFunc("/webhook/", mainHandler(logger))
	http.HandleFunc("/hook/", mainHandler(logger))

	addr := strings.TrimPrefix(port, ":")
	if addr == "" {
		addr = "8080"
	}

	fmt.Printf("Webhook Logger started on http://localhost:%s\n", addr)
	fmt.Printf("Webhook endpoint: http://localhost:%s/webhook/<any-path>\n", addr)

	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
