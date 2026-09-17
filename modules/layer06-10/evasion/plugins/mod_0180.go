package evasion

import (
	"time"
)

type Evasion0180 struct{}

func NewEvasion0180() *Evasion0180 {
	return &Evasion0180{}
}

func (e *Evasion0180) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *Evasion0180) Name() string { return "Evasion0180" }
func (e *Evasion0180) Timestamp() time.Time { return time.Now() }
