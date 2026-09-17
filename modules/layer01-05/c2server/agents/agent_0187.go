package c2server

import (
	"time"
)

type C2ServerAgent0187 struct{}

func NewC2ServerAgent0187() *C2ServerAgent0187 {
	return &C2ServerAgent0187{}
}

func (e *C2ServerAgent0187) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0187) Name() string         { return "C2ServerAgent0187" }
func (e *C2ServerAgent0187) Timestamp() time.Time { return time.Now() }
