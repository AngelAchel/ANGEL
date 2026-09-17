package deser

import (
    "time"
)

type deser0123 struct{}

func Newdeser0123() *deser0123 {
    return &deser0123{}
}

func (e *deser0123) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "deser:done")
    return results, nil
}

func (e *deser0123) Name() string { return "deser0123" }
func (e *deser0123) Timestamp() time.Time { return time.Now() }
