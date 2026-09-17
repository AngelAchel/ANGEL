package c2server

import (
	"time"
)

type C2ServerAgent0178 struct{}

func NewC2ServerAgent0178() *C2ServerAgent0178 {
	return &C2ServerAgent0178{}
}

func (e *C2ServerAgent0178) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0178) Name() string         { return "C2ServerAgent0178" }
func (e *C2ServerAgent0178) Timestamp() time.Time { return time.Now() }
