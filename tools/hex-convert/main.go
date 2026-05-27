// hex-convert converts numbers between hex, decimal, binary, and octal.
//
// Usage:
//
//	hex-convert <number> [--from fmt] [--to fmt]
//	hex-convert all <number> [--from fmt]
//
// Formats: hex, dec, bin, oct
//
// Examples:
//
//	hex-convert FF --from hex
//	hex-convert 255 --to hex
//	hex-convert all 42
//	hex-convert 101010 --from bin --to dec
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	fmtHex = "hex"
	fmtDec = "dec"
	fmtBin = "bin"
	fmtOct = "oct"
)

func detectFormat(input string) string {
	lower := strings.ToLower(strings.TrimSpace(input))

	// Try hex
	if strings.HasPrefix(lower, "0x") || strings.HasPrefix(lower, "0X") {
		return fmtHex
	}
	// Try bin
	if strings.HasPrefix(lower, "0b") || strings.HasPrefix(lower, "0B") {
		return fmtBin
	}
	// Try oct
	if strings.HasPrefix(lower, "0o") || strings.HasPrefix(lower, "0O") {
		return fmtOct
	}

	// Strip prefix if present
	stripped := strings.TrimPrefix(strings.TrimPrefix(lower, "0x"), "0X")
	if stripped != lower {
		return fmtHex
	}
	stripped = strings.TrimPrefix(strings.TrimPrefix(lower, "0b"), "0B")
	if stripped != lower {
		return fmtBin
	}
	stripped = strings.TrimPrefix(strings.TrimPrefix(lower, "0o"), "0O")
	if stripped != lower {
		return fmtOct
	}

	// Pure binary?
	if isPure(lower, "01") {
		return fmtBin
	}
	// Pure octal?
	if isPure(lower, "01234567") {
		// Could be octal or decimal — default to decimal
		return fmtDec
	}
	// Pure hex?
	if isPure(lower, "0123456789abcdefABCDEF") {
		// Could be hex or decimal — default to decimal
		return fmtDec
	}

	return ""
}

func isPure(s, charset string) bool {
	for _, c := range s {
		if !strings.ContainsRune(charset, c) {
			return false
		}
	}
	return len(s) > 0
}

func parseToDecimal(input, fromFmt string) (int64, error) {
	val := strings.TrimSpace(input)
	// Strip sign
	sign := int64(1)
	if strings.HasPrefix(val, "-") {
		sign = -1
		val = val[1:]
	}

	// Strip prefixes
	val = strings.TrimPrefix(strings.ToLower(val), "0x")
	val = strings.TrimPrefix(val, "0b")
	val = strings.TrimPrefix(val, "0o")

	base := 10
	switch fromFmt {
	case fmtHex:
		base = 16
	case fmtBin:
		base = 2
	case fmtOct:
		base = 8
	case fmtDec:
		base = 10
	}

	result, err := strconv.ParseInt(val, base, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid %s number: %s", fromFmt, input)
	}

	return sign * result, nil
}

func formatFromDecimal(dec int64, toFmt string) string {
	n := uint64(dec)
	if dec < 0 {
		n = uint64(-dec)
		prefix := "-"
		switch toFmt {
		case fmtHex:
			return prefix + formatHex(n)
		case fmtBin:
			return prefix + formatBin(n)
		case fmtOct:
			return prefix + formatOct(n)
		}
	}

	switch toFmt {
	case fmtHex:
		return formatHex(n)
	case fmtBin:
		return formatBin(n)
	case fmtOct:
		return formatOct(n)
	case fmtDec:
		return strconv.FormatInt(dec, 10)
	}
	return ""
}

func formatHex(n uint64) string {
	if n == 0 {
		return "0x0"
	}
	hexChars := "0123456789abcdef"
	result := ""
	for n > 0 {
		result = string(hexChars[n&0xf]) + result
		n >>= 4
	}
	return "0x" + result
}

func formatBin(n uint64) string {
	if n == 0 {
		return "0b0"
	}
	result := ""
	for n > 0 {
		result = strconv.Itoa(int(n & 1)) + result
		n >>= 1
	}
	return "0b" + result
}

func formatOct(n uint64) string {
	if n == 0 {
		return "0o0"
	}
	result := ""
	for n > 0 {
		result = strconv.Itoa(int(n & 7)) + result
		n >>= 3
	}
	return "0o" + result
}

func printAll(input string, fromFmt string) {
	dec, err := parseToDecimal(input, fromFmt)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println()
	fmt.Println("  ╔══════════════════════════════════════╗")
	fmt.Println("  ║          NUMBER CONVERSION           ║")
	fmt.Println("  ╠══════════════════════════════════════╣")
	fmt.Printf("  ║  Input: %-28s ║\n", fmt.Sprintf("%s (%s)", input, fromFmt))
	fmt.Println("  ╠══════════════════════════════════════╣")
	fmt.Printf("  ║  HEX:   %-28s ║\n", formatFromDecimal(dec, fmtHex))
	fmt.Printf("  ║  DEC:   %-28s ║\n", formatFromDecimal(dec, fmtDec))
	fmt.Printf("  ║  OCT:   %-28s ║\n", formatFromDecimal(dec, fmtOct))
	fmt.Printf("  ║  BIN:   %-28s ║\n", formatFromDecimal(dec, fmtBin))
	fmt.Println("  ╚══════════════════════════════════════╝")
	fmt.Println()
}

func printConvert(input string, fromFmt, toFmt string) {
	dec, err := parseToDecimal(input, fromFmt)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	result := formatFromDecimal(dec, toFmt)
	fmt.Println(result)
}

func printHelp() {
	fmt.Println("hex-convert — Convert numbers between hex, decimal, binary, and octal")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  hex-convert <number> [--from fmt] [--to fmt]")
	fmt.Println("  hex-convert all <number> [--from fmt]")
	fmt.Println()
	fmt.Println("Formats:")
	fmt.Println("  hex    Hexadecimal (0-9, a-f)")
	fmt.Println("  dec    Decimal (0-9)")
	fmt.Println("  bin    Binary (0-1)")
	fmt.Println("  oct    Octal (0-7)")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  hex-convert FF --from hex")
	fmt.Println("  hex-convert 255 --to hex")
	fmt.Println("  hex-convert all 42")
	fmt.Println("  hex-convert 101010 --from bin --to dec")
	fmt.Println()
	fmt.Println("Auto-detection:")
	fmt.Println("  hex-convert 0xFF       # detected as hex")
	fmt.Println("  hex-convert 0b1010     # detected as binary")
	fmt.Println("  hex-convert 0o77       # detected as octal")
}

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(1)
	}

	if os.Args[1] == "-h" || os.Args[1] == "--help" {
		printHelp()
		os.Exit(0)
	}

	fromFmt := ""
	toFmt := ""
	number := ""
	showAll := false

	i := 1
	for i < len(os.Args) {
		switch os.Args[i] {
		case "--from":
			if i+1 >= len(os.Args) {
				fmt.Fprintf(os.Stderr, "Error: --from requires a format\n")
				os.Exit(1)
			}
			fromFmt = os.Args[i+1]
			i += 2
		case "--to":
			if i+1 >= len(os.Args) {
				fmt.Fprintf(os.Stderr, "Error: --to requires a format\n")
				os.Exit(1)
			}
			toFmt = os.Args[i+1]
			i += 2
		case "all":
			showAll = true
			i++
		default:
			if number == "" {
				number = os.Args[i]
			}
			i++
		}
	}

	if number == "" {
		fmt.Fprintf(os.Stderr, "Error: no number provided\n")
		os.Exit(1)
	}

	// Validate formats
	validFormats := map[string]bool{fmtHex: true, fmtDec: true, fmtBin: true, fmtOct: true}

	if fromFmt == "" {
		fromFmt = detectFormat(number)
		if fromFmt == "" {
			fmt.Fprintf(os.Stderr, "Error: cannot detect format for '%s', use --from\n", number)
			os.Exit(1)
		}
	} else if !validFormats[fromFmt] {
		fmt.Fprintf(os.Stderr, "Error: invalid format '%s' (valid: hex, dec, bin, oct)\n", fromFmt)
		os.Exit(1)
	}

	if showAll {
		printAll(number, fromFmt)
	} else {
		if toFmt == "" {
			toFmt = fmtDec
		}
		if !validFormats[toFmt] {
			fmt.Fprintf(os.Stderr, "Error: invalid format '%s' (valid: hex, dec, bin, oct)\n", toFmt)
			os.Exit(1)
		}
		printConvert(number, fromFmt, toFmt)
	}
}
