package c2server

import (
	"time"
)

type C2ServerAgent0189 struct{}

func NewC2ServerAgent0189() *C2ServerAgent0189 {
	return &C2ServerAgent0189{}
}

func (e *C2ServerAgent0189) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0189) Name() string         { return "C2ServerAgent0189" }
func (e *C2ServerAgent0189) Timestamp() time.Time { return time.Now() }
