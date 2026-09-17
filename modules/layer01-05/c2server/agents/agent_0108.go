package c2server

import (
	"time"
)

type C2ServerAgent0108 struct{}

func NewC2ServerAgent0108() *C2ServerAgent0108 {
	return &C2ServerAgent0108{}
}

func (e *C2ServerAgent0108) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0108) Name() string         { return "C2ServerAgent0108" }
func (e *C2ServerAgent0108) Timestamp() time.Time { return time.Now() }
