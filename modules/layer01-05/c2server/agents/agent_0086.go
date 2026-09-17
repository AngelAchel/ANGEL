package c2server

import (
	"time"
)

type C2ServerAgent0086 struct{}

func NewC2ServerAgent0086() *C2ServerAgent0086 {
	return &C2ServerAgent0086{}
}

func (e *C2ServerAgent0086) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0086) Name() string         { return "C2ServerAgent0086" }
func (e *C2ServerAgent0086) Timestamp() time.Time { return time.Now() }
