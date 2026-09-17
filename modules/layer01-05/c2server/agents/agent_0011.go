package c2server

import (
	"time"
)

type C2ServerAgent0011 struct{}

func NewC2ServerAgent0011() *C2ServerAgent0011 {
	return &C2ServerAgent0011{}
}

func (e *C2ServerAgent0011) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0011) Name() string         { return "C2ServerAgent0011" }
func (e *C2ServerAgent0011) Timestamp() time.Time { return time.Now() }
