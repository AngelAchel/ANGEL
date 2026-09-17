package c2server

import (
	"time"
)

type C2ServerAgent0117 struct{}

func NewC2ServerAgent0117() *C2ServerAgent0117 {
	return &C2ServerAgent0117{}
}

func (e *C2ServerAgent0117) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0117) Name() string         { return "C2ServerAgent0117" }
func (e *C2ServerAgent0117) Timestamp() time.Time { return time.Now() }
