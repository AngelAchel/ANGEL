package deser

import (
    "time"
)

type deser0133 struct{}

func Newdeser0133() *deser0133 {
    return &deser0133{}
}

func (e *deser0133) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "deser:done")
    return results, nil
}

func (e *deser0133) Name() string { return "deser0133" }
func (e *deser0133) Timestamp() time.Time { return time.Now() }
