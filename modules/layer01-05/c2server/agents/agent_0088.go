package c2server

import (
	"time"
)

type C2ServerAgent0088 struct{}

func NewC2ServerAgent0088() *C2ServerAgent0088 {
	return &C2ServerAgent0088{}
}

func (e *C2ServerAgent0088) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0088) Name() string { return "C2ServerAgent0088" }
func (e *C2ServerAgent0088) Timestamp() time.Time { return time.Now() }
