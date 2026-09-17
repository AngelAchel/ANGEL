package deser

import (
    "time"
)

type deser0122 struct{}

func Newdeser0122() *deser0122 {
    return &deser0122{}
}

func (e *deser0122) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "deser:done")
    return results, nil
}

func (e *deser0122) Name() string { return "deser0122" }
func (e *deser0122) Timestamp() time.Time { return time.Now() }
