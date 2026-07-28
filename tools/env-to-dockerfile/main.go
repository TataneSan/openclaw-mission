// env-to-dockerfile converts a .env file into Dockerfile ARG and ENV instructions.
//
// It reads key=value pairs from a .env file and outputs Dockerfile-compatible
// ARG declarations (for build-time) and ENV declarations (for runtime).
//
// Usage:
//
//	env-to-dockerfile [flags] [file]
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	filename := flagArg(0, ".env")
	buildOnly := flagBool("build", false)
	runOnly := flagBool("run", false)

	var f *os.File
	var err error
	if filename == "-" {
		f = os.Stdin
	} else {
		f, err = os.Open(filename)
		if err != nil {
			fatal("open: %v", err)
		}
		defer f.Close()
	}

	scanner := bufio.NewScanner(f)
	var pairs [][]string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		// Strip optional leading/trailing quotes from value.
		key, val, ok := parseEnvLine(line)
		if !ok {
			continue
		}
		pairs = append(pairs, []string{key, val})
	}
	if err := scanner.Err(); err != nil {
		fatal("read: %v", err)
	}

	if !*buildOnly && !*runOnly {
		// Default: print both ARG and ENV.
		fmt.Fprintln(os.Stderr, "# Build-time arguments")
		for _, p := range pairs {
			fmt.Printf("ARG %s\n", sanitizeKey(p[0]))
		}
		fmt.Println()
		fmt.Fprintln(os.Stderr, "# Runtime environment variables")
		for _, p := range pairs {
			fmt.Printf("ENV %s=${%s}\n", sanitizeKey(p[0]), sanitizeKey(p[0]))
		}
	} else if *buildOnly {
		for _, p := range pairs {
			fmt.Printf("ARG %s\n", sanitizeKey(p[0]))
		}
	} else {
		for _, p := range pairs {
			fmt.Printf("ENV %s=%s\n", sanitizeKey(p[0]), quoteVal(p[1]))
		}
	}
}

func parseEnvLine(line string) (string, string, bool) {
	line = strings.TrimSpace(line)
	if line == "" || strings.HasPrefix(line, "#") {
		return "", "", false
	}
	idx := strings.Index(line, "=")
	if idx <= 0 {
		return "", "", false
	}
	key := strings.TrimSpace(line[:idx])
	val := strings.TrimSpace(line[idx+1:])
	// Remove surrounding quotes.
	val = stripQuotes(val)
	if key == "" || !isValidKey(key) {
		return "", "", false
	}
	return key, val, true
}

func stripQuotes(s string) string {
	if len(s) >= 2 {
		if (s[0] == '"' && s[len(s)-1] == '"') || (s[0] == '\'' && s[len(s)-1] == '\'') {
			return s[1 : len(s)-1]
		}
	}
	return s
}

func isValidKey(key string) bool {
	if len(key) == 0 {
		return false
	}
	// Docker ENV/ARG keys: letters, digits, underscores. Must start with letter or underscore.
	for i, c := range key {
		if i == 0 {
			if !((c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || c == '_') {
				return false
			}
		} else {
			if !((c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || c == '_') {
				return false
			}
		}
	}
	return true
}

func sanitizeKey(key string) string {
	var b strings.Builder
	for i, c := range key {
		if i == 0 {
			if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
				b.WriteRune(c)
			} else {
				b.WriteRune('_')
			}
		} else {
			if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || c == '_' {
				b.WriteRune(c)
			} else {
				b.WriteRune('_')
			}
		}
	}
	return b.String()
}

func quoteVal(val string) string {
	if strings.ContainsAny(val, " \"'\n") {
		return "'" + strings.ReplaceAll(val, "'", "'\"'\"'") + "'"
	}
	return val
}

func fatal(msg string, args ...any) {
	fmt.Fprintf(os.Stderr, "env-to-dockerfile: "+msg+"\n", args...)
	os.Exit(1)
}

func flagArg(idx int, def string) string {
	pos := 0
	for i := 1; i < len(os.Args); i++ {
		if strings.HasPrefix(os.Args[i], "-") {
			if i+1 < len(os.Args) && !strings.HasPrefix(os.Args[i+1], "-") {
				i++
			}
			continue
		}
		if pos == idx {
			return os.Args[i]
		}
		pos++
	}
	return def
}

func flagBool(name string, def bool) *bool {
	v := new(bool)
	*v = def
	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == "-"+name || os.Args[i] == "--"+name {
			*v = true
		}
	}
	return v
}
