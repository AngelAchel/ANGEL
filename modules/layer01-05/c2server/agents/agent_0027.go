package c2server

import (
	"time"
)

type C2ServerAgent0027 struct{}

func NewC2ServerAgent0027() *C2ServerAgent0027 {
	return &C2ServerAgent0027{}
}

func (e *C2ServerAgent0027) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0027) Name() string         { return "C2ServerAgent0027" }
func (e *C2ServerAgent0027) Timestamp() time.Time { return time.Now() }
