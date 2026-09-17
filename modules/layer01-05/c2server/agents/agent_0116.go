package c2server

import (
	"time"
)

type C2ServerAgent0116 struct{}

func NewC2ServerAgent0116() *C2ServerAgent0116 {
	return &C2ServerAgent0116{}
}

func (e *C2ServerAgent0116) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0116) Name() string         { return "C2ServerAgent0116" }
func (e *C2ServerAgent0116) Timestamp() time.Time { return time.Now() }
