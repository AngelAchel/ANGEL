package c2server

import (
	"time"
)

type C2ServerAgent0014 struct{}

func NewC2ServerAgent0014() *C2ServerAgent0014 {
	return &C2ServerAgent0014{}
}

func (e *C2ServerAgent0014) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0014) Name() string { return "C2ServerAgent0014" }
func (e *C2ServerAgent0014) Timestamp() time.Time { return time.Now() }
