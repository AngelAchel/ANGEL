package c2server

import (
	"time"
)

type C2ServerAgent0012 struct{}

func NewC2ServerAgent0012() *C2ServerAgent0012 {
	return &C2ServerAgent0012{}
}

func (e *C2ServerAgent0012) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0012) Name() string { return "C2ServerAgent0012" }
func (e *C2ServerAgent0012) Timestamp() time.Time { return time.Now() }
