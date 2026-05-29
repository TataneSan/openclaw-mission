package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// htmlToText converts HTML to plain text
func htmlToText(r io.Reader, width int) string {
	scanner := bufio.NewScanner(r)
	var out strings.Builder
	inPre := false

	for scanner.Scan() {
		data := scanner.Text()
		data = processLine(data, width)
		if out.Len() > 0 && !inPre {
			out.WriteString("\n")
		}
		out.WriteString(data)
	}

	return strings.TrimRight(out.String(), "\n")
}

func processLine(input string, width int) string {
	var result strings.Builder
	inTag := false
	skipText := false
	skipLevel := 0
	currentText := ""

	i := 0
	for i < len(input) {
		ch := input[i]

		if inTag {
			if ch == '>' {
				inTag = false
				i++
				continue
			}
			i++
			continue
		}

		if ch == '<' {
			// Flush current text
			if currentText != "" {
				result.WriteString(currentText)
				currentText = ""
			}

			// Check what tag this is
			remaining := input[i:]
			if len(remaining) > 1 {
				tagLower := ""
				for j := 1; j < len(remaining) && remaining[j] != ' ' && remaining[j] != '>' && remaining[j] != '/'; j++ {
					c := remaining[j]
					if c >= 'A' && c <= 'Z' {
						c += 'a' - 'A'
					}
					tagLower += string(c)
				}

				switch tagLower {
				case "br":
					// handled by flush above
				case "p", "div", "h1", "h2", "h3", "h4", "h5", "h6", "li", "tr":
					if result.Len() > 0 {
						result.WriteString("\n")
					}
				case "pre":
					// handled by flush
				case "script", "style":
					skipText = true
				case "/script", "/style":
					skipText = false
				case "!--":
					skipLevel++
				}

				if strings.HasPrefix(remaining, "<!--") {
					skipLevel++
				}
			}

			inTag = true
			i++
			continue
		}

		// Handle comment end
		if !inTag && !skipText && ch == '-' && i+1 < len(input) && input[i+1] == '-' && i+2 < len(input) && input[i+2] == '>' {
			if skipLevel > 0 {
				skipLevel--
			}
			i += 3
			continue
		}

		if skipText || skipLevel > 0 {
			i++
			continue
		}

		// Handle entities
		if ch == '&' {
			entity := decodeEntity(input[i:])
			currentText += entity
			i += len(entity) + 1 // +1 for the semicolon or end
			// Find the semicolon
			for i < len(input) && input[i] != ';' {
				i++
			}
			if i < len(input) {
				i++ // skip semicolon
			}
			continue
		}

		currentText += string(ch)
		i++
	}

	// Flush remaining text
	if currentText != "" {
		result.WriteString(currentText)
	}

	return result.String()
}

func decodeEntity(s string) string {
	entity := ""
	i := 1 // skip &
	for i < len(s) && i < 10 {
		if s[i] == ';' {
			break
		}
		entity += string(s[i])
		i++
	}

	switch entity {
	case "amp":
		return "&"
	case "lt":
		return "<"
	case "gt":
		return ">"
	case "quot":
		return "\""
	case "apos":
		return "'"
	case "nbsp":
		return " "
	case "copy":
		return "(c)"
	case "reg":
		return "(r)"
	case "mdash":
		return "--"
	case "ndash":
		return "-"
	case "laquo":
		return "<<"
	case "raquo":
		return ">>"
	case "bull":
		return "*"
	default:
		return "&" + entity
	}
}

func printUsage() {
	fmt.Fprintf(os.Stderr, `html-to-text - Convert HTML to plain text

Usage:
  html-to-text [options] [file]

Options:
  -w, --width N   Output line width (default: 0, no wrapping)
  -h, --help      Show this help

Examples:
  html-to-text page.html
  html-to-text -w 80 page.html
  cat page.html | html-to-text
`)
}

func main() {
	width := 0
	file := ""

	args := os.Args[1:]
	i := 0
	for i < len(args) {
		switch args[i] {
		case "-w", "--width":
			i++
			if i >= len(args) {
				fmt.Fprintln(os.Stderr, "error: -w requires a number")
				os.Exit(1)
			}
			fmt.Sscanf(args[i], "%d", &width)
		case "-h", "--help":
			printUsage()
			return
		default:
			if strings.HasPrefix(args[i], "-") {
				fmt.Fprintf(os.Stderr, "error: unknown flag %s\n", args[i])
				os.Exit(1)
			}
			file = args[i]
		}
		i++
	}

	var reader io.Reader
	if file == "" {
		stat, err := os.Stdin.Stat()
		if err != nil || (stat.Mode() & os.ModeCharDevice) != 0 {
			printUsage()
			os.Exit(1)
		}
		reader = os.Stdin
	} else {
		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()
		reader = f
	}

	text := htmlToText(reader, width)
	fmt.Println(text)
}
