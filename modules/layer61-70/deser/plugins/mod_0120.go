package deser

import (
    "time"
)

type deser0120 struct{}

func Newdeser0120() *deser0120 {
    return &deser0120{}
}

func (e *deser0120) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "deser:done")
    return results, nil
}

func (e *deser0120) Name() string { return "deser0120" }
func (e *deser0120) Timestamp() time.Time { return time.Now() }
