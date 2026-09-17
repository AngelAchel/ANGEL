package deser

import (
    "time"
)

type deser0108 struct{}

func Newdeser0108() *deser0108 {
    return &deser0108{}
}

func (e *deser0108) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "deser:done")
    return results, nil
}

func (e *deser0108) Name() string { return "deser0108" }
func (e *deser0108) Timestamp() time.Time { return time.Now() }
