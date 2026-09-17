package deser

import (
    "time"
)

type deser0101 struct{}

func Newdeser0101() *deser0101 {
    return &deser0101{}
}

func (e *deser0101) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "deser:done")
    return results, nil
}

func (e *deser0101) Name() string { return "deser0101" }
func (e *deser0101) Timestamp() time.Time { return time.Now() }
