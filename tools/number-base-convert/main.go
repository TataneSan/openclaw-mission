package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"
)

const charset62 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

var supportedBases = []int{2, 8, 10, 16, 36, 62}

var baseNames = map[int]string{
	2:  "binary",
	8:  "octal",
	10: "decimal",
	16: "hex",
	36: "base36",
	62: "base62",
}

var nameToBase = map[string]int{
	"2":   2, "binary": 2, "bin": 2,
	"8":   8, "octal": 8, "oct": 8,
	"10":  10, "decimal": 10, "dec": 10,
	"16":  16, "hex": 16, "hexadecimal": 16,
	"36":  36, "base36": 36,
	"62":  62, "base62": 62,
}

func parseBase(name string) (int, error) {
	nl := strings.ToLower(name)
	if b, ok := nameToBase[nl]; ok {
		return b, nil
	}
	if b, err := strconv.Atoi(name); err == nil {
		for _, sb := range supportedBases {
			if b == sb {
				return b, nil
			}
		}
	}
	return 0, fmt.Errorf("unsupported base: %s (supported: 2, 8, 10, 16, 36, 62)", name)
}

func autoDetectBase(s string) (int, string, error) {
	neg := false
	if len(s) > 0 && s[0] == '-' {
		neg = true
		s = s[1:]
	}
	lower := strings.ToLower(s)
	var base int
	switch {
	case strings.HasPrefix(lower, "0b"):
		base = 2
		s = s[2:]
	case strings.HasPrefix(lower, "0o"):
		base = 8
		s = s[2:]
	case strings.HasPrefix(lower, "0x"):
		base = 16
		s = s[2:]
	default:
		return 0, "", fmt.Errorf("could not auto-detect base for %q: use --from flag or a prefix (0b, 0o, 0x)", s)
	}
	if neg {
		s = "-" + s
	}
	return base, s, nil
}

func toBigInt(s string, base int) (*big.Int, error) {
	neg := false
	if len(s) > 0 && s[0] == '-' {
		neg = true
		s = s[1:]
	}
	s = strings.TrimSpace(s)
	if s == "" {
		return nil, fmt.Errorf("empty number")
	}

	var val big.Int
	if base == 62 {
		// base62: use our custom charset
		val.SetString("0", 10)
		for _, c := range s {
			val.Mul(&val, big.NewInt(62))
			idx := indexOf(charset62, string(c))
			if idx == -1 {
				return nil, fmt.Errorf("invalid base62 character: %c", c)
			}
			val.Add(&val, big.NewInt(int64(idx)))
		}
	} else {
		if _, ok := val.SetString(s, base); !ok {
			return nil, fmt.Errorf("invalid number in base %d: %s", base, s)
		}
	}

	if neg {
		val.Neg(&val)
	}
	return &val, nil
}

func indexOf(cs, sub string) int {
	for i, c := range cs {
		if string(c) == sub {
			return i
		}
	}
	return -1
}

func fromBigInt(val *big.Int, base int) string {
	if base == 62 {
		return bigToInt62(val)
	}
	return val.Text(base)
}

func bigToInt62(val *big.Int) string {
	neg := false
	if val.Sign() < 0 {
		neg = true
		val = new(big.Int).Abs(val)
	}

	if val.Sign() == 0 {
		return "0"
	}

	base := big.NewInt(62)
	rem := new(big.Int)
	var chars []byte
	for val.Sign() > 0 {
		val.DivMod(val, base, rem)
		chars = append([]byte{charset62[rem.Int64()]}, chars...)
	}

	if neg {
		chars = append([]byte{'-'}, chars...)
	}
	return string(chars)
}

type ConversionResult struct {
	InputNumber string `json:"input_number"`
	InputBase   int    `json:"input_base"`
	OutputNumber string `json:"output_number"`
	OutputBase  int    `json:"output_base"`
}

type BatchResult []ConversionResult

type AllBasesResult struct {
	InputNumber string            `json:"input_number"`
	InputBase   int               `json:"input_base"`
	Conversions map[string]string `json:"conversions"`
}

func convert(s string, fromBase, toBase int) (string, error) {
	val, err := toBigInt(s, fromBase)
	if err != nil {
		return "", err
	}
	return fromBigInt(val, toBase), nil
}

func formatText(results []ConversionResult) {
	for _, r := range results {
		fmt.Printf("%s (base %d) = %s (base %d)\n",
			r.InputNumber, r.InputBase, r.OutputNumber, r.OutputBase)
	}
}

func formatJSON(results []ConversionResult) {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(results)
}

func formatTable(results []ConversionResult) {
	if len(results) == 0 {
		return
	}

	headers := []string{"Input", "From Base", "Output", "To Base"}
	rows := make([][]string, 0, len(results))
	for _, r := range results {
		rows = append(rows, []string{
			r.InputNumber,
			strconv.Itoa(r.InputBase),
			r.OutputNumber,
			strconv.Itoa(r.OutputBase),
		})
	}
	printTable(headers, rows)
}

func printTable(headers []string, rows [][]string) {
	widths := make([]int, len(headers))
	for i, h := range headers {
		widths[i] = len(h)
	}
	for _, row := range rows {
		for i, cell := range row {
			if len(cell) > widths[i] {
				widths[i] = len(cell)
			}
		}
	}

	fmt.Println(border(widths))
	fmt.Println(rowStr(headers, widths))
	fmt.Println(border(widths))
	for _, row := range rows {
		fmt.Println(rowStr(row, widths))
	}
	fmt.Println(border(widths))
}

func border(widths []int) string {
	parts := make([]string, len(widths))
	for i, w := range widths {
		parts[i] = strings.Repeat("-", w+2)
	}
	return "+" + strings.Join(parts, "+") + "+"
}

func rowStr(cells []string, widths []int) string {
	parts := make([]string, len(cells))
	for i, c := range cells {
		parts[i] = " " + c + strings.Repeat(" ", widths[i]-len(c)) + " "
	}
	return "|" + strings.Join(parts, "|") + "|"
}

func formatAllText(input string, fromBase int, val *big.Int) {
	fmt.Printf("Conversions of %s (base %d):\n", input, fromBase)
	for _, b := range supportedBases {
		fmt.Printf("  base %2d (%-7s): %s\n", b, baseNames[b], fromBigInt(val, b))
	}
}

func formatAllJSON(input string, fromBase int, val *big.Int) {
	conversions := make(map[string]string)
	for _, b := range supportedBases {
		conversions[baseNames[b]] = fromBigInt(val, b)
	}
	result := AllBasesResult{
		InputNumber: input,
		InputBase:   fromBase,
		Conversions: conversions,
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(result)
}

func formatAllTable(input string, fromBase int, val *big.Int) {
	headers := []string{"Base", "Name", "Value"}
	rows := make([][]string, 0, len(supportedBases))
	for _, b := range supportedBases {
		rows = append(rows, []string{
			strconv.Itoa(b),
			baseNames[b],
			fromBigInt(val, b),
		})
	}
	fmt.Printf("Conversions of %s (base %d):\n", input, fromBase)
	printTable(headers, rows)
}

func printUsage() {
	fmt.Fprintf(os.Stderr, `Usage: number-base-convert [options] [number]

Convert numbers between different bases.

Commands:
  convert <number> --from <base> --to <base>  Convert a number between bases
  convert <number> --from <base> --all         Show conversion to all supported bases

Options:
  -f, --from <base>   Source base (2, 8, 10, 16, 36, 62)
  -t, --to <base>     Target base (2, 8, 10, 16, 36, 62)
  -a, --all           Show conversion to all supported bases
  -F, --format <fmt>  Output format: text (default), json, table
  -b, --batch         Batch mode: read numbers from stdin, one per line
  -h, --help          Show this help message

Supported bases:
  2 (binary), 8 (octal), 10 (decimal), 16 (hex), 36, 62

Auto-detection:
  Input base can be auto-detected from prefixes:
    0b or 0B  - binary (base 2)
    0o or 0O  - octal (base 8)
    0x or 0X  - hexadecimal (base 16)

Examples:
  number-base-convert 255 --from 10 --to 2
  number-base-convert 0xFF --to 10
  number-base-convert 1010 --from 2 --all
  number-base-convert -42 --from 10 --to 16
  echo -e "255\n4096" | number-base-convert --batch --from 10 --to 2
  number-base-convert 100 --from 10 --to 16 --format json
`)
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		printUsage()
		os.Exit(1)
	}

	// Parse flags
	var number string
	var fromBaseStr, toBaseStr string
	var allBases bool
	var format string = "text"
	var batchMode bool

	i := 0
	for i < len(args) {
		switch args[i] {
		case "-h", "--help":
			printUsage()
			os.Exit(0)
		case "-f", "--from":
			i++
			if i >= len(args) {
				fmt.Fprintf(os.Stderr, "error: --from requires a value\n")
				os.Exit(1)
			}
			fromBaseStr = args[i]
		case "-t", "--to":
			i++
			if i >= len(args) {
				fmt.Fprintf(os.Stderr, "error: --to requires a value\n")
				os.Exit(1)
			}
			toBaseStr = args[i]
		case "-a", "--all":
			allBases = true
		case "-F", "--format":
			i++
			if i >= len(args) {
				fmt.Fprintf(os.Stderr, "error: --format requires a value\n")
				os.Exit(1)
			}
			format = strings.ToLower(args[i])
		case "-b", "--batch":
			batchMode = true
		case "-":
			// Could be a negative number
			if i+1 < len(args) {
				// Check if the next arg looks like it could be part of the number
				// or if this is a standalone negative number
				number = args[i]
			} else {
				number = args[i]
			}
		default:
			if number == "" {
				number = args[i]
			} else {
				// Could be a second number or an unknown flag
				fmt.Fprintf(os.Stderr, "error: unexpected argument: %s\n", args[i])
				os.Exit(1)
			}
		}
		i++
	}

	if format != "text" && format != "json" && format != "table" {
		fmt.Fprintf(os.Stderr, "error: unsupported format: %s (supported: text, json, table)\n", format)
		os.Exit(1)
	}

	if batchMode {
		runBatch(fromBaseStr, toBaseStr, allBases, format)
		return
	}

	if number == "" {
		fmt.Fprintf(os.Stderr, "error: no number provided\n")
		printUsage()
		os.Exit(1)
	}

	if allBases {
		runAll(number, fromBaseStr, format)
	} else {
		if toBaseStr == "" {
			fmt.Fprintf(os.Stderr, "error: --to is required (or use --all)\n")
			os.Exit(1)
		}
		runSingle(number, fromBaseStr, toBaseStr, format)
	}
}

func runSingle(number string, fromBaseStr, toBaseStr, format string) {
	toBase, err := parseBase(toBaseStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	fromBase := 0
	if fromBaseStr != "" {
		fromBase, err = parseBase(fromBaseStr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	} else {
		// Auto-detect
		fromBase, number, err = autoDetectBase(number)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	}

	result, err := convert(number, fromBase, toBase)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	results := []ConversionResult{
		{
			InputNumber:  number,
			InputBase:    fromBase,
			OutputNumber: result,
			OutputBase:   toBase,
		},
	}

	switch format {
	case "json":
		formatJSON(results)
	case "table":
		formatTable(results)
	default:
		formatText(results)
	}
}

func runAll(number string, fromBaseStr, format string) {
	fromBase := 0
	var err error
	if fromBaseStr != "" {
		fromBase, err = parseBase(fromBaseStr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	} else {
		fromBase, number, err = autoDetectBase(number)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	}

	val, err := toBigInt(number, fromBase)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	switch format {
	case "json":
		formatAllJSON(number, fromBase, val)
	case "table":
		formatAllTable(number, fromBase, val)
	default:
		formatAllText(number, fromBase, val)
	}
}

func runBatch(fromBaseStr, toBaseStr string, allBases bool, format string) {
	if !allBases && toBaseStr == "" {
		fmt.Fprintf(os.Stderr, "error: --to is required in batch mode (or use --all)\n")
		os.Exit(1)
	}

	fromBase := 0
	var err error
	if fromBaseStr != "" {
		fromBase, err = parseBase(fromBaseStr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	}

	toBase := 0
	if !allBases && toBaseStr != "" {
		toBase, err = parseBase(toBaseStr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	}

	scanner := bufio.NewScanner(os.Stdin)
	var allResults []ConversionResult
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		num := line
	 fb := fromBase
		if fb == 0 {
			fb, num, err = autoDetectBase(line)
			if err != nil {
				fmt.Fprintf(os.Stderr, "line %d: %v\n", lineNum, err)
				continue
			}
		}

		if allBases {
			val, err := toBigInt(num, fb)
			if err != nil {
				fmt.Fprintf(os.Stderr, "line %d: %v\n", lineNum, err)
				continue
			}
			switch format {
			case "json":
				formatAllJSON(num, fb, val)
			case "table":
				formatAllTable(num, fb, val)
			default:
				formatAllText(num, fb, val)
			}
		} else {
			result, err := convert(num, fb, toBase)
			if err != nil {
				fmt.Fprintf(os.Stderr, "line %d: %v\n", lineNum, err)
				continue
			}
			allResults = append(allResults, ConversionResult{
				InputNumber:  num,
				InputBase:    fb,
				OutputNumber: result,
				OutputBase:   toBase,
			})
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading stdin: %v\n", err)
		os.Exit(1)
	}

	if !allBases && len(allResults) > 0 {
		switch format {
		case "json":
			formatJSON(allResults)
		case "table":
			formatTable(allResults)
		default:
			formatText(allResults)
		}
	}
}

func init() {
	// Ensure supportedBases is sorted
	sort.Ints(supportedBases)
}
