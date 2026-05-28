package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
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
	case int, int64:
		return fmt.Sprintf("%d", val)
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

func yamlToK8sYAML(data map[string]any, kind string) string {
	var sb strings.Builder

	sb.WriteString("apiVersion: v1\n")
	sb.WriteString(fmt.Sprintf("kind: %s\n", kind))
	sb.WriteString("metadata:\n")

	meta := make(map[string]any)
	spec := make(map[string]any)

	for _, k := range sortedKeys(data) {
		if k == "name" || k == "namespace" || k == "labels" || k == "annotations" {
			meta[k] = data[k]
		} else {
			spec[k] = data[k]
		}
	}

	if len(meta) == 0 {
		meta["name"] = "resource"
	}

	for _, k := range sortedKeys(meta) {
		sb.WriteString(fmt.Sprintf("  %s: %s\n", k, toYAMLValue(meta[k], "    ")))
	}

	if len(spec) > 0 {
		sb.WriteString("spec:\n")
		for _, k := range sortedKeys(spec) {
			sb.WriteString(fmt.Sprintf("  %s: %s\n", k, toYAMLValue(spec[k], "    ")))
		}
	}

	return sb.String()
}

var rootCmd = &cobra.Command{
	Use:   "yaml-to-k8s",
	Short: "Convert YAML to Kubernetes YAML manifests",
	Long: `yaml-to-k8s reads a YAML file and generates Kubernetes YAML manifests.

Examples:
  yaml-to-k8s config.yaml
  yaml-to-k8s -k Deployment config.yaml
  yaml-to-k8s -o output.yaml config.yaml`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("exactly one YAML file path required")
		}

		kind, _ := cmd.Flags().GetString("kind")
		outputFile, _ := cmd.Flags().GetString("output")

		if kind == "" {
			kind = "ConfigMap"
		}

		content, err := os.ReadFile(args[0])
		if err != nil {
			return fmt.Errorf("error reading file: %w", err)
		}

		var data map[string]any
		if err := yaml.Unmarshal(content, &data); err != nil {
			return fmt.Errorf("error parsing YAML: %w", err)
		}

		var outputs []string
		for _, key := range sortedKeys(data) {
			if nested, ok := data[key].(map[string]any); ok {
				outputs = append(outputs, yamlToK8sYAML(nested, kind))
			} else {
				outputs = append(outputs, yamlToK8sYAML(map[string]any{"name": key, "value": data[key]}, kind))
			}
		}

		result := strings.Join(outputs, "---\n")

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
	rootCmd.Flags().StringP("kind", "k", "", "Kubernetes resource kind (default: ConfigMap)")
	rootCmd.Flags().StringP("output", "o", "", "Output file (default: stdout)")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
