package c2server

import (
	"time"
)

type C2ServerAgent0172 struct{}

func NewC2ServerAgent0172() *C2ServerAgent0172 {
	return &C2ServerAgent0172{}
}

func (e *C2ServerAgent0172) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0172) Name() string         { return "C2ServerAgent0172" }
func (e *C2ServerAgent0172) Timestamp() time.Time { return time.Now() }
