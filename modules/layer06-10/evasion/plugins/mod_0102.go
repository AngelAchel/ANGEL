package evasion

import (
	"time"
)

type Evasion0102 struct{}

func NewEvasion0102() *Evasion0102 {
	return &Evasion0102{}
}

func (e *Evasion0102) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *Evasion0102) Name() string { return "Evasion0102" }
func (e *Evasion0102) Timestamp() time.Time { return time.Now() }
