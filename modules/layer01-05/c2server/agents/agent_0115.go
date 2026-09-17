package c2server

import (
	"time"
)

type C2ServerAgent0115 struct{}

func NewC2ServerAgent0115() *C2ServerAgent0115 {
	return &C2ServerAgent0115{}
}

func (e *C2ServerAgent0115) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0115) Name() string         { return "C2ServerAgent0115" }
func (e *C2ServerAgent0115) Timestamp() time.Time { return time.Now() }
