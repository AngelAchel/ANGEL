package c2server

import (
	"time"
)

type C2ServerAgent0181 struct{}

func NewC2ServerAgent0181() *C2ServerAgent0181 {
	return &C2ServerAgent0181{}
}

func (e *C2ServerAgent0181) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "c2server:done")
	return results, nil
}

func (e *C2ServerAgent0181) Name() string         { return "C2ServerAgent0181" }
func (e *C2ServerAgent0181) Timestamp() time.Time { return time.Now() }
