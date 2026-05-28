package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type barData struct {
	label string
	value float64
}

func printUsage() {
	fmt.Fprintf(os.Stderr, `Usage: csv-to-graph [options] < input.csv
       csv-to-graph [options] -f file.csv

Generate terminal bar charts from CSV data.

Options:
  -f string    Path to CSV file (default: stdin)
  -col int     Column index for values (0-based, default: 1)
  -label int   Column index for labels (0-based, default: 0)
  -width int   Maximum bar width in characters (default: 50)
  -vertical    Render bars vertically instead of horizontally
  -sort string Sort order: none, asc, desc (default: none)
  -top int     Show only top N bars (0 = all, default: 0)
  -title string Chart title
  -char string Bar character (default: "█")
  -empty string Empty character (default: "░")
  -h, --help   Show this help message

Examples:
  csv-to-graph < data.csv
  csv-to-graph -f sales.csv -col 2 -width 60
  csv-to-graph -f data.csv -sort desc -top 10
  csv-to-graph -f data.csv -vertical -title "Monthly Revenue"
`)
}

func main() {
	var (
		inputPath = flag.String("f", "", "path to CSV file (default: stdin)")
		valCol    = flag.Int("col", 1, "column index for values (0-based)")
		labelCol  = flag.Int("label", 0, "column index for labels (0-based)")
		barWidth  = flag.Int("width", 50, "maximum bar width in characters")
		vertical  = flag.Bool("vertical", false, "render bars vertically")
		sortOrder = flag.String("sort", "none", "sort order: none, asc, desc")
		topN      = flag.Int("top", 0, "show only top N bars (0 = all)")
		title     = flag.String("title", "", "chart title")
		barChar   = flag.String("char", "█", "bar character")
		emptyChar = flag.String("empty", "░", "empty bar character")
	)
	flag.Parse()

	var reader *csv.Reader

	if *inputPath != "" {
		f, err := os.Open(*inputPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()
		reader = csv.NewReader(f)
	} else {
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) != 0 {
			fmt.Fprintf(os.Stderr, "No input: provide a file with -f or pipe CSV to stdin\n")
			os.Exit(1)
		}
		reader = csv.NewReader(os.Stdin)
	}

	reader.FieldsPerRecord = -1 // allow variable fields

	var records [][]string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading CSV: %v\n", err)
			os.Exit(1)
		}
		records = append(records, record)
	}

	if len(records) == 0 {
		fmt.Fprintf(os.Stderr, "Error: empty CSV\n")
		os.Exit(1)
	}

	// Skip header row if it looks like a header
	startIdx := 0
	if len(records) > 0 {
		_, err := strconv.ParseFloat(strings.TrimSpace(records[0][*valCol]), 64)
		if err != nil {
			startIdx = 1 // skip header
		}
	}

	var data []barData
	for i := startIdx; i < len(records); i++ {
		rec := records[i]
		if *valCol >= len(rec) || *labelCol >= len(rec) {
			continue
		}

		valStr := strings.TrimSpace(rec[*valCol])
		label := strings.TrimSpace(rec[*labelCol])

		// Strip common prefixes/suffixes
		valStr = strings.TrimPrefix(valStr, "$")
		valStr = strings.TrimSuffix(valStr, "%")
		valStr = strings.ReplaceAll(valStr, ",", "")

		val, err := strconv.ParseFloat(valStr, 64)
		if err != nil {
			continue // skip non-numeric rows
		}

		if label == "" {
			label = fmt.Sprintf("Row %d", i-startIdx+1)
		}

		data = append(data, barData{label: label, value: val})
	}

	if len(data) == 0 {
		fmt.Fprintf(os.Stderr, "Error: no valid numeric data found in column %d\n", *valCol)
		os.Exit(1)
	}

	// Sort
	switch *sortOrder {
	case "asc":
		sort.Slice(data, func(i, j int) bool { return data[i].value < data[j].value })
	case "desc":
		sort.Slice(data, func(i, j int) bool { return data[i].value > data[j].value })
	}

	// Top N
	if *topN > 0 && *topN < len(data) {
		if *sortOrder == "asc" {
			data = data[:*topN]
		} else {
			data = data[len(data)-*topN:]
			if *sortOrder == "desc" {
				// keep desc order
			} else {
				// reverse to show largest first
				for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
					data[i], data[j] = data[j], data[i]
				}
			}
		}
	}

	// Find max value and max label length
	var maxValue float64
	maxLabelLen := 0
	for _, d := range data {
		if d.value > maxValue {
			maxValue = d.value
		}
		if len(d.label) > maxLabelLen {
			maxLabelLen = len(d.label)
		}
	}

	if maxValue == 0 {
		maxValue = 1
	}

	// Title
	if *title != "" {
		fmt.Printf("\n  %s\n\n", *title)
	}

	if *vertical {
		renderVertical(data, maxValue, maxLabelLen, *barWidth, *barChar, *emptyChar)
	} else {
		renderHorizontal(data, maxValue, maxLabelLen, *barWidth, *barChar, *emptyChar)
	}
}

func renderHorizontal(data []barData, maxValue float64, maxLabelLen, barWidth int, barChar, emptyChar string) {
	for _, d := range data {
		ratio := d.value / maxValue
		filledLen := int(math.Round(ratio * float64(barWidth)))
		if d.value > 0 && filledLen == 0 {
			filledLen = 1
		}
		emptyLen := barWidth - filledLen

		label := d.label
		if len(label) > maxLabelLen {
			label = label[:maxLabelLen]
		} else {
			label = fmt.Sprintf("%-*s", maxLabelLen, label)
		}

		filled := strings.Repeat(barChar, filledLen)
		empty := strings.Repeat(emptyChar, emptyLen)
		valStr := formatValue(d.value)

		fmt.Printf("  %s │%s%s %s\n", label, filled, empty, valStr)
	}
}

func renderVertical(data []barData, maxValue float64, maxLabelLen, barWidth int, barChar, emptyChar string) {
	if len(data) == 0 {
		return
	}

	// Scale: use barWidth as height, but cap for readability
	height := barWidth
	if height > len(data)*2+5 {
		height = len(data)*2 + 5
	}

	// Print bars top to bottom
	for row := height; row >= 1; row-- {
		threshold := (float64(row) / float64(height)) * maxValue
		for _, d := range data {
			space := "  "
			if d.value >= threshold {
				fmt.Printf("%s%s", space, barChar)
			} else {
				fmt.Printf("%s%s", space, emptyChar)
			}
		}
		fmt.Println()
	}

	// Separator
	fmt.Printf("  %s\n", strings.Repeat("─", len(data)*2))

	// Labels
	for i, d := range data {
		label := d.label
		if len(label) > 8 {
			label = label[:8]
		}
		if i == 0 {
			fmt.Printf("  %-8s", label)
		} else {
			fmt.Printf("  %-8s", label)
		}
	}
	fmt.Println()

	// Values
	for _, d := range data {
		valStr := formatValue(d.value)
		fmt.Printf("  %-8s", valStr)
	}
	fmt.Println()
}

func formatValue(v float64) string {
	if v == float64(int64(v)) && abs(v) < 1e15 {
		return fmt.Sprintf("%d", int64(v))
	}
	if abs(v) < 0.01 {
		return fmt.Sprintf("%.4g", v)
	}
	if abs(v) < 1000 {
		return fmt.Sprintf("%.2f", v)
	}
	return fmt.Sprintf("%.2f", v)
}

func abs(v float64) float64 {
	if v < 0 {
		return -v
	}
	return v
}
