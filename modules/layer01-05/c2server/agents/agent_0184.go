package c2server

import (
	"time"
)

type C2ServerAgent0184 struct{}

func NewC2ServerAgent0184() *C2ServerAgent0184 {
	return &C2ServerAgent0184{}
}

func (e *C2ServerAgent0184) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0184) Name() string         { return "C2ServerAgent0184" }
func (e *C2ServerAgent0184) Timestamp() time.Time { return time.Now() }
