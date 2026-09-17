package webmisc

import (
	"testing"
)

func TestXMLParamInject(t *testing.T) {
	x := NewXMLParam()
	if x.Name() != "XMLParam" {
		t.Errorf("expected XMLParam, got %s", x.Name())
	}
}
