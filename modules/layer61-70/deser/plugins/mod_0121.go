package deser

import (
    "time"
)

type deser0121 struct{}

func Newdeser0121() *deser0121 {
    return &deser0121{}
}

func (e *deser0121) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "deser:done")
    return results, nil
}

func (e *deser0121) Name() string { return "deser0121" }
func (e *deser0121) Timestamp() time.Time { return time.Now() }
