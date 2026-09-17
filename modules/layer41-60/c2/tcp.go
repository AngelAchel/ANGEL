package c2

import (
	"time"
)

type TCP struct{}

func NewTCP() *TCP {
	return &TCP{}
}

func (e *TCP) Connect(target string) ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "tcp:connected")
	return results, nil
}

func (e *TCP) Name() string { return "TCP" }
func (e *TCP) Category() C2Category { return CategoryC2 }
func (e *TCP) Timestamp() time.Time { return time.Now() }
