package c2

import (
	"time"
)

type TCP struct{}

func NewTCP() *TCP {
	return &TCP{}
}

func (t *TCP) Connect() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "tcp:done")
	return results, nil
}

func (t *TCP) Name() string         { return "TCP" }
func (t *TCP) Timestamp() time.Time { return time.Now() }
