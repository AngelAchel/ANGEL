package racecond

import (
    "time"
)

type racecond0123 struct{}

func Newracecond0123() *racecond0123 {
    return &racecond0123{}
}

func (e *racecond0123) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "racecond:done")
    return results, nil
}

func (e *racecond0123) Name() string { return "racecond0123" }
func (e *racecond0123) Timestamp() time.Time { return time.Now() }
