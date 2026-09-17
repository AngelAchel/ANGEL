package c2server

import (
	"time"
)

type C2ServerAgent0107 struct{}

func NewC2ServerAgent0107() *C2ServerAgent0107 {
	return &C2ServerAgent0107{}
}

func (e *C2ServerAgent0107) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0107) Name() string { return "C2ServerAgent0107" }
func (e *C2ServerAgent0107) Timestamp() time.Time { return time.Now() }
