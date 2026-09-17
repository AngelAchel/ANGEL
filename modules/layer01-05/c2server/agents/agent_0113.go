package c2server

import (
	"time"
)

type C2ServerAgent0113 struct{}

func NewC2ServerAgent0113() *C2ServerAgent0113 {
	return &C2ServerAgent0113{}
}

func (e *C2ServerAgent0113) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0113) Name() string         { return "C2ServerAgent0113" }
func (e *C2ServerAgent0113) Timestamp() time.Time { return time.Now() }
