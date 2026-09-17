package c2server

import (
	"time"
)

type C2ServerAgent0018 struct{}

func NewC2ServerAgent0018() *C2ServerAgent0018 {
	return &C2ServerAgent0018{}
}

func (e *C2ServerAgent0018) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0018) Name() string { return "C2ServerAgent0018" }
func (e *C2ServerAgent0018) Timestamp() time.Time { return time.Now() }
