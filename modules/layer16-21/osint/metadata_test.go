package osint

import (
	"testing"
)

func TestMetadataGather(t *testing.T) {
	m := NewMetadata()
	if m.Name() != "Metadata" {
		t.Errorf("expected Metadata, got %s", m.Name())
	}
}
