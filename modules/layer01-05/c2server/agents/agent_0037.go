package c2server

import (
	"time"
)

type C2ServerAgent0037 struct{}

func NewC2ServerAgent0037() *C2ServerAgent0037 {
	return &C2ServerAgent0037{}
}

func (e *C2ServerAgent0037) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0037) Name() string         { return "C2ServerAgent0037" }
func (e *C2ServerAgent0037) Timestamp() time.Time { return time.Now() }
