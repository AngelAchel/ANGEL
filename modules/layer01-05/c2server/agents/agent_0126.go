package c2server

import (
	"time"
)

type C2ServerAgent0126 struct{}

func NewC2ServerAgent0126() *C2ServerAgent0126 {
	return &C2ServerAgent0126{}
}

func (e *C2ServerAgent0126) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0126) Name() string { return "C2ServerAgent0126" }
func (e *C2ServerAgent0126) Timestamp() time.Time { return time.Now() }
