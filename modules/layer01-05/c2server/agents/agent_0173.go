package c2server

import (
	"time"
)

type C2ServerAgent0173 struct{}

func NewC2ServerAgent0173() *C2ServerAgent0173 {
	return &C2ServerAgent0173{}
}

func (e *C2ServerAgent0173) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0173) Name() string         { return "C2ServerAgent0173" }
func (e *C2ServerAgent0173) Timestamp() time.Time { return time.Now() }
