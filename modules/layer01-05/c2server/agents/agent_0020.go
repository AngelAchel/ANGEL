package c2server

import (
	"time"
)

type C2ServerAgent0020 struct{}

func NewC2ServerAgent0020() *C2ServerAgent0020 {
	return &C2ServerAgent0020{}
}

func (e *C2ServerAgent0020) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0020) Name() string { return "C2ServerAgent0020" }
func (e *C2ServerAgent0020) Timestamp() time.Time { return time.Now() }
