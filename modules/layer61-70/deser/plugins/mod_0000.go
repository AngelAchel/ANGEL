package deser

import (
    "time"
)

type deser0000 struct{}

func Newdeser0000() *deser0000 {
    return &deser0000{}
}

func (e *deser0000) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "deser:done")
    return results, nil
}

func (e *deser0000) Name() string { return "deser0000" }
func (e *deser0000) Timestamp() time.Time { return time.Now() }
