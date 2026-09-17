package c2server

import (
	"time"
)

type C2ServerAgent0138 struct{}

func NewC2ServerAgent0138() *C2ServerAgent0138 {
	return &C2ServerAgent0138{}
}

func (e *C2ServerAgent0138) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0138) Name() string         { return "C2ServerAgent0138" }
func (e *C2ServerAgent0138) Timestamp() time.Time { return time.Now() }
