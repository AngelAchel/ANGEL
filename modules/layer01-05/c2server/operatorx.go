package c2server

import (
	"time"
)

type Operatorx struct{}

func NewOperatorx() *Operatorx {
	return &Operatorx{}
}

func (e *Operatorx) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "operatorx:done")
	return results, nil
}

func (e *Operatorx) Name() string         { return "Operatorx" }
func (e *Operatorx) Timestamp() time.Time { return time.Now() }
