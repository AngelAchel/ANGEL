package c2server

import (
	"time"
)

type C2ServerAgent0105 struct{}

func NewC2ServerAgent0105() *C2ServerAgent0105 {
	return &C2ServerAgent0105{}
}

func (e *C2ServerAgent0105) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0105) Name() string         { return "C2ServerAgent0105" }
func (e *C2ServerAgent0105) Timestamp() time.Time { return time.Now() }
