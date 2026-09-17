package c2server

import (
	"time"
)

type C2ServerAgent0080 struct{}

func NewC2ServerAgent0080() *C2ServerAgent0080 {
	return &C2ServerAgent0080{}
}

func (e *C2ServerAgent0080) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0080) Name() string { return "C2ServerAgent0080" }
func (e *C2ServerAgent0080) Timestamp() time.Time { return time.Now() }
