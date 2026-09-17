package evasion

import (
	"time"
)

type Evasion0014 struct{}

func NewEvasion0014() *Evasion0014 {
	return &Evasion0014{}
}

func (e *Evasion0014) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *Evasion0014) Name() string { return "Evasion0014" }
func (e *Evasion0014) Timestamp() time.Time { return time.Now() }
