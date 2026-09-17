package c2server

import (
	"time"
)

type C2ServerAgent0160 struct{}

func NewC2ServerAgent0160() *C2ServerAgent0160 {
	return &C2ServerAgent0160{}
}

func (e *C2ServerAgent0160) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0160) Name() string         { return "C2ServerAgent0160" }
func (e *C2ServerAgent0160) Timestamp() time.Time { return time.Now() }
