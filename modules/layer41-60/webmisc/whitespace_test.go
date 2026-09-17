package webmisc

import (
	"testing"
)

func TestWhitespaceObfuscate(t *testing.T) {
	w := NewWhitespace()
	if w.Name() != "Whitespace" {
		t.Errorf("expected Whitespace, got %s", w.Name())
	}
}
