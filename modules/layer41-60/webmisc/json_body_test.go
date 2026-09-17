package webmisc

import (
	"testing"
)

func TestJSONBodyInject(t *testing.T) {
	j := NewJSONBody()
	if j.Name() != "JSONBody" {
		t.Errorf("expected JSONBody, got %s", j.Name())
	}
}
