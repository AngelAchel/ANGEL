package deser

import (
    "time"
)

type deser0080 struct{}

func Newdeser0080() *deser0080 {
    return &deser0080{}
}

func (e *deser0080) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "deser:done")
    return results, nil
}

func (e *deser0080) Name() string { return "deser0080" }
func (e *deser0080) Timestamp() time.Time { return time.Now() }
