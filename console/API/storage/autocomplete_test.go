package storage

import (
	"testing"
)

func TestAutocompleteComplete(t *testing.T) {
	a := NewAutocomplete()
	if a.Name() != "Autocomplete" {
		t.Errorf("expected Autocomplete, got %s", a.Name())
	}
}
