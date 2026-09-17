package c2server

import (
	"time"
)

type C2ServerAgent0006 struct{}

func NewC2ServerAgent0006() *C2ServerAgent0006 {
	return &C2ServerAgent0006{}
}

func (e *C2ServerAgent0006) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0006) Name() string         { return "C2ServerAgent0006" }
func (e *C2ServerAgent0006) Timestamp() time.Time { return time.Now() }
