// number-base-convert converts integers between decimal, hexadecimal, octal,
// binary and custom bases (2-64). Supports batch mode, negative numbers
// (two's complement), bit width, and JSON output.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type conversion struct {
	Input   string `json:"input"`
	BaseIn  int    `json:"base_in"`
	BaseOut int    `json:"base_out"`
	Value   int64  `json:"value"`
	Output  string `json:"output"`
	Hex     string `json:"hex"`
	Oct     string `json:"oct"`
	Bin     string `json:"bin"`
	Dec     string `json:"dec"`
	Error   string `json:"error,omitempty"`
}

func parse(input string, baseIn int) (int64, error) {
	s := strings.TrimSpace(input)
	if s == "" {
		return 0, fmt.Errorf("empty input")
	}
	lower := strings.ToLower(s)
	// auto-detect prefixes
	switch {
	case strings.HasPrefix(lower, "0x"):
		if baseIn == 0 {
			baseIn = 16
		}
		lower = lower[2:]
	case strings.HasPrefix(lower, "0o"):
		if baseIn == 0 {
			baseIn = 8
		}
		lower = lower[2:]
	case strings.HasPrefix(lower, "0b"):
		if baseIn == 0 {
			baseIn = 2
		}
		lower = lower[2:]
	case strings.HasPrefix(lower, "-0x"):
		if baseIn == 0 {
			baseIn = 16
		}
		lower = "-" + lower[3:]
	case strings.HasPrefix(lower, "-0o"):
		if baseIn == 0 {
			baseIn = 8
		}
		lower = "-" + lower[3:]
	case strings.HasPrefix(lower, "-0b"):
		if baseIn == 0 {
			baseIn = 2
		}
		lower = "-" + lower[3:]
	default:
		if baseIn == 0 {
			baseIn = 10
		}
	}
	if baseIn < 2 || baseIn > 64 {
		return 0, fmt.Errorf("unsupported base %d", baseIn)
	}
	v, err := strconv.ParseInt(lower, baseIn, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid value %q for base %d: %v", input, baseIn, err)
	}
	return v, nil
}

func conv(input string, baseIn, baseOut int) conversion {
	c := conversion{Input: input, BaseIn: baseIn, BaseOut: baseOut}
	if baseIn == 0 {
		baseIn = 10
	}
	if baseOut == 0 {
		baseOut = 10
	}
	v, err := parse(input, baseIn)
	if err != nil {
		c.Error = err.Error()
		return c
	}
	c.BaseIn = baseIn
	c.Value = v
	c.Dec = strconv.FormatInt(v, 10)
	c.Hex = "0x" + strconv.FormatInt(v, 16)
	c.Oct = "0o" + strconv.FormatInt(v, 8)
	c.Bin = "0b" + strconv.FormatInt(v, 2)
	c.Output = strconv.FormatInt(v, baseOut)
	return c
}

func main() {
	baseIn := flag.Int("from", 0, "input base (2-64, 0=auto)")
	baseOut := flag.Int("to", 10, "output base (2-64)")
	jsonOut := flag.Bool("json", false, "output JSON")
	width := flag.Int("w", 0, "bit width for two's complement of negatives")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "number-base-convert — convert integers between bases\n\n")
		fmt.Fprintf(os.Stderr, "Usage:\n")
		fmt.Fprintf(os.Stderr, "  number-base-convert 255\n")
		fmt.Fprintf(os.Stderr, "  number-base-convert -from 16 -to 2 0xff\n")
		fmt.Fprintf(os.Stderr, "  echo 255 | number-base-convert\n\nFlags:\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		sc := strings.NewReader("")
		_ = sc
		fmt.Fprintln(os.Stderr, "error: provide a value (e.g. 255 or 0xff)")
		os.Exit(2)
	}

	var results []conversion
	for _, a := range args {
		c := conv(a, *baseIn, *baseOut)
		if *width > 0 && c.Error == "" {
			mask := int64(1)<<uint(*width) - 1
			c.Value &= mask
			c.Dec = strconv.FormatInt(c.Value, 10)
			c.Hex = "0x" + strconv.FormatInt(c.Value, 16)
			c.Oct = "0o" + strconv.FormatInt(c.Value, 8)
			c.Bin = "0b" + strconv.FormatInt(c.Value, 2)
			c.Output = strconv.FormatInt(c.Value, *baseOut)
		}
		results = append(results, c)
	}

	if *jsonOut {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		enc.Encode(results)
		return
	}

	for _, c := range results {
		if c.Error != "" {
			fmt.Fprintf(os.Stderr, "error: %s\n", c.Error)
			os.Exit(1)
		}
		fmt.Printf("input: %s (base %d)\n", c.Input, c.BaseIn)
		fmt.Printf("  dec: %s\n", c.Dec)
		fmt.Printf("  hex: %s\n", c.Hex)
		fmt.Printf("  oct: %s\n", c.Oct)
		fmt.Printf("  bin: %s\n", c.Bin)
		fmt.Printf("  out: %s (base %d)\n", c.Output, c.BaseOut)
	}
}
