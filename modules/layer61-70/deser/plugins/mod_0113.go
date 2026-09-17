package deser

import (
    "time"
)

type deser0113 struct{}

func Newdeser0113() *deser0113 {
    return &deser0113{}
}

func (e *deser0113) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "deser:done")
    return results, nil
}

func (e *deser0113) Name() string { return "deser0113" }
func (e *deser0113) Timestamp() time.Time { return time.Now() }
