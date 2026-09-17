package c2server

import (
	"time"
)

type C2ServerAgent0106 struct{}

func NewC2ServerAgent0106() *C2ServerAgent0106 {
	return &C2ServerAgent0106{}
}

func (e *C2ServerAgent0106) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0106) Name() string         { return "C2ServerAgent0106" }
func (e *C2ServerAgent0106) Timestamp() time.Time { return time.Now() }
