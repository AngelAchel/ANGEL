package deser

import (
    "time"
)

type deser0088 struct{}

func Newdeser0088() *deser0088 {
    return &deser0088{}
}

func (e *deser0088) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "deser:done")
    return results, nil
}

func (e *deser0088) Name() string { return "deser0088" }
func (e *deser0088) Timestamp() time.Time { return time.Now() }
