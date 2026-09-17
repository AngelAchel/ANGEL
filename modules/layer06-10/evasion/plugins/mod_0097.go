package evasion

import (
	"time"
)

type Evasion0097 struct{}

func NewEvasion0097() *Evasion0097 {
	return &Evasion0097{}
}

func (e *Evasion0097) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *Evasion0097) Name() string { return "Evasion0097" }
func (e *Evasion0097) Timestamp() time.Time { return time.Now() }
