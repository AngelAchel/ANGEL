package evasion

import (
	"time"
)

type Evasion0123 struct{}

func NewEvasion0123() *Evasion0123 {
	return &Evasion0123{}
}

func (e *Evasion0123) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *Evasion0123) Name() string { return "Evasion0123" }
func (e *Evasion0123) Timestamp() time.Time { return time.Now() }
