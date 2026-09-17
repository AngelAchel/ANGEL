package c2server

import (
	"time"
)

type C2ServerAgent0033 struct{}

func NewC2ServerAgent0033() *C2ServerAgent0033 {
	return &C2ServerAgent0033{}
}

func (e *C2ServerAgent0033) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0033) Name() string         { return "C2ServerAgent0033" }
func (e *C2ServerAgent0033) Timestamp() time.Time { return time.Now() }
