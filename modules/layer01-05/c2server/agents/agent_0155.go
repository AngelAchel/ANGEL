package c2server

import (
	"time"
)

type C2ServerAgent0155 struct{}

func NewC2ServerAgent0155() *C2ServerAgent0155 {
	return &C2ServerAgent0155{}
}

func (e *C2ServerAgent0155) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0155) Name() string         { return "C2ServerAgent0155" }
func (e *C2ServerAgent0155) Timestamp() time.Time { return time.Now() }
