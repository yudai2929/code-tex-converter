package utils

import "testing"

func TestTrimPath(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Leading and trailing slashes",
			input:    "/path/to/file/",
			expected: "path/to/file",
		},
		{
			name:     "Only leading slash",
			input:    "/path/to/file",
			expected: "path/to/file",
		},
		{
			name:     "Only trailing slash",
			input:    "path/to/file/",
			expected: "path/to/file",
		},
		{
			name:     "No slashes",
			input:    "path/to/file",
			expected: "path/to/file",
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := TrimPath(tc.input)
			if result != tc.expected {
				t.Errorf("TrimPath(%q) = %q; want %q", tc.input, result, tc.expected)
			}
		})
	}
}
