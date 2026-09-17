package c2server

import (
	"time"
)

type C2ServerAgent0021 struct{}

func NewC2ServerAgent0021() *C2ServerAgent0021 {
	return &C2ServerAgent0021{}
}

func (e *C2ServerAgent0021) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0021) Name() string { return "C2ServerAgent0021" }
func (e *C2ServerAgent0021) Timestamp() time.Time { return time.Now() }
