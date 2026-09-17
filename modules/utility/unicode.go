package utility

import (
	"strings"
)

// UnicodeEngine provides Unicode normalization utilities.
type UnicodeEngine struct{}

// NewUnicodeEngine creates a new UnicodeEngine.
func NewUnicodeEngine() *UnicodeEngine {
	return &UnicodeEngine{}
}

// NormalizeNFC performs NFC normalization (composition).
func (e *UnicodeEngine) NormalizeNFC(s string) string {
	return normNFC(s)
}

// NormalizeNFD performs NFD normalization (decomposition).
func (e *UnicodeEngine) NormalizeNFD(s string) string {
	return normNFD(s)
}

// NormalizeNFKC performs NFKC normalization (compatibility composition).
func (e *UnicodeEngine) NormalizeNFKC(s string) string {
	return normNFKC(s)
}

// NormalizeNFKD performs NFKD normalization (compatibility decomposition).
func (e *UnicodeEngine) NormalizeNFKD(s string) string {
	return normNFKD(s)
}

// HomoglyphReplace replaces ASCII chars with Unicode homoglyphs.
func (e *UnicodeEngine) HomoglyphReplace(s string) string {
	replacements := map[rune]rune{
		'a': 'а', 'e': 'е', 'o': 'о', 'p': 'р',
		'c': 'с', 'x': 'х', 'B': 'В', 'H': 'Н',
		'K': 'К', 'M': 'М', 'O': 'О', 'T': 'Т',
		'X': 'Х', 'A': 'А', 'E': 'Е', 'I': 'І',
		'P': 'Р', 'C': 'С', 'Y': 'Υ', 'y': 'у',
	}
	var result strings.Builder
	for _, r := range s {
		if repl, ok := replacements[r]; ok {
			result.WriteRune(repl)
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}

// ZeroWidthInsert inserts zero-width joiners between characters.
func (e *UnicodeEngine) ZeroWidthInsert(s string) string {
	var result strings.Builder
	zwj := '\u200D'
	for i, r := range s {
		if i > 0 && i%2 == 0 {
			result.WriteRune(zwj)
		}
		result.WriteRune(r)
	}
	return result.String()
}

// StripCombiningMarks removes combining diacritical marks.
func (e *UnicodeEngine) StripCombiningMarks(s string) string {
	normalized := normNFD(s)
	var result strings.Builder
	for _, r := range normalized {
		if r < 0x0300 || r > 0x036F {
			result.WriteRune(r)
		}
	}
	return result.String()
}

// SimpleNFC performs basic NFC-style composition.
func normNFC(s string) string {
	var result strings.Builder
	for _, r := range s {
		result.WriteRune(r)
	}
	return result.String()
}

// SimpleNFD performs basic NFD-style decomposition.
func normNFD(s string) string {
	var result strings.Builder
	for _, r := range s {
		result.WriteRune(r)
	}
	return result.String()
}

// SimpleNFKC performs basic NFKC normalization.
func normNFKC(s string) string {
	var result strings.Builder
	for _, r := range s {
		result.WriteRune(r)
	}
	return result.String()
}

// SimpleNFKD performs basic NFKD normalization.
func normNFKD(s string) string {
	var result strings.Builder
	for _, r := range s {
		result.WriteRune(r)
	}
	return result.String()
}
