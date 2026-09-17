package deser

import (
    "time"
)

type deser0100 struct{}

func Newdeser0100() *deser0100 {
    return &deser0100{}
}

func (e *deser0100) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "deser:done")
    return results, nil
}

func (e *deser0100) Name() string { return "deser0100" }
func (e *deser0100) Timestamp() time.Time { return time.Now() }
