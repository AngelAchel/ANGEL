package c2server

import (
	"time"
)

type C2ServerAgent0177 struct{}

func NewC2ServerAgent0177() *C2ServerAgent0177 {
	return &C2ServerAgent0177{}
}

func (e *C2ServerAgent0177) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0177) Name() string { return "C2ServerAgent0177" }
func (e *C2ServerAgent0177) Timestamp() time.Time { return time.Now() }
