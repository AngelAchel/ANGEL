package c2server

import (
	"time"
)

type C2ServerAgent0190 struct{}

func NewC2ServerAgent0190() *C2ServerAgent0190 {
	return &C2ServerAgent0190{}
}

func (e *C2ServerAgent0190) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0190) Name() string { return "C2ServerAgent0190" }
func (e *C2ServerAgent0190) Timestamp() time.Time { return time.Now() }
