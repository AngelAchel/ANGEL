package deser

import (
    "time"
)

type deser0020 struct{}

func Newdeser0020() *deser0020 {
    return &deser0020{}
}

func (e *deser0020) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "deser:done")
    return results, nil
}

func (e *deser0020) Name() string { return "deser0020" }
func (e *deser0020) Timestamp() time.Time { return time.Now() }
