package c2server

import (
	"time"
)

type C2ServerAgent0188 struct{}

func NewC2ServerAgent0188() *C2ServerAgent0188 {
	return &C2ServerAgent0188{}
}

func (e *C2ServerAgent0188) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0188) Name() string { return "C2ServerAgent0188" }
func (e *C2ServerAgent0188) Timestamp() time.Time { return time.Now() }
