package c2server

import (
	"time"
)

type C2ServerAgent0199 struct{}

func NewC2ServerAgent0199() *C2ServerAgent0199 {
	return &C2ServerAgent0199{}
}

func (e *C2ServerAgent0199) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0199) Name() string         { return "C2ServerAgent0199" }
func (e *C2ServerAgent0199) Timestamp() time.Time { return time.Now() }
