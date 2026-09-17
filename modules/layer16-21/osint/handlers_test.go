package osint

import (
	"testing"
)

func TestHandlersHandle(t *testing.T) {
	h := NewHandlers()
	if h.Name() != "Handlers" {
		t.Errorf("expected Handlers, got %s", h.Name())
	}
}
