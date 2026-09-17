package webmisc

import (
	"testing"
)

func TestUARotateRotate(t *testing.T) {
	u := NewUARotate()
	if u.Name() != "UARotate" {
		t.Errorf("expected UARotate, got %s", u.Name())
	}
}
