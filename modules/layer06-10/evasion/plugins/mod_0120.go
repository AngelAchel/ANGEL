package evasion

import (
	"time"
)

type Evasion0120 struct{}

func NewEvasion0120() *Evasion0120 {
	return &Evasion0120{}
}

func (e *Evasion0120) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *Evasion0120) Name() string { return "Evasion0120" }
func (e *Evasion0120) Timestamp() time.Time { return time.Now() }
