package deser

import (
    "time"
)

type deser0117 struct{}

func Newdeser0117() *deser0117 {
    return &deser0117{}
}

func (e *deser0117) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "deser:done")
    return results, nil
}

func (e *deser0117) Name() string { return "deser0117" }
func (e *deser0117) Timestamp() time.Time { return time.Now() }
