package deser

import (
    "time"
)

type deser0001 struct{}

func Newdeser0001() *deser0001 {
    return &deser0001{}
}

func (e *deser0001) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "deser:done")
    return results, nil
}

func (e *deser0001) Name() string { return "deser0001" }
func (e *deser0001) Timestamp() time.Time { return time.Now() }
