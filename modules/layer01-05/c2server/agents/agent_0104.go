package c2server

import (
	"time"
)

type C2ServerAgent0104 struct{}

func NewC2ServerAgent0104() *C2ServerAgent0104 {
	return &C2ServerAgent0104{}
}

func (e *C2ServerAgent0104) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0104) Name() string { return "C2ServerAgent0104" }
func (e *C2ServerAgent0104) Timestamp() time.Time { return time.Now() }
