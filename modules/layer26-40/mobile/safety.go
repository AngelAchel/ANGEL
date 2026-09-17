package mobile

import (
	"time"
)

type Safety struct{}

func NewSafety() *Safety {
	return &Safety{}
}

func (s *Safety) Check() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "safety:done")
	return results, nil
}

func (s *Safety) Name() string { return "Safety" }
func (s *Safety) Timestamp() time.Time { return time.Now() }
