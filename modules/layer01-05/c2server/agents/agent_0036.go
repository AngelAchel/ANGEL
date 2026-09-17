package c2server

import (
	"time"
)

type C2ServerAgent0036 struct{}

func NewC2ServerAgent0036() *C2ServerAgent0036 {
	return &C2ServerAgent0036{}
}

func (e *C2ServerAgent0036) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0036) Name() string { return "C2ServerAgent0036" }
func (e *C2ServerAgent0036) Timestamp() time.Time { return time.Now() }
