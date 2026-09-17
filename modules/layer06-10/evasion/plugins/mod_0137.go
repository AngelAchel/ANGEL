package evasion

import (
	"time"
)

type Evasion0137 struct{}

func NewEvasion0137() *Evasion0137 {
	return &Evasion0137{}
}

func (e *Evasion0137) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *Evasion0137) Name() string { return "Evasion0137" }
func (e *Evasion0137) Timestamp() time.Time { return time.Now() }
