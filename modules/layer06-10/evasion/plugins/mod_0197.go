package evasion

import (
	"time"
)

type Evasion0197 struct{}

func NewEvasion0197() *Evasion0197 {
	return &Evasion0197{}
}

func (e *Evasion0197) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *Evasion0197) Name() string { return "Evasion0197" }
func (e *Evasion0197) Timestamp() time.Time { return time.Now() }
