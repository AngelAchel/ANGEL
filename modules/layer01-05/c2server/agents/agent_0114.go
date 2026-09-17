package c2server

import (
	"time"
)

type C2ServerAgent0114 struct{}

func NewC2ServerAgent0114() *C2ServerAgent0114 {
	return &C2ServerAgent0114{}
}

func (e *C2ServerAgent0114) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0114) Name() string { return "C2ServerAgent0114" }
func (e *C2ServerAgent0114) Timestamp() time.Time { return time.Now() }
