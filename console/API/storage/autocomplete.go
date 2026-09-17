package storage

import (
	"time"
)

type Autocomplete struct{}

func NewAutocomplete() *Autocomplete {
	return &Autocomplete{}
}

func (a *Autocomplete) Complete() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "autocomplete:done")
	return results, nil
}

func (a *Autocomplete) Name() string         { return "Autocomplete" }
func (a *Autocomplete) Timestamp() time.Time { return time.Now() }
