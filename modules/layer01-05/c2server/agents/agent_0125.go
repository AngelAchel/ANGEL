package c2server

import (
	"time"
)

type C2ServerAgent0125 struct{}

func NewC2ServerAgent0125() *C2ServerAgent0125 {
	return &C2ServerAgent0125{}
}

func (e *C2ServerAgent0125) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0125) Name() string         { return "C2ServerAgent0125" }
func (e *C2ServerAgent0125) Timestamp() time.Time { return time.Now() }
