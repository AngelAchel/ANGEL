package c2server

import (
	"time"
)

type C2ServerAgent0046 struct{}

func NewC2ServerAgent0046() *C2ServerAgent0046 {
	return &C2ServerAgent0046{}
}

func (e *C2ServerAgent0046) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0046) Name() string { return "C2ServerAgent0046" }
func (e *C2ServerAgent0046) Timestamp() time.Time { return time.Now() }
