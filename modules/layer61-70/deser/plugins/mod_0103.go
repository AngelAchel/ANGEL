package deser

import (
    "time"
)

type deser0103 struct{}

func Newdeser0103() *deser0103 {
    return &deser0103{}
}

func (e *deser0103) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "deser:done")
    return results, nil
}

func (e *deser0103) Name() string { return "deser0103" }
func (e *deser0103) Timestamp() time.Time { return time.Now() }
