package deser

import (
    "time"
)

type deser0127 struct{}

func Newdeser0127() *deser0127 {
    return &deser0127{}
}

func (e *deser0127) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "deser:done")
    return results, nil
}

func (e *deser0127) Name() string { return "deser0127" }
func (e *deser0127) Timestamp() time.Time { return time.Now() }
