package deser

import (
    "time"
)

type deser0098 struct{}

func Newdeser0098() *deser0098 {
    return &deser0098{}
}

func (e *deser0098) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "deser:done")
    return results, nil
}

func (e *deser0098) Name() string { return "deser0098" }
func (e *deser0098) Timestamp() time.Time { return time.Now() }
