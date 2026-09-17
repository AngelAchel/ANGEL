package c2server

import (
	"time"
)

type C2ServerAgent0028 struct{}

func NewC2ServerAgent0028() *C2ServerAgent0028 {
	return &C2ServerAgent0028{}
}

func (e *C2ServerAgent0028) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0028) Name() string         { return "C2ServerAgent0028" }
func (e *C2ServerAgent0028) Timestamp() time.Time { return time.Now() }
