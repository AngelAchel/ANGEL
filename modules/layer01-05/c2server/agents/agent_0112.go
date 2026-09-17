package c2server

import (
	"time"
)

type C2ServerAgent0112 struct{}

func NewC2ServerAgent0112() *C2ServerAgent0112 {
	return &C2ServerAgent0112{}
}

func (e *C2ServerAgent0112) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0112) Name() string         { return "C2ServerAgent0112" }
func (e *C2ServerAgent0112) Timestamp() time.Time { return time.Now() }
