package c2server

import (
	"time"
)

type C2ServerAgent0082 struct{}

func NewC2ServerAgent0082() *C2ServerAgent0082 {
	return &C2ServerAgent0082{}
}

func (e *C2ServerAgent0082) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0082) Name() string { return "C2ServerAgent0082" }
func (e *C2ServerAgent0082) Timestamp() time.Time { return time.Now() }
