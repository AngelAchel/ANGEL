package deser

import (
    "time"
)

type deser0012 struct{}

func Newdeser0012() *deser0012 {
    return &deser0012{}
}

func (e *deser0012) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "deser:done")
    return results, nil
}

func (e *deser0012) Name() string { return "deser0012" }
func (e *deser0012) Timestamp() time.Time { return time.Now() }
