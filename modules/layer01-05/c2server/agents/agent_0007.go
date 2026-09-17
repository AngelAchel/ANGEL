package c2server

import (
	"time"
)

type C2ServerAgent0007 struct{}

func NewC2ServerAgent0007() *C2ServerAgent0007 {
	return &C2ServerAgent0007{}
}

func (e *C2ServerAgent0007) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0007) Name() string         { return "C2ServerAgent0007" }
func (e *C2ServerAgent0007) Timestamp() time.Time { return time.Now() }
