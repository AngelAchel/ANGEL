package webmisc

import (
	"testing"
)

func TestUnicodeNormalize(t *testing.T) {
	u := NewUnicode()
	if u.Name() != "Unicode" {
		t.Errorf("expected Unicode, got %s", u.Name())
	}
}
