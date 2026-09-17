package c2server

import (
	"time"
)

type C2ServerAgent0103 struct{}

func NewC2ServerAgent0103() *C2ServerAgent0103 {
	return &C2ServerAgent0103{}
}

func (e *C2ServerAgent0103) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0103) Name() string { return "C2ServerAgent0103" }
func (e *C2ServerAgent0103) Timestamp() time.Time { return time.Now() }
