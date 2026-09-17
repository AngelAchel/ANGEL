package deser

import (
    "time"
)

type deser0083 struct{}

func Newdeser0083() *deser0083 {
    return &deser0083{}
}

func (e *deser0083) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "deser:done")
    return results, nil
}

func (e *deser0083) Name() string { return "deser0083" }
func (e *deser0083) Timestamp() time.Time { return time.Now() }
