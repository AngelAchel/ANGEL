package c2server

import (
	"time"
)

type C2ServerAgent0085 struct{}

func NewC2ServerAgent0085() *C2ServerAgent0085 {
	return &C2ServerAgent0085{}
}

func (e *C2ServerAgent0085) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0085) Name() string         { return "C2ServerAgent0085" }
func (e *C2ServerAgent0085) Timestamp() time.Time { return time.Now() }
