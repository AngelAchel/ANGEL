package c2server

import (
	"time"
)

type C2ServerAgent0062 struct{}

func NewC2ServerAgent0062() *C2ServerAgent0062 {
	return &C2ServerAgent0062{}
}

func (e *C2ServerAgent0062) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0062) Name() string         { return "C2ServerAgent0062" }
func (e *C2ServerAgent0062) Timestamp() time.Time { return time.Now() }
