package c2server

import (
	"time"
)

type C2ServerAgent0102 struct{}

func NewC2ServerAgent0102() *C2ServerAgent0102 {
	return &C2ServerAgent0102{}
}

func (e *C2ServerAgent0102) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0102) Name() string         { return "C2ServerAgent0102" }
func (e *C2ServerAgent0102) Timestamp() time.Time { return time.Now() }
