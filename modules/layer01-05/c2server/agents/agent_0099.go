package c2server

import (
	"time"
)

type C2ServerAgent0099 struct{}

func NewC2ServerAgent0099() *C2ServerAgent0099 {
	return &C2ServerAgent0099{}
}

func (e *C2ServerAgent0099) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0099) Name() string { return "C2ServerAgent0099" }
func (e *C2ServerAgent0099) Timestamp() time.Time { return time.Now() }
