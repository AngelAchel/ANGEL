package c2server

import (
	"time"
)

type C2ServerAgent0008 struct{}

func NewC2ServerAgent0008() *C2ServerAgent0008 {
	return &C2ServerAgent0008{}
}

func (e *C2ServerAgent0008) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0008) Name() string { return "C2ServerAgent0008" }
func (e *C2ServerAgent0008) Timestamp() time.Time { return time.Now() }
