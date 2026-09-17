package c2server

import (
	"time"
)

type C2ServerAgent0096 struct{}

func NewC2ServerAgent0096() *C2ServerAgent0096 {
	return &C2ServerAgent0096{}
}

func (e *C2ServerAgent0096) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0096) Name() string         { return "C2ServerAgent0096" }
func (e *C2ServerAgent0096) Timestamp() time.Time { return time.Now() }
