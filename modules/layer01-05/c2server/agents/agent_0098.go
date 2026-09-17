package c2server

import (
	"time"
)

type C2ServerAgent0098 struct{}

func NewC2ServerAgent0098() *C2ServerAgent0098 {
	return &C2ServerAgent0098{}
}

func (e *C2ServerAgent0098) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0098) Name() string         { return "C2ServerAgent0098" }
func (e *C2ServerAgent0098) Timestamp() time.Time { return time.Now() }
