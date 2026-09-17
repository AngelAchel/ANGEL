package c2server

import (
	"time"
)

type C2ServerAgent0133 struct{}

func NewC2ServerAgent0133() *C2ServerAgent0133 {
	return &C2ServerAgent0133{}
}

func (e *C2ServerAgent0133) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0133) Name() string         { return "C2ServerAgent0133" }
func (e *C2ServerAgent0133) Timestamp() time.Time { return time.Now() }
