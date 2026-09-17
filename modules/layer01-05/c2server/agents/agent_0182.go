package c2server

import (
	"time"
)

type C2ServerAgent0182 struct{}

func NewC2ServerAgent0182() *C2ServerAgent0182 {
	return &C2ServerAgent0182{}
}

func (e *C2ServerAgent0182) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0182) Name() string { return "C2ServerAgent0182" }
func (e *C2ServerAgent0182) Timestamp() time.Time { return time.Now() }
