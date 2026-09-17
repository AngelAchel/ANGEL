package c2server

import (
	"time"
)

type C2ServerAgent0087 struct{}

func NewC2ServerAgent0087() *C2ServerAgent0087 {
	return &C2ServerAgent0087{}
}

func (e *C2ServerAgent0087) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0087) Name() string         { return "C2ServerAgent0087" }
func (e *C2ServerAgent0087) Timestamp() time.Time { return time.Now() }
