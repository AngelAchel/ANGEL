package c2server

import (
	"time"
)

type C2ServerAgent0000 struct{}

func NewC2ServerAgent0000() *C2ServerAgent0000 {
	return &C2ServerAgent0000{}
}

func (e *C2ServerAgent0000) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0000) Name() string { return "C2ServerAgent0000" }
func (e *C2ServerAgent0000) Timestamp() time.Time { return time.Now() }
