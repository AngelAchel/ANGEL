package evasion

import (
	"time"
)

type Evasion0027 struct{}

func NewEvasion0027() *Evasion0027 {
	return &Evasion0027{}
}

func (e *Evasion0027) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *Evasion0027) Name() string { return "Evasion0027" }
func (e *Evasion0027) Timestamp() time.Time { return time.Now() }
