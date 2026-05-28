package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestValidateValidJSON(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "test.json")
	os.WriteFile(path, []byte(`{"name": "test", "version": 1}`), 0644)

	result := validateFile(path, false)
	if !result.Valid {
		t.Errorf("expected valid JSON, got errors: %v", result.Errors)
	}
	if result.Format != "json" {
		t.Errorf("expected format json, got %s", result.Format)
	}
}

func TestValidateInvalidJSON(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "test.json")
	os.WriteFile(path, []byte(`{"name": "test",}`), 0644)

	result := validateFile(path, false)
	if result.Valid {
		t.Error("expected invalid JSON")
	}
	if len(result.Errors) == 0 {
		t.Error("expected errors")
	}
}

func TestValidateValidYAML(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "test.yaml")
	os.WriteFile(path, []byte("name: test\nversion: 1\n"), 0644)

	result := validateFile(path, false)
	if !result.Valid {
		t.Errorf("expected valid YAML, got errors: %v", result.Errors)
	}
}

func TestValidateYAMLStrictTabs(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "test.yaml")
	os.WriteFile(path, []byte("name: test\n\tbad: tab\n"), 0644)

	result := validateFile(path, true)
	if len(result.Errors) == 0 {
		t.Error("expected tab error in strict mode")
	}
}

func TestValidateValidTOML(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "test.toml")
	os.WriteFile(path, []byte("[section]\nkey = \"value\"\n"), 0644)

	result := validateFile(path, false)
	if !result.Valid {
		t.Errorf("expected valid TOML, got errors: %v", result.Errors)
	}
}

func TestValidateInvalidTOML(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "test.toml")
	os.WriteFile(path, []byte("[not closed\n"), 0644)

	result := validateFile(path, false)
	if result.Valid {
		t.Error("expected invalid TOML")
	}
}

func TestDetectFormat(t *testing.T) {
	tests := []struct {
		path     string
		expected string
	}{
		{"file.json", "json"},
		{"file.yaml", "yaml"},
		{"file.yml", "yaml"},
		{"file.toml", "toml"},
		{"file.ini", "ini"},
		{"file.cfg", "ini"},
		{"file.conf", "ini"},
		{"file.txt", ""},
	}

	for _, tt := range tests {
		if got := detectFormat(tt.path); got != tt.expected {
			t.Errorf("detectFormat(%q) = %q, want %q", tt.path, got, tt.expected)
		}
	}
}
