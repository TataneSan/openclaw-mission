package inferrer

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// DataType represents an inferred column type.
type DataType string

const (
	TypeUnknown  DataType = "unknown"
	TypeString   DataType = "string"
	TypeInteger  DataType = "integer"
	TypeFloat    DataType = "float"
	TypeBoolean  DataType = "boolean"
	TypeDate     DataType = "date"
	TypeDateTime DataType = "datetime"
	TypeEmail    DataType = "email"
	TypeURL      DataType = "url"
	TypeIP       DataType = "ip"
)

// ColumnSchema holds the inferred schema for a single column.
type ColumnSchema struct {
	Name       string   `json:"name"`
	Type       DataType `json:"type"`
	Confidence float64  `json:"confidence"`
	Nullable   bool     `json:"nullable"`
	Samples    []string `json:"samples,omitempty"`
	NullCount  int      `json:"null_count"`
	TotalCount int      `json:"total_count"`
}

// Infer analyzes column values and returns the inferred schema.
func Infer(name string, values []string) *ColumnSchema {
	total := len(values)
	if total == 0 {
		return &ColumnSchema{
			Name:       name,
			Type:       TypeUnknown,
			Confidence: 0,
			Nullable:   true,
		}
	}

	nullCount := 0
	nonEmpty := make([]string, 0, total)
	for _, v := range values {
		trimmed := strings.TrimSpace(v)
		if trimmed == "" {
			nullCount++
		} else {
			nonEmpty = append(nonEmpty, trimmed)
		}
	}

	if len(nonEmpty) == 0 {
		return &ColumnSchema{
			Name:       name,
			Type:       TypeString,
			Confidence: 0,
			Nullable:   true,
			NullCount:  nullCount,
			TotalCount: total,
		}
	}

	// Try each type and pick the best match
	bestType := TypeString
	bestConfidence := 0.0

	candidates := []struct {
		fn       func([]string) (DataType, float64)
		priority int // lower = checked first, higher confidence boost
	}{
		{inferBoolean, 1},
		{inferInteger, 2},
		{inferFloat, 3},
		{inferDate, 4},
		{inferDateTime, 5},
		{inferEmail, 6},
		{inferURL, 7},
		{inferIP, 8},
	}

	for _, c := range candidates {
		t, conf := c.fn(nonEmpty)
		if t != TypeUnknown && conf > bestConfidence {
			bestType = t
			bestConfidence = conf
		}
	}

	// Samples: up to 5 representative values
	samples := nonEmpty
	if len(samples) > 5 {
		samples = samples[:5]
	}

	return &ColumnSchema{
		Name:       name,
		Type:       bestType,
		Confidence: roundTo2(bestConfidence * 100),
		Nullable:   nullCount > 0,
		Samples:    samples,
		NullCount:  nullCount,
		TotalCount: total,
	}
}

func inferBoolean(values []string) (DataType, float64) {
	match := 0
	for _, v := range values {
		lower := strings.ToLower(v)
		if lower == "true" || lower == "false" || lower == "yes" || lower == "no" || lower == "1" || lower == "0" || lower == "on" || lower == "off" {
			match++
		}
	}
	if match == len(values) && len(values) > 0 {
		return TypeBoolean, 1.0
	}
	return TypeUnknown, 0
}

func inferInteger(values []string) (DataType, float64) {
	match := 0
	for _, v := range values {
		if _, err := strconv.ParseInt(v, 10, 64); err == nil {
			match++
		}
	}
	if match == len(values) && len(values) > 0 {
		return TypeInteger, 1.0
	}
	ratio := float64(match) / float64(len(values))
	if ratio >= 0.8 {
		return TypeInteger, ratio
	}
	return TypeUnknown, 0
}

func inferFloat(values []string) (DataType, float64) {
	match := 0
	for _, v := range values {
		if _, err := strconv.ParseFloat(v, 64); err == nil {
			match++
		}
	}
	if match == len(values) && len(values) > 0 {
		return TypeFloat, 1.0
	}
	ratio := float64(match) / float64(len(values))
	if ratio >= 0.8 {
		return TypeFloat, ratio
	}
	return TypeUnknown, 0
}

var dateFormats = []string{
	"2006-01-02",
	"01/02/2006",
	"02/01/2006",
	"2006/01/02",
	"02-01-2006",
	"01-02-2006",
	"20060102",
	"Jan 2, 2006",
	"2 Jan 2006",
	"January 2, 2006",
	"2 January 2006",
	"02.01.2006",
	"01.02.2006",
}

func inferDate(values []string) (DataType, float64) {
	match := 0
	for _, v := range values {
		if tryParseDate(v) {
			match++
		}
	}
	if match == len(values) && len(values) > 0 {
		return TypeDate, 1.0
	}
	ratio := float64(match) / float64(len(values))
	if ratio >= 0.8 {
		return TypeDate, ratio
	}
	return TypeUnknown, 0
}

func tryParseDate(s string) bool {
	for _, layout := range dateFormats {
		if _, err := time.Parse(layout, s); err == nil {
			return true
		}
	}
	return false
}

var dateTimeFormats = []string{
	"2006-01-02T15:04:05Z",
	"2006-01-02T15:04:05",
	"2006-01-02 15:04:05",
	"01/02/2006 15:04:05",
	"2006-01-02T15:04:05-07:00",
	"2006-01-02T15:04:05.000Z",
	"2006-01-02 15:04",
	"01/02/2006 15:04",
	"Jan 2, 2006 15:04:05",
	"2006-01-02 15:04:05.000",
}

func inferDateTime(values []string) (DataType, float64) {
	match := 0
	for _, v := range values {
		if tryParseDateTime(v) {
			match++
		}
	}
	if match == len(values) && len(values) > 0 {
		return TypeDateTime, 1.0
	}
	ratio := float64(match) / float64(len(values))
	if ratio >= 0.8 {
		return TypeDateTime, ratio
	}
	return TypeUnknown, 0
}

func tryParseDateTime(s string) bool {
	for _, layout := range dateTimeFormats {
		if _, err := time.Parse(layout, s); err == nil {
			return true
		}
	}
	return false
}

func inferEmail(values []string) (DataType, float64) {
	match := 0
	for _, v := range values {
		if lookLikeEmail(v) {
			match++
		}
	}
	if match == len(values) && len(values) > 0 {
		return TypeEmail, 1.0
	}
	ratio := float64(match) / float64(len(values))
	if ratio >= 0.8 {
		return TypeEmail, ratio
	}
	return TypeUnknown, 0
}

func lookLikeEmail(s string) bool {
	if !strings.Contains(s, "@") {
		return false
	}
	parts := strings.SplitN(s, "@", 2)
	if len(parts) != 2 {
		return false
	}
	local := parts[0]
	domain := parts[1]
	if len(local) == 0 || len(domain) == 0 {
		return false
	}
	if !strings.Contains(domain, ".") {
		return false
	}
	return true
}

func inferURL(values []string) (DataType, float64) {
	match := 0
	for _, v := range values {
		if lookLikeURL(v) {
			match++
		}
	}
	if match == len(values) && len(values) > 0 {
		return TypeURL, 1.0
	}
	ratio := float64(match) / float64(len(values))
	if ratio >= 0.8 {
		return TypeURL, ratio
	}
	return TypeUnknown, 0
}

func lookLikeURL(s string) bool {
	return strings.HasPrefix(s, "http://") || strings.HasPrefix(s, "https://") || strings.HasPrefix(s, "www.")
}

func inferIP(values []string) (DataType, float64) {
	match := 0
	for _, v := range values {
		if lookLikeIP(v) {
			match++
		}
	}
	if match == len(values) && len(values) > 0 {
		return TypeIP, 1.0
	}
	ratio := float64(match) / float64(len(values))
	if ratio >= 0.8 {
		return TypeIP, ratio
	}
	return TypeUnknown, 0
}

func lookLikeIP(s string) bool {
	parts := strings.Split(s, ".")
	if len(parts) != 4 {
		return false
	}
	for _, p := range parts {
		n, err := strconv.Atoi(p)
		if err != nil || n < 0 || n > 255 {
			return false
		}
	}
	return true
}

func roundTo2(f float64) float64 {
	return float64(int(f*100+0.5)) / 100
}

// SchemaError is returned when schema inference fails.
type SchemaError struct {
	Column string
	Cause  error
}

func (e SchemaError) Error() string {
	return fmt.Sprintf("schema inference failed for column %q: %v", e.Column, e.Cause)
}
