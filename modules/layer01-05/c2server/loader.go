package c2server

import (
	"time"
)

type Loader struct{}

func NewLoader() *Loader {
	return &Loader{}
}

func (e *Loader) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "loader:done")
	return results, nil
}

func (e *Loader) Name() string         { return "Loader" }
func (e *Loader) Timestamp() time.Time { return time.Now() }
