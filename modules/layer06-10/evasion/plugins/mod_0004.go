package evasion

import (
	"time"
)

type Evasion0004 struct{}

func NewEvasion0004() *Evasion0004 {
	return &Evasion0004{}
}

func (e *Evasion0004) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *Evasion0004) Name() string { return "Evasion0004" }
func (e *Evasion0004) Timestamp() time.Time { return time.Now() }
