package c2server

import (
	"time"
)

type C2ServerAgent0198 struct{}

func NewC2ServerAgent0198() *C2ServerAgent0198 {
	return &C2ServerAgent0198{}
}

func (e *C2ServerAgent0198) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0198) Name() string         { return "C2ServerAgent0198" }
func (e *C2ServerAgent0198) Timestamp() time.Time { return time.Now() }
