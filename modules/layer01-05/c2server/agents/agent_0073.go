package c2server

import (
	"time"
)

type C2ServerAgent0073 struct{}

func NewC2ServerAgent0073() *C2ServerAgent0073 {
	return &C2ServerAgent0073{}
}

func (e *C2ServerAgent0073) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0073) Name() string { return "C2ServerAgent0073" }
func (e *C2ServerAgent0073) Timestamp() time.Time { return time.Now() }
