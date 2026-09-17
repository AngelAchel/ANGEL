package deser

import (
    "time"
)

type deser0102 struct{}

func Newdeser0102() *deser0102 {
    return &deser0102{}
}

func (e *deser0102) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "deser:done")
    return results, nil
}

func (e *deser0102) Name() string { return "deser0102" }
func (e *deser0102) Timestamp() time.Time { return time.Now() }
