package racecond

import (
    "time"
)

type racecond0120 struct{}

func Newracecond0120() *racecond0120 {
    return &racecond0120{}
}

func (e *racecond0120) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "racecond:done")
    return results, nil
}

func (e *racecond0120) Name() string { return "racecond0120" }
func (e *racecond0120) Timestamp() time.Time { return time.Now() }
