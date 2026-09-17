package osint

import (
	"testing"
)

func TestReportsGenerate(t *testing.T) {
	r := NewReports()
	if r.Name() != "Reports" {
		t.Errorf("expected Reports, got %s", r.Name())
	}
}
