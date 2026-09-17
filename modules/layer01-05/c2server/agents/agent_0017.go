package c2server

import (
	"time"
)

type C2ServerAgent0017 struct{}

func NewC2ServerAgent0017() *C2ServerAgent0017 {
	return &C2ServerAgent0017{}
}

func (e *C2ServerAgent0017) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0017) Name() string         { return "C2ServerAgent0017" }
func (e *C2ServerAgent0017) Timestamp() time.Time { return time.Now() }
