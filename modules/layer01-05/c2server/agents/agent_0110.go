package c2server

import (
	"time"
)

type C2ServerAgent0110 struct{}

func NewC2ServerAgent0110() *C2ServerAgent0110 {
	return &C2ServerAgent0110{}
}

func (e *C2ServerAgent0110) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0110) Name() string         { return "C2ServerAgent0110" }
func (e *C2ServerAgent0110) Timestamp() time.Time { return time.Now() }
