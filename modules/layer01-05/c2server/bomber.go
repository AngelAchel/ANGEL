package c2server

import (
	"time"
)

type Bomber struct{}

func NewBomber() *Bomber {
	return &Bomber{}
}

func (e *Bomber) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "bomber:done")
	return results, nil
}

func (e *Bomber) Name() string { return "Bomber" }
func (e *Bomber) Timestamp() time.Time { return time.Now() }
