package evasion

import (
	"time"
)

type Evasion0087 struct{}

func NewEvasion0087() *Evasion0087 {
	return &Evasion0087{}
}

func (e *Evasion0087) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *Evasion0087) Name() string { return "Evasion0087" }
func (e *Evasion0087) Timestamp() time.Time { return time.Now() }
