package evasion

import (
	"time"
)

type Evasion0173 struct{}

func NewEvasion0173() *Evasion0173 {
	return &Evasion0173{}
}

func (e *Evasion0173) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *Evasion0173) Name() string { return "Evasion0173" }
func (e *Evasion0173) Timestamp() time.Time { return time.Now() }
