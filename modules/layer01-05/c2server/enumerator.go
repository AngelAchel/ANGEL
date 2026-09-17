package c2server

import (
	"time"
)

type Enumerator struct{}

func NewEnumerator() *Enumerator {
	return &Enumerator{}
}

func (e *Enumerator) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "enumerator:done")
	return results, nil
}

func (e *Enumerator) Name() string         { return "Enumerator" }
func (e *Enumerator) Timestamp() time.Time { return time.Now() }
