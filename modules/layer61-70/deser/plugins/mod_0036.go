package deser

import (
    "time"
)

type deser0036 struct{}

func Newdeser0036() *deser0036 {
    return &deser0036{}
}

func (e *deser0036) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "deser:done")
    return results, nil
}

func (e *deser0036) Name() string { return "deser0036" }
func (e *deser0036) Timestamp() time.Time { return time.Now() }
