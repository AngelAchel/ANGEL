package c2server

import (
	"time"
)

type C2ServerAgent0084 struct{}

func NewC2ServerAgent0084() *C2ServerAgent0084 {
	return &C2ServerAgent0084{}
}

func (e *C2ServerAgent0084) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0084) Name() string         { return "C2ServerAgent0084" }
func (e *C2ServerAgent0084) Timestamp() time.Time { return time.Now() }
