package evasion

import (
	"time"
)

type Evasion0101 struct{}

func NewEvasion0101() *Evasion0101 {
	return &Evasion0101{}
}

func (e *Evasion0101) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *Evasion0101) Name() string { return "Evasion0101" }
func (e *Evasion0101) Timestamp() time.Time { return time.Now() }
