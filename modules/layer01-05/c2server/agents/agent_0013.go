package c2server

import (
	"time"
)

type C2ServerAgent0013 struct{}

func NewC2ServerAgent0013() *C2ServerAgent0013 {
	return &C2ServerAgent0013{}
}

func (e *C2ServerAgent0013) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0013) Name() string { return "C2ServerAgent0013" }
func (e *C2ServerAgent0013) Timestamp() time.Time { return time.Now() }
