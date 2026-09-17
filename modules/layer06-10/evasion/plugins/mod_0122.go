package evasion

import (
	"time"
)

type Evasion0122 struct{}

func NewEvasion0122() *Evasion0122 {
	return &Evasion0122{}
}

func (e *Evasion0122) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *Evasion0122) Name() string { return "Evasion0122" }
func (e *Evasion0122) Timestamp() time.Time { return time.Now() }
