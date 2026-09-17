package c2server

import (
	"time"
)

type C2ServerAgent0001 struct{}

func NewC2ServerAgent0001() *C2ServerAgent0001 {
	return &C2ServerAgent0001{}
}

func (e *C2ServerAgent0001) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0001) Name() string { return "C2ServerAgent0001" }
func (e *C2ServerAgent0001) Timestamp() time.Time { return time.Now() }
