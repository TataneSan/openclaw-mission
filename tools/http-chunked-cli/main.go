package main

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	var (
		method  = flag.String("method", "GET", "HTTP method")
		url     = flag.String("url", "", "Target URL (required)")
		timeout = flag.Duration("timeout", 60*time.Second, "Request timeout")
		chunked = flag.Bool("chunked", false, "Show each chunk as it arrives")
		skipVerify = flag.Bool("insecure", false, "Skip TLS verification")
		output  = flag.String("output", "full", "Output: full, headers, body, status, json, chunks")
		userAgent = flag.String("user-agent", "http-chunked-cli/1.0.0", "User-Agent header")
		data    = flag.String("data", "", "Request body")
	)
	flag.Parse()

	if *url == "" {
		fmt.Fprintln(os.Stderr, "error: -url is required")
		printUsage()
		os.Exit(1)
	}

	// Build request
	var body io.Reader
	if *data != "" {
		body = strings.NewReader(*data)
	}

	req, err := http.NewRequest(*method, *url, body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	req.Header.Set("User-Agent", *userAgent)
	req.Header.Set("Accept-Encoding", "identity")

	// Parse -H flags
	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]
		if arg == "-H" && i+1 < len(os.Args) {
			setHeader(req, os.Args[i+1])
			i++
		} else if strings.HasPrefix(arg, "-H=") {
			setHeader(req, strings.TrimPrefix(arg, "-H="))
		}
	}

	// Build client
	client := &http.Client{
		Timeout: *timeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: *skipVerify},
		},
	}

	start := time.Now()
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	firstByte := time.Since(start)

	// Handle chunked output mode
	if *chunked || *output == "chunks" {
		handleChunked(resp, start)
		return
	}

	// Standard output modes
	bodyBytes, err := io.ReadAll(resp.Body)
	duration := time.Since(start)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading body: %v\n", err)
		os.Exit(1)
	}

	switch *output {
	case "status":
		fmt.Println(resp.Status)
	case "headers":
		for k, vv := range resp.Header {
			for _, v := range vv {
				fmt.Printf("%s: %s\n", k, v)
			}
		}
	case "body":
		os.Stdout.Write(bodyBytes)
	case "json":
		printJSON(resp, bodyBytes, duration, firstByte)
	case "full":
		printFull(resp, bodyBytes, duration, firstByte)
	}
}

func handleChunked(resp *http.Response, startTime time.Time) {
	fmt.Printf("HTTP/%s %s\n", resp.Proto, resp.Status)
	for k, vv := range resp.Header {
		for _, v := range vv {
			fmt.Printf("%s: %s\n", k, v)
		}
	}
	fmt.Println()

	isChunked := resp.Header.Get("Transfer-Encoding") == "chunked"
	if isChunked {
		fmt.Fprintln(os.Stderr, "[chunked transfer encoding detected]")
	}

	reader := resp.Body
	chunkNum := 0
	var totalBytes int64
	var mu sync.Mutex

	// Use a TeeReader approach to track bytes while streaming
	buf := make([]byte, 4096)
	for {
		n, err := reader.Read(buf)
		if n > 0 {
			chunkNum++
			mu.Lock()
			totalBytes += int64(n)
			mu.Unlock()
			elapsed := time.Since(startTime)
			fmt.Fprintf(os.Stderr, "[chunk %d] %d bytes (total: %d, elapsed: %v)\n",
				chunkNum, n, totalBytes, elapsed)
			os.Stdout.Write(buf[:n])
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	}

	duration := time.Since(startTime)
	fmt.Fprintf(os.Stderr, "\n[done] %d chunks, %d bytes total, %v\n", chunkNum, totalBytes, duration)
}

func setHeader(req *http.Request, val string) {
	idx := strings.Index(val, ":")
	if idx == -1 {
		return
	}
	k := strings.TrimSpace(val[:idx])
	v := strings.TrimSpace(val[idx+1:])
	if k != "" {
		req.Header.Set(k, v)
	}
}

func printJSON(resp *http.Response, body []byte, duration, firstByte time.Duration) {
	result := map[string]interface{}{
		"status":       resp.Status,
		"statusCode":   resp.StatusCode,
		"headers":      make(map[string][]string),
		"bodySize":     len(body),
		"duration":     duration.String(),
		"firstByte":    firstByte.String(),
		"transferEnc":  resp.Header.Get("Transfer-Encoding"),
		"contentEnc":   resp.Header.Get("Content-Encoding"),
	}
	for k, vv := range resp.Header {
		result["headers"].(map[string][]string)[k] = vv
	}

	var jsonBody interface{}
	if len(body) > 0 {
		if err := json.Unmarshal(body, &jsonBody); err == nil {
			result["body"] = jsonBody
		} else {
			result["body"] = string(body)
		}
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(result)
}

func printFull(resp *http.Response, body []byte, duration, firstByte time.Duration) {
	fmt.Printf("HTTP/%s %s\n", resp.Proto, resp.Status)
	for k, vv := range resp.Header {
		for _, v := range vv {
			fmt.Printf("%s: %s\n", k, v)
		}
	}
	fmt.Println()
	os.Stdout.Write(body)
	fmt.Fprintf(os.Stderr, "\n[timing] total: %v, first byte: %v, body: %d bytes\n",
		duration, firstByte, len(body))
}

func printUsage() {
	fmt.Fprintln(os.Stderr, `
http-chunked-cli - HTTP client with chunked transfer encoding support

Usage:
  http-chunked-cli -url URL [options]

Options:
  -url string      Target URL (required)
  -method string   HTTP method (default "GET")
  -data string     Request body
  -H string        Header: Key: Value (repeatable)
  -timeout         Request timeout (default 60s)
  -chunked         Show each chunk as it arrives
  -output string   Output mode: full, headers, body, status, json, chunks (default "full")
  -insecure        Skip TLS verification
  -user-agent      User-Agent header

Examples:
  http-chunked-cli -url https://api.example.com/stream -chunked
  http-chunked-cli -url https://api.example.com/data -output json
  http-chunked-cli -url http://localhost:8080/sse -method GET -chunked
  http-chunked-cli -url http://localhost:8080/api -method POST -data '{"key":"value"}' -H "Content-Type: application/json"
`)
}

// Ensure bufio.Writer is used for flushing
var _ = bufio.Writer{}
