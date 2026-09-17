package c2server

import (
	"time"
)

type C2ServerAgent0171 struct{}

func NewC2ServerAgent0171() *C2ServerAgent0171 {
	return &C2ServerAgent0171{}
}

func (e *C2ServerAgent0171) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0171) Name() string { return "C2ServerAgent0171" }
func (e *C2ServerAgent0171) Timestamp() time.Time { return time.Now() }
