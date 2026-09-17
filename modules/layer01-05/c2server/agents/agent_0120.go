package c2server

import (
	"time"
)

type C2ServerAgent0120 struct{}

func NewC2ServerAgent0120() *C2ServerAgent0120 {
	return &C2ServerAgent0120{}
}

func (e *C2ServerAgent0120) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0120) Name() string         { return "C2ServerAgent0120" }
func (e *C2ServerAgent0120) Timestamp() time.Time { return time.Now() }
