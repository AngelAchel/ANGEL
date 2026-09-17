package c2server

import (
	"time"
)

type C2ServerAgent0060 struct{}

func NewC2ServerAgent0060() *C2ServerAgent0060 {
	return &C2ServerAgent0060{}
}

func (e *C2ServerAgent0060) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0060) Name() string         { return "C2ServerAgent0060" }
func (e *C2ServerAgent0060) Timestamp() time.Time { return time.Now() }
