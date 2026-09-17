package c2server

import (
	"time"
)

type C2ServerAgent0197 struct{}

func NewC2ServerAgent0197() *C2ServerAgent0197 {
	return &C2ServerAgent0197{}
}

func (e *C2ServerAgent0197) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0197) Name() string         { return "C2ServerAgent0197" }
func (e *C2ServerAgent0197) Timestamp() time.Time { return time.Now() }
