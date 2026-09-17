package deser

import (
    "time"
)

type deser0081 struct{}

func Newdeser0081() *deser0081 {
    return &deser0081{}
}

func (e *deser0081) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "deser:done")
    return results, nil
}

func (e *deser0081) Name() string { return "deser0081" }
func (e *deser0081) Timestamp() time.Time { return time.Now() }
