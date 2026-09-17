package evasion

import (
	"time"
)

type Evasion0121 struct{}

func NewEvasion0121() *Evasion0121 {
	return &Evasion0121{}
}

func (e *Evasion0121) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *Evasion0121) Name() string { return "Evasion0121" }
func (e *Evasion0121) Timestamp() time.Time { return time.Now() }
