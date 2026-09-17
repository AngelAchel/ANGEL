package c2server

import (
	"time"
)

type C2ServerAgent0005 struct{}

func NewC2ServerAgent0005() *C2ServerAgent0005 {
	return &C2ServerAgent0005{}
}

func (e *C2ServerAgent0005) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0005) Name() string { return "C2ServerAgent0005" }
func (e *C2ServerAgent0005) Timestamp() time.Time { return time.Now() }
