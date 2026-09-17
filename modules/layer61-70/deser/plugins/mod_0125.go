package deser

import (
    "time"
)

type deser0125 struct{}

func Newdeser0125() *deser0125 {
    return &deser0125{}
}

func (e *deser0125) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "deser:done")
    return results, nil
}

func (e *deser0125) Name() string { return "deser0125" }
func (e *deser0125) Timestamp() time.Time { return time.Now() }
