package evasion

import (
	"time"
)

type Evasion0105 struct{}

func NewEvasion0105() *Evasion0105 {
	return &Evasion0105{}
}

func (e *Evasion0105) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *Evasion0105) Name() string { return "Evasion0105" }
func (e *Evasion0105) Timestamp() time.Time { return time.Now() }
