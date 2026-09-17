package c2server

import (
	"time"
)

type C2ServerAgent0121 struct{}

func NewC2ServerAgent0121() *C2ServerAgent0121 {
	return &C2ServerAgent0121{}
}

func (e *C2ServerAgent0121) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0121) Name() string { return "C2ServerAgent0121" }
func (e *C2ServerAgent0121) Timestamp() time.Time { return time.Now() }
