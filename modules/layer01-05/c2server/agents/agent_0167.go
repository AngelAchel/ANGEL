package c2server

import (
	"time"
)

type C2ServerAgent0167 struct{}

func NewC2ServerAgent0167() *C2ServerAgent0167 {
	return &C2ServerAgent0167{}
}

func (e *C2ServerAgent0167) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0167) Name() string         { return "C2ServerAgent0167" }
func (e *C2ServerAgent0167) Timestamp() time.Time { return time.Now() }
