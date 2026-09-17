package c2server

import (
	"time"
)

type C2ServerAgent0030 struct{}

func NewC2ServerAgent0030() *C2ServerAgent0030 {
	return &C2ServerAgent0030{}
}

func (e *C2ServerAgent0030) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0030) Name() string         { return "C2ServerAgent0030" }
func (e *C2ServerAgent0030) Timestamp() time.Time { return time.Now() }
