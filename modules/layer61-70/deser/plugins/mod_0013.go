package deser

import (
    "time"
)

type deser0013 struct{}

func Newdeser0013() *deser0013 {
    return &deser0013{}
}

func (e *deser0013) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "deser:done")
    return results, nil
}

func (e *deser0013) Name() string { return "deser0013" }
func (e *deser0013) Timestamp() time.Time { return time.Now() }
