package sqlinject

import (
	"time"
)

type TimeBased struct{}

func NewTimeBased() *TimeBased {
	return &TimeBased{}
}

func (t *TimeBased) Inject() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "time_based:injected")
	return results, nil
}

func (t *TimeBased) Name() string { return "TimeBased" }
func (t *TimeBased) Timestamp() time.Time { return time.Now() }
