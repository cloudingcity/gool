package text

import (
	"strings"
)

// Escape escapes quotes and backslashes in text for JSON string usage
// Example: {"key": "say "value""} -> {\"key\": \"say \"value\"\"}
func Escape(s string) string {
	// First escape backslashes, then escape quotes
	s = strings.ReplaceAll(s, "\\", "\\\\")
	s = strings.ReplaceAll(s, "\"", "\\\"")
	return s
}

// Unescape unescapes quotes and backslashes in text
// Example: {\"key\": \"say \"value\"\"} -> {"key": "say "value""}
func Unescape(s string) string {
	// First unescape quotes, then unescape backslashes
	s = strings.ReplaceAll(s, "\\\"", "\"")
	s = strings.ReplaceAll(s, "\\\\", "\\")
	return s
}
