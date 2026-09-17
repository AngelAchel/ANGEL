package c2server

import (
	"time"
)

type C2ServerAgent0123 struct{}

func NewC2ServerAgent0123() *C2ServerAgent0123 {
	return &C2ServerAgent0123{}
}

func (e *C2ServerAgent0123) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0123) Name() string         { return "C2ServerAgent0123" }
func (e *C2ServerAgent0123) Timestamp() time.Time { return time.Now() }
