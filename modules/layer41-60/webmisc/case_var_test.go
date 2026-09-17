package webmisc

import (
	"testing"
)

func TestCaseVarVary(t *testing.T) {
	c := NewCaseVar()
	if c.Name() != "CaseVar" {
		t.Errorf("expected CaseVar, got %s", c.Name())
	}
}
