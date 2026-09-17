package c2server

import (
	"time"
)

type C2ServerAgent0100 struct{}

func NewC2ServerAgent0100() *C2ServerAgent0100 {
	return &C2ServerAgent0100{}
}

func (e *C2ServerAgent0100) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0100) Name() string { return "C2ServerAgent0100" }
func (e *C2ServerAgent0100) Timestamp() time.Time { return time.Now() }
