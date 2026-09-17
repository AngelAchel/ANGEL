package c2server

import (
	"time"
)

type C2ServerAgent0196 struct{}

func NewC2ServerAgent0196() *C2ServerAgent0196 {
	return &C2ServerAgent0196{}
}

func (e *C2ServerAgent0196) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0196) Name() string         { return "C2ServerAgent0196" }
func (e *C2ServerAgent0196) Timestamp() time.Time { return time.Now() }
