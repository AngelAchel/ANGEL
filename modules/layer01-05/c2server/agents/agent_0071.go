package c2server

import (
	"time"
)

type C2ServerAgent0071 struct{}

func NewC2ServerAgent0071() *C2ServerAgent0071 {
	return &C2ServerAgent0071{}
}

func (e *C2ServerAgent0071) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0071) Name() string         { return "C2ServerAgent0071" }
func (e *C2ServerAgent0071) Timestamp() time.Time { return time.Now() }
