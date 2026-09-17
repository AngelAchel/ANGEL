package c2server

import (
	"time"
)

type C2ServerAgent0180 struct{}

func NewC2ServerAgent0180() *C2ServerAgent0180 {
	return &C2ServerAgent0180{}
}

func (e *C2ServerAgent0180) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0180) Name() string         { return "C2ServerAgent0180" }
func (e *C2ServerAgent0180) Timestamp() time.Time { return time.Now() }
