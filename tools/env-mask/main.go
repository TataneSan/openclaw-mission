// env-mask masks sensitive values in .env files.
//
// It reads .env files and replaces sensitive values with masked
// placeholders like [MASKED] while preserving the key names and
// file structure (comments, blank lines, etc.).
//
// Usage:
//
//	env-mask .env
//	env-mask --keys PASSWORD,SECRET .env
//	env-mask --pattern '.*(PASS|SECRET|KEY|TOKEN).*' .env
//	env-mask --mask '***' .env
//	env-mask --format json .env
//	env-mask --strip .env
package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Default sensitive key patterns (case-insensitive).
var defaultPatterns = []string{
	"password",
	"passwd",
	"secret",
	"token",
	"api_key",
	"apikey",
	"api-key",
	"private_key",
	"private-key",
	"access_key",
	"access-key",
	"auth",
	"credential",
	"credentials",
	"secret_key",
	"secret-key",
	"db_pass",
	"db_password",
	"database_password",
	"mysql_password",
	"postgres_password",
	"redis_password",
	"smtp_password",
	"encryption_key",
	"signing_key",
	"jwt_secret",
	"session_secret",
	"cookie_secret",
	"aws_secret",
	"gcp_key",
	"azure_key",
}

type maskedLine struct {
	Raw      string `json:"-"`
	Key      string `json:"key,omitempty"`
	Value    string `json:"value,omitempty"`
	Masked   bool   `json:"masked"`
	IsComment bool  `json:"is_comment,omitempty"`
	IsBlank  bool   `json:"is_blank,omitempty"`
}

func main() {
	fs := flag.NewFlagSet("env-mask", flag.ExitOnError)
	keysFlag := fs.String("keys", "", "comma-separated list of exact key names to mask")
	patternFlag := fs.String("pattern", "", "regex pattern to match sensitive key names (overrides defaults)")
	useDefaultsFlag := fs.Bool("defaults", true, "use built-in sensitive key patterns")
	maskFlag := fs.String("mask", "[MASKED]", "replacement string for masked values")
	formatFlag := fs.String("format", "env", "output format: env, json")
	stripFlag := fs.Bool("strip", false, "strip values entirely (set to empty string)")
	showFlag := fs.Bool("show", false, "show which keys were masked (summary mode)")
	versionFlag := fs.Bool("version", false, "print version")

	fs.Parse(os.Args[1:])

	if *versionFlag {
		fmt.Println("env-mask v1.0.0")
		os.Exit(0)
	}

	args := fs.Args()
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "usage: env-mask [flags] <file.env>")
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "Masks sensitive values in .env files.")
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "Examples:")
		fmt.Fprintln(os.Stderr, "  env-mask .env                        # mask defaults")
		fmt.Fprintln(os.Stderr, "  env-mask --keys PASSWORD,SECRET .env # mask specific keys")
		fmt.Fprintln(os.Stderr, "  env-mask --pattern '.*TOKEN.*' .env  # mask by regex")
		fmt.Fprintln(os.Stderr, "  env-mask --mask '***' .env           # custom mask string")
		fmt.Fprintln(os.Stderr, "  env-mask --format json .env          # JSON output")
		fmt.Fprintln(os.Stderr, "  env-mask --strip .env                # set values to empty")
		fmt.Fprintln(os.Stderr, "  env-mask --show .env                 # show summary only")
		fs.PrintDefaults()
		os.Exit(1)
	}

	filename := args[0]

	// Build patterns to match.
	var patterns []*regexp.Regexp
	if *patternFlag != "" {
		re, err := regexp.Compile(`(?i)` + *patternFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: invalid pattern: %v\n", err)
			os.Exit(1)
		}
		patterns = append(patterns, re)
	} else if *useDefaultsFlag {
		for _, p := range defaultPatterns {
			re := regexp.MustCompile(`(?i)^` + regexp.QuoteMeta(p) + `$`)
			patterns = append(patterns, re)
		}
	}

	// Additional exact keys.
	var exactKeys map[string]bool
	if *keysFlag != "" {
		exactKeys = make(map[string]bool)
		for _, k := range strings.Split(*keysFlag, ",") {
			exactKeys[strings.TrimSpace(k)] = true
		}
	}

	// Read file.
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	lines, keysMasked := parseAndMask(file, patterns, exactKeys, *maskFlag, *stripFlag)

	if *showFlag {
		printSummary(keysMasked, filename)
		return
	}

	if *formatFlag == "json" {
		printJSON(lines)
	} else {
		for _, line := range lines {
			fmt.Println(line.Raw)
		}
	}
}

func isSensitive(key string, patterns []*regexp.Regexp, exactKeys map[string]bool) bool {
	// Check exact keys first.
	if exactKeys != nil {
		for ek := range exactKeys {
			if strings.EqualFold(key, ek) {
				return true
			}
		}
	}

	// Check patterns.
	for _, re := range patterns {
		if re.MatchString(key) {
			return true
		}
	}

	return false
}

func parseAndMask(file *os.File, patterns []*regexp.Regexp, exactKeys map[string]bool, mask string, strip bool) ([]maskedLine, []string) {
	var lines []maskedLine
	var masked []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		raw := scanner.Text()
		trimmed := strings.TrimSpace(raw)

		// Blank line.
		if trimmed == "" {
			lines = append(lines, maskedLine{Raw: raw, IsBlank: true})
			continue
		}

		// Comment line.
		if strings.HasPrefix(trimmed, "#") || strings.HasPrefix(trimmed, ";") {
			lines = append(lines, maskedLine{Raw: raw, IsComment: true})
			continue
		}

		// Key=Value or KEY="value" or KEY='value'.
		if eqIdx := strings.Index(raw, "="); eqIdx > 0 {
			key := strings.TrimSpace(raw[:eqIdx])
			rest := raw[eqIdx+1:]
			value := unquote(strings.TrimSpace(rest))

			if isSensitive(key, patterns, exactKeys) {
				masked = append(masked, key)
				replacement := mask
				if strip {
					replacement = ""
				}
				lines = append(lines, maskedLine{
					Raw:   key + "=" + replacement,
					Key:   key,
					Value: replacement,
					Masked: true,
				})
			} else {
				lines = append(lines, maskedLine{
					Raw:   raw,
					Key:   key,
					Value: value,
					Masked: false,
				})
			}
		} else {
			// Not a key=value line, keep as-is.
			lines = append(lines, maskedLine{Raw: raw})
		}
	}

	return lines, masked
}

func unquote(s string) string {
	if len(s) >= 2 {
		if (s[0] == '"' && s[len(s)-1] == '"') || (s[0] == '\'' && s[len(s)-1] == '\'') {
			return s[1 : len(s)-1]
		}
	}
	return s
}

func printSummary(masked []string, filename string) {
	dir, base := filepath.Split(filename)
	if dir == "" {
		base = filepath.Base(filename)
	}

	fmt.Printf("File: %s\n", base)
	fmt.Printf("Sensitive keys found: %d\n", len(masked))
	fmt.Println()
	if len(masked) > 0 {
		fmt.Println("Masked keys:")
		for _, k := range masked {
			fmt.Printf("  - %s\n", k)
		}
	}
}

func printJSON(lines []maskedLine) {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	encoder.Encode(lines)
}
