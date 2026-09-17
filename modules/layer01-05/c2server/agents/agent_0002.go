package c2server

import (
	"time"
)

type C2ServerAgent0002 struct{}

func NewC2ServerAgent0002() *C2ServerAgent0002 {
	return &C2ServerAgent0002{}
}

func (e *C2ServerAgent0002) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0002) Name() string { return "C2ServerAgent0002" }
func (e *C2ServerAgent0002) Timestamp() time.Time { return time.Now() }
