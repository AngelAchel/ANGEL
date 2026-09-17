package c2server

import (
	"time"
)

type C2ServerAgent0186 struct{}

func NewC2ServerAgent0186() *C2ServerAgent0186 {
	return &C2ServerAgent0186{}
}

func (e *C2ServerAgent0186) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0186) Name() string { return "C2ServerAgent0186" }
func (e *C2ServerAgent0186) Timestamp() time.Time { return time.Now() }
