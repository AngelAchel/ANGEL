package evasion

import (
	"time"
)

type Evasion0117 struct{}

func NewEvasion0117() *Evasion0117 {
	return &Evasion0117{}
}

func (e *Evasion0117) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *Evasion0117) Name() string { return "Evasion0117" }
func (e *Evasion0117) Timestamp() time.Time { return time.Now() }
