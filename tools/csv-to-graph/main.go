package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

func sortedKeys(m map[string]any) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func toYAMLValue(v any, indent string) string {
	next := indent + "  "
	switch val := v.(type) {
	case nil:
		return "null"
	case bool:
		return fmt.Sprintf("%v", val)
	case float64:
		if val == float64(int64(val)) {
			return fmt.Sprintf("%d", int64(val))
		}
		return fmt.Sprintf("%g", val)
	case string:
		if strings.ContainsAny(val, ":#{}[]|>&*!%@`\"'") || strings.HasPrefix(val, " ") || strings.HasSuffix(val, " ") {
			escaped := strings.ReplaceAll(val, `"`, `\"`)
			return fmt.Sprintf(`"%s"`, escaped)
		}
		return val
	case map[string]any:
		var lines []string
		for _, k := range sortedKeys(val) {
			lines = append(lines, fmt.Sprintf("%s%s: %s", indent, k, toYAMLValue(val[k], next)))
		}
		return strings.Join(lines, "\n")
	case []any:
		var lines []string
		for _, item := range val {
			switch iv := item.(type) {
			case map[string]any:
				lines = append(lines, fmt.Sprintf("%s- %s", indent, toYAMLValue(iv, next+"  ")))
			default:
				lines = append(lines, fmt.Sprintf("%s- %s", indent, toYAMLValue(item, next)))
			}
		}
		return strings.Join(lines, "\n")
	default:
		return fmt.Sprintf("%v", v)
	}
}

func readCSV(path string) ([][]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	return reader.ReadAll()
}

func csvToK8sYAML(rows [][]string, kind string) ([]string, error) {
	if len(rows) < 2 {
		return nil, fmt.Errorf("CSV must have a header row and at least one data row")
	}

	headers := rows[0]
	var outputs []string

	for _, row := range rows[1:] {
		if len(row) != len(headers) {
			continue
		}

		meta := make(map[string]any)
		spec := make(map[string]any)

		for i, h := range headers {
			val := row[i]
			switch h {
			case "name", "namespace", "labels", "annotations":
				meta[h] = val
			default:
				spec[h] = val
			}
		}

		if _, ok := meta["name"]; !ok {
			if len(spec) > 0 {
				firstKey := sortedKeys(spec)[0]
				meta["name"] = spec[firstKey]
			} else {
				meta["name"] = "resource"
			}
		}

		var sb strings.Builder
		sb.WriteString("apiVersion: v1\n")
		sb.WriteString(fmt.Sprintf("kind: %s\n", kind))
		sb.WriteString("metadata:\n")
		for _, k := range sortedKeys(meta) {
			sb.WriteString(fmt.Sprintf("  %s: %s\n", k, toYAMLValue(meta[k], "    ")))
		}
		if len(spec) > 0 {
			sb.WriteString("spec:\n")
			for _, k := range sortedKeys(spec) {
				sb.WriteString(fmt.Sprintf("  %s: %s\n", k, toYAMLValue(spec[k], "    ")))
			}
		}
		outputs = append(outputs, sb.String())
	}

	return outputs, nil
}

var rootCmd = &cobra.Command{
	Use:   "csv-to-graph",
	Short: "Generate a graph representation from CSV data",
	Long: `csv-to-graph reads a CSV file with source and target columns and generates
a graph representation in various formats (DOT, JSON, adjacency list).

Examples:
  csv-to-graph edges.csv
  csv-to-graph -f dot -o graph.dot edges.csv
  csv-to-graph -s source -t target edges.csv`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("exactly one CSV file path required")
		}

		format, _ := cmd.Flags().GetString("format")
		sourceCol, _ := cmd.Flags().GetString("source")
		targetCol, _ := cmd.Flags().GetString("target")
		outputFile, _ := cmd.Flags().GetString("output")

		rows, err := readCSV(args[0])
		if err != nil {
			return fmt.Errorf("error reading CSV: %w", err)
		}

		if len(rows) < 2 {
			return fmt.Errorf("CSV must have a header row and at least one data row")
		}

		headers := rows[0]
		srcIdx, tgtIdx := 0, 1

		if sourceCol != "" {
			for i, h := range headers {
				if h == sourceCol {
					srcIdx = i
					break
				}
			}
		}
		if targetCol != "" {
			for i, h := range headers {
				if h == targetCol {
					tgtIdx = i
					break
				}
			}
		}

		type Edge struct {
			Source string `json:"source"`
			Target string `json:"target"`
		}

		edges := make([]Edge, 0, len(rows)-1)
		nodes := make(map[string]struct{})

		for _, row := range rows[1:] {
			if len(row) <= srcIdx || len(row) <= tgtIdx {
				continue
			}
			src, tgt := row[srcIdx], row[tgtIdx]
			if src != "" && tgt != "" {
				edges = append(edges, Edge{Source: src, Target: tgt})
				nodes[src] = struct{}{}
				nodes[tgt] = struct{}{}
			}
		}

		var result string
		switch format {
		case "json":
			result = "{\"nodes\": ["
			nodeList := make([]string, 0, len(nodes))
			for n := range nodes {
				nodeList = append(nodeList, n)
			}
			sort.Strings(nodeList)
			for i, n := range nodeList {
				if i > 0 {
					result += ", "
				}
				result += fmt.Sprintf("\"%s\"", n)
			}
			result += "], \"edges\": ["
			for i, e := range edges {
				if i > 0 {
					result += ", "
				}
				result += fmt.Sprintf("{\"source\": \"%s\", \"target\": \"%s\"}", e.Source, e.Target)
			}
			result += "]}"

		case "dot":
			var sb strings.Builder
			sb.WriteString("digraph G {\n")
			sb.WriteString("  rankdir=LR;\n")
			for _, e := range edges {
				sb.WriteString(fmt.Sprintf("  \"%s\" -> \"%s\";\n", e.Source, e.Target))
			}
			sb.WriteString("}\n")
			result = sb.String()

		default:
			var sb strings.Builder
			sb.WriteString("Adjacency List:\n")
			sb.WriteString(strings.Repeat("-", 40) + "\n")
			adj := make(map[string][]string)
			for _, e := range edges {
				adj[e.Source] = append(adj[e.Source], e.Target)
			}
			srcs := make([]string, 0, len(adj))
			for s := range adj {
				srcs = append(srcs, s)
			}
			sort.Strings(srcs)
			for _, s := range srcs {
				sb.WriteString(fmt.Sprintf("%s -> %s\n", s, strings.Join(adj[s], ", ")))
			}
			sb.WriteString(strings.Repeat("-", 40) + "\n")
			sb.WriteString(fmt.Sprintf("Nodes: %d, Edges: %d\n", len(nodes), len(edges)))
			result = sb.String()
		}

		if outputFile != "" {
			if err := os.WriteFile(outputFile, []byte(result), 0644); err != nil {
				return fmt.Errorf("error writing output: %w", err)
			}
			fmt.Printf("Written to %s\n", outputFile)
		} else {
			fmt.Print(result)
		}

		return nil
	},
}

func init() {
	rootCmd.Flags().StringP("format", "f", "adjacency", "Output format: adjacency, dot, json")
	rootCmd.Flags().StringP("source", "s", "", "Source column name (default: first column)")
	rootCmd.Flags().StringP("target", "t", "", "Target column name (default: second column)")
	rootCmd.Flags().StringP("output", "o", "", "Output file (default: stdout)")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
