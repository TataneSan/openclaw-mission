// markdown-mermaid-gen: generate Mermaid diagrams from structured data.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

type config struct {
	input   string
	format  string
	typ     string
	output  string
	verbose bool
}

func parseFlags() config {
	format := flag.String("f", "auto", "input format: auto, json, csv")
	typ := flag.String("t", "flowchart", "diagram type: flowchart, sequence, state, class")
	output := flag.String("o", "", "output file (default: stdout)")
	verbose := flag.Bool("v", false, "verbose output")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: markdown-mermaid-gen [flags] INPUT\n\n")
		fmt.Fprintf(os.Stderr, "Generate Mermaid diagrams from structured data.\n\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nInput formats:\n")
		fmt.Fprintf(os.Stderr, "  JSON: array of nodes [{id, label, children: [ids]}] or edges [{from, to, label}]\n")
		fmt.Fprintf(os.Stderr, "  CSV:  from,to,label columns\n\n")
		fmt.Fprintf(os.Stderr, "Examples:\n")
		fmt.Fprintf(os.Stderr, "  markdown-mermaid-gen nodes.json\n")
		fmt.Fprintf(os.Stderr, "  markdown-mermaid-gen -t sequence sequence.json\n")
		fmt.Fprintf(os.Stderr, "  markdown-mermaid-gen -f csv edges.csv\n")
	}

	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "error: exactly one input file required")
		flag.Usage()
		os.Exit(1)
	}

	return config{
		input:   flag.Arg(0),
		format:  *format,
		typ:     *typ,
		output:  *output,
		verbose: *verbose,
	}
}

type node struct {
	ID       string   `json:"id"`
	Label    string   `json:"label"`
	Children []string `json:"children"`
}

type edge struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Label string `json:"label"`
}

type seqMsg struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Msg   string `json:"msg"`
	Type  string `json:"type"`
}

type classDef struct {
	Name   string   `json:"name"`
	Fields []string `json:"fields"`
	Meths  []string `json:"methods"`
	Extends string  `json:"extends"`
}

func sanitizeID(id string) string {
	var b strings.Builder
	for i, r := range id {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9' && i > 0) || r == '_' {
			b.WriteRune(r)
		} else {
			b.WriteRune('_')
		}
	}
	s := b.String()
	if s == "" || (s[0] >= '0' && s[0] <= '9') {
		s = "n" + s
	}
	return s
}

func escapeLabel(s string) string {
	return strings.ReplaceAll(s, "\"", "\\\"")
}

func generateFlowchart(data []byte) (string, error) {
	var nodes []node
	if err := json.Unmarshal(data, &nodes); err != nil {
		return "", err
	}

	var sb strings.Builder
	sb.WriteString("```mermaid\nflowchart TD\n")

	seen := make(map[string]bool)
	for _, n := range nodes {
		if !seen[n.ID] {
			label := n.Label
			if label == "" {
				label = n.ID
			}
			sb.WriteString(fmt.Sprintf("    %s[\"%s\"]\n", sanitizeID(n.ID), escapeLabel(label)))
			seen[n.ID] = true
		}
		for _, child := range n.Children {
			sb.WriteString(fmt.Sprintf("    %s --> %s\n", sanitizeID(n.ID), sanitizeID(child)))
		}
	}

	sb.WriteString("```\n")
	return sb.String(), nil
}

func generateSequence(data []byte) (string, error) {
	var msgs []seqMsg
	if err := json.Unmarshal(data, &msgs); err != nil {
		return "", err
	}

	var sb strings.Builder
	sb.WriteString("```mermaid\nsequenceDiagram\n")

	actors := make(map[string]bool)
	for _, m := range msgs {
		actors[m.From] = true
		actors[m.To] = true
	}

	for actor := range actors {
		sb.WriteString(fmt.Sprintf("    participant %s\n", sanitizeID(actor)))
	}
	sb.WriteString("\n")

	for _, m := range msgs {
		arrow := "->>"
		if m.Type == "reply" {
			arrow = "-->>"
		} else if m.Type == "async" {
			arrow = "-->"
		}
		sb.WriteString(fmt.Sprintf("    %s %s %s: %s\n",
			sanitizeID(m.From), arrow, sanitizeID(m.To), escapeLabel(m.Msg)))
	}

	sb.WriteString("```\n")
	return sb.String(), nil
}

func generateState(data []byte) (string, error) {
	var nodes []node
	if err := json.Unmarshal(data, &nodes); err != nil {
		return "", err
	}

	var sb strings.Builder
	sb.WriteString("```mermaid\nstateDiagram-v2\n")

	for _, n := range nodes {
		label := n.Label
		if label == "" {
			label = n.ID
		}
		sb.WriteString(fmt.Sprintf("    [%s] : %s\n", sanitizeID(n.ID), escapeLabel(label)))
		for _, child := range n.Children {
			sb.WriteString(fmt.Sprintf("    [%s] --> [%s]\n", sanitizeID(n.ID), sanitizeID(child)))
		}
	}

	sb.WriteString("```\n")
	return sb.String(), nil
}

func generateClass(data []byte) (string, error) {
	var classes []classDef
	if err := json.Unmarshal(data, &classes); err != nil {
		return "", err
	}

	var sb strings.Builder
	sb.WriteString("```mermaid\nclassDiagram\n")

	for _, c := range classes {
		sb.WriteString(fmt.Sprintf("    class %s {\n", sanitizeID(c.Name)))
		for _, f := range c.Fields {
			sb.WriteString(fmt.Sprintf("        %s\n", escapeLabel(f)))
		}
		for _, m := range c.Meths {
			sb.WriteString(fmt.Sprintf("        %s()\n", escapeLabel(m)))
		}
		sb.WriteString("    }\n")
		if c.Extends != "" {
			sb.WriteString(fmt.Sprintf("    %s <|-- %s\n", sanitizeID(c.Extends), sanitizeID(c.Name)))
		}
	}

	sb.WriteString("```\n")
	return sb.String(), nil
}

func generateFromCSV(data []byte, typ string) (string, error) {
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	if len(lines) < 1 {
		return "", fmt.Errorf("empty CSV")
	}

	var sb strings.Builder
	switch typ {
	case "flowchart":
		sb.WriteString("```mermaid\nflowchart TD\n")
	case "sequence":
		sb.WriteString("```mermaid\nsequenceDiagram\n")
	case "state":
		sb.WriteString("```mermaid\nstateDiagram-v2\n")
	case "class":
		sb.WriteString("```mermaid\nclassDiagram\n")
	}

	skipHeader := false
	if len(lines) > 0 {
		first := strings.ToLower(strings.TrimSpace(lines[0]))
		if strings.Contains(first, "from") && strings.Contains(first, "to") {
			skipHeader = true
		}
	}

	for i := 0; i < len(lines); i++ {
		if i == 0 && skipHeader {
			continue
		}
		line := strings.TrimSpace(lines[i])
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, ",", 3)
		if len(parts) < 2 {
			continue
		}
		from := strings.TrimSpace(parts[0])
		to := strings.TrimSpace(parts[1])
		label := ""
		if len(parts) > 2 {
			label = strings.TrimSpace(parts[2])
		}

		switch typ {
		case "flowchart":
			if label != "" {
				sb.WriteString(fmt.Sprintf("    %s -->|%s| %s\n", sanitizeID(from), escapeLabel(label), sanitizeID(to)))
			} else {
				sb.WriteString(fmt.Sprintf("    %s --> %s\n", sanitizeID(from), sanitizeID(to)))
			}
		case "sequence":
			sb.WriteString(fmt.Sprintf("    %s ->> %s: %s\n", sanitizeID(from), sanitizeID(to), escapeLabel(label)))
		case "state":
			if label != "" {
				sb.WriteString(fmt.Sprintf("    [%s] --> [%s] : %s\n", sanitizeID(from), sanitizeID(to), escapeLabel(label)))
			} else {
				sb.WriteString(fmt.Sprintf("    [%s] --> [%s]\n", sanitizeID(from), sanitizeID(to)))
			}
		case "class":
			sb.WriteString(fmt.Sprintf("    %s --> %s\n", sanitizeID(from), sanitizeID(to)))
		}
	}

	sb.WriteString("```\n")
	return sb.String(), nil
}

func main() {
	cfg := parseFlags()

	data, err := os.ReadFile(cfg.input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: read input: %v\n", err)
		os.Exit(1)
	}

	format := cfg.format
	if format == "auto" {
		switch {
		case strings.HasSuffix(cfg.input, ".json"):
			format = "json"
		case strings.HasSuffix(cfg.input, ".csv"):
			format = "csv"
		default:
			format = "json"
		}
	}

	var result string
	switch format {
	case "csv":
		result, err = generateFromCSV(data, cfg.typ)
	case "json":
		switch cfg.typ {
		case "flowchart":
			result, err = generateFlowchart(data)
		case "sequence":
			result, err = generateSequence(data)
		case "state":
			result, err = generateState(data)
		case "class":
			result, err = generateClass(data)
		default:
			err = fmt.Errorf("unknown diagram type: %s", cfg.typ)
		}
	default:
		err = fmt.Errorf("unknown format: %s", format)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	if cfg.output != "" {
		if err := os.WriteFile(cfg.output, []byte(result), 0644); err != nil {
			fmt.Fprintf(os.Stderr, "error: write output: %v\n", err)
			os.Exit(1)
		}
		if cfg.verbose {
			fmt.Fprintf(os.Stderr, "wrote %s (%d bytes)\n", cfg.output, len(result))
		}
	} else {
		fmt.Print(result)
	}
}
