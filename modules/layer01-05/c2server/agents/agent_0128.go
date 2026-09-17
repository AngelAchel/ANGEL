package c2server

import (
	"time"
)

type C2ServerAgent0128 struct{}

func NewC2ServerAgent0128() *C2ServerAgent0128 {
	return &C2ServerAgent0128{}
}

func (e *C2ServerAgent0128) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0128) Name() string         { return "C2ServerAgent0128" }
func (e *C2ServerAgent0128) Timestamp() time.Time { return time.Now() }
