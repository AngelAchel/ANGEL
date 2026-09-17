package c2server

import (
	"time"
)

type C2ServerAgent0122 struct{}

func NewC2ServerAgent0122() *C2ServerAgent0122 {
	return &C2ServerAgent0122{}
}

func (e *C2ServerAgent0122) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0122) Name() string         { return "C2ServerAgent0122" }
func (e *C2ServerAgent0122) Timestamp() time.Time { return time.Now() }
