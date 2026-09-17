package c2server

import (
	"time"
)

type C2ServerAgent0166 struct{}

func NewC2ServerAgent0166() *C2ServerAgent0166 {
	return &C2ServerAgent0166{}
}

func (e *C2ServerAgent0166) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0166) Name() string         { return "C2ServerAgent0166" }
func (e *C2ServerAgent0166) Timestamp() time.Time { return time.Now() }
