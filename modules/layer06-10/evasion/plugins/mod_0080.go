package evasion

import (
	"time"
)

type Evasion0080 struct{}

func NewEvasion0080() *Evasion0080 {
	return &Evasion0080{}
}

func (e *Evasion0080) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *Evasion0080) Name() string { return "Evasion0080" }
func (e *Evasion0080) Timestamp() time.Time { return time.Now() }
