package sqlinject

import (
	"time"
)

type Stacked struct{}

func NewStacked() *Stacked {
	return &Stacked{}
}

func (s *Stacked) Inject() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "stacked:done")
	return results, nil
}

func (s *Stacked) Name() string         { return "Stacked" }
func (s *Stacked) Timestamp() time.Time { return time.Now() }
