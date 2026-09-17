package c2server

import (
	"time"
)

type C2ServerAgent0003 struct{}

func NewC2ServerAgent0003() *C2ServerAgent0003 {
	return &C2ServerAgent0003{}
}

func (e *C2ServerAgent0003) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0003) Name() string { return "C2ServerAgent0003" }
func (e *C2ServerAgent0003) Timestamp() time.Time { return time.Now() }
