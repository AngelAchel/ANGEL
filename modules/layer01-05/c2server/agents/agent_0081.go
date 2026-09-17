package c2server

import (
	"time"
)

type C2ServerAgent0081 struct{}

func NewC2ServerAgent0081() *C2ServerAgent0081 {
	return &C2ServerAgent0081{}
}

func (e *C2ServerAgent0081) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0081) Name() string         { return "C2ServerAgent0081" }
func (e *C2ServerAgent0081) Timestamp() time.Time { return time.Now() }
