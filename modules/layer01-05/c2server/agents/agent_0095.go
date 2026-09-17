package c2server

import (
	"time"
)

type C2ServerAgent0095 struct{}

func NewC2ServerAgent0095() *C2ServerAgent0095 {
	return &C2ServerAgent0095{}
}

func (e *C2ServerAgent0095) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0095) Name() string         { return "C2ServerAgent0095" }
func (e *C2ServerAgent0095) Timestamp() time.Time { return time.Now() }
