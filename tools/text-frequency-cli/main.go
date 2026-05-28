package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

type WordCount struct {
	Word  string
	Count int
}

func analyze(input string) []WordCount {
	counts := make(map[string]int)
	re := regexp.MustCompile(`[a-zA-Z]+`)

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllString(line, -1)
		for _, word := range matches {
			counts[strings.ToLower(word)]++
		}
	}

	var result []WordCount
	for word, count := range counts {
		result = append(result, WordCount{word, count})
	}
	sort.Slice(result, func(i, j int) bool {
		if result[i].Count != result[j].Count {
			return result[i].Count > result[j].Count
		}
		return result[i].Word < result[j].Word
	})
	return result
}

func printResults(words []WordCount, top int) {
	if top > len(words) {
		top = len(words)
	}

	if len(words) == 0 {
		fmt.Println("No words found.")
		return
	}

	maxCount := words[0].Count
	maxLen := 0
	for _, w := range words[:top] {
		if len(w.Word) > maxLen {
			maxLen = len(w.Word)
		}
	}

	totalWords := 0
	for _, w := range words {
		totalWords += w.Count
	}

	fmt.Printf("Word Frequency Analysis\n")
	fmt.Printf("Total words: %d | Unique words: %d\n", totalWords, len(words))
	fmt.Println(strings.Repeat("-", 50))

	for _, w := range words[:top] {
		bar := strings.Repeat("#", (w.Count*40)/maxCount)
		pct := float64(w.Count) / float64(totalWords) * 100
		fmt.Printf("  %*d  %-*s  %s (%.1f%%)\n",
			len(fmt.Sprintf("%d", totalWords)), w.Count, maxLen, w.Word, bar, pct)
	}

	if top < len(words) {
		fmt.Printf("\n  ... and %d more unique words\n", len(words)-top)
	}
}

func printHelp() {
	fmt.Println("text-frequency-cli - Word frequency analyzer")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  text-frequency-cli [options] [file]")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -n, --top N    Show top N words (default: 20)")
	fmt.Println("  -h, --help     Show this help message")
	fmt.Println()
	fmt.Println("If no file is provided, reads from stdin.")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  text-frequency-cli document.txt")
	fmt.Println("  text-frequency-cli -n 10 report.md")
	fmt.Println("  cat file.txt | text-frequency-cli")
	fmt.Println("  text-frequency-cli -n 5 < input.txt")
}

func main() {
	top := 20
	args := []string{}

	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]
		switch {
		case arg == "-h", arg == "--help":
			printHelp()
			return
		case arg == "-n", arg == "--top":
			if i+1 < len(os.Args) {
				fmt.Sscanf(os.Args[i+1], "%d", &top)
				i++
			}
		case strings.HasPrefix(arg, "-n") || strings.HasPrefix(arg, "--top"):
			fmt.Sscanf(arg[2:], "%d", &top)
		default:
			args = append(args, arg)
		}
	}

	var input string
	if len(args) > 0 {
		data, err := os.ReadFile(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
			os.Exit(1)
		}
		input = string(data)
	} else {
		buf := new(strings.Builder)
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			buf.WriteString(scanner.Text())
			buf.WriteString("\n")
		}
		input = buf.String()
	}

	words := analyze(input)
	printResults(words, top)
}
