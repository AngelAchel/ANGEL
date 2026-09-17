package evasion

import (
	"time"
)

type Evasion0028 struct{}

func NewEvasion0028() *Evasion0028 {
	return &Evasion0028{}
}

func (e *Evasion0028) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *Evasion0028) Name() string { return "Evasion0028" }
func (e *Evasion0028) Timestamp() time.Time { return time.Now() }
