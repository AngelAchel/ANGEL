package evasion

import (
	"time"
)

type Evasion0020 struct{}

func NewEvasion0020() *Evasion0020 {
	return &Evasion0020{}
}

func (e *Evasion0020) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *Evasion0020) Name() string { return "Evasion0020" }
func (e *Evasion0020) Timestamp() time.Time { return time.Now() }
