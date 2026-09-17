package webmisc

import (
	"testing"
)

func TestMultipartUpload(t *testing.T) {
	m := NewMultipart()
	if m.Name() != "Multipart" {
		t.Errorf("expected Multipart, got %s", m.Name())
	}
}
