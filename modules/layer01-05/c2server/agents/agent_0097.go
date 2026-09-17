package c2server

import (
	"time"
)

type C2ServerAgent0097 struct{}

func NewC2ServerAgent0097() *C2ServerAgent0097 {
	return &C2ServerAgent0097{}
}

func (e *C2ServerAgent0097) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0097) Name() string { return "C2ServerAgent0097" }
func (e *C2ServerAgent0097) Timestamp() time.Time { return time.Now() }
