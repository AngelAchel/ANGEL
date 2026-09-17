package deser

import (
    "time"
)

type deser0017 struct{}

func Newdeser0017() *deser0017 {
    return &deser0017{}
}

func (e *deser0017) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "deser:done")
    return results, nil
}

func (e *deser0017) Name() string { return "deser0017" }
func (e *deser0017) Timestamp() time.Time { return time.Now() }
