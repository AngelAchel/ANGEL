package deser

import (
    "time"
)

type deser0110 struct{}

func Newdeser0110() *deser0110 {
    return &deser0110{}
}

func (e *deser0110) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "deser:done")
    return results, nil
}

func (e *deser0110) Name() string { return "deser0110" }
func (e *deser0110) Timestamp() time.Time { return time.Now() }
