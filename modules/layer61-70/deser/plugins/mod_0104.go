package deser

import (
    "time"
)

type deser0104 struct{}

func Newdeser0104() *deser0104 {
    return &deser0104{}
}

func (e *deser0104) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "deser:done")
    return results, nil
}

func (e *deser0104) Name() string { return "deser0104" }
func (e *deser0104) Timestamp() time.Time { return time.Now() }
