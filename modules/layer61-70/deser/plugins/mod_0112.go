package deser

import (
    "time"
)

type deser0112 struct{}

func Newdeser0112() *deser0112 {
    return &deser0112{}
}

func (e *deser0112) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "deser:done")
    return results, nil
}

func (e *deser0112) Name() string { return "deser0112" }
func (e *deser0112) Timestamp() time.Time { return time.Now() }
