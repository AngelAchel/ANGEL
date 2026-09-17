package c2server

import (
	"time"
)

type C2ServerAgent0101 struct{}

func NewC2ServerAgent0101() *C2ServerAgent0101 {
	return &C2ServerAgent0101{}
}

func (e *C2ServerAgent0101) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0101) Name() string         { return "C2ServerAgent0101" }
func (e *C2ServerAgent0101) Timestamp() time.Time { return time.Now() }
