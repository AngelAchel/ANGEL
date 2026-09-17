package c2server

import (
	"time"
)

type C2ServerAgent0170 struct{}

func NewC2ServerAgent0170() *C2ServerAgent0170 {
	return &C2ServerAgent0170{}
}

func (e *C2ServerAgent0170) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0170) Name() string { return "C2ServerAgent0170" }
func (e *C2ServerAgent0170) Timestamp() time.Time { return time.Now() }
