package evasion

import (
	"time"
)

type Evasion0107 struct{}

func NewEvasion0107() *Evasion0107 {
	return &Evasion0107{}
}

func (e *Evasion0107) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *Evasion0107) Name() string { return "Evasion0107" }
func (e *Evasion0107) Timestamp() time.Time { return time.Now() }
