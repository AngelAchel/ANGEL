package c2server

import (
	"time"
)

type C2ServerAgent0174 struct{}

func NewC2ServerAgent0174() *C2ServerAgent0174 {
	return &C2ServerAgent0174{}
}

func (e *C2ServerAgent0174) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0174) Name() string         { return "C2ServerAgent0174" }
func (e *C2ServerAgent0174) Timestamp() time.Time { return time.Now() }
