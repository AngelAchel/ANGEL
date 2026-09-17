package c2server

import (
	"time"
)

type C2ServerAgent0146 struct{}

func NewC2ServerAgent0146() *C2ServerAgent0146 {
	return &C2ServerAgent0146{}
}

func (e *C2ServerAgent0146) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0146) Name() string         { return "C2ServerAgent0146" }
func (e *C2ServerAgent0146) Timestamp() time.Time { return time.Now() }
