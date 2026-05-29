// env-mask masks sensitive values in .env files.
//
// It detects sensitive keys by pattern matching (passwords, secrets, tokens, keys)
// and replaces their values with a masked placeholder.
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Default sensitive key patterns.
var defaultPatterns = []string{
	"password", "passwd", "pass",
	"secret", "token", "api_key", "apikey", "api-key",
	"private_key", "privatekey", "private-key",
	"access_key", "accesskey", "access-key",
	"secret_key", "secretkey", "secret-key",
	"auth", "credential", "credentials",
	"encryption_key", "encryptionkey", "encryption-key",
	"signing_key", "signingkey", "signing-key",
	"jwt_secret", "jwtsecret",
	"db_password", "database_password",
	"smtp_password", "email_password",
	"aws_secret", "gcp_secret", "azure_secret",
	"ssh_key", "ssh-key",
	"cookie_secret", "session_secret",
	"master_key", "masterkey",
	"client_secret", "clientsecret",
}

// buildPattern compiles sensitive key patterns into a single regex.
func buildPattern(patterns []string) (*regexp.Regexp, error) {
	var parts []string
	for _, p := range patterns {
		escaped := regexp.QuoteMeta(p)
		parts = append(parts, escaped)
	}
	re, err := regexp.Compile(`(?i)^(?:` + strings.Join(parts, "|") + `)`)
	return re, err
}

// MaskStyle defines how values are masked.
type MaskStyle int

const (
	MaskAsterisk MaskStyle = iota // ****
	MaskDots                      // ••••
	MaskPlaceholder               // <masked>
	MaskStar4                     // ****
)

func maskValue(style MaskStyle) string {
	switch style {
	case MaskAsterisk, MaskStar4:
		return "****"
	case MaskDots:
		return "\u2022\u2022\u2022\u2022"
	case MaskPlaceholder:
		return "<masked>"
	default:
		return "****"
	}
}

func processFile(path string, patterns []string, style MaskStyle, whitelist []string) error {
	re, err := buildPattern(patterns)
	if err != nil {
		return fmt.Errorf("invalid pattern: %w", err)
	}

	whitelistSet := make(map[string]bool)
	for _, w := range whitelist {
		whitelistSet[strings.ToUpper(strings.TrimSpace(w))] = true
	}

	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	mask := maskValue(style)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)

		// Skip empty lines and comments
		if trimmed == "" || strings.HasPrefix(trimmed, "#") {
			fmt.Println(line)
			continue
		}

		// Parse KEY=VALUE
		eqIdx := strings.Index(line, "=")
		if eqIdx < 0 {
			// Try export KEY=VALUE
			if strings.HasPrefix(strings.TrimSpace(line), "export ") {
				rest := strings.TrimSpace(line[len("export "):])
				eqIdx2 := strings.Index(rest, "=")
				if eqIdx2 < 0 {
					fmt.Println(line)
					continue
				}
				key := rest[:eqIdx2]
				upperKey := strings.ToUpper(key)
				if whitelistSet[upperKey] {
					fmt.Println(line)
					continue
				}
				if re.MatchString(key) {
					fmt.Printf("export %s=%s\n", key, mask)
				} else {
					fmt.Println(line)
				}
			} else {
				fmt.Println(line)
			}
			continue
		}

		key := line[:eqIdx]
		upperKey := strings.ToUpper(key)

		// Check whitelist
		if whitelistSet[upperKey] {
			fmt.Println(line)
			continue
		}

		// Check if sensitive
		if re.MatchString(key) {
			fmt.Printf("%s=%s\n", key, mask)
		} else {
			fmt.Println(line)
		}
	}

	return scanner.Err()
}

func processStdin(patterns []string, style MaskStyle, whitelist []string) error {
	re, err := buildPattern(patterns)
	if err != nil {
		return fmt.Errorf("invalid pattern: %w", err)
	}

	whitelistSet := make(map[string]bool)
	for _, w := range whitelist {
		whitelistSet[strings.ToUpper(strings.TrimSpace(w))] = true
	}

	mask := maskValue(style)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)

		if trimmed == "" || strings.HasPrefix(trimmed, "#") {
			fmt.Println(line)
			continue
		}

		eqIdx := strings.Index(line, "=")
		if eqIdx < 0 {
			if strings.HasPrefix(strings.TrimSpace(line), "export ") {
				rest := strings.TrimSpace(line[len("export "):])
				eqIdx2 := strings.Index(rest, "=")
				if eqIdx2 < 0 {
					fmt.Println(line)
					continue
				}
				key := rest[:eqIdx2]
				if whitelistSet[strings.ToUpper(key)] {
					fmt.Println(line)
					continue
				}
				if re.MatchString(key) {
					fmt.Printf("export %s=%s\n", key, mask)
				} else {
					fmt.Println(line)
				}
			} else {
				fmt.Println(line)
			}
			continue
		}

		key := line[:eqIdx]
		if whitelistSet[strings.ToUpper(key)] {
			fmt.Println(line)
			continue
		}
		if re.MatchString(key) {
			fmt.Printf("%s=%s\n", key, mask)
		} else {
			fmt.Println(line)
		}
	}

	return scanner.Err()
}

func printUsage() {
	fmt.Println(`env-mask — mask sensitive values in .env files

USAGE:
    env-mask [OPTIONS] [FILE]
    cat .env | env-mask [OPTIONS]

OPTIONS:
    -p, --pattern PATTERNS   Comma-separated sensitive key patterns
                              (default: password,secret,token,api_key,private_key,
                               access_key,secret_key,auth,credential,encryption_key,
                               signing_key,jwt_secret,db_password,smtp_password,
                               aws_secret,ssh_key,cookie_secret,master_key,
                               client_secret)
    -w, --whitelist KEYS     Comma-separated keys to skip masking
    -s, --style STYLE        Mask style: asterisk (****), dots (••••), placeholder (<masked>)
                              (default: asterisk)
    -h, --help               Show this help

EXAMPLES:
    # Mask sensitive values in .env file
    env-mask .env

    # Mask with custom style
    env-mask -s placeholder .env

    # Pipe from stdin
    cat .env | env-mask

    # Whitelist specific keys
    env-mask -w "DB_PASSWORD,API_KEY" .env

    # Custom patterns
    env-mask -p "secret,key,token" .env

NOTES:
    - Pattern matching is case-insensitive
    - Comments and empty lines are preserved
    - Lines without '=' are passed through unchanged
    - Supports both KEY=VALUE and export KEY=VALUE formats`)
}

func main() {
	var patternsStr, whitelistStr, styleStr string
	file := ".env"
	args := os.Args[1:]

	i := 0
	for i < len(args) {
		switch args[i] {
		case "-h", "--help":
			printUsage()
			os.Exit(0)
		case "-p", "--pattern":
			i++
			if i >= len(args) {
				fmt.Fprintln(os.Stderr, "error: --pattern requires a value")
				os.Exit(2)
			}
			patternsStr = args[i]
		case "-w", "--whitelist":
			i++
			if i >= len(args) {
				fmt.Fprintln(os.Stderr, "error: --whitelist requires a value")
				os.Exit(2)
			}
			whitelistStr = args[i]
		case "-s", "--style":
			i++
			if i >= len(args) {
				fmt.Fprintln(os.Stderr, "error: --style requires a value")
				os.Exit(2)
			}
			styleStr = args[i]
		default:
			if !strings.HasPrefix(args[i], "-") {
				file = args[i]
			} else {
				fmt.Fprintf(os.Stderr, "error: unknown flag %s\n", args[i])
				os.Exit(2)
			}
		}
		i++
	}

	// Patterns
	var patterns []string
	if patternsStr != "" {
		patterns = strings.Split(patternsStr, ",")
		for i, p := range patterns {
			patterns[i] = strings.TrimSpace(p)
		}
	} else {
		patterns = defaultPatterns
	}

	// Whitelist
	var whitelist []string
	if whitelistStr != "" {
		whitelist = strings.Split(whitelistStr, ",")
		for i, w := range whitelist {
			whitelist[i] = strings.TrimSpace(w)
		}
	}

	// Style
	style := MaskAsterisk
	switch strings.ToLower(styleStr) {
	case "dots":
		style = MaskDots
	case "placeholder":
		style = MaskPlaceholder
	case "asterisk", "":
		style = MaskAsterisk
	default:
		fmt.Fprintf(os.Stderr, "error: invalid style %q, use: asterisk, dots, placeholder\n", styleStr)
		os.Exit(2)
	}

	// Check if stdin is a pipe
	stat, err := os.Stdin.Stat()
	if err == nil && stat.Mode()&os.ModeNamedPipe != 0 {
		if err := processStdin(patterns, style, whitelist); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(2)
		}
		os.Exit(0)
	}

	// Process file
	if err := processFile(file, patterns, style, whitelist); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(2)
	}
}
