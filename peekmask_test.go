package peekmask

import "testing"

func TestMasker_Mask(t *testing.T) {
	tests := []struct {
		name        string
		masker      *Masker
		input       string
		expected    string
		description string
	}{
		{
			name:        "Basic masking",
			masker:      New('*', 2, 2, 0.3),
			input:       "1234567890",
			expected:    "12******90",
			description: "Keep first 2 and last 2 characters, mask the rest",
		},
		{
			name:        "Japanese text masking",
			masker:      New('*', 2, 2, 0.3),
			input:       "こんにちは世界です",
			expected:    "こん*****です",
			description: "Correctly mask Japanese characters",
		},
		{
			name:        "Short string (mask all)",
			masker:      New('*', 2, 2, 0.3),
			input:       "abc",
			expected:    "***",
			description: "Mask all when prefix+suffix >= length",
		},
		{
			name:        "Empty string",
			masker:      New('*', 2, 2, 0.3),
			input:       "",
			expected:    "",
			description: "Return empty string as is",
		},
		{
			name:        "Below minimum mask ratio (mask all)",
			masker:      New('*', 4, 4, 0.5),
			input:       "1234567890",
			expected:    "**********",
			description: "Mask all when ratio 0.2 < 0.5",
		},
		{
			name:        "Custom mask character",
			masker:      New('■', 1, 1, 0.3),
			input:       "abcdef",
			expected:    "a■■■■f",
			description: "Use custom mask character",
		},
		{
			name:        "Keep prefix only",
			masker:      New('*', 3, 0, 0.3),
			input:       "1234567890",
			expected:    "123*******",
			description: "Keep only first 3 characters",
		},
		{
			name:        "Keep suffix only",
			masker:      New('*', 0, 3, 0.3),
			input:       "1234567890",
			expected:    "*******890",
			description: "Keep only last 3 characters",
		},
		{
			name:        "Emoji string",
			masker:      New('*', 2, 2, 0.3),
			input:       "😀😁😂😃😄😅😆😇",
			expected:    "😀😁****😆😇",
			description: "Correctly handle emoji as UTF-8",
		},
		{
			name:        "Mixed character string",
			masker:      New('X', 3, 3, 0.2),
			input:       "abc123あいう",
			expected:    "abcXXXあいう",
			description: "Mix of alphanumeric and Japanese characters",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.masker.Mask(tt.input)
			if result != tt.expected {
				t.Errorf("%s\nInput: %q\nExpected: %q\nGot: %q",
					tt.description, tt.input, tt.expected, result)
			}
		})
	}
}

func TestMask(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Mask with default settings",
			input:    "1234567890",
			expected: "12******90",
		},
		{
			name:     "Japanese text with default settings",
			input:    "メールアドレス",
			expected: "メー***レス",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Mask(tt.input, 2, 2)
			if result != tt.expected {
				t.Errorf("Input: %q, Expected: %q, Got: %q", tt.input, tt.expected, result)
			}
		})
	}
}

func TestDefault(t *testing.T) {
	masker := Default()
	if masker.MaskChar != '*' {
		t.Errorf("Default MaskChar is incorrect: %c", masker.MaskChar)
	}
	if masker.PrefixLength != 2 {
		t.Errorf("Default PrefixLength is incorrect: %d", masker.PrefixLength)
	}
	if masker.SuffixLength != 2 {
		t.Errorf("Default SuffixLength is incorrect: %d", masker.SuffixLength)
	}
	if masker.MinMaskRatio != 0.3 {
		t.Errorf("Default MinMaskRatio is incorrect: %f", masker.MinMaskRatio)
	}
}

// Benchmark
func BenchmarkMask(b *testing.B) {
	masker := New('*', 2, 2, 0.3)
	input := "これは日本語を含む比較的長い文字列のテストです1234567890"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		masker.Mask(input)
	}
}
