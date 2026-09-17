package evasion

import (
	"time"
)

type Evasion0130 struct{}

func NewEvasion0130() *Evasion0130 {
	return &Evasion0130{}
}

func (e *Evasion0130) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *Evasion0130) Name() string { return "Evasion0130" }
func (e *Evasion0130) Timestamp() time.Time { return time.Now() }
