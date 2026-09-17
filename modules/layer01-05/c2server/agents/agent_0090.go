package c2server

import (
	"time"
)

type C2ServerAgent0090 struct{}

func NewC2ServerAgent0090() *C2ServerAgent0090 {
	return &C2ServerAgent0090{}
}

func (e *C2ServerAgent0090) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0090) Name() string         { return "C2ServerAgent0090" }
func (e *C2ServerAgent0090) Timestamp() time.Time { return time.Now() }
