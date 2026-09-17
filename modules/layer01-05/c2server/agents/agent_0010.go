package c2server

import (
	"time"
)

type C2ServerAgent0010 struct{}

func NewC2ServerAgent0010() *C2ServerAgent0010 {
	return &C2ServerAgent0010{}
}

func (e *C2ServerAgent0010) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0010) Name() string { return "C2ServerAgent0010" }
func (e *C2ServerAgent0010) Timestamp() time.Time { return time.Now() }
