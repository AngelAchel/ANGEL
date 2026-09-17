package storage

import (
	"testing"
)

func TestModelsGet(t *testing.T) {
	m := NewModels()
	if m.Name() != "Models" {
		t.Errorf("expected Models, got %s", m.Name())
	}
}
