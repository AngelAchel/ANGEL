package c2server

import (
	"time"
)

type C2ServerAgent0130 struct{}

func NewC2ServerAgent0130() *C2ServerAgent0130 {
	return &C2ServerAgent0130{}
}

func (e *C2ServerAgent0130) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0130) Name() string         { return "C2ServerAgent0130" }
func (e *C2ServerAgent0130) Timestamp() time.Time { return time.Now() }
