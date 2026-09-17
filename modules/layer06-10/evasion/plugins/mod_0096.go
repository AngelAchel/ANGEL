package evasion

import (
	"time"
)

type Evasion0096 struct{}

func NewEvasion0096() *Evasion0096 {
	return &Evasion0096{}
}

func (e *Evasion0096) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *Evasion0096) Name() string { return "Evasion0096" }
func (e *Evasion0096) Timestamp() time.Time { return time.Now() }
