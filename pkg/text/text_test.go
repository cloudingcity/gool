package text

import "testing"

func TestEscape(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "basic quotes",
			input:    `{"key": "say "value""}`,
			expected: `{\"key\": \"say \"value\"\"}`,
		},
		{
			name:     "backslashes",
			input:    `path\to\file`,
			expected: `path\\to\\file`,
		},
		{
			name:     "mixed quotes and backslashes",
			input:    `C:\Program Files\test "app"`,
			expected: `C:\\Program Files\\test \"app\"`,
		},
		{
			name:     "already escaped",
			input:    `{\"key\": \"value\"}`,
			expected: `{\\\"key\\\": \\\"value\\\"}`,
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "no special characters",
			input:    "hello world",
			expected: "hello world",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Escape(tt.input)
			if result != tt.expected {
				t.Errorf("Escape(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestUnescape(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "basic quotes",
			input:    `{\"key\": \"say \"value\"\"}`,
			expected: `{"key": "say "value""}`,
		},
		{
			name:     "backslashes",
			input:    `path\\to\\file`,
			expected: `path\to\file`,
		},
		{
			name:     "mixed quotes and backslashes",
			input:    `C:\\Program Files\\test \"app\"`,
			expected: `C:\Program Files\test "app"`,
		},
		{
			name:     "double escaped",
			input:    `{\\\"key\\\": \\\"value\\\"}`,
			expected: `{\"key\": \"value\"}`,
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "no special characters",
			input:    "hello world",
			expected: "hello world",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Unescape(tt.input)
			if result != tt.expected {
				t.Errorf("Unescape(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestEscapeUnescape(t *testing.T) {
	// Test that escape and unescape are inverse operations for basic cases
	testCases := []string{
		`{"key": "value"}`,
		`hello "world"`,
		`path\to\file`,
		`mixed "quotes" and \backslashes\`,
		"simple text",
		"",
	}

	for _, input := range testCases {
		t.Run("roundtrip_"+input, func(t *testing.T) {
			escaped := Escape(input)
			unescaped := Unescape(escaped)
			if unescaped != input {
				t.Errorf("Roundtrip failed: input=%q, escaped=%q, unescaped=%q", input, escaped, unescaped)
			}
		})
	}
}
