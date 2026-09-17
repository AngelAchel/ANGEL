package evasion

import (
	"time"
)

type Evasion0103 struct{}

func NewEvasion0103() *Evasion0103 {
	return &Evasion0103{}
}

func (e *Evasion0103) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *Evasion0103) Name() string { return "Evasion0103" }
func (e *Evasion0103) Timestamp() time.Time { return time.Now() }
