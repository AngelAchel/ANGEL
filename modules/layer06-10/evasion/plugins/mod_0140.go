package evasion

import (
	"time"
)

type Evasion0140 struct{}

func NewEvasion0140() *Evasion0140 {
	return &Evasion0140{}
}

func (e *Evasion0140) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *Evasion0140) Name() string { return "Evasion0140" }
func (e *Evasion0140) Timestamp() time.Time { return time.Now() }
