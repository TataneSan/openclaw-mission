package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type ColStats struct {
	Name   string  `json:"column"`
	Count  int     `json:"count"`
	Nulls  int     `json:"nulls"`
	Min    float64 `json:"min,omitempty"`
	Max    float64 `json:"max,omitempty"`
	Mean   float64 `json:"mean,omitempty"`
	Median float64 `json:"median,omitempty"`
	StdDev float64 `json:"stddev,omitempty"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: csv-column-stats <file.csv>")
		os.Exit(1)
	}

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.LazyQuotes = true
	reader.TrimLeadingSpace = true

	headers, err := reader.Read()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading headers: %v\n", err)
		os.Exit(1)
	}

	numCols := len(headers)
	var allValues [][]string
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
		allValues = append(allValues, record)
	}

	stats := make([]ColStats, numCols)
	for i, h := range headers {
		stats[i].Name = strings.TrimSpace(h)
		stats[i].Count = len(allValues)
		var nums []float64
		for _, row := range allValues {
			val := ""
			if i < len(row) {
				val = strings.TrimSpace(row[i])
			}
			if val == "" {
				stats[i].Nulls++
				continue
			}
			if f, err := strconv.ParseFloat(val, 64); err == nil {
				nums = append(nums, f)
			}
		}
		if len(nums) > 0 {
			sort.Float64s(nums)
			stats[i].Min = nums[0]
			stats[i].Max = nums[len(nums)-1]
			sum := 0.0
			for _, n := range nums {
				sum += n
			}
			stats[i].Mean = sum / float64(len(nums))
			mid := len(nums) / 2
			if len(nums)%2 == 0 {
				stats[i].Median = (nums[mid-1] + nums[mid]) / 2
			} else {
				stats[i].Median = nums[mid]
			}
			if len(nums) > 1 {
				variance := 0.0
				for _, n := range nums {
					diff := n - stats[i].Mean
					variance += diff * diff
				}
				variance /= float64(len(nums) - 1)
				stats[i].StdDev = math.Sqrt(variance)
			}
		}
	}

	if len(os.Args) >= 3 && os.Args[2] == "--json" {
		enc, _ := json.MarshalIndent(stats, "", "  ")
		fmt.Println(string(enc))
	} else {
		fmt.Printf("%-20s %7s %5s %10s %10s %10s %10s %10s\n",
			"COLUMN", "COUNT", "NULLS", "MIN", "MAX", "MEAN", "MEDIAN", "STDDEV")
		fmt.Println(strings.Repeat("-", 92))
		for _, s := range stats {
			min, max, mean, median, stddev := "-", "-", "-", "-", "-"
			if s.Min != 0 || s.Max != 0 || s.Mean != 0 {
				min = fmt.Sprintf("%.2f", s.Min)
				max = fmt.Sprintf("%.2f", s.Max)
				mean = fmt.Sprintf("%.2f", s.Mean)
				median = fmt.Sprintf("%.2f", s.Median)
				if s.StdDev > 0 {
					stddev = fmt.Sprintf("%.2f", s.StdDev)
				}
			}
			fmt.Printf("%-20s %7d %5d %10s %10s %10s %10s %10s\n",
				s.Name, s.Count, s.Nulls, min, max, mean, median, stddev)
		}
	}
}
