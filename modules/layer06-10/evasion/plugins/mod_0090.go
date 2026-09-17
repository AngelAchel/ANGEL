package evasion

import (
	"time"
)

type Evasion0090 struct{}

func NewEvasion0090() *Evasion0090 {
	return &Evasion0090{}
}

func (e *Evasion0090) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *Evasion0090) Name() string { return "Evasion0090" }
func (e *Evasion0090) Timestamp() time.Time { return time.Now() }
