package peekmask

import (
	"strings"
	"unicode/utf8"
)

// Masker holds the configuration for masking strings
type Masker struct {
	// MaskChar is the character used for masking
	MaskChar rune
	// PrefixLength is the number of characters to keep from the start
	PrefixLength int
	// SuffixLength is the number of characters to keep from the end
	SuffixLength int
	// MinMaskRatio is the minimum masking ratio (0.0-1.0)
	// If the masking ratio falls below this value, all characters will be masked
	MinMaskRatio float64
}

// New creates a new Masker with the specified configuration
func New(maskChar rune, prefixLength, suffixLength int, minMaskRatio float64) *Masker {
	return &Masker{
		MaskChar:     maskChar,
		PrefixLength: prefixLength,
		SuffixLength: suffixLength,
		MinMaskRatio: minMaskRatio,
	}
}

// Default returns a Masker with default settings
// MaskChar: '*', PrefixLength: 2, SuffixLength: 2, MinMaskRatio: 0.3
func Default() *Masker {
	return &Masker{
		MaskChar:     '*',
		PrefixLength: 2,
		SuffixLength: 2,
		MinMaskRatio: 0.3,
	}
}

// defaultMasker is a shared instance for the global Mask function
var defaultMasker = Default()

// Mask masks the given string
func (m *Masker) Mask(s string) string {
	return maskString(s, m.MaskChar, m.PrefixLength, m.SuffixLength, m.MinMaskRatio)
}

func Mask(s string, prefixLength, suffixLength int) string {
	return maskString(s, defaultMasker.MaskChar, prefixLength, suffixLength, defaultMasker.MinMaskRatio)
}

func maskString(s string, maskChar rune, prefixLength, suffixLength int, minMaskRatio float64) string {
	// Count runes without allocating the slice yet
	length := utf8.RuneCountInString(s)

	// Handle empty string
	if length == 0 {
		return s
	}

	// If prefix and suffix lengths exceed the string length, mask everything
	if prefixLength+suffixLength >= length {
		return strings.Repeat(string(maskChar), length)
	}

	// Calculate masking ratio
	maskCount := length - prefixLength - suffixLength
	maskRatio := float64(maskCount) / float64(length)

	// If mask ratio is below minimum, mask everything
	if maskRatio < minMaskRatio {
		return strings.Repeat(string(maskChar), length)
	}

	// Only allocate runes slice when we actually need to mask
	runes := []rune(s)
	result := make([]rune, length)
	copy(result[:prefixLength], runes[:prefixLength])
	for i := 0; i < maskCount; i++ {
		result[prefixLength+i] = maskChar
	}
	copy(result[length-suffixLength:], runes[length-suffixLength:])

	return string(result)
}
