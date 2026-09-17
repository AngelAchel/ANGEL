package deser

import (
    "time"
)

type deser0096 struct{}

func Newdeser0096() *deser0096 {
    return &deser0096{}
}

func (e *deser0096) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "deser:done")
    return results, nil
}

func (e *deser0096) Name() string { return "deser0096" }
func (e *deser0096) Timestamp() time.Time { return time.Now() }
