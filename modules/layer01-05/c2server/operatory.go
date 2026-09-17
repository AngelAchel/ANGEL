package c2server

import (
	"time"
)

type Operatory struct{}

func NewOperatory() *Operatory {
	return &Operatory{}
}

func (e *Operatory) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "operatory:done")
	return results, nil
}

func (e *Operatory) Name() string         { return "Operatory" }
func (e *Operatory) Timestamp() time.Time { return time.Now() }
