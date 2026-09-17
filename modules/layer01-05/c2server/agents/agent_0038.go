package c2server

import (
	"time"
)

type C2ServerAgent0038 struct{}

func NewC2ServerAgent0038() *C2ServerAgent0038 {
	return &C2ServerAgent0038{}
}

func (e *C2ServerAgent0038) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0038) Name() string         { return "C2ServerAgent0038" }
func (e *C2ServerAgent0038) Timestamp() time.Time { return time.Now() }
