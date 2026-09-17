package c2server

import (
	"time"
)

type C2ServerAgent0193 struct{}

func NewC2ServerAgent0193() *C2ServerAgent0193 {
	return &C2ServerAgent0193{}
}

func (e *C2ServerAgent0193) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0193) Name() string         { return "C2ServerAgent0193" }
func (e *C2ServerAgent0193) Timestamp() time.Time { return time.Now() }
