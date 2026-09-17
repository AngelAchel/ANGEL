package utility

import (
	"strings"
)

// WhitespaceEngine provides whitespace obfuscation utilities.
type WhitespaceEngine struct{}

// NewWhitespaceEngine creates a new WhitespaceEngine.
func NewWhitespaceEngine() *WhitespaceEngine {
	return &WhitespaceEngine{}
}

// Obfuscate inserts whitespace variations into a string.
func (e *WhitespaceEngine) Obfuscate(s string) string {
	var result strings.Builder
	for i, r := range s {
		if i > 0 && i%3 == 0 {
			result.WriteString("\u00A0") // non-breaking space
		}
		result.WriteRune(r)
	}
	return result.String()
}

// TabInject inserts tab characters at regular intervals.
func (e *WhitespaceEngine) TabInject(s string, interval int) string {
	if interval <= 0 {
		interval = 5
	}
	var result strings.Builder
	for i, r := range s {
		if i > 0 && i%interval == 0 {
			result.WriteString("\t")
		}
		result.WriteRune(r)
	}
	return result.String()
}

// ZeroWidthEncode encodes a binary string using zero-width characters.
func (e *WhitespaceEngine) ZeroWidthEncode(bits string) string {
	zwj := "\u200D"   // zero-width joiner
	zwnj := "\u200C"  // zero-width non-joiner
	zwsp := "\u200B"  // zero-width space

	var result strings.Builder
	for _, bit := range bits {
		switch bit {
		case '0':
			result.WriteString(zwj)
		case '1':
			result.WriteString(zwnj)
		default:
			result.WriteString(zwsp)
		}
	}
	return result.String()
}

// ZeroWidthDecode decodes a zero-width encoded string back to bits.
func (e *WhitespaceEngine) ZeroWidthDecode(s string) string {
	zwj := "\u200D"
	zwnj := "\u200C"

	var bits strings.Builder
	for _, r := range s {
		switch string(r) {
		case zwj:
			bits.WriteString("0")
		case zwnj:
			bits.WriteString("1")
		}
	}
	return bits.String()
}

// LineBreakObfuscate inserts line breaks at intervals.
func (e *WhitespaceEngine) LineBreakObfuscate(s string, interval int) string {
	if interval <= 0 {
		interval = 10
	}
	var result strings.Builder
	for i, r := range s {
		if i > 0 && i%interval == 0 {
			result.WriteString("\n")
		}
		result.WriteRune(r)
	}
	return result.String()
}

// MixedWhitespace mixes various whitespace types for obfuscation.
func (e *WhitespaceEngine) MixedWhitespace(s string) string {
	whitespaces := []string{" ", "\t", "\n", "\r", "\u00A0", "\u2002", "\u2003", "\u2009"}
	var result strings.Builder
	for i, r := range s {
		result.WriteRune(r)
		if i > 0 && i%4 == 0 {
			idx := i % len(whitespaces)
			result.WriteString(whitespaces[idx])
		}
	}
	return result.String()
}