package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

func parseEnv(content []byte) map[string]string {
	result := make(map[string]string)
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			value = strings.Trim(value, "\"")
			value = strings.Trim(value, "'")
			if key != "" {
				result[key] = value
			}
		}
	}
	return result
}

var rootCmd = &cobra.Command{
	Use:   "env-to-k8s-secret",
	Short: "Convert a .env file to a Kubernetes Secret YAML",
	Long: `env-to-k8s-secret reads a .env file and generates a Kubernetes Secret manifest
with base64-encoded values.

Examples:
  env-to-k8s-secret .env
  env-to-k8s-secret -n my-secret .env
  env-to-k8s-secret -n my-secret -s default .env
  env-to-k8s-secret -o secret.yaml .env`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("exactly one .env file path required")
		}

		name, _ := cmd.Flags().GetString("name")
		namespace, _ := cmd.Flags().GetString("namespace")
		outputFile, _ := cmd.Flags().GetString("output")

		if name == "" {
			name = "env-secret"
		}

		content, err := os.ReadFile(args[0])
		if err != nil {
			return fmt.Errorf("error reading file: %w", err)
		}

		env := parseEnv(content)
		keys := make([]string, 0, len(env))
		for k := range env {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		var sb strings.Builder
		sb.WriteString("apiVersion: v1\n")
		sb.WriteString("kind: Secret\n")
		sb.WriteString("metadata:\n")
		sb.WriteString(fmt.Sprintf("  name: %s\n", name))
		if namespace != "" {
			sb.WriteString(fmt.Sprintf("  namespace: %s\n", namespace))
		}
		sb.WriteString("type: Opaque\n")
		sb.WriteString("data:\n")

		for _, key := range keys {
			encoded := base64.StdEncoding.EncodeToString([]byte(env[key]))
			sb.WriteString(fmt.Sprintf("  %s: %s\n", key, encoded))
		}

		if outputFile != "" {
			if err := os.WriteFile(outputFile, []byte(sb.String()), 0644); err != nil {
				return fmt.Errorf("error writing output: %w", err)
			}
			fmt.Printf("Written to %s\n", outputFile)
		} else {
			fmt.Print(sb.String())
		}

		return nil
	},
}

func init() {
	rootCmd.Flags().StringP("name", "n", "", "Secret name (default: env-secret)")
	rootCmd.Flags().StringP("namespace", "s", "", "Kubernetes namespace")
	rootCmd.Flags().StringP("output", "o", "", "Output file (default: stdout)")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
