package evasion

import (
	"time"
)

type Evasion0008 struct{}

func NewEvasion0008() *Evasion0008 {
	return &Evasion0008{}
}

func (e *Evasion0008) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *Evasion0008) Name() string { return "Evasion0008" }
func (e *Evasion0008) Timestamp() time.Time { return time.Now() }
