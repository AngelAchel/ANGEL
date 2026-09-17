package c2server

import (
	"time"
)

type C2ServerAgent0137 struct{}

func NewC2ServerAgent0137() *C2ServerAgent0137 {
	return &C2ServerAgent0137{}
}

func (e *C2ServerAgent0137) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0137) Name() string         { return "C2ServerAgent0137" }
func (e *C2ServerAgent0137) Timestamp() time.Time { return time.Now() }
