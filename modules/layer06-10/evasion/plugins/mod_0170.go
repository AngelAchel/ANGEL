package evasion

import (
	"time"
)

type Evasion0170 struct{}

func NewEvasion0170() *Evasion0170 {
	return &Evasion0170{}
}

func (e *Evasion0170) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *Evasion0170) Name() string { return "Evasion0170" }
func (e *Evasion0170) Timestamp() time.Time { return time.Now() }
